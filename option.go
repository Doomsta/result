package result

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	Unwrap() T
	UnwrapOr(or T) T
}

type Some[T any] struct {
	v *T
}

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
