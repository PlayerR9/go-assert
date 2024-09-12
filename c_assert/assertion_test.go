package c_assert

import (
	"testing"

	gda "github.com/PlayerR9/go-debug/assert"
)

func TestAssertion(t *testing.T) {
	foo := "foo"

	var actual_err gda.ErrAssertFail

	res := AssertThat(NewVariable("foo"), NewOrderedAssert(foo).In("bar", "foo", "baz")).Not().Error()
	if res == nil {
		t.Errorf("expected an error but got none")
	} else {
		tmp, ok := res.(*gda.ErrAssertFail)
		if !ok {
			t.Errorf("expected an error of type *ErrAssertFail but got %T", res)
		}

		actual_err = *tmp
	}

	msg_to_check := "expected \"foo\" to not be one of {bar, baz, foo}; got foo instead"

	if actual_err.Msg != msg_to_check {
		t.Errorf("expected error %q but got %q instead", msg_to_check, actual_err.Msg)
	}
}
