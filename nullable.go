package result

type Nullable[T any] struct {
	Valid bool
	Val   T
}

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		Val:   value,
		Valid: true,
	}
}

func (n Nullable[T]) ToOption() Option[T] {
	if n.Valid {
		return NewSome(n.Val)
	}
	return NewNone[T]()
}

func zeroValue[T any]() T {
	var zero T
	return zero
}
