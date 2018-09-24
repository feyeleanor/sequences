package sequences

import R "reflect"

func eachBoolFunction(seq func(int) bool, f interface{}) (e error) {
	switch f := f.(type) {
	case func(bool):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, bool):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan bool:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan bool:
		eachRawIndex(func(i int) {
			writeBoolChannels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex64Function(seq func(int) complex64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex64):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, complex64):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan complex64:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan complex64:
		eachRawIndex(func(i int) {
			writeComplex64Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachComplex128Function(seq func(int) complex128, f interface{}) (e error) {
	switch f := f.(type) {
	case func(complex128):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, complex128):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan complex128:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan complex128:
		eachRawIndex(func(i int) {
			writeComplex128Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachErrorFunction(seq func(int) error, f interface{}) (e error) {
	switch f := f.(type) {
	case func(error):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, error):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan error:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan error:
		eachRawIndex(func(i int) {
			writeErrorChannels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat32Function(seq func(int) float32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(float32):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, float32):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan float32:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan float32:
		eachRawIndex(func(i int) {
			writeFloat32Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFloat64Function(seq func(int) float64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(float64):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, float64):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan float64:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan float64:
		eachRawIndex(func(i int) {
			writeFloat64Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachIntFunction(seq func(int) int, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, int):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan int:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan int:
		eachRawIndex(func(i int) {
			writeIntChannels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt8Function(seq func(int) int8, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int8):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, int8):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan int8:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan int8:
		eachRawIndex(func(i int) {
			writeInt8Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt16Function(seq func(int) int16, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int16):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, int16):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan int16:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan int16:
		eachRawIndex(func(i int) {
			writeInt16Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt32Function(seq func(int) int32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int32):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, int32):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan int32:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan int32:
		eachRawIndex(func(i int) {
			writeInt32Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInt64Function(seq func(int) int64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(int64):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, int64):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan int64:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan int64:
		eachRawIndex(func(i int) {
			writeInt64Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachStringFunction(seq func(int) string, f interface{}) (e error) {
	switch f := f.(type) {
	case func(string):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, string):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan string:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan string:
		eachRawIndex(func(i int) {
			writeStringChannels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintFunction(seq func(int) uint, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, uint):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan uint:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan uint:
		eachRawIndex(func(i int) {
			writeUintChannels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint8Function(seq func(int) uint8, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint8):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, uint8):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan uint8:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan uint8:
		eachRawIndex(func(i int) {
			writeUint8Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint16Function(seq func(int) uint16, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint16):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, uint16):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan uint16:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan uint16:
		eachRawIndex(func(i int) {
			writeUint16Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint32Function(g func(int) uint32, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint32):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, uint32):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan uint32:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan uint32:
		eachRawIndex(func(i int) {
			writeUint32Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUint64Function(g func(int) uint64, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uint64):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, uint64):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan uint64:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan uint64:
		eachRawIndex(func(i int) {
			writeUint64Channels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachUintptrFunction(g func(int) uintptr, f interface{}) (e error) {
	switch f := f.(type) {
	case func(uintptr):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, uintptr):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan uintptr:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan uintptr:
		eachRawIndex(func(i int) {
			writeUintptrChannels(f, seq(i))
		})

	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachInterfaceFunction(seq func(int) interface{}, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i))
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(seq(i)))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(seq(i)))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- R.ValueOf(seq(i))
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachRValueFunction(seq func(int) R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachRawIndex(func(i int) {
			f(seq(i).Interface())
		})
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, seq(i).Interface())
		})
	case chan interface{}:
		eachRawIndex(func(i int) {
			f <- seq(i).Interface()
		})
	case []chan interface{}:
		eachRawIndex(func(i int) {
			writeInterfaceChannels(f, seq(i).Interface())
		})

	case func(R.Value):
		eachRawIndex(func(i int) {
			f(seq(i))
		})
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, seq(i))
		})
	case chan R.Value:
		eachRawIndex(func(i int) {
			f <- seq(i)
		})
	case []chan R.Value:
		eachRawIndex(func(i int) {
			writeRValueChannels(f, seq(i))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func eachFunction(seq R.Value, f interface{}) (e error) {
	switch f := f.(type) {
	case func(interface{}):
		eachIndex(seq, func(v R.Value) {
			f(v.Interface())
		})
	case func(int, interface{}):
		eachIndexWithIndex(seq, func(i int, v R.Value) {
			f(i, v.Interface())
		})
	case chan interface{}:
		eachIndex(seq, func(v R.Value) {
			f <- v.Interface()
		})
	case []chan interface{}:
		eachIndex(seq, func(v R.Value) {
			writeInterfaceChannels(f, v.Interface())
		})

	case func(R.Value):
		eachIndex(seq, func(v R.Value) {
			f(v)
		})
	case func(int, R.Value):
		eachIndexWithIndex(seq, func(i int, v R.Value) {
			f(i, v)
		})
	case chan R.Value:
		eachIndex(seq, func(v R.Value) {
			f <- v
		})
	case []chan R.Value:
		eachIndex(seq, func(v R.Value) {
			writeRValueChannels(f, v)
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
