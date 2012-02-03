package sequences

import(
	R "reflect"
	"testing"
)

func TestStepPrimitive(t *testing.T) {
	var count	int

	Offset := func(s interface{}, inc int) (r int) {
		if inc < 0 {
			r = Len(s) + inc * (count + 1)
		} else {
			r = -1 + inc * (count + 1)
		}
		return
	}

	Sublen := func(s interface{}, inc int) (r int) {
		if inc < 0 {
			r = Len(s) / -inc
		} else {
			r = Len(s) / inc
		}
		return
	}

	ConfirmStep := func(s interface{}, inc int, f interface{}) {
		count = 0
		switch {
		case !Step(s, inc, f):			t.Fatalf("%T increment %v: failed to perform iteration %v", s, inc, f)
		case count != Sublen(s, inc):	t.Fatalf("%T increment %v: total iterations should be %v but are %v", s, inc, Sublen(s, inc), count)
		}
	}

	ConfirmVariadicStep := func(s interface{}, inc int, f interface{}) {
		count = 0
		switch {
		case !Step(s, inc, f):		t.Logf("%T increment %v: failed to perform iteration %v", s, inc, f)
		case count != 1:			t.Fatalf("%T increment %v: total iterations should be 1 but are %v", s, inc, count)
		}
	}

	ConfirmStepBuiltin := func(s interface{}, inc int) {
		ConfirmStep(s, inc, func(v interface{}) {
			if v != AtOffset(s, Offset(s, inc)) {
				panic(s)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v interface{}) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(k, v interface{}) {
			switch {
			case k != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), k)
			case v != AtOffset(s, k.(int)):		t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, inc, func(v ...interface{}) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice of %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v R.Value) {
			if v.Interface() != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v.Interface())
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v R.Value) {
			switch {
			case i != Offset(s, inc):				t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v.Interface() != AtOffset(s, i):	t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, count, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(k interface{}, v R.Value) {
			switch {
			case k != Offset(s, inc):					t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), k)
			case v.Interface() != AtOffset(s, k.(int)):	t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, k, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(k, v R.Value) {
			switch {
			case k.Interface() != Offset(s, inc):				t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), k)
			case v.Interface() != AtOffset(s, int(k.Int())):	t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, inc, func(v ...R.Value) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})
	}

	ConfirmStepBool := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...bool) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v bool) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v bool) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepBool([]bool{true, false, false, true, true}, -5)
	ConfirmStepBool([]bool{true, false, false, true, true}, -4)
	ConfirmStepBool([]bool{true, false, false, true, true}, -3)
	ConfirmStepBool([]bool{true, false, false, true, true}, -2)
	ConfirmStepBool([]bool{true, false, false, true, true}, -1)
	ConfirmStepBool([]bool{true, false, false, true, true}, 1)
	ConfirmStepBool([]bool{true, false, false, true, true}, 2)
	ConfirmStepBool([]bool{true, false, false, true, true}, 3)
	ConfirmStepBool([]bool{true, false, false, true, true}, 4)
	ConfirmStepBool([]bool{true, false, false, true, true}, 5)

	ConfirmStepComplex64 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...complex64) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v complex64) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v complex64) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, -5)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, -4)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, -3)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, -2)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, -1)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, 1)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, 2)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, 3)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, 4)
	ConfirmStepComplex64([]complex64{0, 1, 2, 3, 4}, 5)

	ConfirmStepComplex128 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...complex128) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v complex128) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v complex128) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, -5)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, -4)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, -3)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, -2)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, -1)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, 1)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, 2)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, 3)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, 4)
	ConfirmStepComplex128([]complex128{0, 1, 2, 3, 4}, 5)

	ConfirmStepError := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...error) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v error) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v error) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, -5)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, -4)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, -3)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, -2)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, -1)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, 1)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, 2)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, 3)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, 4)
	ConfirmStepError([]error{Error(0), Error(1), Error(2), Error(3), Error(4)}, 5)

	ConfirmStepFloat32 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...float32) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v float32) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v float32) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, -5)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, -4)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, -3)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, -2)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, -1)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, 1)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, 2)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, 3)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, 4)
	ConfirmStepFloat32([]float32{0, 1, 2, 3, 4}, 5)

	ConfirmStepFloat64 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...float64) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v float64) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v float64) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, -5)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, -4)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, -3)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, -2)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, -1)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, 1)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, 2)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, 3)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, 4)
	ConfirmStepFloat64([]float64{0, 1, 2, 3, 4}, 5)

	ConfirmStepInt := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v int) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v int) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, -5)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, -4)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, -3)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, -2)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, -1)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, 1)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, 2)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, 3)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, 4)
	ConfirmStepInt([]int{0, 1, 2, 3, 4}, 5)

	ConfirmStepInt8 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...int8) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v int8) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v int8) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, -5)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, -4)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, -3)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, -2)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, -1)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, 1)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, 2)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, 3)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, 4)
	ConfirmStepInt8([]int8{0, 1, 2, 3, 4}, 5)
	
	ConfirmStepInt16 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...int16) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v int16) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v int16) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, -5)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, -4)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, -3)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, -2)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, -1)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, 1)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, 2)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, 3)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, 4)
	ConfirmStepInt16([]int16{0, 1, 2, 3, 4}, 5)

	ConfirmStepInt32 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...int32) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v int32) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v int32) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, -5)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, -4)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, -3)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, -2)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, -1)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, 1)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, 2)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, 3)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, 4)
	ConfirmStepInt32([]int32{0, 1, 2, 3, 4}, 5)

	ConfirmStepInt64 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...int64) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v int64) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v int64) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, -5)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, -4)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, -3)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, -2)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, -1)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, 1)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, 2)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, 3)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, 4)
	ConfirmStepInt64([]int64{0, 1, 2, 3, 4}, 5)

	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, -5)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, -4)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, -3)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, -2)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, -1)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, 1)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, 2)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, 3)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, 4)
	ConfirmStepBuiltin([]interface{}{0, 1, 2, 3, 4}, 5)
	
	ConfirmStepString := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...string) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v string) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v string) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, -5)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, -4)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, -3)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, -2)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, -1)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, 1)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, 2)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, 3)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, 4)
	ConfirmStepString([]string{"A", "B", "C", "D", "E"}, 5)
	
	ConfirmStepUint := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...uint) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v uint) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v uint) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, -5)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, -4)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, -3)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, -2)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, -1)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, 1)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, 2)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, 3)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, 4)
	ConfirmStepUint([]uint{0, 1, 2, 3, 4}, 5)

	ConfirmStepUint8 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...uint8) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v uint8) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v uint8) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, -5)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, -4)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, -3)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, -2)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, -1)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, 1)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, 2)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, 3)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, 4)
	ConfirmStepUint8([]uint8{0, 1, 2, 3, 4}, 5)
	
	ConfirmStepUint16 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...uint16) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v uint16) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v uint16) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, -5)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, -4)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, -3)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, -2)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, -1)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, 1)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, 2)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, 3)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, 4)
	ConfirmStepUint16([]uint16{0, 1, 2, 3, 4}, 5)

	ConfirmStepUint32 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...uint32) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v uint32) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v uint32) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, -5)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, -4)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, -3)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, -2)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, -1)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, 1)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, 2)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, 3)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, 4)
	ConfirmStepUint32([]uint32{0, 1, 2, 3, 4}, 5)

	ConfirmStepUint64 := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...uint64) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v uint64) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v uint64) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, -5)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, -4)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, -3)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, -2)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, -1)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, 1)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, 2)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, 3)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, 4)
	ConfirmStepUint64([]uint64{0, 1, 2, 3, 4}, 5)

	ConfirmStepUintptr := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...uintptr) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v uintptr) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v uintptr) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, -5)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, -4)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, -3)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, -2)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, -1)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, 1)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, 2)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, 3)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, 4)
	ConfirmStepUintptr([]uintptr{0, 1, 2, 3, 4}, 5)

	ConfirmStepRValueSlice := func(s []R.Value, inc int) {
		ConfirmStep(s, inc, func(v interface{}) {
			if v != s[Offset(s, inc)].Interface() {
				panic(s)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v interface{}) {
			switch {
			case i != Offset(s, inc):				t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != s[i].Interface():				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(k, v interface{}) {
			switch {
			case k != Offset(s, inc):				t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), k)
			case v != s[k.(int)].Interface():		t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, inc, func(v ...interface{}) {
			switch {
			case count != 0:						t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):			t.Fatalf("%T increment %v: variadic slice of %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v R.Value) {
			if v.Interface() != s[Offset(s, inc)].Interface() {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v.Interface())
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v R.Value) {
			switch {
			case i != Offset(s, inc):					t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v.Interface() != s[i].Interface():		t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, count, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(k interface{}, v R.Value) {
			switch {
			case k != Offset(s, inc):							t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), k)
			case v.Interface() != s[k.(int)].Interface():		t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, k, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(k, v R.Value) {
			switch {
			case k.Interface() != Offset(s, inc):				t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), k)
			case v.Interface() != s[int(k.Int())].Interface():	t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, k, v)
			}
			count++
		})

		ConfirmVariadicStep(s, inc, func(v ...R.Value) {
			switch {
			case count != 0:							t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):				t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})
	}
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, -5)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, -4)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, -3)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, -2)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, -1)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 1)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 2)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 3)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 4)
	ConfirmStepRValueSlice([]R.Value{R.ValueOf(0), R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 5)

	ConfirmStepRValue := func(s interface{}, inc int) {
		ConfirmStepBuiltin(s, inc)
		ConfirmVariadicStep(s, inc, func(v ...int) {
			switch {
			case count != 0:					t.Fatalf("%T increment %v: variadic function erroneously called %v times", s, inc, count)
			case len(v) != Sublen(s, inc):		t.Fatalf("%T increment %v: variadic slice %v erroneously passed as %v", s, inc, s, v)
			}
			count++
		})

		ConfirmStep(s, inc, func(v int) {
			if v != AtOffset(s, Offset(s, inc)) {
				t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, Offset(s, inc), v)
			}
			count++
		})

		ConfirmStep(s, inc, func(i int, v int) {
			switch {
			case i != Offset(s, inc):			t.Fatalf("%T increment %v: index %v erroneously reported as %v", s, inc, Offset(s, inc), i)
			case v != AtOffset(s, i):			t.Fatalf("%T increment %v: element %v erroneously reported as %v", s, inc, i, v)
			}
			count++
		})
	}
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), -5)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), -4)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), -3)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), -2)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), -1)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), 1)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), 2)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), 3)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), 4)
	ConfirmStepRValue(R.ValueOf([]int{0, 1, 2, 3, 4}), 5)
}



//	Write tests for the 0 case which should always fail to iterate
//	ConfirmStepBool([]bool{true, false, false, true, true}, 0)