package sequences

import R "reflect"

func eachBoolSlice(seq []bool, f interface{}) (e error) {
	switch f := f.(type) {
	case func(bool):
		for _, v := range seq {
			f(v)
		}
	case func(int, bool):
		for i, v := range seq {
			f(i, v)
		}
	case chan bool:
		for _, v := range seq {
			f <- v
		}
	case []chan bool:
		for _, v := range seq {
			writeBoolChannels(v)
		}

	case func(interface{}):
		for _, v := range seq {
			f(v)
		}
	case func(int, interface{}):
		for i, v := range seq {
			f(i, v)
		}
	case chan interface{}:
		for _, v := range seq {
			f <- v
		}
	case []chan interface{}:
		for _, v := range seq {
			writeInterfaceChannels(v)
		}

	case func(R.Value):
		for _, v := range seq {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for i, v := range seq {
			f(i, R.ValueOf(v))
		}
	case chan R.Value:
		for _, v := range seq {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for _, v := range seq {
			writeRValueChannels(v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex64Slice(seq []complex64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex64):
		for _, v := range seq {
			f(v)
		}
	case func(int, complex64):
		for i, v := range seq {
			f(i, v)
		}
	case chan complex64:
		for _, v := range seq {
			f <- v
		}
	case []chan complex64:
		for _, v := range seq {
			writeComplex64Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex128Slice(seq []complex128, f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex128):
		for _, v := range seq {
			f(v)
		}
	case func(int, complex128):
		for i, v := range seq {
			f(i, v)
		}
	case chan complex128:
		for _, v := range seq {
			f <- v
		}
	case []chan complex128:
		for _, v := range seq {
			writeComplex128Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachErrorSlice(seq []error, f interface{}) (e error) {
	switch f := f.(type) {
	case func(error):
		for _, v := range seq {
			f(v)
		}
	case func(int, error):
		for i, v := range seq {
			f(i, v)
		}
	case chan error:
		for _, v := range seq {
			f <- v
		}
	case []chan error:
		for _, v := range seq {
			writeErrorChannels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat32Slice(seq []float32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(float32):
		for _, v := range seq {
			f(v)
		}
	case func(int, float32):
		for i, v := range seq {
			f(i, v)
		}
	case chan float32:
		for _, v := range seq {
			f <- v
		}
	case []chan float32:
		for _, v := range seq {
			writeFloat32Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat64Slice(seq []float64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(float64):
		for _, v := range seq {
			f(v)
		}
	case func(int, float64):
		for i, v := range seq {
			f(i, v)
		}
	case chan float64:
		for _, v := range seq {
			f <- v
		}
	case []chan float64:
		for _, v := range seq {
			writeFloat64Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachIntSlice(seq []int, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int):
		for _, v := range seq {
			f(v)
		}
	case func(int, int):
		for i, v := range seq {
			f(i, v)
		}
	case chan int:
		for _, v := range seq {
			f <- v
		}
	case []chan int:
		for _, v := range seq {
			writeIntChannels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt8Slice(seq []int8, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int8):
		for _, v := range seq {
			f(v)
		}
	case func(int, int8):
		for i, v := range seq {
			f(i, v)
		}
	case chan int8:
		for _, v := range seq {
			f <- v
		}
	case []chan int8:
		for _, v := range seq {
			writeInt8Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt16Slice(seq []int16, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int16):
		for _, v := range seq {
			f(v)
		}
	case func(int, int16):
		for i, v := range seq {
			f(i, v)
		}
	case chan int16:
		for _, v := range seq {
			f <- v
		}
	case []chan int16:
		for _, v := range seq {
			writeInt16Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt32Slice(seq []int32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int32):
		for _, v := range seq {
			f(v)
		}
	case func(int, int32):
		for i, v := range seq {
			f(i, v)
		}
	case chan int32:
		for _, v := range seq {
			f <- v
		}
	case []chan int32:
		for _, v := range seq {
			writeInt32Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt64Slice(seq []int64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int64):
		for _, v := range seq {
			f(v)
		}
	case func(int, int64):
		for i, v := range seq {
			f(i, v)
		}
	case chan int64:
		for _, v := range seq {
			f <- v
		}
	case []chan int64:
		for _, v := range seq {
			writeInt64Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInterfaceSlice(seq []interface{}, f interface{}) (e error) {
	switch f := f.(type) {
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachStringSlice(seq []string, f interface{}) (e error) {
	switch f := f.(type) {
	case func(string):
		for _, v := range seq {
			f(v)
		}
	case func(int, string):
		for i, v := range seq {
			f(i, v)
		}
	case chan string:
		for _, v := range seq {
			f <- v
		}
	case []chan string:
		for _, v := range seq {
			writeStringChannels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintSlice(seq []uint, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint):
		for _, v := range seq {
			f(v)
		}
	case func(int, uint):
		for i, v := range seq {
			f(i, v)
		}
	case chan uint:
		for _, v := range seq {
			f <- v
		}
	case []chan uint:
		for _, v := range seq {
			writeUintChannels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint8Slice(seq []uint8, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint8):
		for _, v := range seq {
			f(v)
		}
	case func(int, uint8):
		for i, v := range seq {
			f(i, v)
		}
	case chan uint8:
		for _, v := range seq {
			f <- v
		}
	case []chan uint8:
		for _, v := range seq {
			writeUint8Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint16Slice(seq []uint16, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint16):
		for _, v := range seq {
			f(v)
		}
	case func(int, uint16):
		for i, v := range seq {
			f(i, v)
		}
	case chan uint16:
		for _, v := range seq {
			f <- v
		}
	case []chan uint16:
		for _, v := range seq {
			writeUint16Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint32Slice(seq []uint32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint32):
		for _, v := range seq {
			f(v)
		}
	case func(int, uint32):
		for i, v := range seq {
			f(i, v)
		}
	case chan uint32:
		for _, v := range seq {
			f <- v
		}
	case []chan uint32:
		for _, v := range seq {
			writeUint32Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint64Slice(seq []uint64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint64):
		for _, v := range seq {
			f(v)
		}
	case func(int, uint64):
		for i, v := range seq {
			f(i, v)
		}
	case chan uint64:
		for _, v := range seq {
			f <- v
		}
	case []chan uint64:
		for _, v := range seq {
			writeUint64Channels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintptrSlice(seq []uintptr, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uintptr):
		for _, v := range seq {
			f(v)
		}
	case func(int, uintptr):
		for i, v := range seq {
			f(i, v)
		}
	case chan uintptr:
		for _, v := range seq {
			f <- v
		}
	case []chan uintptr:
		for _, v := range seq {
			writeUintptrChannels(v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachRValueSlice(seq []R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		for _, v := range seq {
			f(v.Interface())
		}
	case func(int, interface{}):
		for i, v := range seq {
			f(i, v.Interface())
		}
	case chan interface{}:
		for _, v := range seq {
			f <- v.Interface()
		}
	case []chan interface{}:
		for _, v := range seq {
			writeInterfaceChannels(v.Interface())
		}

	case func(R.Value):
		for _, v := range seq {
			f(v)
		}
	case func(int, R.Value):
		for i, v := range seq {
			f(i, v)
		}
	case chan R.Value:
		for _, v := range seq {
			f <- v
		}
	case []chan R.Value:
		for _, v := range seq {
			writeRValueChannels(v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachSlice(s R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachRValueSliceElement(func(i int, v R.Value) {
			f(v.Interface())
		}, s.Len())
	case func(int, interface{}):
		eachRValueSliceElement(func(i int, v R.Value) {
			f(i, v.Interface())
		}, s.Len())
	case chan interface{}:
		eachRValueSliceElement(func(i int, v R.Value) {
			f <- v.Interface()
		}, s.Len())
	case []chan interface{}:
		eachRValueSliceElement(func(i int, v R.Value) {
			writeInterfaceChannels(v.Interface())
		}, s.Len())

	case func(R.Value):
		eachRValueSliceElement(func(i int, v R.Value) {
			f(v)
		}, s.Len())
	case func(int, R.Value):
		eachRValueSliceElement(func(i int, v R.Value) {
			f(i, v)
		}, s.Len())
	case chan R.Value:
		eachRValueSliceElement(func(i int, v R.Value) {
			f <- v
		}, s.Len())
	case []chan R.Value:
		eachRValueSliceElement(func(i int, v R.Value) {
			writeRValueChannels(v)
		}, s.Len())
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
