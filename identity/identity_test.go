package identity

import "testing"

func TestOf(t *testing.T) {
	identity := Of[string]("hello")
	result := identity.Unwrap()
	if result != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result)
	}
}

func TestBind(t *testing.T) {
	identity1 := Of[string]("hello")
	identity2 := identity1.Bind(func(s string) Identity[string] {
		return Of[string](s + " world")
	})
	result := identity2.Unwrap()
	if result != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", result)
	}
}

func TestValue(t *testing.T) {
	identity := Of[string]("hello")
	result := identity.Unwrap()
	if result != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result)
	}
}
