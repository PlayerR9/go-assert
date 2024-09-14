package assert

import (
	"fmt"

	gda "github.com/PlayerR9/go-debug/assert"
)

// AssertDeref tries to dereference an element and panics if it is nil.
//
// Parameters:
//   - elem: the element to dereference.
//   - param_name: the name of the parameter.
//
// Returns:
//   - T: the dereferenced element.
func AssertDeref[T any](elem *T, is_param bool, name string) T {
	if elem != nil {
		return *elem
	}

	var msg string

	if is_param {
		msg = "parameter (" + name + ")"
	} else {
		msg = "variable (" + name + ")"
	}

	msg += " expected to not be nil"

	panic(gda.NewErrAssertFailed(msg))
}

// AssertTypeOf panics if the element is not of the expected type.
//
// Parameters:
//   - elem: the element to check.
//   - var_name: the name of the variable.
//   - allow_nil: if the element can be nil.
func AssertTypeOf[T any](elem any, target string, allow_nil bool) {
	if elem == nil {
		if !allow_nil {
			msg := fmt.Sprintf("expected %q to be of type %T, got nil instead", target, *new(T))

			panic(gda.NewErrAssertFailed(msg))
		}

		return
	}

	_, ok := elem.(T)
	if !ok {
		msg := fmt.Sprintf("expected %q to be of type %T, got %T instead", target, *new(T), elem)

		panic(gda.NewErrAssertFailed(msg))
	}
}

// AssertConv tries to convert an element to the expected type and panics if it is not possible.
//
// Parameters:
//   - elem: the element to check.
//   - var_name: the name of the variable.
//
// Returns:
//   - T: the converted element.
func AssertConv[T any](elem any, target string) T {
	if elem == nil {
		msg := fmt.Sprintf("expected %q to be of type %T, got nil instead", target, *new(T))

		panic(gda.NewErrAssertFailed(msg))
	}

	res, ok := elem.(T)
	if !ok {
		msg := fmt.Sprintf("expected %q to be of type %T, got %T instead", target, *new(T), elem)

		panic(gda.NewErrAssertFailed(msg))
	}

	return res
}
