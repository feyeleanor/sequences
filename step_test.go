package sequences

import(
	R "reflect"
	"testing"
)

func TestStepChannelBool(t *testing.T) {
	var count		int
	var sequence	func() chan bool
	var elements	[]interface{}

	s := []bool{true, false, false, true, true}
	sequence = func() (r chan bool) {
		r = make(chan bool)
		go func() {
			for _, v := range s {
				r <- v
			}
			close(r)
		}()
		return
	}

	Sublen := func(span int) (r int) {
		l := len(s)
		if span < 0 {
			r = l / -span
		} else {
			r = l / span
		}
		if l % span != 0 {
			r++
		}
		return
	}
	

	ConfirmStep := func(s interface{}, span, limit int, f interface{}) {
		count = 0
		switch count, e := Step(s, 0, span, limit, f); {
		case e != nil:
			t.Fatalf("%T span %v: iteration failed with error %v", s, span, e)
		case count != Sublen(span):
			t.Fatalf("%T span %v: total iterations should be %v but are %v", s, span, Sublen(span), count)
		}
	}

	ConfirmVariadicStep := func(s interface{}, span, limit int, f interface{}) {
		count = 0
		switch count, e := Step(s, 0, span, limit, f); {
		case e != nil:
			t.Fatalf("%T span %v: iteration failed with error %v", s, span, e)
		case count != 1:
			t.Fatalf("%T span %v: total iterations should be 1 but are %v", s, span, count)
		}
	}

	ConfirmChannelStep := func(seq chan bool, span, limit int, f interface{}) {
		elements = make([]interface{}, 0, 0)
		l := Sublen(span)
		done := make(chan bool)
		switch r := f.(type) {
		case chan bool:
			go func() {
				for p := l; p > 0; p-- {
					v := <-r
					elements = append(elements, v)
				}
				close(done)
			}()
		case chan interface{}:
			go func() {
				for p := l; p > 0; p-- {
					v := <-r
					elements = append(elements, v)
				}
				close(done)
			}()
		case chan R.Value:
			go func() {
				for p := l; p > 0; p-- {
					v := <-r
					elements = append(elements, v)
				}
				close(done)
			}()
		}
		n, e := Step(seq, 0, span, limit, f)
		for _ = range done {}
		switch {
		case e != nil:
			t.Fatalf("%T span %v: iteration failed with error %v", seq, span, e)
		case n != l:
			t.Fatalf("%T span %v: Each() should return %v but is %v", seq, span, l, n)
		case len(elements) != l:
			t.Fatalf("%T span %v: total iterations should be %v but are %v", seq, span, l, len(elements))
		}
	}

	ConfirmVariadicChannelStep := func(seq chan bool, span, limit int, f interface{}) {
		switch r := f.(type) {
		case []chan bool:
			for p, c := range r {
				go func(p int, c chan bool) {
					n := 0
					x := 0
					for v := range c {
						if v != s[n] {
							t.Fatalf("worker %v [%T span %v]: item %v should be %v but is %v", p, seq, span, x, s[n], v)
						}
						n += span
						x++
					}
				}(p, c)
			}
		case []chan interface{}:
			for p, c := range r {
				go func(p int, c chan interface{}) {
					n := 0
					x := 0
					for v := range c {
						if v != s[n] {
							t.Fatalf("worker %v [%T span %v]: item %v should be %v but is %v", p, seq, span, x, s[n], v)
						}
						n += span
						x++
					}
				}(p, c)
			}
		case []chan R.Value:
			for p, c := range r {
				go func(p int, c chan R.Value) {
					n := 0
					x := 0
					for v := range c {
						if v.Interface() != s[n] {
							t.Fatalf("worker %v [%T span %v]: item %v should be %v but is %v", p, seq, span, x, s[n], v.Interface())
						}
						n += span
						x++
					}
				}(p, c)
			}
		}
		n, e := Step(seq, 0, span, limit, f)
		switch {
		case e != nil:
			t.Fatalf("%T span %v: iteration failed with error %v", seq, span, e)
		case n != Sublen(span):
			t.Fatalf("%T span %v: Each() should return %v but is %v", seq, span, Sublen(span), n)
		}
	}

	Confirm := func(seq func() chan bool, span, limit int) {
		//	The channel produces a value only once so AtOffset can't be used to confirm that value is correct
		ConfirmStep(seq(), span, limit, func(v bool) {
			count++
		})

		ConfirmStep(seq(), span, limit, func(v interface{}) {
			count++
		})

		ConfirmStep(seq(), span, limit, func(i int, v bool) {
			if offset := span * count; i != span * count {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, i)
			}
			count++
		})

		ConfirmStep(seq(), span, limit, func(i int, v interface{}) {
			if offset := span * count; i != span * count {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, i)
			}
			count++
		})

		ConfirmStep(seq(), span, limit, func(k, v interface{}) {
			if offset := span * count; k != span * count {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, k)
			}
			count++
		})

		ConfirmVariadicStep(seq(), span, limit, func(v ...interface{}) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", seq, span, count)
			case len(v) != Sublen(span):
				t.Fatalf("%T span %v: variadic channel of %v items erroneously passed as %v", seq, span, len(s), v)
			}
			count++
		})

		ConfirmStep(seq(), span, limit, func(v R.Value) {
			count++
		})

		ConfirmStep(seq(), span, limit, func(i int, v R.Value) {
			if offset := span * count; i != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, i)
			}
			count++
		})

		ConfirmStep(seq(), span, limit, func(k interface{}, v R.Value) {
			if offset := span * count; k != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, k)
			}
			count++
		})

		ConfirmStep(seq(), span, limit, func(k, v R.Value) {
			if offset := span * count; k.Interface() != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, k)
			}
			count++
		})

		ConfirmVariadicStep(seq(), span, limit, func(v ...R.Value) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", seq, span, count)
			case len(v) != Sublen(span):
				t.Fatalf("%T span %v: variadic channel %v items erroneously passed as %v", seq, span, len(s), v)
			}
			count++
		})

		ConfirmChannelStep(seq(), span, limit, make(chan bool))
		ConfirmChannelStep(seq(), span, limit, make(chan interface{}))
		ConfirmChannelStep(seq(), span, limit, make(chan R.Value))
		ConfirmVariadicChannelStep(seq(), span, limit, []chan bool{make(chan bool), make(chan bool), make(chan bool)})
		ConfirmVariadicChannelStep(seq(), span, limit, []chan interface{}{make(chan interface{}), make(chan interface{}), make(chan interface{})})
		ConfirmVariadicChannelStep(seq(), span, limit, []chan R.Value{make(chan R.Value), make(chan R.Value), make(chan R.Value)})
	}

	Confirm(sequence, 1, -1)
	Confirm(sequence, 2, -1)
	Confirm(sequence, 3, -1)
	Confirm(sequence, 4, -1)
	Confirm(sequence, 5, -1)
	Confirm(sequence, 6, -1)
}

/*
	ConfirmStepBool(sequence, -6, 0)
	ConfirmStepBool(sequence, -5, 0)
	ConfirmStepBool(sequence, -4, 0)
	ConfirmStepBool(sequence, -3, 0)
	ConfirmStepBool(sequence, -2, 0)
	ConfirmStepBool(sequence, -1, 0)
	ConfirmStepBool(sequence, 1, -1, 5)
	ConfirmStepBool(sequence, 2, -1, 3)
	ConfirmStepBool(sequence, 3, 2)
	ConfirmStepBool(sequence, 4, 1)
	ConfirmStepBool(sequence, 5, 1)
	ConfirmStepBool(sequence, 6, 1)
*/

func TestStepIndexables(t *testing.T) {
	var count		int
	var sequence	interface{}

	Offset := func(s interface{}, span int) (r int) {
		if span < 0 {
			r, _ = Len(s)
			r += -1 + (span * count)
		} else {
			r = span * count
		}
		return
	}

	Sublen := func(s interface{}, span int) (r int) {
		l, _ := Len(s)
		if span < 0 {
			r = l / -span
		} else {
			r = l / span
		}
		if l % span != 0 {
			r++
		}
		return
	}

	ConfirmStep := func(s interface{}, span int, f interface{}) {
		count = 0
		switch count, e := Step(s, 0, span, -1, f); {
		case e != nil:
			t.Fatalf("%T span %v: iteration failed with error %v", s, span, e)
		case count != Sublen(s, span):
			t.Fatalf("%T span %v: total iterations should be %v but are %v", s, span, Sublen(s, span), count)
		}
	}

	ConfirmVariadicStep := func(s interface{}, span int, f interface{}) {
		count = 0
		switch count, e := Step(s, 0, span, -1, f); {
		case e != nil:
			t.Fatalf("%T span %v: iteration failed with error %v", s, span, e)
		case count != 1:
			t.Fatalf("%T span %v: total iterations should be 1 but are %v", s, span, count)
		}
	}

	ConfirmStepIndexable := func(s interface{}, span int) {
		ConfirmStep(s, span, func(v interface{}) {
			if v != AtOffset(s, Offset(s, span)) {
				panic(Offset(s, span))
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v interface{}) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})

		ConfirmStep(s, span, func(k, v interface{}) {
			switch {
			case k != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), k)
			case v != AtOffset(s, k.(int)):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, span, func(v ...interface{}) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice of %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v R.Value) {
			if v.Interface() != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v.Interface())
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v R.Value) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v.Interface() != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, count, v)
			}
			count++
		})

		ConfirmStep(s, span, func(k interface{}, v R.Value) {
			switch {
			case k != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), k)
			case v.Interface() != AtOffset(s, k.(int)):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		ConfirmStep(s, span, func(k, v R.Value) {
			switch {
			case k.Interface() != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), k)
			case v.Interface() != AtOffset(s, int(k.Int())):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, span, func(v ...R.Value) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})
	}


	ConfirmStepBool := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...bool) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v bool) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v bool) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []bool{true, false, false, true, true}
	ConfirmStepBool(sequence, -6)
	ConfirmStepBool(sequence, -5)
	ConfirmStepBool(sequence, -4)
	ConfirmStepBool(sequence, -3)
	ConfirmStepBool(sequence, -2)
	ConfirmStepBool(sequence, -1)
	ConfirmStepBool(sequence, 1)
	ConfirmStepBool(sequence, 2)
	ConfirmStepBool(sequence, 3)
	ConfirmStepBool(sequence, 4)
	ConfirmStepBool(sequence, 5)
	ConfirmStepBool(sequence, 6)

	ConfirmStepComplex64 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...complex64) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v complex64) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v complex64) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []complex64{0, 1, 2, 3, 4}
	ConfirmStepComplex64(sequence, -6)
	ConfirmStepComplex64(sequence, -5)
	ConfirmStepComplex64(sequence, -4)
	ConfirmStepComplex64(sequence, -3)
	ConfirmStepComplex64(sequence, -2)
	ConfirmStepComplex64(sequence, -1)
	ConfirmStepComplex64(sequence, 1)
	ConfirmStepComplex64(sequence, 2)
	ConfirmStepComplex64(sequence, 3)
	ConfirmStepComplex64(sequence, 4)
	ConfirmStepComplex64(sequence, 5)
	ConfirmStepComplex64(sequence, 6)

	ConfirmStepComplex128 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...complex128) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v complex128) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v complex128) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []complex128{0, 1, 2, 3, 4}
	ConfirmStepComplex128(sequence, -6)
	ConfirmStepComplex128(sequence, -5)
	ConfirmStepComplex128(sequence, -4)
	ConfirmStepComplex128(sequence, -3)
	ConfirmStepComplex128(sequence, -2)
	ConfirmStepComplex128(sequence, -1)
	ConfirmStepComplex128(sequence, 1)
	ConfirmStepComplex128(sequence, 2)
	ConfirmStepComplex128(sequence, 3)
	ConfirmStepComplex128(sequence, 4)
	ConfirmStepComplex128(sequence, 5)
	ConfirmStepComplex128(sequence, 6)

	ConfirmStepError := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...error) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v error) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v error) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []error{Error(0), Error(1), Error(2), Error(3), Error(4)}
	ConfirmStepError(sequence, -6)
	ConfirmStepError(sequence, -5)
	ConfirmStepError(sequence, -4)
	ConfirmStepError(sequence, -3)
	ConfirmStepError(sequence, -2)
	ConfirmStepError(sequence, -1)
	ConfirmStepError(sequence, 1)
	ConfirmStepError(sequence, 2)
	ConfirmStepError(sequence, 3)
	ConfirmStepError(sequence, 4)
	ConfirmStepError(sequence, 5)
	ConfirmStepError(sequence, 6)

	ConfirmStepFloat32 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...float32) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v float32) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v float32) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []float32{0, 1, 2, 3, 4}
	ConfirmStepFloat32(sequence, -6)
	ConfirmStepFloat32(sequence, -5)
	ConfirmStepFloat32(sequence, -4)
	ConfirmStepFloat32(sequence, -3)
	ConfirmStepFloat32(sequence, -2)
	ConfirmStepFloat32(sequence, -1)
	ConfirmStepFloat32(sequence, 1)
	ConfirmStepFloat32(sequence, 2)
	ConfirmStepFloat32(sequence, 3)
	ConfirmStepFloat32(sequence, 4)
	ConfirmStepFloat32(sequence, 5)
	ConfirmStepFloat32(sequence, 6)

	ConfirmStepFloat64 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...float64) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v float64) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v float64) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []float64{0, 1, 2, 3, 4}
	ConfirmStepFloat64(sequence, -6)
	ConfirmStepFloat64(sequence, -5)
	ConfirmStepFloat64(sequence, -4)
	ConfirmStepFloat64(sequence, -3)
	ConfirmStepFloat64(sequence, -2)
	ConfirmStepFloat64(sequence, -1)
	ConfirmStepFloat64(sequence, 1)
	ConfirmStepFloat64(sequence, 2)
	ConfirmStepFloat64(sequence, 3)
	ConfirmStepFloat64(sequence, 4)
	ConfirmStepFloat64(sequence, 5)
	ConfirmStepFloat64(sequence, 6)

	ConfirmStepInt := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...int) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v int) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v int) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []int{0, 1, 2, 3, 4}
	ConfirmStepInt(sequence, -6)
	ConfirmStepInt(sequence, -5)
	ConfirmStepInt(sequence, -4)
	ConfirmStepInt(sequence, -3)
	ConfirmStepInt(sequence, -2)
	ConfirmStepInt(sequence, -1)
	ConfirmStepInt(sequence, 1)
	ConfirmStepInt(sequence, 2)
	ConfirmStepInt(sequence, 3)
	ConfirmStepInt(sequence, 4)
	ConfirmStepInt(sequence, 5)
	ConfirmStepInt(sequence, 6)

	ConfirmStepInt8 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...int8) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v int8) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v int8) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []int8{0, 1, 2, 3, 4}
	ConfirmStepInt8(sequence, -6)
	ConfirmStepInt8(sequence, -5)
	ConfirmStepInt8(sequence, -4)
	ConfirmStepInt8(sequence, -3)
	ConfirmStepInt8(sequence, -2)
	ConfirmStepInt8(sequence, -1)
	ConfirmStepInt8(sequence, 1)
	ConfirmStepInt8(sequence, 2)
	ConfirmStepInt8(sequence, 3)
	ConfirmStepInt8(sequence, 4)
	ConfirmStepInt8(sequence, 5)
	ConfirmStepInt8(sequence, 6)
	
	ConfirmStepInt16 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...int16) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v int16) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v int16) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []int16{0, 1, 2, 3, 4}
	ConfirmStepInt16(sequence, -6)
	ConfirmStepInt16(sequence, -5)
	ConfirmStepInt16(sequence, -4)
	ConfirmStepInt16(sequence, -3)
	ConfirmStepInt16(sequence, -2)
	ConfirmStepInt16(sequence, -1)
	ConfirmStepInt16(sequence, 1)
	ConfirmStepInt16(sequence, 2)
	ConfirmStepInt16(sequence, 3)
	ConfirmStepInt16(sequence, 4)
	ConfirmStepInt16(sequence, 5)
	ConfirmStepInt16(sequence, 6)

	ConfirmStepInt32 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...int32) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v int32) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v int32) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []int32{0, 1, 2, 3, 4}
	ConfirmStepInt32(sequence, -6)
	ConfirmStepInt32(sequence, -5)
	ConfirmStepInt32(sequence, -4)
	ConfirmStepInt32(sequence, -3)
	ConfirmStepInt32(sequence, -2)
	ConfirmStepInt32(sequence, -1)
	ConfirmStepInt32(sequence, 1)
	ConfirmStepInt32(sequence, 2)
	ConfirmStepInt32(sequence, 3)
	ConfirmStepInt32(sequence, 4)
	ConfirmStepInt32(sequence, 5)
	ConfirmStepInt32(sequence, 6)

	ConfirmStepInt64 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...int64) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v int64) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v int64) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []int64{0, 1, 2, 3, 4}
	ConfirmStepInt64(sequence, -6)
	ConfirmStepInt64(sequence, -5)
	ConfirmStepInt64(sequence, -4)
	ConfirmStepInt64(sequence, -3)
	ConfirmStepInt64(sequence, -2)
	ConfirmStepInt64(sequence, -1)
	ConfirmStepInt64(sequence, 1)
	ConfirmStepInt64(sequence, 2)
	ConfirmStepInt64(sequence, 3)
	ConfirmStepInt64(sequence, 4)
	ConfirmStepInt64(sequence, 5)
	ConfirmStepInt64(sequence, 6)

	sequence = []interface{}{0, 1, 2, 3, 4}
	ConfirmStepIndexable(sequence, -6)
	ConfirmStepIndexable(sequence, -5)
	ConfirmStepIndexable(sequence, -4)
	ConfirmStepIndexable(sequence, -3)
	ConfirmStepIndexable(sequence, -2)
	ConfirmStepIndexable(sequence, -1)
	ConfirmStepIndexable(sequence, 1)
	ConfirmStepIndexable(sequence, 2)
	ConfirmStepIndexable(sequence, 3)
	ConfirmStepIndexable(sequence, 4)
	ConfirmStepIndexable(sequence, 5)
	ConfirmStepIndexable(sequence, 6)
	
	ConfirmStepString := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...string) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v string) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v string) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []string{"A", "B", "C", "D", "E"}
	ConfirmStepString(sequence, -6)
	ConfirmStepString(sequence, -5)
	ConfirmStepString(sequence, -4)
	ConfirmStepString(sequence, -3)
	ConfirmStepString(sequence, -2)
	ConfirmStepString(sequence, -1)
	ConfirmStepString(sequence, 1)
	ConfirmStepString(sequence, 2)
	ConfirmStepString(sequence, 3)
	ConfirmStepString(sequence, 4)
	ConfirmStepString(sequence, 5)
	ConfirmStepString(sequence, 6)
	
	ConfirmStepUint := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...uint) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v uint) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v uint) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []uint{0, 1, 2, 3, 4}
	ConfirmStepUint(sequence, -6)
	ConfirmStepUint(sequence, -5)
	ConfirmStepUint(sequence, -4)
	ConfirmStepUint(sequence, -3)
	ConfirmStepUint(sequence, -2)
	ConfirmStepUint(sequence, -1)
	ConfirmStepUint(sequence, 1)
	ConfirmStepUint(sequence, 2)
	ConfirmStepUint(sequence, 3)
	ConfirmStepUint(sequence, 4)
	ConfirmStepUint(sequence, 5)
	ConfirmStepUint(sequence, 6)

	ConfirmStepUint8 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...uint8) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v uint8) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v uint8) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []uint8{0, 1, 2, 3, 4}
	ConfirmStepUint8(sequence, -6)
	ConfirmStepUint8(sequence, -5)
	ConfirmStepUint8(sequence, -4)
	ConfirmStepUint8(sequence, -3)
	ConfirmStepUint8(sequence, -2)
	ConfirmStepUint8(sequence, -1)
	ConfirmStepUint8(sequence, 1)
	ConfirmStepUint8(sequence, 2)
	ConfirmStepUint8(sequence, 3)
	ConfirmStepUint8(sequence, 4)
	ConfirmStepUint8(sequence, 5)
	ConfirmStepUint8(sequence, 6)
	
	ConfirmStepUint16 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...uint16) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v uint16) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v uint16) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []uint16{0, 1, 2, 3, 4}
	ConfirmStepUint16(sequence, -6)
	ConfirmStepUint16(sequence, -5)
	ConfirmStepUint16(sequence, -4)
	ConfirmStepUint16(sequence, -3)
	ConfirmStepUint16(sequence, -2)
	ConfirmStepUint16(sequence, -1)
	ConfirmStepUint16(sequence, 1)
	ConfirmStepUint16(sequence, 2)
	ConfirmStepUint16(sequence, 3)
	ConfirmStepUint16(sequence, 4)
	ConfirmStepUint16(sequence, 5)
	ConfirmStepUint16(sequence, 6)

	ConfirmStepUint32 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...uint32) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v uint32) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v uint32) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []uint32{0, 1, 2, 3, 4}
	ConfirmStepUint32(sequence, -6)
	ConfirmStepUint32(sequence, -5)
	ConfirmStepUint32(sequence, -4)
	ConfirmStepUint32(sequence, -3)
	ConfirmStepUint32(sequence, -2)
	ConfirmStepUint32(sequence, -1)
	ConfirmStepUint32(sequence, 1)
	ConfirmStepUint32(sequence, 2)
	ConfirmStepUint32(sequence, 3)
	ConfirmStepUint32(sequence, 4)
	ConfirmStepUint32(sequence, 5)
	ConfirmStepUint32(sequence, 6)

	ConfirmStepUint64 := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...uint64) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v uint64) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v uint64) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []uint64{0, 1, 2, 3, 4}
	ConfirmStepUint64(sequence, -6)
	ConfirmStepUint64(sequence, -5)
	ConfirmStepUint64(sequence, -4)
	ConfirmStepUint64(sequence, -3)
	ConfirmStepUint64(sequence, -2)
	ConfirmStepUint64(sequence, -1)
	ConfirmStepUint64(sequence, 1)
	ConfirmStepUint64(sequence, 2)
	ConfirmStepUint64(sequence, 3)
	ConfirmStepUint64(sequence, 4)
	ConfirmStepUint64(sequence, 5)
	ConfirmStepUint64(sequence, 6)

	ConfirmStepUintptr := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...uintptr) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v uintptr) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v uintptr) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = []uintptr{0, 1, 2, 3, 4}
	ConfirmStepUintptr(sequence, -6)
	ConfirmStepUintptr(sequence, -5)
	ConfirmStepUintptr(sequence, -4)
	ConfirmStepUintptr(sequence, -3)
	ConfirmStepUintptr(sequence, -2)
	ConfirmStepUintptr(sequence, -1)
	ConfirmStepUintptr(sequence, 1)
	ConfirmStepUintptr(sequence, 2)
	ConfirmStepUintptr(sequence, 3)
	ConfirmStepUintptr(sequence, 4)
	ConfirmStepUintptr(sequence, 5)
	ConfirmStepUintptr(sequence, 6)

	ConfirmStepRValueSlice := func(s interface{}, span int) {
		ConfirmStep(s, span, func(v interface{}) {
			if v != AtOffset(s, Offset(s, span)).(R.Value).Interface() {
				panic(s)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v interface{}) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i).(R.Value).Interface():
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})

		ConfirmStep(s, span, func(k, v interface{}) {
			switch {
			case k != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), k)
			case v != AtOffset(s, k.(int)).(R.Value).Interface():
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, span, func(v ...interface{}) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice of %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v R.Value) {
			if v.Interface() != AtOffset(s, Offset(s, span)).(R.Value).Interface() {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v.Interface())
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v R.Value) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v.Interface() != AtOffset(s, i).(R.Value).Interface():
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, count, v)
			}
			count++
		})

		ConfirmStep(s, span, func(k interface{}, v R.Value) {
			switch {
			case k != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), k)
			case v.Interface() != AtOffset(s, k.(int)).(R.Value).Interface():
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		ConfirmStep(s, span, func(k, v R.Value) {
			switch {
			case k.Interface() != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), k)
			case v.Interface() != AtOffset(s, int(k.Int())).(R.Value).Interface():
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, span, func(v ...R.Value) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})
	}
	sequence = []R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}
	ConfirmStepRValueSlice(sequence, -6)
	ConfirmStepRValueSlice(sequence, -5)
	ConfirmStepRValueSlice(sequence, -4)
	ConfirmStepRValueSlice(sequence, -3)
	ConfirmStepRValueSlice(sequence, -2)
	ConfirmStepRValueSlice(sequence, -1)
	ConfirmStepRValueSlice(sequence, 1)
	ConfirmStepRValueSlice(sequence, 2)
	ConfirmStepRValueSlice(sequence, 3)
	ConfirmStepRValueSlice(sequence, 4)
	ConfirmStepRValueSlice(sequence, 5)
	ConfirmStepRValueSlice(sequence, 6)

	ConfirmStepRValue := func(s interface{}, span int) {
		ConfirmStepIndexable(s, span)
		ConfirmVariadicStep(s, span, func(v ...int) {
			switch {
			case count != 0:
				t.Fatalf("%T span %v: variadic function erroneously called %v times", s, span, count)
			case len(v) != Sublen(s, span):
				t.Fatalf("%T span %v: variadic slice %v erroneously passed as %v", s, span, s, v)
			}
			count++
		})

		ConfirmStep(s, span, func(v int) {
			if v != AtOffset(s, Offset(s, span)) {
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, Offset(s, span), v)
			}
			count++
		})

		ConfirmStep(s, span, func(i int, v int) {
			switch {
			case i != Offset(s, span):
				t.Fatalf("%T span %v: index %v erroneously reported as %v", s, span, Offset(s, span), i)
			case v != AtOffset(s, i):
				t.Fatalf("%T span %v: element %v erroneously reported as %v", s, span, i, v)
			}
			count++
		})
	}
	sequence = R.ValueOf([]int{0, 1, 2, 3, 4})
	ConfirmStepRValue(sequence, -6)
	ConfirmStepRValue(sequence, -5)
	ConfirmStepRValue(sequence, -4)
	ConfirmStepRValue(sequence, -3)
	ConfirmStepRValue(sequence, -2)
	ConfirmStepRValue(sequence, -1)
	ConfirmStepRValue(sequence, 1)
	ConfirmStepRValue(sequence, 2)
	ConfirmStepRValue(sequence, 3)
	ConfirmStepRValue(sequence, 4)
	ConfirmStepRValue(sequence, 5)
	ConfirmStepRValue(sequence, 6)

	sequence = indexable_slice{0, 1, 2, 3, 4}
	ConfirmStepIndexable(sequence, -6)
	ConfirmStepIndexable(sequence, -5)
	ConfirmStepIndexable(sequence, -4)
	ConfirmStepIndexable(sequence, -3)
	ConfirmStepIndexable(sequence, -2)
	ConfirmStepIndexable(sequence, -1)
	ConfirmStepIndexable(sequence, 1)
	ConfirmStepIndexable(sequence, 2)
	ConfirmStepIndexable(sequence, 3)
	ConfirmStepIndexable(sequence, 4)
	ConfirmStepIndexable(sequence, 5)
	ConfirmStepIndexable(sequence, 6)

	sequence = enumerable_slice{0, 1, 2, 3, 4}
	ConfirmStepIndexable(sequence, -6)
	ConfirmStepIndexable(sequence, -5)
	ConfirmStepIndexable(sequence, -4)
	ConfirmStepIndexable(sequence, -3)
	ConfirmStepIndexable(sequence, -2)
	ConfirmStepIndexable(sequence, -1)
	ConfirmStepIndexable(sequence, 1)
	ConfirmStepIndexable(sequence, 2)
	ConfirmStepIndexable(sequence, 3)
	ConfirmStepIndexable(sequence, 4)
	ConfirmStepIndexable(sequence, 5)
	ConfirmStepIndexable(sequence, 6)

	sequence = indexable_function(func(i int) interface{} { return i })
	ConfirmStepIndexable(sequence, -6)
	ConfirmStepIndexable(sequence, -5)
	ConfirmStepIndexable(sequence, -4)
	ConfirmStepIndexable(sequence, -3)
	ConfirmStepIndexable(sequence, -2)
	ConfirmStepIndexable(sequence, -1)
	ConfirmStepIndexable(sequence, 1)
	ConfirmStepIndexable(sequence, 2)
	ConfirmStepIndexable(sequence, 3)
	ConfirmStepIndexable(sequence, 4)
	ConfirmStepIndexable(sequence, 5)
	ConfirmStepIndexable(sequence, 6)
}



//	Write tests for the 0 case which should always fail to iterate
//	ConfirmStepBool([]bool{true, false, false, true, true}, 0)