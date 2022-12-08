package io

import (
	"testing"
)

func TestOf(t *testing.T) {
	var printed bool
	io := Of[string]("hello", func() { printed = true })
	result := io.Run()
	if result != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result)
	}
	if !printed {
		t.Error("Expected side effect to be executed")
	}
}

func TestBind(t *testing.T) {
	io1 := Of[string]("hello", func() {})
	io2 := io1.Bind(func(s string) IO[string] {
		return Of[string](s+" world", func() {})
	})
	result := io2.Run()
	if result != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", result)
	}
}

func TestRun(t *testing.T) {
	var printed bool
	io := Of[string]("hello", func() { printed = true })
	result := io.Run()
	if result != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result)
	}
	if !printed {
		t.Error("Expected side effect to be executed")
	}
}
