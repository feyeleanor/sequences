package sequences

import (
	R "reflect"
	"testing"
)

func TestReduceBoolFunction(t *testing.T) {
	seq := func(s ...bool) func(int) bool {
		return func(index int) bool {
			return s[index]
		}
	}

	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceBoolIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), false, false)
	ConfirmReduce(seq(), true, true)
	ConfirmReduce(seq(false), false, false)
	ConfirmReduce(seq(false), true, false)
	ConfirmReduce(seq(true), false, false)
	ConfirmReduce(seq(true), true, true)
	ConfirmReduce(seq(false, true), true, false)
	ConfirmReduce(seq(true, false), true, false)
	ConfirmReduce(seq(true, true), true, true)
	ConfirmReduce(seq(true, true, true), true, true)
	ConfirmReduce(seq(true, true, false, true), true, false)
}

func TestReduceComplex64Function(t *testing.T) {
	seq := func(s ...complex64) func(int) complex64 {
		return func(index int) complex64 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result complex64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceComplex64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 1, 1)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 1, 1)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 1, 2)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceComplex128Function(t *testing.T) {
	seq := func(s ...complex128) func(int) complex128 {
		return func(index int) complex128 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result complex128) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceComplex128Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 1, 1)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 1, 1)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 1, 2)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceErrorFunction(t *testing.T) {
	seq := func(s ...error) func(int) error {
		return func(index int) error {
			return s[index]
		}
	}

	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceErrorIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), Error(0), Error(0))
	ConfirmReduce(seq(), Error(8), Error(8))
	ConfirmReduce(seq(Error(0)), Error(0), Error(0))
	ConfirmReduce(seq(Error(0)), Error(8), Error(8))
	ConfirmReduce(seq(Error(1)), Error(0), Error(1))
	ConfirmReduce(seq(Error(1)), Error(8), Error(9))
	ConfirmReduce(seq(Error(0), Error(1)), Error(0), Error(1))
	ConfirmReduce(seq(Error(0), Error(1)), Error(8), Error(9))
	ConfirmReduce(seq(Error(0), Error(1), Error(2)), Error(0), Error(3))
	ConfirmReduce(seq(Error(0), Error(1), Error(2)), Error(8), Error(11))
	ConfirmReduce(seq(Error(0), Error(1), Error(2), Error(4)), Error(0), Error(7))
	ConfirmReduce(seq(Error(0), Error(1), Error(2), Error(4)), Error(8), Error(15))
}

func TestReduceFloat32Function(t *testing.T) {
	seq := func(s ...float32) func(int) float32 {
		return func(index int) float32 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result float32) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceFloat32Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 1, 1)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 1, 1)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 1, 2)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceFloat64Function(t *testing.T) {
	seq := func(s ...float64) func(int) float64 {
		return func(index int) float64 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result float64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceFloat64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 1, 1)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 1, 1)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 1, 2)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceIntFunction(t *testing.T) {
	seq := func(s ...int) func(int) int {
		return func(index int) int {
			return s[index]
		}
	}

	ConfirmReduce := func(o, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceIntIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceInt8Function(t *testing.T) {
	seq := func(s ...int8) func(int) int8 {
		return func(index int) int8 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result int8) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt8Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceInt16Function(t *testing.T) {
	seq := func(s ...int16) func(int) int16 {
		return func(index int) int16 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result int16) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt16Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceInt32Function(t *testing.T) {
	seq := func(s ...int32) func(int) int32 {
		return func(index int) int32 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result int32) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt32Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceInt64Function(t *testing.T) {
	seq := func(s ...int64) func(int) int64 {
		return func(index int) int64 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result int64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceInt64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceInterfaceFunction(t *testing.T) {
	seq := func(s ...interface{}) func(int) interface{} {
		return func(index int) interface{} {
			return s[index]
		}
	}

	ConfirmReduce := func(iterators []interface{}, o interface{}, seed, result interface{}) {
		defer reportReductionError(t, o, seed)
		for n, f := range iterators {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(), 0, 0)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(), 10, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(0), 0, 0)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(0), 10, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1), 0, 1)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1), 10, 11)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1, 2), 0, 3)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1, 2), 10, 13)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1, 2, 3), 0, 6)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1, 2, 3), 10, 16)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(ReduceInterfaceIteratorsInt(), seq(1, 2, 3, 4), 10, 20)

	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(), false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(), true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(false), false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(false), true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(true), false, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(true), true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(false, true), true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(true, false), true, false)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(true, true), true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(true, true, true), true, true)
	ConfirmReduce(ReduceInterfaceIteratorsBool(), seq(true, true, false, true), true, false)
}

func TestReduceStringFunction(t *testing.T) {
	seq := func(s ...string) func(int) string {
		return func(index int) string {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result string) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceStringIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), "", "")
	ConfirmReduce(seq(), "Z", "Z")
	ConfirmReduce(seq(""), "", "")
	ConfirmReduce(seq(""), "Z", "Z")
	ConfirmReduce(seq("A"), "", "A")
	ConfirmReduce(seq("A"), "Z", "ZA")
	ConfirmReduce(seq("A", "B"), "", "AB")
	ConfirmReduce(seq("A", "B"), "Z", "ZAB")
	ConfirmReduce(seq("A", "B", "C"), "", "ABC")
	ConfirmReduce(seq("A", "B", "C"), "Z", "ZABC")
	ConfirmReduce(seq("A", "B", "C", "D"), "", "ABCD")
	ConfirmReduce(seq("A", "B", "C", "D"), "Z", "ZABCD")
}

func TestReduceUintFunction(t *testing.T) {
	seq := func(s ...uint) func(int) uint {
		return func(index int) uint {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result uint) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUintIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceUint8Function(t *testing.T) {
	seq := func(s ...uint8) func(int) uint8 {
		return func(index int) uint8 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result uint8) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint8Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceUint16Function(t *testing.T) {
	seq := func(s ...uint16) func(int) uint16 {
		return func(index int) uint16 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result uint16) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint16Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceUint32Function(t *testing.T) {
	seq := func(s ...uint32) func(int) uint32 {
		return func(index int) uint32 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result uint32) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint32Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceUint64Function(t *testing.T) {
	seq := func(s ...uint64) func(int) uint64 {
		return func(index int) uint64 {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result uint64) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUint64Iterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceUintptrFunction(t *testing.T) {
	seq := func(s ...uintptr) func(int) uintptr {
		return func(index int) uintptr {
			return s[index]
		}
	}

	ConfirmReduce := func(o interface{}, seed, result uintptr) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceUintptrIterators() {
			if v := Reduce(o, seed, f); v != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v)
			}
		}
	}

	ConfirmReduce(seq(), 0, 0)
	ConfirmReduce(seq(), 10, 10)
	ConfirmReduce(seq(0), 0, 0)
	ConfirmReduce(seq(0), 10, 10)
	ConfirmReduce(seq(1), 0, 1)
	ConfirmReduce(seq(1), 10, 11)
	ConfirmReduce(seq(1, 2), 0, 3)
	ConfirmReduce(seq(1, 2), 10, 13)
	ConfirmReduce(seq(1, 2, 3), 0, 6)
	ConfirmReduce(seq(1, 2, 3), 10, 16)
	ConfirmReduce(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduce(seq(1, 2, 3, 4), 10, 20)
}

func TestReduceRValueFunction(t *testing.T) {
	seq := func(s ...interface{}) func(int) R.Value {
		return func(index int) R.Value {
			return R.ValueOf(s[index])
		}
	}

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

	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(), 0, 0)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(), 10, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(0), 0, 0)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(0), 10, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1), 0, 1)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1), 10, 11)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1, 2), 0, 3)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1, 2), 10, 13)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3), 0, 6)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3), 10, 16)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3, 4), 0, 10)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3, 4), 10, 20)

	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(), 0, 0)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(), 10, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(0), 0, 0)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(0), 10, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1), 0, 1)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1), 10, 11)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1, 2), 0, 3)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1, 2), 10, 13)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3), 0, 6)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3), 10, 16)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3, 4), 0, 10)
	ConfirmReduceValueSeed(ReduceRValueIteratorsInt(), seq(1, 2, 3, 4), 10, 20)

	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(), false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(), true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(false), false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(false), true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(true), false, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(true), true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(false, true), true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(true, false), true, false)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(true, true), true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(true, true, true), true, true)
	ConfirmReduceInterfaceSeed(ReduceRValueIteratorsBool(), seq(true, true, false, true), true, false)

	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(), false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(), true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(false), false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(false), true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(true), false, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(true), true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(false, true), true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(true, false), true, false)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(true, true), true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(true, true, true), true, true)
	ConfirmReduceValueSeed(ReduceRValueIteratorsBool(), seq(true, true, false, true), true, false)
}

func TestReduceFunction(t *testing.T) {
	seq := func(s ...UDT) func(int) UDT {
		return func(index int) UDT {
			return s[index]
		}
	}

	ConfirmReduceInterfaceSeed := func(o interface{}, seed, result UDT) {
		defer reportReductionError(t, o, seed)
		for n, f := range ReduceRValueChannel() {
			if v := Reduce(o, seed, f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", o, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceValueSeed := func(o func(int) UDT, seed, result UDT) {
		defer reportReductionError(t, o, seed)
		ov := R.ValueOf(o)
		for n, f := range ReduceRValueChannel() {
			if v := Reduce(ov, R.ValueOf(seed), f); v.(R.Value).Interface() != result {
				t.Fatalf("Reduce(%v, %v, iterators[%v]): should have returned %v but instead returned %v", ov, seed, n, result, v.(R.Value).Interface())
			}
		}
	}

	ConfirmReduceInterfaceSeed(seq(), 0, 0)
	ConfirmReduceInterfaceSeed(seq(), 10, 10)
	ConfirmReduceInterfaceSeed(seq(0), 0, 0)
	ConfirmReduceInterfaceSeed(seq(0), 10, 10)
	ConfirmReduceInterfaceSeed(seq(1), 0, 1)
	ConfirmReduceInterfaceSeed(seq(1), 10, 11)
	ConfirmReduceInterfaceSeed(seq(1, 2), 0, 3)
	ConfirmReduceInterfaceSeed(seq(1, 2), 10, 13)
	ConfirmReduceInterfaceSeed(seq(1, 2, 3), 0, 6)
	ConfirmReduceInterfaceSeed(seq(1, 2, 3), 10, 16)
	ConfirmReduceInterfaceSeed(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduceInterfaceSeed(seq(1, 2, 3, 4), 10, 20)

	ConfirmReduceValueSeed(seq(), 0, 0)
	ConfirmReduceValueSeed(seq(), 10, 10)
	ConfirmReduceValueSeed(seq(0), 0, 0)
	ConfirmReduceValueSeed(seq(0), 10, 10)
	ConfirmReduceValueSeed(seq(1), 0, 1)
	ConfirmReduceValueSeed(seq(1), 10, 11)
	ConfirmReduceValueSeed(seq(1, 2), 0, 3)
	ConfirmReduceValueSeed(seq(1, 2), 10, 13)
	ConfirmReduceValueSeed(seq(1, 2, 3), 0, 6)
	ConfirmReduceValueSeed(seq(1, 2, 3), 10, 16)
	ConfirmReduceValueSeed(seq(1, 2, 3, 4), 0, 10)
	ConfirmReduceValueSeed(seq(1, 2, 3, 4), 10, 20)
}
