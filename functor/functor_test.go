package functor

import (
	"github.com/medmouine/gomad/identity"
	"strconv"
	"testing"
)

func TestLift(t *testing.T) {
	intToString := func(i int) string {
		return strconv.Itoa(i)
	}
	got := Lift(intToString)

	intFunctor := Functor[int]{
		val: 1,
	}

	if got(intFunctor).val != "1" {
		t.Errorf("Lift(intToString) failed")
	}

	got2 := Lift(identity.Identity[int])

	if got2(intFunctor).val != 1 {
		t.Errorf("Lift(identity.Identity) failed")
	}
}

func TestMap(t *testing.T) {
	intToString := func(i int) string {
		return strconv.Itoa(i)
	}

	intFunctor := Functor[int]{
		val: 1,
	}

	got := Map(intFunctor, intToString)

	if got.val != "1" {
		t.Errorf("Map(intToString) failed")
	}

	got2 := Map(intFunctor, identity.Identity[int])

	if got2.val != 1 {
		t.Errorf("Map(identity.Identity) failed")
	}
}
