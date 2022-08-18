package either_test

import (
	"reflect"
	"testing"

	"github.com/medmouine/gomad/either"
)

func TestRight(t *testing.T) {
	t.Parallel()

	got := either.Right("foo")
	if !reflect.DeepEqual(got.IsLeft(), false) {
		t.Errorf("Right() = %v, want %v", got.IsLeft(), false)
	}

	if !reflect.DeepEqual(*got.Right(), "foo") {
		t.Errorf("Right() = %v, want %v", *got.Right(), "foo")
	}

	if !reflect.DeepEqual(got.IsRight(), true) {
		t.Errorf("Right() = %v, want %v", got.IsRight(), true)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Left() on Right did not panic")
		}

		t.Logf("PASS: Left() on Right panic msg: %v", err)
	}()
	either.Right("foo").Left()
}
