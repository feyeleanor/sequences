package sequences

import "R" reflect

type Collectable interface {
	Collect(interface{}) (interface{}, error)
}

func Collect(seq, f interface{}) (r interface{}, e error) {
	switch seq := seq.(type) {
	case Collectable:
		r, e = seq.Collect(f)

	case Enumerable:
		panic()
		//	TODO: enumerable doesn't currently provide sufficient information to perform a collection using the provided function f
		//	TODO: think of how enumerable might be changed to remedy this
		//
		//	TODO: a simple copy takes the form
		//
		//		r = make([]interface{}, 0)
		//		seq.Each(func(v interface{} {
		//			r = append(r, v)
		//		})
		//
		//	TODO: the problem is how to get something like
		//
		//		r = make([]interface{}, 0)
		//		seq.Each(func(v interface{} {
		//			r = append(r, f(v))
		//		})
		//
		//	TODO: where f is the function passed to Collect as an interface{} value
		//	TODO: perhaps the case for Indexable might hold some clues?

	case Indexable:
		switch f := f.(type) {
		case func(interface{}):
			r = make([]interface{}, 0, seq.Len())
			e = Each(seq, func(v interface{}) {
				r = append(r, f(v))
			})
		case func(R.Value):
			r = make([]interface{}, 0, seq.Len())
			e = Each(seq, func(v R.Value) {
				r = append(r, f(v))
			})
		case func(int, interface{}):
			r = make([]interface{}, 0, seq.Len())
			e = Each(seq, func(i int, v interface{}) {
				r = append(r, f(i, v))
			})
		case func(int, R.Value):
			r = make([]interface{}, 0, seq.Len())
			e = Each(seq, func(i int, v R.Value) {
				r = append(r, f(i, v))
			})
		default:
			e = NOT_AN_ENUMERATOR
		}

	case []bool:
		e = eachBoolSlice(seq, f)
	case map[int]bool:
		e = eachBoolMap(seq, f)
	case chan bool:
		e = eachBoolChannel(seq, f)
	case func(int) bool:
		e = eachBoolFunction(seq, f)
	case func(int) (bool, bool):
		e = eachBoolBoundedFunction(seq, f)

	case []complex64:
		e = eachComplex64Slice(seq, f)
	case map[int]complex64:
		e = eachComplex64Map(seq, f)
	case chan complex64:
		e = eachComplex64Channel(seq, f)
	case func(int) complex64:
		e = eachComplex64Function(seq, f)
	case func(int) (complex64, bool):
		e = eachComplex64BoundedFunction(seq, f)

	case []complex128:
		e = eachComplex128Slice(seq, f)
	case map[int]complex128:
		e = eachComplex128Map(seq, f)
	case chan complex128:
		e = eachComplex128Channel(seq, f)
	case func(int) complex128:
		e = eachComplex128Function(seq, f)
	case func(int) (complex128, bool):
		e = eachComplex128BoundedFunction(seq, f)

	case []error:
		e = eachErrorSlice(seq, f)
	case map[int]error:
		e = eachErrorMap(seq, f)
	case chan error:
		e = eachErrorChannel(seq, f)
	case func(int) error:
		e = eachErrorFunction(seq, f)
	case func(int) (error, bool):
		e = eachErrorBoundedFunction(seq, f)

	case []float32:
		e = eachFloat32Slice(seq, f)
	case map[int]float32:
		e = eachFloat32Map(seq, f)
	case chan float32:
		e = eachFloat32Channel(seq, f)
	case func(int) float32:
		e = eachFloat32Function(seq, f)
	case func(int) (float32, bool):
		e = eachFloat32BoundedFunction(seq, f)

	case []float64:
		e = eachFloat64Slice(seq, f)
	case map[int]float64:
		e = eachFloat64Map(seq, f)
	case chan float64:
		e = eachFloat64Channel(seq, f)
	case func(int) float64:
		e = eachFloat64Function(seq, f)
	case func(int) (float64, bool):
		e = eachFloat64BoundedFunction(seq, f)

	case []int:
		e = eachIntSlice(seq, f)
	case map[int]int:
		e = eachIntMap(seq, f)
	case chan int:
		e = eachIntChannel(seq, f)
	case func(int) int:
		e = eachIntFunction(seq, f)
	case func(int) (int, bool):
		e = eachIntBoundedFunction(seq, f)

	case []int8:
		e = eachInt8Slice(seq, f)
	case map[int]int8:
		e = eachInt8Map(seq, f)
	case chan int8:
		e = eachInt8Channel(seq, f)
	case func(int) int8:
		e = eachInt8Function(seq, f)
	case func(int) (int8, bool):
		e = eachInt8BoundedFunction(seq, f)

	case []int16:
		e = eachInt16Slice(seq, f)
	case map[int]int16:
		e = eachInt16Map(seq, f)
	case chan int16:
		e = eachInt16Channel(seq, f)
	case func(int) int16:
		e = eachInt16Function(seq, f)
	case func(int) (int16, bool):
		e = eachInt16BoundedFunction(seq, f)

	case []int32:
		e = eachInt32Slice(seq, f)
	case map[int]int32:
		e = eachInt32Map(seq, f)
	case chan int32:
		e = eachInt32Channel(seq, f)
	case func(int) int32:
		e = eachInt32Function(seq, f)
	case func(int) (int32, bool):
		e = eachInt32BoundedFunction(seq, f)

	case []int64:
		e = eachInt64Slice(seq, f)
	case map[int]int64:
		e = eachInt64Map(seq, f)
	case chan int64:
		e = eachInt64Channel(seq, f)
	case func(int) int64:
		e = eachInt64Function(seq, f)
	case func(int) (int64, bool):
		e = eachInt64BoundedFunction(seq, f)

	case []interface{}:
		e = eachInterfaceSlice(seq, f)
	case map[int]interface{}:
		e = eachInterfaceMap(seq, f)
	case chan interface{}:
		e = eachInterfaceChannel(seq, f)
	case func(int) interface{}:
		e = eachInterfaceFunction(seq, f)
	case func(int) (interface{}, bool):
		e = eachInterfaceBoundedFunction(seq, f)

	case []string:
		e = eachStringSlice(seq, f)
	case map[int]string:
		e = eachStringMap(seq, f)
	case chan string:
		e = eachStringChannel(seq, f)
	case func(int) string:
		e = eachStringFunction(seq, f)
	case func(int) (string, bool):
		e = eachStringBoundedFunction(seq, f)

	case []uint:
		e = eachUintSlice(seq, f)
	case map[int]uint:
		e = eachUintMap(seq, f)
	case chan uint:
		e = eachUintChannel(seq, f)
	case func(int) uint:
		e = eachUintFunction(seq, f)
	case func(int) (uint, bool):
		e = eachUintBoundedFunction(seq, f)

	case []uint8:
		e = eachUint8Slice(seq, f)
	case map[int]uint8:
		e = eachUint8Map(seq, f)
	case chan uint8:
		e = eachUint8Channel(seq, f)
	case func(int) uint8:
		e = eachUint8Function(seq, f)
	case func(int) (uint8, bool):
		e = eachUint8BoundedFunction(seq, f)

	case []uint16:
		e = eachUint16Slice(seq, f)
	case map[int]uint16:
		e = eachUint16Map(seq, f)
	case chan uint16:
		e = eachUint16Channel(seq, f)
	case func(int) uint16:
		e = eachUint16Function(seq, f)
	case func(int) (uint16, bool):
		e = eachUint16BoundedFunction(seq, f)

	case []uint32:
		e = eachUint32Slice(seq, f)
	case map[int]uint32:
		e = eachUint32Map(seq, f)
	case chan uint32:
		e = eachUint32Channel(seq, f)
	case func(int) uint32:
		e = eachUint32Function(seq, f)
	case func(int) (uint32, bool):
		e = eachUint32BoundedFunction(seq, f)

	case []uint64:
		e = eachUint64Slice(seq, f)
	case map[int]uint64:
		e = eachUint64Map(seq, f)
	case chan uint64:
		e = eachUint64Channel(seq, f)
	case func(int) uint64:
		e = eachUint64Function(seq, f)
	case func(int) (uint64, bool):
		e = eachUint64BoundedFunction(seq, f)

	case []uintptr:
		e = eachUintptrSlice(seq, f)
	case map[int]uintptr:
		e = eachUintptrMap(seq, f)
	case chan uintptr:
		e = eachUintptrChannel(seq, f)
	case func(int) uintptr:
		e = eachUintptrFunction(seq, f)
	case func(int) (uintptr, bool):
		e = eachUintptrBoundedFunction(seq, f)

	case []R.Value:
		e = eachRValueSlice(seq, f)
	case map[int]R.Value:
		e = eachRValueMap(seq, f)
	case chan R.Value:
		e = eachRValueChannel(seq, f)
	case func(int) R.Value:
		e = eachRValueFunction(seq, f)
	case func(int) (R.Value, bool):
		e = eachRValueBoundedFunction(seq, f)

	case R.Value:
		switch seq.Kind() {
		case R.Slice:
			e = eachSlice(seq, f)
		case R.Chan:
			e = eachChannel(seq, f)
		case R.Map:
			e = eachMap(seq, f)
		case R.Func:
			e = eachFunction(seq, f)
		default:
			e = Each(R.ValueOf(seq), f)
		}
	default:
		e = Each(R.ValueOf(seq), f)
	}
	return
}
