package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertFail)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertFail but got %T", r)
		}

		if err.Msg != "foo is not \"bar\"" {
			t.Errorf("Expected 'foo is not \"bar\"' but got %s", err.Msg)
		}
	}()

	foo := "foo"

	AssertF(foo == "bar", "foo is not %q", "bar")
}

func TestAssertF(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertFail)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertionFail but got %T", r)
		}

		if err.Msg != "foo is not \"bar\"" {
			t.Errorf("Expected 'foo is not bar' but got %s", err.Msg)
		}
	}()

	foo := "foo"
	bar := "bar"

	AssertF(foo == bar, "%s is not %s", foo, bar)
}

func TestAssertErr(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertFail)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertFail but got %T", r)
		}

		if err.Msg != "function my_function(foo, \"bar\") returned the error: <err>" {
			t.Errorf("Expected 'function my_function(foo, \"bar\") returned the error: <err>' but got %s", err.Msg)
		}
	}()

	foo := "foo"

	my_function := func(left, right string) error {
		t.Logf("my_function(%s, %s)", left, right)

		return errors.New("<err>")
	}

	err := my_function(foo, "bar")
	AssertErr(err, "my_function(foo, %q)", "bar")
}

func TestAssertOk(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertFail)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertFail but got %T", r)
		}

		if err.Msg != "function my_function(foo, \"bar\") returned false while true was expected" {
			t.Errorf("Expected 'function my_function(foo, \"bar\") returned false while true was expected' but got %s", err.Msg)
		}
	}()

	foo := "foo"

	my_function := func(left, right string) bool {
		t.Logf("my_function(%s, %s)", left, right)

		return false
	}

	ok := my_function(foo, "bar")
	AssertOk(ok, "my_function(foo, %q)", "bar")
}
