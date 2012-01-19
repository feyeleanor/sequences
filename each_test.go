package gosequence

import(
	"strconv"
	"testing"
)

func TestEachIterable(t *testing.T) {
	var count	int

	ConfirmEach := func(s iterable_slice, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):	t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	S := iterable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})
}

func TestEachIndexable(t *testing.T) {
	var count	int
	var S		Indexable

	ConfirmEach := func(s Indexable, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != s.Len():	t.Fatalf("total iterations should be %v but are %v", s.Len(), count)
		}
	}

	S = indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	S = indexable_function(
		func(i int) interface{} {
			return i
		})
	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(S, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})
}

func TestEachSlice(t *testing.T) {
	var count	int

	ConfirmEach := func(s []int, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):			t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):		t.Fatalf("total iterations should be %v but is %v", len(s), count)
		}
	}

	I := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmEach(I, func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	ConfirmEach(I, func(i int, v interface{}) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(I, func(k, v interface{}) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(I, func(i int, v int) {
		switch {
		case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})

	ConfirmEach(I, func(k, v int) {
		switch {
		case k != count:			t.Fatalf("index %v erroneously reported as %v", count, k)
		case v != count:			t.Fatalf("element %v erroneously reported as %v", count, v)
		}
		count++
	})
}

func TestEachMap(t *testing.T) {
	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
	Each(M, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(M, func(k, v interface{}) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})

	Each(M, func(k, v int) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})

	Each(M, func(k interface{}, v int) {
		if k != v {
			t.Fatalf("index %v erroneously reported as %v", k, v)
		}
	})
}

func TestEachChannel(t *testing.T) {
	var index	int

	ConfirmEach := func(s []int, f interface{}) {
		var v		int
		var count	int

		C := make(chan int)
		go func() {
			for count, v = range s {
				C <- v
			}
			close(C)
		}()

		index = 0
		Each(C, f)
		if count != len(s) - 1 {
			t.Fatalf("total iterations should be %v but is %v", len(s) - 1, count)
		}
	}

	S := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	ConfirmEach(S, func(v interface{}) {
		if v != S[index] {
			t.Fatalf("index %v erroneously reported as %v", index, v)
		}
		index++
	})

	ConfirmEach(S, func(i int, v interface{}) {
		switch {
		case i != index:		t.Fatalf("index %v erroneously reported as %v", index, i)
		case v != S[i]:			t.Fatalf("value %v erroneously reported as %v", S[9], v)
		}
		index++
	})
}

func TestEachFunction(t *testing.T) {
	F1 := func(v interface{}) (r interface{}, ok bool) {
		if v.(int) < 10 {
			r = v
			ok = true
		}
		return
	}

	count := 0
	Each(F1, func(v interface{}) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	Each(F1, func(i interface{}, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F1, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	F2 := func(v int) (r int, ok bool) {
		if v < 10 {
			r = v
			ok = true
		}
		return
	}

	count = 0
	Each(F2, func(v int) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(F2, func(v interface{}) {
		if v != count {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	Each(F2, func(i interface{}, v int) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F2, func(i int, v interface{}) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F2, func(i int, v int) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	F3 := func(v int) (r string, ok bool) {
		if v < 10 {
			r = strconv.Itoa(v)
			ok = true
		}
		return
	}
	
	count = 0
	Each(F3, func(v string) {
		if v != strconv.Itoa(count) {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})

	count = 0
	Each(F3, func(v interface{}) {
		if v != strconv.Itoa(count) {
			t.Fatalf("index %v erroneously reported as %v", count, v)
		}
		count++
	})
	
	Each(F3, func(i interface{}, v string) {
		if i != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})

	Each(F3, func(i int, v string) {
		if strconv.Itoa(i) != v {
			t.Fatalf("index %v erroneously reported as %v", i, v)
		}
	})
}

func TestCycleIterable(t *testing.T) {
	ConfirmCycle := func(s iterable_slice, c int) {
		iterations := 0
		Cycle(s, c, func(i interface{}) {
			iterations++
		})
		if expected := c * len(s); iterations != expected {
			t.Fatalf("cycle(%v): iteration count should be %v but is %v", c, expected, iterations)
		}
	}

	S := iterable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmCycle(S, 1)
	ConfirmCycle(S, 2)
	ConfirmCycle(S, 3)
}

func TestCycleIndexable(t *testing.T) {
	ConfirmCycle := func(s indexable_slice, c int) {
		iterations := 0
		Cycle(s, c, func(i interface{}) {
			iterations++
		})
		if expected := c * len(s); iterations != expected {
			t.Fatalf("cycle(%v): iteration count should be %v but is %v", c, expected, iterations)
		}
	}

	S := indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmCycle(S, 1)
	ConfirmCycle(S, 2)
	ConfirmCycle(S, 3)
}

func TestCycleSlice(t *testing.T) {
	ConfirmCycle := func(s []int, c int) {
		iterations := 0
		Cycle(s, c, func(i interface{}) {
			iterations++
		})
		if expected := c * len(s); iterations != expected {
			t.Fatalf("cycle(%v): iteration count should be %v but is %v", c, expected, iterations)
		}
	}

	S := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ConfirmCycle(S, 1)
	ConfirmCycle(S, 2)
	ConfirmCycle(S, 3)
}