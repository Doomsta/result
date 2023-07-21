package result

type Result[T any] interface {
	IsOk() bool
	IsErr() bool
	Error() error
	Unwrap() T
}

type OK[T any] struct {
	value *T
}

type Failed[T any] struct {
	err error
}

func Ok[T any](value T) OK[T] {
	return OK[T]{
		value: &value,
	}
}

func Err[T any](err error) Failed[T] {
	return Failed[T]{
		err: err,
	}
}

func AsResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

func (r OK[T]) IsOk() bool {
	return true
}

func (r Failed[T]) IsOk() bool {
	return false
}

func (r OK[T]) IsErr() bool {
	return false
}

func (r Failed[T]) IsErr() bool {
	return true
}

func (r OK[T]) Error() error {
	return nil
}

func (r Failed[T]) Error() error {
	return r.err
}

func (r Failed[T]) Unwrap() T {
	panic("unwrap error:" + r.err.Error())
}

func (r OK[T]) Unwrap() T {
	return *r.value
}
