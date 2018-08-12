package sequences

import "testing"

func TestCycle(t *testing.T) {
	ConfirmCycle := func(s interface{}, l int) {
		var c int
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("cycle([%T], %v): failed with error %v", s, c, e)
			}
		}()
		for c = 1; c < 5; c++ {
			i := 0
			r := c * l
			if Cycle(s, c, func(interface{}) { i++ }); i != r {
				t.Fatalf("cycle([%T], %v): iteration count should be %v but is %v", s, c, r, i)
			}
		}
	}

	ConfirmCycle(enumerable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	ConfirmCycle(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	ConfirmCycle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
}
