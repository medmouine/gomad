package magma

import "testing"

func TestConcat(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}

	testCases := []struct {
		desc   string
		values []int
		x      int
		want   int
	}{
		{
			desc:   "empty set",
			values: []int{},
			x:      0,
			want:   0,
		},
		{
			desc:   "single-element set",
			values: []int{1},
			x:      0,
			want:   1,
		},
		{
			desc:   "multi-element set",
			values: []int{1, 2, 3, 4, 5},
			x:      0,
			want:   15,
		},
	}

	// run the tests
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// create a Magma instance
			mg := Of(tc.values, op)

			// test the Concat method
			got := mg.Concat(tc.x)
			if got != tc.want {
				t.Errorf("Concat(%v) = %v; want %v", tc.x, got, tc.want)
			}
		})
	}
}
