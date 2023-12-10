package functional

// Basic errorable monad
// works with single return value
// even if error
type Maybe[T any] struct {
	value T
	err   string
}

func wrap_err[T any](value T, err string) *Maybe[T] {
	return &Maybe[T]{
		value: value,
		err:   err,
	}
}

func wrap[T any](value T) *Maybe[T] {
	return &Maybe[T]{
		value: value,
		err:   "",
	}
}

func (m *Maybe[T]) unwrap() T {
	return m.value
}

// TODO look at standard implementation https://github.com/OlegStotsky/go-monads/blob/master/maybe/maybe.go
