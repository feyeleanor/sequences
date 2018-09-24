package sequences

import (
	R "reflect"
	"sort"
)

func eachBoolMap(seq map[int]bool, f interface{}) (e error) {
	switch f := f.(type) {
	case func(bool):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, bool):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan bool:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan bool:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeBoolChannels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex64Map(seq map[int]complex64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, complex64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan complex64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan complex64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeComplex64Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex128Map(seq map[int]complex128, f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex128):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, complex128):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan complex128:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan complex128:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeComplex128Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachErrorMap(seq map[int]error, f interface{}) (e error) {
	switch f := f.(type) {
	case func(error):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, error):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan error:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan error:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeErrorChannels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat32Map(seq map[int]float32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(float32):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, float32):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan float32:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan float32:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeFloat32Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat64Map(seq map[int]float64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(float64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, float64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan float64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan float64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeFloat64Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachIntMap(seq map[int]int, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, int):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan int:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan int:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeIntChannels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt8Map(seq map[int]int8, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int8):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, int8):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan int8:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan int8:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInt8Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt16Map(seq map[int]int16, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int16):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, int16):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan int16:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan int16:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInt16Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt32Map(seq map[int]int32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int32):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, int32):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan int32:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan int32:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInt32Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt64Map(seq map[int]int64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, int64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan int64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan int64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInt64Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachStringMap(seq map[int]string, f interface{}) (e error) {
	switch f := f.(type) {
	case func(string):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, string):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan string:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan string:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeStringChannels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintMap(seq map[int]uint, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, uint):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan uint:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan uint:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeUintChannels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint8Map(seq map[int]uint8, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint8):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, uint8):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan uint8:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan uint8:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeUint8Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint16Map(seq map[int]uint16, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint16):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, uint16):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan uint16:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan uint16:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeUint16Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint32Map(seq map[int]uint32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint32):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, uint32):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan uint32:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan uint32:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeUint32Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint64Map(seq map[int]uint64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, uint64):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan uint64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan uint64:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeUint64Channels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintptrMap(seq map[int]uintptr, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uintptr):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, uintptr):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan uintptr:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan uintptr:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeUintptrChannels(f, v)
			return ok
		}, len(seq))

	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInterfaceMap(seq map[int]interface{}, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v)
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(R.ValueOf(v))
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, R.ValueOf(v))
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- R.ValueOf(v)
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachRValueMap(seq map[int]R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v.Interface())
			return ok
		}, len(seq))
	case func(int, interface{}):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v.Interface())
			return ok
		}, len(seq))
	case chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v.Interface()
			return ok
		}, len(seq))
	case []chan interface{}:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeInterfaceChannels(f, v.Interface())
			return ok
		}, len(seq))

	case func(R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(v)
			return ok
		}, len(seq))
	case func(int, R.Value):
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f(i, v)
			return ok
		}, len(seq))
	case chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			f <- v
			return ok
		}, len(seq))
	case []chan R.Value:
		eachMapElement(func(i int) bool {
			v, ok := seq[i]
			writeRValueChannels(f, v)
			return ok
		}, len(seq))
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachMap(seq R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachRValueMapElement(func(i int, v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		eachRValueMapElement(func(i int, v R.Value) {
			f(i, v.Interface())
		})
	case chan interface{}:
		eachRValueMapElement(func(i int, v R.Value) {
			f <- v.Interface()
		})
	case []chan interface{}:
		eachRValueMapElement(func(i int, v R.Value) {
			writeInterfaceChannels(v.Interface())
		})

	case func(R.Value):
		eachRValueMapElement(func(i int, v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		eachRValueMapElement(func(i int, v R.Value) {
			f(i, v)
		})
	case chan R.Value:
		eachRValueMapElement(func(i int, v R.Value) {
			f <- v
		})
	case []chan R.Value:
		eachRValueMapElement(func(i int, v R.Value) {
			writeRValueChannels(v)
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
