package sequences

import R "reflect"

func eachBoolBoundedFunction(seq func(int) (bool, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(bool):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, bool):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan bool:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan bool:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeBoolChannels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex64BoundedFunction(seq func(int) (complex64, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, complex64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan complex64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan complex64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeComplex64Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex128BoundedFunction(seq func(int) (complex128, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex128):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, complex128):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan complex128:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan complex128:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeComplex128Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachErrorBoundedFunction(seq func(int) (error, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(error):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, error):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan error:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan error:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeErrorChannels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat32BoundedFunction(seq func(int) (float32, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(float32):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, float32):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan float32:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan float32:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeFloat32Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat64BoundedFunction(seq func(int) (float64, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(float64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, float64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan float64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan float64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeFloat64Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachIntBoundedFunction(seq func(int) (int, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(int):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, int):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan int:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan int:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeIntChannels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt8BoundedFunction(seq func(int) (int8, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(int8):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, int8):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan int8:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan int8:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInt8Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt16BoundedFunction(seq func(int) (int16, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(int16):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, int16):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan int16:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan int16:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInt16Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt32BoundedFunction(seq func(int) (int32, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(int32):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, int32):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan int32:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan int32:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInt32Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt64BoundedFunction(seq func(int) (int64, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(int64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, int64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan int64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan int64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInt64Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachStringBoundedFunction(seq func(int) (string, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(string):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, string):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan string:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan string:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeStringChannels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintBoundedFunction(seq func(int) (uint, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, uint):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan uint:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan uint:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeUintChannels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint8BoundedFunction(seq func(int) (uint8, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint8):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, uint8):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan uint8:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan uint8:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeUint8Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint16BoundedFunction(seq func(int) (uint16, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint16):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, uint16):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan uint16:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan uint16:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeUint16Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint32BoundedFunction(g func(int) (uint32, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint32):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, uint32):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan uint32:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan uint32:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeUint32Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint64BoundedFunction(g func(int) (uint64, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, uint64):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan uint64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan uint64:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeUint64Channels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintptrBoundedFunction(g func(int) (uintptr, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(uintptr):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, uintptr):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan uintptr:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan uintptr:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeUintptrChannels(f, v)
			return ok
		})

	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInterfaceBoundedFunction(seq func(int) (interface{}, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachRValueBoundedFunction(seq func(int) (R.Value, bool), f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

//	COMBINE THIS WITH eachFunction

func eachBoundedFunction(seq R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionInterfaceKeyed(seq, i)
			f(v)
			return ok
		})
	case func(int, interface{}):
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionInterfaceKeyed(seq, i)
			f(i, v)
			return ok
		})
	case chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionInterfaceKeyed(seq, i)
			f <- v
			return ok
		})
	case []chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionKeyed(seq, i)
			writeRValueChannels(f, v)
			return ok
		})

	case func(R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionKeyed(seq, i)
			f(v)
			return ok
		})
	case func(int, R.Value):
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionKeyed(seq, i)
			f(i, v)
			return ok
		})
	case chan R.Value:
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionKeyed(seq, i)
			f <- v
			return ok
		})
	case []chan interface{}:
		eachFunctionElement(func(i int) bool {
			v, ok := readFunctionInterfaceKeyed(seq, i)
			writeInterfaceChannels(f, v)
			return ok
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
