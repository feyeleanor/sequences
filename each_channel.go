package sequences

import R "reflect"

func eachBoolChannel(c chan bool, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(bool):
		for v := range c {
			f(v)
		}
	case func(int, bool):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan bool:
		for v := range c {
			f <- v
		}
	case []chan bool:
		for v := range c {
			writeBoolChannels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex64Channel(c chan complex64, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(complex64):
		for v := range c {
			f(v)
		}
	case func(int, complex64):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan complex64:
		for v := range c {
			f <- v
		}
	case []chan complex64:
		for v := range c {
			writeComplex64Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex128Channel(c chan complex128, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(complex128):
		for v := range c {
			f(v)
		}
	case func(int, complex128):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan complex128:
		for v := range c {
			f <- v
		}
	case []chan complex128:
		for v := range c {
			writeComplex128Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachErrorChannel(c chan error, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(error):
		for v := range c {
			f(v)
		}
	case func(int, error):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan error:
		for v := range c {
			f <- v
		}
	case []chan error:
		for v := range c {
			writeErrorChannels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat32Channel(c chan float32, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(float32):
		for v := range c {
			f(v)
		}
	case func(int, float32):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan float32:
		for v := range c {
			f <- v
		}
	case []chan float32:
		for v := range c {
			writeFloat32Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat64Channel(c chan float64, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(float64):
		for v := range c {
			f(v)
		}
	case func(int, float64):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan float64:
		for v := range c {
			f <- v
		}
	case []chan float64:
		for v := range c {
			writeFloat64Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachIntChannel(c chan int, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(int):
		for v := range c {
			f(v)
		}
	case func(int, int):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan int:
		for v := range c {
			f <- v
		}
	case []chan int:
		for v := range c {
			writeIntChannels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt8Channel(c chan int8, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(int8):
		for v := range c {
			f(v)
		}
	case func(int, int8):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan int8:
		for v := range c {
			f <- v
		}
	case []chan int8:
		for v := range c {
			writeInt8Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt16Channel(c chan int16, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(int16):
		for v := range c {
			f(v)
		}
	case func(int, int16):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan int16:
		for v := range c {
			f <- v
		}
	case []chan int16:
		for v := range c {
			writeInt16Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt32Channel(c chan int32, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(int32):
		for v := range c {
			f(v)
		}
	case func(int, int32):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan int32:
		for v := range c {
			f <- v
		}
	case []chan int32:
		for v := range c {
			writeInt32Channels(f, v)
		}

	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt64Channel(c chan int64, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(int64):
		for v := range c {
			f(v)
		}
	case func(int, int64):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan int64:
		for v := range c {
			f <- v
		}
	case []chan int64:
		for v := range c {
			writeInt64Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachStringChannel(c chan string, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(string):
		for v := range c {
			f(v)
		}
	case func(int, string):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan string:
		for v := range c {
			f <- v
		}
	case []chan string:
		for v := range c {
			writeStringChannels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintChannel(c chan uint, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(uint):
		for v := range c {
			f(v)
		}
	case func(int, uint):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan uint:
		for v := range c {
			f <- v
		}
	case []chan uint:
		for v := range c {
			writeUintChannels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint8Channel(c chan uint8, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(uint8):
		for v := range c {
			f(v)
		}
	case func(int, uint8):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan uint8:
		for v := range c {
			f <- v
		}
	case []chan uint8:
		for v := range c {
			writeUint8Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint16Channel(c chan uint16, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(uint16):
		for v := range c {
			f(v)
		}
	case func(int, uint16):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan uint16:
		for v := range c {
			f <- v
		}
	case []chan uint16:
		for v := range c {
			writeUint16Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint32Channel(c chan uint32, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(uint32):
		for v := range c {
			f(v)
		}
	case func(int, uint32):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan uint32:
		for v := range c {
			f <- v
		}
	case []chan uint32:
		for v := range c {
			writeUint32Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint64Channel(c chan uint64, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(uint64):
		for v := range c {
			f(v)
		}
	case func(int, uint64):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan uint64:
		for v := range c {
			f <- v
		}
	case []chan uint64:
		for v := range c {
			writeUint64Channels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintptrChannel(c chan uintptr, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(uintptr):
		for v := range c {
			f(v)
		}
	case func(int, uintptr):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan uintptr:
		for v := range c {
			f <- v
		}
	case []chan uintptr:
		for v := range c {
			writeUintptrChannels(f, v)
		}

	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInterfaceChannel(c chan interface{}, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(interface{}):
		for v := range c {
			f(v)
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(R.ValueOf(v))
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, R.ValueOf(v))
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- R.ValueOf(v)
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachRValueChannel(c chan R.Value, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(interface{}):
		for v := range c {
			f(v.Interface())
		}
	case func(int, interface{}):
		for v := range c {
			f(cursor, v.Interface())
			cursor++
		}
	case chan interface{}:
		for v := range c {
			f <- v.Interface()
		}
	case []chan interface{}:
		for v := range c {
			writeInterfaceChannels(f, v)
		}

	case func(R.Value):
		for v := range c {
			f(v)
		}
	case func(int, R.Value):
		for v := range c {
			f(cursor, v)
			cursor++
		}
	case chan R.Value:
		for v := range c {
			f <- v
		}
	case []chan R.Value:
		for v := range c {
			writeRValueChannels(f, v)
		}
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachChannel(c R.Value, f interface{}) (e error) {
	var cursor int

	switch f := f.(type) {
	case func(interface{}):
		eachRecv(seq, func(v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		eachRecv(seq, func(v R.Value) {
			f(cursor, v.Interface())
			cursor++
		})
	case chan interface{}:
		eachRecv(seq, func(v R.Value) {
			f <- v.Interface()
		})
	case []chan interface{}:
		eachRecv(seq, func(v R.Value) {
			writeInterfaceChannels(f, v.Interface())
		})

	case func(R.Value):
		eachRecv(seq, func(v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		eachRecv(seq, func(v R.Value) {
			f(cursor, v)
			cursor++
		})
	case chan R.Value:
		eachRecv(seq, func(v R.Value) {
			f <- v
		})
	case []chan R.Value:
		eachRecv(seq, func(v R.Value) {
			writeRValueChannels(f, v)
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
