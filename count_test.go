package sequences

import "testing"

func TestCountSlice(t *testing.T) {
	conditions := []func(interface{}) bool{GreaterThanZero, EqualZero, LessThanZero}

	ConfirmCount := func(o interface{}, r ...int) {
		i := 0
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("Count(%v, conditions[%v]) failed with error %v", o, i, e)
			}
		}()
		for _, c := range conditions {
			if x := Count(o, c); x != r[i] {
				t.Fatalf("Count(%v, %v) should be %v but is %v", o, interface{}(c), r[i], x)
			}
			i++
		}
	}

	ConfirmCount([]int{}, 0, 0, 0)
	ConfirmCount([]int{0}, 0, 1, 0)
	ConfirmCount([]int{1}, 1, 0, 0)
	ConfirmCount([]int{0, 1}, 1, 1, 0)
	ConfirmCount([]int{0, -1}, 0, 1, 1)
	ConfirmCount([]int{1, 2}, 2, 0, 0)
	ConfirmCount([]int{-2, -1}, 0, 0, 2)
	ConfirmCount([]int{-2, -1, 0, 1, 2}, 2, 1, 2)

	ConfirmCount([]interface{}{"test"}, 0, 0, 0)
	ConfirmCount([]interface{}{0}, 0, 1, 0)
	ConfirmCount([]interface{}{1}, 1, 0, 0)
	ConfirmCount([]interface{}{"test", 1}, 1, 0, 0)
	ConfirmCount([]interface{}{0, 1}, 1, 1, 0)
	ConfirmCount([]interface{}{1, 2}, 2, 0, 0)
	ConfirmCount([]interface{}{-2, -1}, 0, 0, 2)
	ConfirmCount([]interface{}{-2, -1, 0, 1, 2}, 2, 1, 2)
}

func TestCountFunction(t *testing.T) {
	t.Fatalf("Fix Tests")
	conditions := []func(interface{}) bool{GreaterThanZero, EqualZero, LessThanZero}
	limit := 10

	ConfirmCount := func(o interface{}, r ...int) {
		i := 0
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("Count(%v, conditions[%v]) failed with error %v", o, i, e)
			}
		}()
		for _, c := range conditions {
			if x := Count(o, c); x != r[i] {
				t.Fatalf("Count(%v, %v) should be %v but is %v", o, interface{}(c), r[i], x)
			}
			i++
		}
	}

	ConfirmCount(func(i int) (v int, ok bool) {
		return 0, false
	}, 0, 0, 0)

	ConfirmCount(func(i int) (v int, ok bool) {
		return i, i < limit
	}, limit-1, 1, 0)

	/*	ConfirmEachFunction := func(F interface{}) {
			ConfirmCount(F, func(v interface{}) bool {})

			ConfirmCount(F, func(i int, v interface{}) {
				switch {
				case i != count:
					t.Fatalf("index %v erroneously reported as %v", count, i)
				case v != i:
					t.Fatalf("value %v erroneously reported as %v", i, v)
				}
				count++
			})

			ConfirmCount(F, func(i, v interface{}) {
				switch {
				case i != count:
					t.Fatalf("index %v erroneously reported as %v", count, i)
				case v != i:
					t.Fatalf("value %v erroneously reported as %v", i, v)
				}
				count++
			})

			ConfirmCount(F, func(v R.Value) {
				if v.Interface() != count {
					t.Fatalf("value %v erroneously reported as %v", count, v.Interface())
				}
				count++
			})

			ConfirmCount(F, func(i int, v R.Value) {
				switch {
				case i != count:
					t.Fatalf("index %v erroneously reported as %v", count, i)
				case v.Interface() != i:
					t.Fatalf("value %v erroneously reported as %v", i, v.Interface())
				}
				count++
			})

			ConfirmCount(F, func(i interface{}, v R.Value) {
				switch {
				case i != count:
					t.Fatalf("index %v erroneously reported as %v", count, i)
				case v.Interface() != i:
					t.Fatalf("value %v erroneously reported as %v", i, v.Interface())
				}
				count++
			})

			ConfirmCount(F, func(i, v R.Value) {
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
			ConfirmCount(F, func(v int) {
				if v != count {
					t.Fatalf("value %v erroneously reported as %v", count, v)
				}
				count++
			})

			ConfirmCount(F, func(i, v int) {
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
