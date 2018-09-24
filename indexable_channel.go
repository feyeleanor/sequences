package sequences

import R "reflect"

func atOffsetBoolChannel(seq chan bool, i int, f interface{}) {
	switch f := f.(type) {
	case func(bool):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan bool:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan bool:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetComplex64Channel(seq chan complex64, i int, f interface{}) {
	switch f := f.(type) {
	case func(complex64):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan complex64:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan complex64:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetComplex128Channel(seq chan complex128, int i, f interface{}) {
	switch f := f.(type) {
	case func(complex128):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan complex128:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan complex128:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetErrorChannel(seq chan error, i int, f interface{}) {
	switch f := f.(type) {
	case func(error):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan error:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan error:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetFloat32Channel(seq chan float32, i int, f interface{}) {
	switch f := f.(type) {
	case func(float32):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan float32:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan float32:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetFloat64Channel(seq chan float64, i int, f interface{}) {
	switch f := f.(type) {
	case func(float64):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan float64:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan float64:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetIntChannel(seq chan int, i int, f interface{}) {
	switch f := f.(type) {
	case func(int):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan int:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan int:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetInt8Channel(seq chan int8, i int, f interface{}) {
	switch f := f.(type) {
	case func(int8):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan int8:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan int8:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetInt16Channel(seq chan int16, i int, f interface{}) {
	switch f := f.(type) {
	case func(int16):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan int16:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan int16:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetInt32Channel(seq chan int32, i int, f interface{}) {
	switch f := f.(type) {
	case func(int32):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan int32:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan int32:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetInt64Channel(seq chan int64, i int, f interface{}) {
	switch f := f.(type) {
	case func(int64):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan int64:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan int64:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetStringChannel(seq chan string, i int, f interface{}) {
	switch f := f.(type) {
	case func(string):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan string:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan string:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetUintChannel(seq chan uint, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan uint:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan uint:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetUint8Channel(seq chan uint8, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint8):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan uint8:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan uint8:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetUint16Channel(seq chan uint16, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint16):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan uint16:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan uint16:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetUint32Channel(seq chan uint32, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint32):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan uint32:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan uint32:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetUint64Channel(seq chan uint64, i int, f interface{}) {
	switch f := f.(type) {
	case func(uint64):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan uint64:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan uint64:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetUintptrChannel(seq chan uintptr, i int, f interface{}) {
	switch f := f.(type) {
	case func(uintptr):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan uintptr:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan uintptr:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetInterfaceChannel(seq chan interface{}, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}
	case []chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetRValueChannel(seq chan R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		for n, v := range seq {
			if n == i {
				f(v)
				break
			}
		}
	case chan interface{}:
		for n, v := range seq {
			if n == i {
				f <- v
				break
			}
		}
	case []chan interface{}:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- v
				}
				break
			}
		}

	case func(R.Value):
		for n, v := range seq {
			if n == i {
				f(R.ValueOf(v))
				break
			}
		}
	case chan R.Value:
		for n, v := range seq {
			if n == i {
				f <- R.ValueOf(v)
				break
			}
		}

	case chan R.Value:
		for n, v := range seq {
			if n == i {
				for c := range f {
					c <- R.ValueOf(v)
				}
				break
			}
		}
	}
}

func atOffsetChannel(seq R.Value, i int, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):
		for v, ok := seq.Recv(); ok; v, ok = seq.Recv() {
			if i--; i == 0 {
				f(v.Interface())
				break
			}
		}
	case chan interface{}:
		for v, ok := seq.Recv(); ok; v, ok = seq.Recv() {
			if i--; i == 0 {
				f <- v.Interface()
			}
		}
	case []chan interface{}:
		for v, ok := seq.Recv(); ok; v, ok = seq.Recv() {
			if i--; i == 0 {
				for c := range f {
					c <- v.Interface()
				}
			}
		}

	case func(R.Value):
		for v, ok := seq.Recv(); ok; v, ok = seq.Recv() {
			if i--; i == 0 {
				f(v)
				break
			}
		}
	case chan R.Value:
		for v, ok := seq.Recv(); ok; v, ok = seq.Recv() {
			if i--; i == 0 {
				f <- v
				break
			}
		}
	case []chan R.Value:
		for n, v := range c {
			if i--; i == 0 {
				for c := range f {
					c <- v
				}
				break
			}
		}
	}
}
