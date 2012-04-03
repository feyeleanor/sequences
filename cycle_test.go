package sequences

import "testing"

func TestCycle(t *testing.T) {
	ConfirmCycle := func(s interface{}, l int) {
		var c	int
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("cycle([%T], %v): failed with error %v", s, c, e)
			}
		}()
		for c = 1; c < 5; c++ {
			iterations := 0
			expected := c * l
			switch cycles := Cycle(s, c, func(i interface{}) { iterations++ }); {
			case cycles != c:				t.Fatalf("cycle([%T], %v): cycle count should be %v but is %v", s, c, c, cycles)
			case iterations != expected:	t.Fatalf("cycle([%T], %v): iteration count should be %v but is %v", s, c, expected, iterations)
			}
		}
	}

	ConfirmCycle(enumerable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	ConfirmCycle(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	ConfirmCycle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
}