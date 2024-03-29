package collection_utils

import "testing"

func TestFoldLeft(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}

	testCases := []struct {
		desc string
		coll []int
		x    int
		want int
	}{
		{
			desc: "empty collection",
			coll: []int{},
			x:    0,
			want: 0,
		},
		{
			desc: "single-element collection",
			coll: []int{1},
			x:    0,
			want: 1,
		},
		{
			desc: "multi-element collection",
			coll: []int{1, 2, 3, 4, 5},
			x:    0,
			want: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := FoldLeft(tc.coll, op, tc.x)
			if got != tc.want {
				t.Errorf("FoldLeft(%v, %v, %v) = %v; want %v", tc.coll, "op", tc.x, got, tc.want)
			}
		})
	}
}

func TestFoldRight(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}

	testCases := []struct {
		desc string
		coll []int
		x    int
		want int
	}{
		{
			desc: "empty collection",
			coll: []int{},
			x:    0,
			want: 0,
		},
		{
			desc: "single-element collection",
			coll: []int{1},
			x:    0,
			want: 1,
		},
		{
			desc: "multi-element collection",
			coll: []int{1, 2, 3, 4, 5},
			x:    0,
			want: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := FoldRight(tc.coll, op, tc.x)
			if got != tc.want {
				t.Errorf("FoldRight(%v, %v, %v) = %v; want %v", tc.coll, "op", tc.x, got, tc.want)
			}
		})
	}
}
