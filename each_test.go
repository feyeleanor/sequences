package sequences

import(
	R "reflect"
	"testing"
)

func TestEachPrimitive(t *testing.T) {
	var count	int

	ConfirmEach := func(s, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != Len(s):	t.Fatalf("total iterations should be %v but are %v", Len(s), count)
		}
	}

	ConfirmVariadicEach := func(s, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Logf("failed to perform iteration %v over %v", f, s)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachBuiltin := func(s interface{}) {
		ConfirmEach(s, func(v interface{}) {
			if v != AtOffset(s, count) {
				panic(s)
			}
			count++
		})

		ConfirmEach(s, func(i int, v interface{}) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v interface{}) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v != AtOffset(s, count):		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v R.Value) {
			if v.Interface() != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(s, func(i int, v R.Value) {
			switch {
			case i != count:						t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != AtOffset(s, i):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k interface{}, v R.Value) {
			switch {
			case k != count:							t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != AtOffset(s, count):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v R.Value) {
			switch {
			case k.Interface() != count:				t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != AtOffset(s, count):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})
	}

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...bool) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v bool) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v bool) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]bool{true, false, false, true, true})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...complex64) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v complex64) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v complex64) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]complex64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...complex128) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v complex128) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v complex128) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]complex128{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...error) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v error) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v error) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]error{Error(0), Error(1), Error(2), Error(3), Error(4), Error(5), Error(6), Error(7), Error(8), Error(9)})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...float32) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v float32) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v float32) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...float64) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v float64) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v float64) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i, v int) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int8) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int8) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v int8) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int16) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int16) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v int16) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int32) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int32) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v int32) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int64) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int64) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v int64) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	ConfirmEachBuiltin([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...string) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v string) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v string) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]string{"A", "B", "C", "D", "E", "F"})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v uint) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v uint) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint8) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v uint8) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v uint8) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint16) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v uint16) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v uint16) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint32) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v uint32) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v uint32) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uint64) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v uint64) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v uint64) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...uintptr) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v uintptr) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v uintptr) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}([]uintptr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	func(s []R.Value) {
		ConfirmEach(s, func(v interface{}) {
			if v != s[count].Interface() {
				panic(s)
			}
			count++
		})

		ConfirmEach(s, func(i int, v interface{}) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s[i].Interface():			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v interface{}) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v != s[count].Interface():		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v R.Value) {
			if v.Interface() != s[count].Interface() {
				t.Fatalf("element %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(s, func(i int, v R.Value) {
			switch {
			case i != count:							t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != s[i].Interface():		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k interface{}, v R.Value) {
			switch {
			case k != count:							t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s[count].Interface():	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v R.Value) {
			switch {
			case k.Interface() != count:				t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s[count].Interface():	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})
	}([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4), R.ValueOf(5), R.ValueOf(6), R.ValueOf(7), R.ValueOf(8), R.ValueOf(9)})

	func(s interface{}) {
		ConfirmEachBuiltin(s)
		ConfirmVariadicEach(s, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int) {
			if v != AtOffset(s, count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i, v int) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != AtOffset(s, i):			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}(R.ValueOf([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

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

func TestEachEnumerable(t *testing.T) {
	var count	int

	ConfirmEach := func(s enumerable_slice, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):	t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	S := enumerable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
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

	ConfirmEach := func(s Indexable, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != Len(s):	t.Fatalf("total iterations should be %v but are %v", Len(s), count)
		}
	}

	ConfirmVariadicEach := func(s Indexable, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachIndexable := func(s Indexable) {
		ConfirmEach(s, func(v interface{}) {
			if v != s.AtOffset(count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v interface{}) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s.AtOffset(i):					t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v interface{}) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v != s.AtOffset(count):				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v R.Value) {
			if v.Interface() != s.AtOffset(count) {
				t.Fatalf("element %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(s, func(i int, v R.Value) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != s.AtOffset(i):		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k interface{}, v R.Value) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s.AtOffset(count):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v R.Value) {
			switch {
			case k.Interface() != count:		t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s.AtOffset(count):	t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int) {
			if v != s.AtOffset(count) {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i, v int) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s.AtOffset(i):					t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}

	ConfirmEachIndexable(indexable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachIndexable(indexable_function(func(i int) interface{} { return i }))
}

func TestEachMappable(t *testing.T) {
	var count	int

	ConfirmEach := func(m Mappable, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != Len(m):	t.Fatalf("total iterations should be %v but are %v", Len(m), count)
		}
	}

	ConfirmVariadicEach := func(m Mappable, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachMappable := func(m Mappable) {
		ConfirmEach(m, func(v interface{}) {
			count++
		})

		ConfirmEach(m, func(i int, v interface{}) {
			if v != m.StoredAs(i) {
				t.Fatalf("element %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(m, func(k, v interface{}) {
			if v != m.StoredAs(k) {
				t.Fatalf("element %v erroneously reported as %v", k, v)
			}
			count++
		})

		ConfirmVariadicEach(m, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != Len(m):				t.Fatalf("variadic slice %v erroneously passed as %v", m, v)
			}
			count++
		})

		ConfirmEach(m, func(v R.Value) {
			count++
		})

		ConfirmEach(m, func(i int, v R.Value) {
			if v.Interface() != m.StoredAs(i) {
				t.Fatalf("element %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(m, func(k interface{}, v R.Value) {
			if v.Interface() != m.StoredAs(k) {
				t.Fatalf("element %v erroneously reported as %v", k, v)
			}
			count++
		})

		ConfirmEach(m, func(k, v R.Value) {
			if v.Interface() != m.StoredAs(k.Interface()) {
				t.Fatalf("element %v erroneously reported as %v", k.Interface(), v.Interface())
			}
			count++
		})

		ConfirmVariadicEach(m, func(v ...R.Value) {
			switch {
			case count == 1:					t.Fatalf("variadic function erroneously called more than once")
			case len(v) != Len(m):				t.Fatalf("variadic slice %v erroneously passed as %v", m, v)
			}
			count++
		})

		ConfirmVariadicEach(m, func(v ...int) {
			switch {
			case count == 1:					t.Fatalf("variadic function erroneously called more than once")
			case len(v) != Len(m):				t.Fatalf("variadic slice %v erroneously passed as %v", m, v)
			}
			count++
		})

		ConfirmEach(m, func(v int) {
			count++
		})

		ConfirmEach(m, func(i, v int) {
			if v != m.StoredAs(i) {
				t.Fatalf("element %v erroneously reported as %v", i, v)
			}
			count++
		})
	}

	ConfirmEachMappable(mappable_slice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ConfirmEachMappable(mappable_function(func(i int) interface{} { return i }))
	ConfirmEachMappable(mappable_map{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9})

	m := mappable_string_map{"0": 0, "1": 1, "2": 2, "3": 3}
	ConfirmEach(m, func(k string, v interface{}) {
		if v != m.StoredAs(k) {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(m, func(k string, v R.Value) {
		if v.Interface() != m.StoredAs(k) {
			t.Fatalf("element %v erroneously reported as %v", k, v.Interface())
		}
		count++
	})
}

func TestEachSlice(t *testing.T) {
	var count	int

	ConfirmEach := func(s []int, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):	t.Fatalf("total iterations should be %v but are %v", len(s), count)
		}
	}

	ConfirmVariadicEach := func(s []int, f interface{}) {
		count = 0
		switch {
		case !Each(s, f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}


	ConfirmEachSlice := func(s []int) {
		ConfirmEach(s, func(v interface{}) {
			if v != s[count] {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i int, v interface{}) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s[i]:						t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v interface{}) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v != s[count]:					t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v R.Value) {
			if v.Interface() != s[count] {
				t.Fatalf("element %v erroneously reported as %v", count, v.Interface())
			}
			count++
		})

		ConfirmEach(s, func(i int, v R.Value) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != s[i]:			t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k interface{}, v R.Value) {
			switch {
			case k != count:					t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s[count]:		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(k, v R.Value) {
			switch {
			case k.Interface() != count:		t.Fatalf("index %v erroneously reported as %v", count, k)
			case v.Interface() != s[count]:		t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmVariadicEach(s, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != len(s):				t.Fatalf("variadic slice %v erroneously passed as %v", s, v)
			}
			count++
		})

		ConfirmEach(s, func(v int) {
			if v != s[count] {
				t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(s, func(i, v int) {
			switch {
			case i != count:					t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != s[i]:						t.Fatalf("element %v erroneously reported as %v", count, v)
			}
			count++
		})
	}

	ConfirmEachSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestEachMap(t *testing.T) {
	var count	int

	ConfirmEach := func(m, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != Len(m):	t.Fatalf("total iterations should be %v but are %v", Len(m), count)
		}
	}

	ConfirmVariadicEach := func(m, f interface{}) {
		count = 0
		switch {
		case !Each(m, f):		t.Fatalf("failed to perform iteration %v over %v", f, m)
		case count != 1:		t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	M := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}

	ConfirmEach(M, func(v interface{}) {
		count++
	})

	ConfirmEach(M, func(i int, v interface{}) {
		if v != M[i] {
			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
	})

	ConfirmEach(M, func(k, v interface{}) {
		if v != M[k.(int)] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmVariadicEach(M, func(v ...interface{}) {
		switch {
		case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
		case len(v) != Len(M):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", M, v)
		}
		count++
	})

	ConfirmEach(M, func(v R.Value) {
		count++
	})

	ConfirmEach(M, func(i int, v R.Value) {
		if v.Interface() != M[i] {
			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
	})

	ConfirmEach(M, func(k interface{}, v R.Value) {
		if v.Interface() != M[k.(int)] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(M, func(k, v R.Value) {
		if v.Interface() != M[int(k.Int())] {
			t.Fatalf("element %v erroneously reported as %v", k.Interface(), v.Interface())
		}
		count++
	})

	ConfirmVariadicEach(M, func(v ...R.Value) {
		switch {
		case count == 1:					t.Fatalf("variadic function erroneously called more than once")
		case len(v) != Len(M):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", M, v)
		}
		count++
	})

	ConfirmVariadicEach(M, func(v ...int) {
		switch {
		case count == 1:					t.Fatalf("variadic function erroneously called more than once")
		case len(v) != Len(M):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", M, v)
		}
		count++
	})

	ConfirmEach(M, func(v int) {
		count++
	})

	ConfirmEach(M, func(i, v int) {
		if v != M[i] {
			t.Fatalf("element %v erroneously reported as %v", i, v)
		}
		count++
	})

	MS := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	ConfirmEach(MS, func(k string, v interface{}) {
		if v != MS[k] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})

	ConfirmEach(MS, func(k string, v R.Value) {
		if v.Interface() != MS[k] {
			t.Fatalf("element %v erroneously reported as %v", k, v)
		}
		count++
	})
}

func TestEachChannel(t *testing.T) {
	var count	int
	var index	int

	Generate := func(s []int) (c chan int) {
		c = make(chan int)
		go func() {
			for _, v := range s {
				c <- v
				count++
			}
			close(c)
		}()
		return c
	}

	ConfirmEach := func(s []int, f interface{}) {
		count = 0
		index = 0
		switch {
		case !Each(Generate(s), f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):			t.Fatalf("total items produced should be %v but are %v", len(s), count)
		case index != len(s):			t.Fatalf("total iterations should be %v but are %v", len(s), index)
		}
	}

	ConfirmVariadicEach := func(s []int, f interface{}) {
		count = 0
		index = 0
		switch {
		case !Each(Generate(s), f):		t.Fatalf("failed to perform iteration %v over %v", f, s)
		case count != len(s):			t.Fatalf("total items produced should be %v but are %v", len(s), count)
		case index != 1:				t.Fatalf("total iterations should be 1 but are %v", index)
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
		case v != S[i]:			t.Fatalf("value %v erroneously reported as %v", S[i], v)
		}
		index++
	})

	ConfirmEach(S, func(i, v interface{}) {
		switch {
		case i != index:		t.Fatalf("index %v erroneously reported as %v", index, i)
		case v != S[i.(int)]:	t.Fatalf("value %v erroneously reported as %v", S[i.(int)], v)
		}
		index++
	})

	ConfirmVariadicEach(S, func(v ...interface{}) {
		switch {
		case index != 0:					t.Fatalf("variadic function erroneously called %v times", index)
		case len(v) != Len(S):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", S, v)
		}
		index++
	})

	ConfirmEach(S, func(v R.Value) {
		if v.Interface() != S[index] {
			t.Fatalf("index %v erroneously reported as %v", index, v.Interface())
		}
		index++
	})

	ConfirmEach(S, func(i int, v R.Value) {
		switch {
		case i != index:					t.Fatalf("index %v erroneously reported as %v", index, i)
		case v.Interface() != S[i]:			t.Fatalf("value %v erroneously reported as %v", S[i], v.Interface())
		}
		index++
	})

	ConfirmEach(S, func(i interface{}, v R.Value) {
		switch {
		case i != index:					t.Fatalf("index %v erroneously reported as %v", index, i)
		case v.Interface() != S[i.(int)]:	t.Fatalf("value %v erroneously reported as %v", S[i.(int)], v.Interface())
		}
		index++
	})

	ConfirmEach(S, func(i, v R.Value) {
		switch x := int(i.Int()); {
		case x != index:				t.Fatalf("index %v erroneously reported as %v", index, x)
		case v.Interface() != S[x]:		t.Fatalf("value %v erroneously reported as %v", S[x], v.Interface())
		}
		index++
	})

	ConfirmVariadicEach(S, func(v ...R.Value) {
		switch {
		case index != 0:					t.Fatalf("variadic function erroneously called %v times", index)
		case len(v) != Len(S):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", S, v)
		}
		index++
	})

	ConfirmVariadicEach(S, func(v ...int) {
		switch {
		case index != 0:					t.Fatalf("variadic function erroneously called %v times", index)
		case len(v) != Len(S):				t.Fatalf("variadic slice of keys for %v erroneously passed as %v", S, v)
		}
		index++
	})

	ConfirmEach(S, func(v int) {
		if v != S[index] {
			t.Fatalf("index %v erroneously reported as %v", index, v)
		}
		index++
	})

	ConfirmEach(S, func(i, v int) {
		switch {
		case i != index:					t.Fatalf("index %v erroneously reported as %v", index, i)
		case v != S[i]:						t.Fatalf("value %v erroneously reported as %v", S[i], v)
		}
		index++
	})
}

func TestEachFunction(t *testing.T) {
	limit := 10
	count := 0
	value := 0

	ConfirmEach := func(F, f interface{}) {
		count = 0
		value = 0
		switch {
		case !Each(F, f):			t.Fatalf("failed to perform iteration %v over %v", f, F)
		case value != limit:		t.Fatalf("total items produced should be %v but are %v", limit, value)
		case count != limit:		//t.Fatalf("total iterations should be %v but are %v", limit, count)
									panic(count)
		}
	}

	ConfirmVariadicEach := func(F, f interface{}) {
		count = 0
		value = 0
		switch {
		case !Each(F, f):			t.Fatalf("failed to perform iteration %v over %v", f, F)
		case count != 1:			t.Fatalf("total iterations should be 1 but are %v", count)
		}
	}

	ConfirmEachFunction := func(F interface{}) {
		ConfirmEach(F, func(v interface{}) {
			if v != count {
				t.Fatalf("value %v erroneously reported as %v", count, v)
			}
			count++
		})

		ConfirmEach(F, func(i int, v interface{}) {
			switch {
			case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != i:				t.Fatalf("value %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmEach(F, func(i, v interface{}) {
			switch {
			case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != i:				t.Fatalf("value %v erroneously reported as %v", i, v)
			}
			count++
		})

		ConfirmVariadicEach(F, func(v ...interface{}) {
			switch {
			case count != 0:			t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != limit:		//t.Fatalf("variadic slice generator %v erroneously passed as %v", F, v)
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
			case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != i:		t.Fatalf("value %v erroneously reported as %v", i, v.Interface())
			}
			count++
		})

		ConfirmEach(F, func(i interface{}, v R.Value) {
			switch {
			case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
			case v.Interface() != i:	t.Fatalf("value %v erroneously reported as %v", i, v.Interface())
			}
			count++
		})

		ConfirmEach(F, func(i, v R.Value) {
			switch {
			case i.Interface() != count:			t.Fatalf("index %v erroneously reported as %v", count, i.Interface())
			case v.Interface() != i.Interface():	t.Fatalf("value %v erroneously reported as %v", i.Interface(), v.Interface())
			}
			count++
		})

		ConfirmVariadicEach(F, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != limit:				t.Fatalf("variadic slice generator %v erroneously passed as %v", F, v)
			}
			count++
		})
	}

	ConfirmEachIntFunction := func(F interface{}) {
		ConfirmEachFunction(F)

		ConfirmVariadicEach(F, func(v ...int) {
			switch {
			case count != 0:			t.Fatalf("variadic function erroneously called %v times", count)
			case len(v) != limit:		t.Fatalf("variadic slice generator %v erroneously passed as %v", F, v)
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
			case i != count:			t.Fatalf("index %v erroneously reported as %v", count, i)
			case v != i:				t.Fatalf("value %v erroneously reported as %v", i, v)
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
}