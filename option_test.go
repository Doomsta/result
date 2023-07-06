package result

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOption(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() Option[int]
		expected bool
		value    int
	}{
		{
			name:     "some(1)",
			fn:       demoFuncSome[int](1),
			expected: true,
			value:    1,
		}, {
			name:     "some(2)",
			fn:       demoFuncSome[int](2),
			expected: true,
			value:    2,
		}, {
			name:     "none",
			fn:       demoFuncNone[int](),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch v := test.fn().(type) {
			case Some[int]:
				if !test.expected {
					panic("can't happen")
				}
				assert.True(t, v.IsSome())
				assert.False(t, v.IsNone())
				assert.NotPanics(t, func() {
					v.Unwrap()
				})
				assert.Equal(t, test.value, v.Unwrap())
			case None[int]:
				if test.expected {
					panic("can't happen")
				}
				assert.False(t, v.IsSome())
				assert.True(t, v.IsNone())
				assert.Panics(t, func() {
					v.Unwrap()
				})
			}
		})
	}
}

func demoFuncSome[T any](v T) func() Option[T] {
	return func() Option[T] {
		return NewSome[T](v)
	}
}

func demoFuncNone[T any]() func() Option[T] {
	return func() Option[T] {
		return NewNone[T]()
	}
}
