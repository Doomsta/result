package result

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleResultSwitch() {
	fn := func() Result[string] {
		return Ok[string]("hello")
	}

	switch v := fn().(type) {
	case Failed[string]:
		fmt.Println(v.Error())
	case OK[string]:
		fmt.Println(v.Unwrap())
	}
	// Output: hello
}

func ExampleResultIf() {
	fn := func() Result[string] {
		return Ok[string]("hello")
	}

	if v := fn(); v.IsOk() {
		fmt.Println(v.Unwrap())
	}
	// Output: hello
}

func ExampleResultTypeAssert() {
	fn := func() Result[string] {
		return Ok[string]("hello")
	}

	if v, ok := fn().(OK[string]); ok {
		fmt.Println(v.Unwrap())
	}
	// Output: hello
}

func TestResult(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() Result[int]
		expected int
		error    bool
	}{
		{
			name:  "ok",
			fn:    demoOKFunc[int](1),
			error: false,
		},
		{
			name:  "err",
			fn:    demoFailedFunc[int](errors.New("failed")),
			error: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch v := test.fn().(type) {
			case OK[int]:
				if test.error {
					panic("invalid test case")
				}
				assert.True(t, v.IsOk())
				assert.False(t, v.IsErr())
				assert.NoError(t, v.Error(), test.error)
			case Failed[int]:
				if !test.error {
					panic("invalid test case")
				}
				assert.False(t, v.IsOk())
				assert.True(t, v.IsErr())
				assert.Error(t, v.Error(), test.fn().Error())
			}
		})
	}
}

func demoOKFunc[T any](v T) func() Result[T] {
	return func() Result[T] {
		return Ok[T](v)
	}
}

func demoFailedFunc[T any](err error) func() Result[T] {
	return func() Result[T] {
		return Err[T](err)
	}
}
