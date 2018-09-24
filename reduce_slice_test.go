package sequences

import (
	R "reflect"
	"testing"
)

func TestReduceBoolSlice(t *testing.T) {
	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceBoolIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]bool{}, false, false)
	ConfirmReduce([]bool{}, true, true)
	ConfirmReduce([]bool{false}, false, false)
	ConfirmReduce([]bool{false}, true, false)
	ConfirmReduce([]bool{true}, false, false)
	ConfirmReduce([]bool{true}, true, true)
	ConfirmReduce([]bool{false, true}, true, false)
	ConfirmReduce([]bool{true, false}, true, false)
	ConfirmReduce([]bool{true, true}, true, true)
	ConfirmReduce([]bool{true, true, true}, true, true)
	ConfirmReduce([]bool{true, true, false, true}, true, false)
}

func TestReduceComplex64Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result complex64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceComplex64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]complex64{}, 0, 0)
	ConfirmReduce([]complex64{}, 1, 1)
	ConfirmReduce([]complex64{0}, 0, 0)
	ConfirmReduce([]complex64{0}, 1, 1)
	ConfirmReduce([]complex64{1}, 0, 1)
	ConfirmReduce([]complex64{1}, 1, 2)
	ConfirmReduce([]complex64{1, 2}, 0, 3)
	ConfirmReduce([]complex64{1, 2}, 10, 13)
	ConfirmReduce([]complex64{1, 2, 3}, 0, 6)
	ConfirmReduce([]complex64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]complex64{1, 2, 3, 4}, 10, 20)
}

func TestReduceComplex128Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result complex128) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceComplex128Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]complex128{}, 0, 0)
	ConfirmReduce([]complex128{}, 1, 1)
	ConfirmReduce([]complex128{0}, 0, 0)
	ConfirmReduce([]complex128{0}, 1, 1)
	ConfirmReduce([]complex128{1}, 0, 1)
	ConfirmReduce([]complex128{1}, 1, 2)
	ConfirmReduce([]complex128{1, 2}, 0, 3)
	ConfirmReduce([]complex128{1, 2}, 10, 13)
	ConfirmReduce([]complex128{1, 2, 3}, 0, 6)
	ConfirmReduce([]complex128{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]complex128{1, 2, 3, 4}, 10, 20)
}

func TestReduceErrorSlice(t *testing.T) {
	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceErrorIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]error{}, Error(0), Error(0))
	ConfirmReduce([]error{}, Error(8), Error(8))
	ConfirmReduce([]error{Error(0)}, Error(0), Error(0))
	ConfirmReduce([]error{Error(0)}, Error(8), Error(8))
	ConfirmReduce([]error{Error(1)}, Error(0), Error(1))
	ConfirmReduce([]error{Error(1)}, Error(8), Error(9))
	ConfirmReduce([]error{Error(0), Error(1)}, Error(0), Error(1))
	ConfirmReduce([]error{Error(0), Error(1)}, Error(8), Error(9))
	ConfirmReduce([]error{Error(0), Error(1), Error(2)}, Error(0), Error(3))
	ConfirmReduce([]error{Error(0), Error(1), Error(2)}, Error(8), Error(11))
	ConfirmReduce([]error{Error(0), Error(1), Error(2), Error(4)}, Error(0), Error(7))
	ConfirmReduce([]error{Error(0), Error(1), Error(2), Error(4)}, Error(8), Error(15))
}

func TestReduceFloat32Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result float32) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceFloat32Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]float32{}, 0, 0)
	ConfirmReduce([]float32{}, 1, 1)
	ConfirmReduce([]float32{0}, 0, 0)
	ConfirmReduce([]float32{0}, 1, 1)
	ConfirmReduce([]float32{1}, 0, 1)
	ConfirmReduce([]float32{1}, 1, 2)
	ConfirmReduce([]float32{1, 2}, 0, 3)
	ConfirmReduce([]float32{1, 2}, 10, 13)
	ConfirmReduce([]float32{1, 2, 3}, 0, 6)
	ConfirmReduce([]float32{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]float32{1, 2, 3, 4}, 10, 20)
}

func TestReduceFloat64Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result float64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceFloat64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]float64{}, 0, 0)
	ConfirmReduce([]float64{}, 1, 1)
	ConfirmReduce([]float64{0}, 0, 0)
	ConfirmReduce([]float64{0}, 1, 1)
	ConfirmReduce([]float64{1}, 0, 1)
	ConfirmReduce([]float64{1}, 1, 2)
	ConfirmReduce([]float64{1, 2}, 0, 3)
	ConfirmReduce([]float64{1, 2}, 10, 13)
	ConfirmReduce([]float64{1, 2, 3}, 0, 6)
	ConfirmReduce([]float64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]float64{1, 2, 3, 4}, 10, 20)
}

func TestReduceIntSlice(t *testing.T) {
	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceIntIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int{}, 0, 0)
	ConfirmReduce([]int{}, 10, 10)
	ConfirmReduce([]int{0}, 0, 0)
	ConfirmReduce([]int{0}, 10, 10)
	ConfirmReduce([]int{1}, 0, 1)
	ConfirmReduce([]int{1}, 10, 11)
	ConfirmReduce([]int{1, 2}, 0, 3)
	ConfirmReduce([]int{1, 2}, 10, 13)
	ConfirmReduce([]int{1, 2, 3}, 0, 6)
	ConfirmReduce([]int{1, 2, 3}, 10, 16)
	ConfirmReduce([]int{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt8Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result int8) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt8Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int8{}, 0, 0)
	ConfirmReduce([]int8{}, 10, 10)
	ConfirmReduce([]int8{0}, 0, 0)
	ConfirmReduce([]int8{0}, 10, 10)
	ConfirmReduce([]int8{1}, 0, 1)
	ConfirmReduce([]int8{1}, 10, 11)
	ConfirmReduce([]int8{1, 2}, 0, 3)
	ConfirmReduce([]int8{1, 2}, 10, 13)
	ConfirmReduce([]int8{1, 2, 3}, 0, 6)
	ConfirmReduce([]int8{1, 2, 3}, 10, 16)
	ConfirmReduce([]int8{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int8{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt16Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result int16) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt16Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int16{}, 0, 0)
	ConfirmReduce([]int16{}, 10, 10)
	ConfirmReduce([]int16{0}, 0, 0)
	ConfirmReduce([]int16{0}, 10, 10)
	ConfirmReduce([]int16{1}, 0, 1)
	ConfirmReduce([]int16{1}, 10, 11)
	ConfirmReduce([]int16{1, 2}, 0, 3)
	ConfirmReduce([]int16{1, 2}, 10, 13)
	ConfirmReduce([]int16{1, 2, 3}, 0, 6)
	ConfirmReduce([]int16{1, 2, 3}, 10, 16)
	ConfirmReduce([]int16{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int16{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt32Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result int32) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt32Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int32{}, 0, 0)
	ConfirmReduce([]int32{}, 10, 10)
	ConfirmReduce([]int32{0}, 0, 0)
	ConfirmReduce([]int32{0}, 10, 10)
	ConfirmReduce([]int32{1}, 0, 1)
	ConfirmReduce([]int32{1}, 10, 11)
	ConfirmReduce([]int32{1, 2}, 0, 3)
	ConfirmReduce([]int32{1, 2}, 10, 13)
	ConfirmReduce([]int32{1, 2, 3}, 0, 6)
	ConfirmReduce([]int32{1, 2, 3}, 10, 16)
	ConfirmReduce([]int32{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int32{1, 2, 3, 4}, 10, 20)
}

func TestReduceInt64Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result int64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]int64{}, 0, 0)
	ConfirmReduce([]int64{}, 10, 10)
	ConfirmReduce([]int64{0}, 0, 0)
	ConfirmReduce([]int64{0}, 10, 10)
	ConfirmReduce([]int64{1}, 0, 1)
	ConfirmReduce([]int64{1}, 10, 11)
	ConfirmReduce([]int64{1, 2}, 0, 3)
	ConfirmReduce([]int64{1, 2}, 10, 13)
	ConfirmReduce([]int64{1, 2, 3}, 0, 6)
	ConfirmReduce([]int64{1, 2, 3}, 10, 16)
	ConfirmReduce([]int64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]int64{1, 2, 3, 4}, 10, 20)
}

func TestReduceInterfaceSlice(t *testing.T) {
	ConfirmReduce := func(iterators []interface{}, o interface{}, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{}, 0, 0)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{}, 10, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{0}, 0, 0)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{0}, 10, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1}, 0, 1)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1}, 10, 11)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1, 2}, 0, 3)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1, 2}, 10, 13)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1, 2, 3}, 0, 6)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1, 2, 3}, 10, 16)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1, 2, 3, 4}, 0, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), []interface{}{1, 2, 3, 4}, 10, 20)

	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{}, false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{}, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{false}, false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{false}, true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{true}, false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{true}, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{false, true}, true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{true, false}, true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{true, true}, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{true, true, true}, true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), []interface{}{true, true, false, true}, true, false)
}

func TestReduceStringSlice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result string) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceStringIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]string{}, "", "")
	ConfirmReduce([]string{}, "Z", "Z")
	ConfirmReduce([]string{""}, "", "")
	ConfirmReduce([]string{""}, "Z", "Z")
	ConfirmReduce([]string{"A"}, "", "A")
	ConfirmReduce([]string{"A"}, "Z", "ZA")
	ConfirmReduce([]string{"A", "B"}, "", "AB")
	ConfirmReduce([]string{"A", "B"}, "Z", "ZAB")
	ConfirmReduce([]string{"A", "B", "C"}, "", "ABC")
	ConfirmReduce([]string{"A", "B", "C"}, "Z", "ZABC")
	ConfirmReduce([]string{"A", "B", "C", "D"}, "", "ABCD")
	ConfirmReduce([]string{"A", "B", "C", "D"}, "Z", "ZABCD")
}

func TestReduceUintSlice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result uint) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUintIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint{}, 0, 0)
	ConfirmReduce([]uint{}, 10, 10)
	ConfirmReduce([]uint{0}, 0, 0)
	ConfirmReduce([]uint{0}, 10, 10)
	ConfirmReduce([]uint{1}, 0, 1)
	ConfirmReduce([]uint{1}, 10, 11)
	ConfirmReduce([]uint{1, 2}, 0, 3)
	ConfirmReduce([]uint{1, 2}, 10, 13)
	ConfirmReduce([]uint{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint8Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result uint8) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint8Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint8{}, 0, 0)
	ConfirmReduce([]uint8{}, 10, 10)
	ConfirmReduce([]uint8{0}, 0, 0)
	ConfirmReduce([]uint8{0}, 10, 10)
	ConfirmReduce([]uint8{1}, 0, 1)
	ConfirmReduce([]uint8{1}, 10, 11)
	ConfirmReduce([]uint8{1, 2}, 0, 3)
	ConfirmReduce([]uint8{1, 2}, 10, 13)
	ConfirmReduce([]uint8{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint8{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint8{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint8{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint16Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result uint16) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint16Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint16{}, 0, 0)
	ConfirmReduce([]uint16{}, 10, 10)
	ConfirmReduce([]uint16{0}, 0, 0)
	ConfirmReduce([]uint16{0}, 10, 10)
	ConfirmReduce([]uint16{1}, 0, 1)
	ConfirmReduce([]uint16{1}, 10, 11)
	ConfirmReduce([]uint16{1, 2}, 0, 3)
	ConfirmReduce([]uint16{1, 2}, 10, 13)
	ConfirmReduce([]uint16{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint16{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint16{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint16{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint32Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result uint32) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint32Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint32{}, 0, 0)
	ConfirmReduce([]uint32{}, 10, 10)
	ConfirmReduce([]uint32{0}, 0, 0)
	ConfirmReduce([]uint32{0}, 10, 10)
	ConfirmReduce([]uint32{1}, 0, 1)
	ConfirmReduce([]uint32{1}, 10, 11)
	ConfirmReduce([]uint32{1, 2}, 0, 3)
	ConfirmReduce([]uint32{1, 2}, 10, 13)
	ConfirmReduce([]uint32{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint32{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint32{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint32{1, 2, 3, 4}, 10, 20)
}

func TestReduceUint64Slice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result uint64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uint64{}, 0, 0)
	ConfirmReduce([]uint64{}, 10, 10)
	ConfirmReduce([]uint64{0}, 0, 0)
	ConfirmReduce([]uint64{0}, 10, 10)
	ConfirmReduce([]uint64{1}, 0, 1)
	ConfirmReduce([]uint64{1}, 10, 11)
	ConfirmReduce([]uint64{1, 2}, 0, 3)
	ConfirmReduce([]uint64{1, 2}, 10, 13)
	ConfirmReduce([]uint64{1, 2, 3}, 0, 6)
	ConfirmReduce([]uint64{1, 2, 3}, 10, 16)
	ConfirmReduce([]uint64{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uint64{1, 2, 3, 4}, 10, 20)
}

func TestReduceUintptrSlice(t *testing.T) {
	ConfirmReduce := func(o interface{}, seed, result uintptr) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUintptrIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce([]uintptr{}, 0, 0)
	ConfirmReduce([]uintptr{}, 10, 10)
	ConfirmReduce([]uintptr{0}, 0, 0)
	ConfirmReduce([]uintptr{0}, 10, 10)
	ConfirmReduce([]uintptr{1}, 0, 1)
	ConfirmReduce([]uintptr{1}, 10, 11)
	ConfirmReduce([]uintptr{1, 2}, 0, 3)
	ConfirmReduce([]uintptr{1, 2}, 10, 13)
	ConfirmReduce([]uintptr{1, 2, 3}, 0, 6)
	ConfirmReduce([]uintptr{1, 2, 3}, 10, 16)
	ConfirmReduce([]uintptr{1, 2, 3, 4}, 0, 10)
	ConfirmReduce([]uintptr{1, 2, 3, 4}, 10, 20)
}

func TestReduceRValueSlice(t *testing.T) {
	ConfirmReduceInterfaceSeed := func(iterators []interface{}, o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(iterators []interface{}, o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{}, 0, 0)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{}, 10, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(0)}, 0, 0)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(0)}, 10, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1)}, 0, 1)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1)}, 10, 11)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2)}, 0, 3)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2)}, 10, 13)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, 0, 6)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, 10, 16)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 0, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 10, 20)

	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{}, 0, 0)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{}, 10, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(0)}, 0, 0)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(0)}, 10, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1)}, 0, 1)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1)}, 10, 11)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2)}, 0, 3)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2)}, 10, 13)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, 0, 6)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3)}, 10, 16)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 0, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), []R.Value{R.ValueOf(1), R.ValueOf(2), R.ValueOf(3), R.ValueOf(4)}, 10, 20)

	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{}, false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{}, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(false)}, false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(false)}, true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true)}, false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true)}, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(false), R.ValueOf(true)}, true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(false)}, true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(true)}, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(true)}, true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(false), R.ValueOf(true)}, true, false)

	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{}, false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{}, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(false)}, false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(false)}, true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true)}, false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true)}, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(false), R.ValueOf(true)}, true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(false)}, true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(true)}, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(true)}, true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), []R.Value{R.ValueOf(true), R.ValueOf(true), R.ValueOf(false), R.ValueOf(true)}, true, false)
}

func TestReduceSlice(t *testing.T) {
	s := func(v ...UDT) []UDT {
		return v
	}

	ConfirmReduceInterfaceSeed := func(o interface{}, seed, result UDT) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceRValueChannel() {
			if v := Reduce(o, seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(o []UDT, seed, result UDT) {
		defer reportReductionError(t, o, seed)
		ov := R.ValueOf(o)
		for n, f := range ReduceRValueChannel() {
			if v := Reduce(ov, R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", ov, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(s(), 0, 0)
	ConfirmReduceInterfaceSeed(s(), 10, 10)
	ConfirmReduceInterfaceSeed(s(0), 0, 0)
	ConfirmReduceInterfaceSeed(s(0), 10, 10)
	ConfirmReduceInterfaceSeed(s(1), 0, 1)
	ConfirmReduceInterfaceSeed(s(1), 10, 11)
	ConfirmReduceInterfaceSeed(s(1, 2), 0, 3)
	ConfirmReduceInterfaceSeed(s(1, 2), 10, 13)
	ConfirmReduceInterfaceSeed(s(1, 2, 3), 0, 6)
	ConfirmReduceInterfaceSeed(s(1, 2, 3), 10, 16)
	ConfirmReduceInterfaceSeed(s(1, 2, 3, 4), 0, 10)
	ConfirmReduceInterfaceSeed(s(1, 2, 3, 4), 10, 20)

	ConfirmReduceValueSeed(s(), 0, 0)
	ConfirmReduceValueSeed(s(), 10, 10)
	ConfirmReduceValueSeed(s(0), 0, 0)
	ConfirmReduceValueSeed(s(0), 10, 10)
	ConfirmReduceValueSeed(s(1), 0, 1)
	ConfirmReduceValueSeed(s(1), 10, 11)
	ConfirmReduceValueSeed(s(1, 2), 0, 3)
	ConfirmReduceValueSeed(s(1, 2), 10, 13)
	ConfirmReduceValueSeed(s(1, 2, 3), 0, 6)
	ConfirmReduceValueSeed(s(1, 2, 3), 10, 16)
	ConfirmReduceValueSeed(s(1, 2, 3, 4), 0, 10)
	ConfirmReduceValueSeed(s(1, 2, 3, 4), 10, 20)
}
