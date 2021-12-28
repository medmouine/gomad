package either

import (
	"reflect"
	"testing"
)

func TestRight(t *testing.T) {
	got := Right("foo")

	if !reflect.DeepEqual(got.IsLeft(), false) {
		t.Errorf("Right() = %v, want %v", got.IsLeft(), false)
	}
	if !reflect.DeepEqual(got.Right(), "foo") {
		t.Errorf("Right() = %v, want %v", got.Right(), "foo")
	}
	if !reflect.DeepEqual(got.IsRight(), true) {
		t.Errorf("Right() = %v, want %v", got.IsRight(), true)
	}
}
