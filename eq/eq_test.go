package eq

import "testing"

func TestFromEquals(t *testing.T) {
	intEq := func(x int, y int) bool { return x == y }

	got := FromEquals(intEq)

	if got.Equals(1, 1) != true {
		t.Errorf("Expected 1 == 1 to be true")
	}
}

type Person struct {
	Name string
}

func Test_eq_Equals(t *testing.T) {
	intEq := func(x int, y int) bool { return x == y }
	got := FromEquals(intEq)

	if got.Equals(1, 2) {
		t.Errorf("eq.Equals() = %v, want %v", got.Equals(1, 2), false)
	}
	if !got.Equals(1, 1) {
		t.Errorf("eq.Equals() = %v, want %v", got.Equals(1, 1), true)
	}

	strEq := func(x string, y string) bool { return x == y }
	got2 := FromEquals(strEq)

	if got2.Equals("a", "b") {
		t.Errorf("eq.Equals() = %v, want %v", got2.Equals("a", "b"), false)
	}
	if !got2.Equals("a", "a") {
		t.Errorf("eq.Equals() = %v, want %v", got2.Equals("a", "a"), true)
	}

	personEq := func(x Person, y Person) bool { return x == y || x.Name == y.Name }
	got3 := FromEquals(personEq)

	if got3.Equals(Person{"a"}, Person{"b"}) {
		t.Errorf("eq.Equals() = %v, want %v", got3.Equals(Person{"a"}, Person{"b"}), false)
	}

	if !got3.Equals(Person{"a"}, Person{"a"}) {
		t.Errorf("eq.Equals() = %v, want %v", got3.Equals(Person{"a"}, Person{"a"}), true)
	}
}
