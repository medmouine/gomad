package either_test

import (
	"reflect"
	"testing"

	"github.com/medmouine/gomad/either"
)

func TestLeft(t *testing.T) {
	t.Parallel()

	got := either.Left("foo")
	if !reflect.DeepEqual(got.IsLeft(), true) {
		t.Errorf("Left() = %v, want %v", got.IsLeft(), true)
	}

	if !reflect.DeepEqual(*got.Left(), "foo") {
		t.Errorf("Left() = %v, want %v", *got.Left(), "foo")
	}

	if !reflect.DeepEqual(got.IsRight(), false) {
		t.Errorf("Left() = %v, want %v", got.IsRight(), false)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Right() on Left did not panic")
		}

		t.Logf("PASS: Right() on left panic msg: %v", err)
	}()
	either.Left("foo").Right()
}
