package sequences

import (
	R "reflect"
	"testing"
)

func TestLen(t *testing.T) {
	ConfirmLen := func(o interface{}, r int) {
		if l := Len(o); l != r {
			t.Fatalf("Len(%v) should be %v but is %v", o, r, l)
		}
	}
	ConfirmLen(0, 1)
	ConfirmLen(([]int)(nil), 0)
	ConfirmLen([]int{}, 0)
	ConfirmLen([]int{0}, 1)
	ConfirmLen([]int{0, 1, 2}, 3)
	ConfirmLen(R.ValueOf(0), 1)
	ConfirmLen(R.ValueOf(([]int)(nil)), 0)
	ConfirmLen(mappable_slice{0, 1, 2, 3, 4}, 5)
	ConfirmLen(map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}, 5)
}

func TestCap(t *testing.T) {
	ConfirmCap := func(o interface{}, r int) {
		if c := Cap(o); c != r {
			t.Fatalf("Cap(%v) should be %v but is %v", o, r, c)
		}
	}
	ConfirmCap(nil, 0)
	ConfirmCap(0, 1)
	ConfirmCap(([]int)(nil), 0)
	ConfirmCap([]int{}, 0)
	ConfirmCap([]int{0}, 1)
	ConfirmCap([]int{0, 1, 2}, 3)
	ConfirmCap(R.ValueOf(0), 1)
	ConfirmCap(R.ValueOf(([]int)(nil)), 0)
	ConfirmCap(mappable_slice{0, 1, 2, 3, 4}, 5)
	ConfirmCap(map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}, 0)
}