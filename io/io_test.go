package io

import (
	"reflect"
	"testing"
)

func TestOf(t *testing.T) {
	got := Of("foo")

	if !reflect.DeepEqual(got.Call(), "foo") {
		t.Errorf("Of(\"foo\") = %v, want %v", got.Call(), "foo")
	}
}

func TestFrom(t *testing.T) {
	fn := func() string {
		return "foo"
	}
	got := From(fn)

	if !reflect.DeepEqual(got.Call(), "foo") {
		t.Errorf("From() = %v, want %v", got, "foo")
	}
}

func TestMap(t *testing.T) {
	fn1 := func() string {
		return "foo"
	}
	fn2 := func(s string) int {
		return len(s)
	}
	got := Map(From(fn1), fn2)

	if !reflect.DeepEqual(got.Call(), 3) {
		t.Errorf("Map() = %v, want %v", got.Call(), 3)
	}
}

func Test_io_Map(t *testing.T) {
	fn1 := func() string {
		return "foo"
	}
	fn2 := func(s string) string {
		return s + "bar"
	}
	got := From(fn1).Map(fn2)

	if !reflect.DeepEqual(got.Call(), "foobar") {
		t.Errorf("Map() = %v, want %v", got.Call(), "foobar")
	}
}
