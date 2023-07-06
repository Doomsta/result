package result

// Option is a type that represents an optional value:
// every Option is either Some and contains a value, or None, and does not.
// Option types have a number of uses:
//   - Initial values
//   - Return value for otherwise reporting a generic error, where None is returned on error
//   - Optional struct fields
//   - Optional function arguments
//   - Nullable pointers
type Option[T any] interface {
	option() // dummy method to prevent user from implementing this interface

	IsSome() bool
	IsNone() bool
	Unwrap() T
	UnwrapOr(or T) T
}

// Some value of type T
type Some[T any] struct {
	v *T
}

// None no value. Panics on unwrap
type None[T any] struct{}

var _ Option[int] = Some[int]{}
var _ Option[int] = None[int]{}

func NewSome[T any](value T) Some[T] {
	return Some[T]{
		v: &value,
	}
}

func NewNone[T any]() None[T] {
	return None[T]{}
}

func (o Some[T]) option() {}
func (o None[T]) option() {}

func (o Some[T]) IsSome() bool {
	return true
}
func (o None[T]) IsSome() bool {
	return false
}

func (o Some[T]) IsNone() bool {
	return false
}
func (o None[T]) IsNone() bool {
	return true
}

func (o Some[T]) Unwrap() T {
	return *o.v
}
func (o None[T]) Unwrap() T {
	panic("can't unwrap none value")
}

func (o Some[T]) UnwrapOr(T) T {
	return *o.v
}

func (o None[T]) UnwrapOr(or T) T {
	return or
}
