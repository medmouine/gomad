package either

import (
	"reflect"
	"testing"
)

func TestLeft(t *testing.T) {
	got := Left("foo")

	if !reflect.DeepEqual(got.IsLeft(), true) {
		t.Errorf("Left() = %v, want %v", got.IsLeft(), true)
	}
	if !reflect.DeepEqual(got.Left(), "foo") {
		t.Errorf("Left() = %v, want %v", got.Left(), "foo")
	}
	if !reflect.DeepEqual(got.IsRight(), false) {
		t.Errorf("Left() = %v, want %v", got.IsRight(), false)
	}
}
