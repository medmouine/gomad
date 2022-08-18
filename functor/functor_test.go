package functor_test

import (
	"strconv"
	"testing"

	. "github.com/medmouine/gomad/functor"
	"github.com/medmouine/gomad/identity"
)

func TestLift(t *testing.T) {
	t.Parallel()

	intToString := strconv.Itoa
	got := Lift(intToString)

	intFunctor := Functor[int]{
		Val: 1,
	}

	if got(intFunctor).Val != "1" {
		t.Errorf("Lift(intToString) failed")
	}

	got2 := Lift(identity.Identity[int])

	if got2(intFunctor).Val != 1 {
		t.Errorf("Lift(identity.Identity) failed")
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	intToString := strconv.Itoa

	intFunctor := Functor[int]{
		Val: 1,
	}

	got := Map(intFunctor, intToString)

	if got.Val != "1" {
		t.Errorf("Map(intToString) failed")
	}

	got2 := Map(intFunctor, identity.Identity[int])

	if got2.Val != 1 {
		t.Errorf("Map(identity.Identity) failed")
	}
}

func TestFunctor_Map(t *testing.T) {
	t.Parallel()

	intFunctor := Functor[int]{
		Val: 1,
	}

	got := intFunctor.Map(func(i int) int {
		return 5
	})

	if got.Val != 5 {
		t.Errorf("Map(intToString) failed")
	}
}
