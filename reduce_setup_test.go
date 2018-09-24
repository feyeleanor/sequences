package sequences

import (
	R "reflect"
	"testing"
)

func reportReductionError(t *testing.T, o, seed interface{}) {
	if e := recover(); e != nil {
		panic(e)
		t.Fatalf("Reduce(%v, %v, f) iteration failed with error %v", o, seed, e)
	}
}

func ReduceBoolIterators() []interface{} {
	return []interface{}{
		func(seed, v bool) bool {
			return seed && v
		},
		func(seed, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
		func(seed bool, index int, v bool) bool {
			return seed && v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
	}
}

func ReduceComplex64Iterators() []interface{} {
	return []interface{}{
		func(seed, v complex64) complex64 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(complex64) + v.(complex64)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Complex() + v.Complex())
		},
		func(seed complex64, index int, v complex64) complex64 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(complex64) + v.(complex64)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Complex() + v.Complex())
		},
	}
}

func ReduceComplex128Iterators() []interface{} {
	return []interface{}{
		func(seed, v complex128) complex128 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(complex128) + v.(complex128)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Complex() + v.Complex())
		},
		func(seed complex128, index int, v complex128) complex128 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(complex128) + v.(complex128)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Complex() + v.Complex())
		},
	}
}

func ReduceErrorIterators() []interface{} {
	return []interface{}{
		func(seed, v error) error {
			return seed.(Error) | v.(Error)
		},
		func(seed, v interface{}) interface{} {
			return seed.(Error) | v.(Error)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Interface().(Error) | v.Interface().(Error))
		},
		func(seed error, index int, v error) error {
			return seed.(Error) | v.(Error)
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(Error) | v.(Error)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Interface().(Error) | v.Interface().(Error))
		},
	}
}

func ReduceFloat32Iterators() []interface{} {
	return []interface{}{
		func(seed, v float32) float32 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(float32) + v.(float32)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Float() + v.Float())
		},
		func(seed float32, index int, v float32) float32 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(float32) + v.(float32)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Float() + v.Float())
		},
	}
}

func ReduceFloat64Iterators() []interface{} {
	return []interface{}{
		func(seed, v float64) float64 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(float64) + v.(float64)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Float() + v.Float())
		},
		func(seed float64, index int, v float64) float64 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(float64) + v.(float64)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Float() + v.Float())
		},
	}
}

func ReduceIntIterators() []interface{} {
	return []interface{}{
		func(seed, v int) int {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
		func(seed int, index int, v int) int {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
	}
}

func ReduceInt8Iterators() []interface{} {
	return []interface{}{
		func(seed, v int8) int8 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(int8) + v.(int8)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(int8(seed.Int()) + int8(v.Int()))
		},
		func(seed int8, index int, v int8) int8 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int8) + v.(int8)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(int8(seed.Int()) + int8(v.Int()))
		},
	}
}

func ReduceInt16Iterators() []interface{} {
	return []interface{}{
		func(seed, v int16) int16 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(int16) + v.(int16)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(int16(seed.Int()) + int16(v.Int()))
		},
		func(seed int16, index int, v int16) int16 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int16) + v.(int16)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(int16(seed.Int()) + int16(v.Int()))
		},
	}
}

func ReduceInt32Iterators() []interface{} {
	return []interface{}{
		func(seed, v int32) int32 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(int32) + v.(int32)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(int32(seed.Int()) + int32(v.Int()))
		},
		func(seed int32, index int, v int32) int32 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int32) + v.(int32)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(int32(seed.Int()) + int32(v.Int()))
		},
	}
}

func ReduceInt64Iterators() []interface{} {
	return []interface{}{
		func(seed, v int64) int64 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(int64) + v.(int64)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
		func(seed int64, index int, v int64) int64 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int64) + v.(int64)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Int() + v.Int())
		},
	}
}

func ReduceInterfaceIteratorsInt() []interface{} {
	return []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(int(seed.Int() + v.Int()))
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(int(seed.Int() + v.Int()))
		},
	}
}

func ReduceInterfaceIteratorsBool() []interface{} {
	return []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
	}
}

func ReduceStringIterators() []interface{} {
	return []interface{}{
		func(seed, v string) string {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(string) + v.(string)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.String() + v.String())
		},
		func(seed string, index int, v string) string {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(string) + v.(string)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.String() + v.String())
		},
	}
}

func ReduceUintIterators() []interface{} {
	return []interface{}{
		func(seed, v uint) uint {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(uint) + v.(uint)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(uint(seed.Uint() + v.Uint()))
		},
		func(seed uint, index int, v uint) uint {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(uint) + v.(uint)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(uint(seed.Uint() + v.Uint()))
		},
	}
}

func ReduceUint8Iterators() []interface{} {
	return []interface{}{
		func(seed, v uint8) uint8 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(uint8) + v.(uint8)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(uint8(seed.Uint()) + uint8(v.Uint()))
		},
		func(seed uint8, index int, v uint8) uint8 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(uint8) + v.(uint8)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(uint8(seed.Uint()) + uint8(v.Uint()))
		},
	}
}

func ReduceUint16Iterators() []interface{} {
	return []interface{}{
		func(seed, v uint16) uint16 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(uint16) + v.(uint16)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(uint16(seed.Uint()) + uint16(v.Uint()))
		},
		func(seed uint16, index int, v uint16) uint16 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(uint16) + v.(uint16)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(uint16(seed.Uint()) + uint16(v.Uint()))
		},
	}
}

func ReduceUint32Iterators() []interface{} {
	return []interface{}{
		func(seed, v uint32) uint32 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(uint32) + v.(uint32)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(uint32(seed.Uint()) + uint32(v.Uint()))
		},
		func(seed uint32, index int, v uint32) uint32 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(uint32) + v.(uint32)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(uint32(seed.Uint()) + uint32(v.Uint()))
		},
	}
}

func ReduceUint64Iterators() []interface{} {
	return []interface{}{
		func(seed, v uint64) uint64 {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(uint64) + v.(uint64)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Uint() + v.Uint())
		},
		func(seed uint64, index int, v uint64) uint64 {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(uint64) + v.(uint64)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Uint() + v.Uint())
		},
	}
}

func ReduceUintptrIterators() []interface{} {
	return []interface{}{
		func(seed, v uintptr) uintptr {
			return seed + v
		},
		func(seed, v interface{}) interface{} {
			return seed.(uintptr) + v.(uintptr)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(uintptr(seed.Uint()) + uintptr(v.Uint()))
		},
		func(seed uintptr, index int, v uintptr) uintptr {
			return seed + v
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(uintptr) + v.(uintptr)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(uintptr(seed.Uint()) + uintptr(v.Uint()))
		},
	}
}

func ReduceRValueIteratorsInt() []interface{} {
	return []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(int(seed.Int() + v.Int()))
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(int) + v.(int)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(int(seed.Int() + v.Int()))
		},
	}
}

func ReduceRValueIteratorsBool() []interface{} {
	return []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
	}
}

func ReduceRValueChannel() []interface{} {
	return []interface{}{
		func(seed, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
		func(seed interface{}, index int, v interface{}) interface{} {
			return seed.(bool) && v.(bool)
		},
		func(seed R.Value, index int, v R.Value) R.Value {
			return R.ValueOf(seed.Bool() && v.Bool())
		},
	}
}
