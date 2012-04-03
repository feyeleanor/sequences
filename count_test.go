package sequences

import (
	"testing"
)

func TestCountSlice(t *testing.T) {
	ConfirmCount := func(o, f, r interface{}) {
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("Count(%v, f) failed with error %v", o, e)
			}
		}()
		if x := Each(o, f); x != r {
			t.Fatalf("Count(%v, f) should be %v but is %v", o, r, x)
		}
	}

	f := func(i int) bool {
		return i > 0
	}

	ConfirmCount([]int{}, f, 0)
	ConfirmCount([]int{0}, f, 0)
	ConfirmCount([]int{1}, f, 1)
	ConfirmCount([]int{0, 1}, f, 1)
	ConfirmCount([]int{1, 2}, f, 2)

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmCount([]interface{}{"test"}, IsPositive, 0)
	ConfirmCount([]interface{}{1}, IsPositive, 1)
	ConfirmCount([]interface{}{"test", 1}, IsPositive, 1)
	ConfirmCount([]interface{}{0, 1}, IsPositive, 1)
	ConfirmCount([]interface{}{1, 2}, IsPositive, 2)
}


func TestCountFunction(t *testing.T) {
	t.Fatalf("Write tests")
/*	limit := 10
	value := 0

	ConfirmEach := func(F, f interface{}) {
		value = 0
		switch count, ok := Count(F, f); {
		case !ok:
			t.Fatalf("failed to perform iteration %v over %v", f, F)
		case value != limit:
			t.Fatalf("total items produced should be %v but are %v", limit, value)
		case count != limit:
			//t.Fatalf("total iterations should be %v but are %v", limit, count)
			panic(count)
		}
	}

	ConfirmEachFunction := func(F interface{}) {
		ConfirmEach(F, func(v interface{}) bool {})

		ConfirmEach(F, func(i int, v interface{}) {
			switch {
			case i != count:
				t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != i:
				t.Fatalf("value %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(F, func(i, v interface{}) {
			switch {
			case i != count:
				t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != i:
				t.Fatalf("value %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(F, func(v R.Value) {
			if v.Interface() != count {
				t.Fatalf("value %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(F, func(i int, v R.Value) {
			switch {
			case i != count:
				t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != i:
				t.Fatalf("value %v erroneously reported as %v", i, v.Interface())
			}
			count++
		})

		ConfirmEach(F, func(i interface{}, v R.Value) {
			switch {
			case i != count:
				t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != i:
				t.Fatalf("value %v erroneously reported as %v", i, v.Interface())
			}
			count++
		})

		ConfirmEach(F, func(i, v R.Value) {
			switch {
			case i.Interface() != count:
				t.Fatalf("index %v erroneously reported as %v", count, i.Interface())
			case v.Interface() != i.Interface():
				t.Fatalf("value %v erroneously reported as %v", i.Interface(), v.Interface())
			}
			count++
		})
	}

	ConfirmEachIntFunction := func(F interface{}) {
		ConfirmEachFunction(F)
		ConfirmEach(F, func(v int) {
			if v != count {
				t.Fatalf("value %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(F, func(i, v int) {
			switch {
			case i != count:
				t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != i:
				t.Fatalf("value %v erroneously reported as %v", i, v)
			}
			count++
		})
	}

	ConfirmEachFunction(func() (r interface{}, finished bool) {
		r = value
		if finished = value > 9; !finished {
			value++
		}
		return
	})

	ConfirmEachIntFunction(func() (r int, finished bool) {
		r = value
		if finished = value > 9; !finished {
			value++
		}
		return
	})
*/
}