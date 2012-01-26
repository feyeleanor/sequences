package sequences

import "testing"

func TestCycle(t *testing.T) {
	ConfirmCycle := func(s interface{}, l int) {
		for c := 1; c < 5; c++ {
			iterations := 0
			expected := c * l
			switch cycles, ok := Cycle(s, c, func(i interface{}) { iterations++ }); {
			case !ok:						t.Fatalf("cycle([%T], %v): iteration failed to complete", s, c)
			case cycles != c:				t.Fatalf("cycle([%T], %v): cycle count should be %v but is %v", s, c, c, cycles)
			case iterations != expected:	t.Fatalf("cycle([%T], %v): iteration count should be %v but is %v", s, c, expected, iterations)
			}
		}
	}

	ConfirmCycle(iterable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	ConfirmCycle(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	ConfirmCycle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
}