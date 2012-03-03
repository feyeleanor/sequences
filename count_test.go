package sequences

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	ConfirmCount := func(o, r interface{}, f func(interface{}) bool) {
		switch x, e := Count(o, f); {
		case e != nil:
			t.Fatalf("Count(%v, f) failed with error %v", o, e)
		case !reflect.DeepEqual(r, x):
			t.Fatalf("Count(%v, f) should be %v but is %v", o, r, x)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmCount([]int{}, 0, IsPositive)
	ConfirmCount([]interface{}{}, 0, IsPositive)

	ConfirmCount([]int{0}, 0, IsPositive)
	ConfirmCount([]int{1}, 1, IsPositive)
	ConfirmCount([]interface{}{"test"}, 0, IsPositive)
	ConfirmCount([]interface{}{1}, 1, IsPositive)


	ConfirmCount([]int{0, 1}, 1, IsPositive)
	ConfirmCount([]int{1, 2}, 2, IsPositive)
	ConfirmCount([]interface{}{"test", 1}, 1, IsPositive)
	ConfirmCount([]interface{}{1, 2}, 2, IsPositive)
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

	ConfirmVariadicEach := func(F, f interface{}) {
		value = 0
		switch count, ok := Count(F, f); {
		case !ok:
			t.Fatalf("failed to perform iteration %v over %v", f, F)
		case count != 1:
			t.Fatalf("total iterations should be 1 but are %v", count)
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

		ConfirmVariadicEach(F, func(v ...interface{}) {
			switch {
			case count != 0:
				t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != limit:
				//t.Fatalf("variadic slice generator %v erroneously passed as %v", F, v)
				panic(v)
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

		ConfirmVariadicEach(F, func(v ...R.Value) {
			switch {
			case count != 0:
				t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != limit:
				t.Fatalf("variadic slice generator %v erroneously passed as %v", F, v)
			}
			count++
		})
	}

	ConfirmEachIntFunction := func(F interface{}) {
		ConfirmEachFunction(F)

		ConfirmVariadicEach(F, func(v ...int) {
			switch {
			case count != 0:
				t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != limit:
				t.Fatalf("variadic slice generator %v erroneously passed as %v", F, v)
			}
			count++
		})

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