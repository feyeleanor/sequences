package sequences

import R "reflect"

func collectBoolChannel(seq chan bool, f interface{}) (r chan bool, e error) {
	switch f := f.(type) {
	case func(bool) bool:
		r = make(chan bool)
		go func(c chan bool) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, bool) bool:
		r = make(chan bool, 0)
		i := 0
		go func(c chan bool) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectComplex64Channel(seq chan complex64, f interface{}) (r chan complex64, e error) {
	switch f := f.(type) {
	case func(complex64) complex64:
		r = make(chan complex64)
		go func(c chan complex64) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, complex64) complex64:
		r = make(chan complex64, 0)
		i := 0
		go func(c chan complex64) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectComplex128Channel(seq chan complex128, f interface{}) (r chan complex128, e error) {
	switch f := f.(type) {
	case func(complex128) complex128:
		r = make(chan complex128)
		go func(c chan complex128) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, complex128) complex128:
		r = make(chan complex128, 0)
		i := 0
		go func(c chan complex128) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectErrorChannel(seq chan error, f interface{}) (r chan error, e error) {
	switch f := f.(type) {
	case func(error) error:
		r = make(chan error)
		go func(c chan error) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, error) error:
		r = make(chan error, 0)
		i := 0
		go func(c chan error) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFloat32Channel(seq chan float32, f interface{}) (r chan float32, e error) {
	switch f := f.(type) {
	case func(float32) float32:
		r = make(chan float32)
		go func(c chan float32) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, float32) float32:
		r = make(chan float32, 0)
		i := 0
		go func(c chan float32) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFloat64Channel(seq chan float64, f interface{}) (r chan float64, e error) {
	switch f := f.(type) {
	case func(float64) float64:
		r = make(chan float64)
		go func(c chan float64) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, float64) float64:
		r = make(chan float64, 0)
		i := 0
		go func(c chan float64) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectIntChannel(seq chan int, f interface{}) (r chan int, e error) {
	switch f := f.(type) {
	case func(int) int:
		r = make(chan int)
		go func(c chan int) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, int) int:
		r = make(chan int, 0)
		i := 0
		go func(c chan int) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt8Channel(seq chan int8, f interface{}) (r chan int8, e error) {
	switch f := f.(type) {
	case func(int8) int8:
		r = make(chan int8)
		go func(c chan int8) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, int8) int8:
		r = make(chan int8, 0)
		i := 0
		go func(c chan int8) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt16Channel(seq chan int16, f interface{}) (r chan int16, e error) {
	switch f := f.(type) {
	case func(int16) int16:
		r = make(chan int16)
		go func(c chan int16) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, int16):
		r = make(chan int16, 0)
		i := 0
		go func(c chan int16) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt32Channel(seq chan int32, f interface{}) (r chan int32, e error) {
	switch f := f.(type) {
	case func(int32) int32:
		r = make(chan int32)
		go func(c chan int32) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, int32):
		r = make(chan int32, 0)
		i := 0
		go func(c chan int32) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt64Channel(seq chan int64, f interface{}) (r chan int64, e error) {
	switch f := f.(type) {
	case func(int64) int64:
		r = make(chan int64)
		go func(c chan int64) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, int64):
		r = make(chan int64, 0)
		i := 0
		go func(c chan int64) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInterfaceChannel(seq chan interface{}, f interface{}) (r chan interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectStringChannel(seq chan string, f interface{}) (r chan string, e error) {
	switch f := f.(type) {
	case func(string) string:
		r = make(chan string)
		go func(c chan string) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, string) string:
		r = make(chan string, 0)
		i := 0
		go func(c chan string) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUintChannel(seq chan uint, f interface{}) (r chan uint, e error) {
	switch f := f.(type) {
	case func(uint) uint:
		r = make(chan uint)
		go func(c chan uint) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, uint) uint:
		r = make(chan uint, 0)
		i := 0
		go func(c chan uint) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint8Channel(seq chan uint8, f interface{}) (r chan uint8, e error) {
	switch f := f.(type) {
	case func(uint8) uint8:
		r = make(chan uint8)
		go func(c chan uint8) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, uint8) uint8:
		r = make(chan uint8, 0)
		i := 0
		go func(c chan uint8) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint16Channel(seq chan uint16, f interface{}) (r chan uint16, e error) {
	switch f := f.(type) {
	case func(uint16) uint16:
		r = make(chan uint16)
		go func(c chan uint16) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, uint16) uint16:
		r = make(chan uint16, 0)
		i := 0
		go func(c chan uint16) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint32Channel(seq chan uint32, f interface{}) (r chan uint32, e error) {
	switch f := f.(type) {
	case func(uint32) uint32:
		r = make(chan uint32)
		go func(c chan uint32) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, uint32) uint32:
		r = make(chan uint32, 0)
		i := 0
		go func(c chan uint32) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint64Channel(seq chan uint64, f interface{}) (r chan uint64, e error) {
	switch f := f.(type) {
	case func(uint64) uint64:
		r = make(chan uint64)
		go func(c chan uint64) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, uint64) uint64:
		r = make(chan uint64, 0)
		i := 0
		go func(c chan uint64) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUintptrChannel(seq chan uintptr, f interface{}) (r chan uintptr, e error) {
	switch f := f.(type) {
	case func(uintptr) uintptr:
		r = make(chan uintptr)
		go func(c chan uintptr) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v)
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(R.ValueOf(v))
			}
			close(c)
		}(r)
	case func(int, uintptr) uintptr:
		r = make(chan uintptr, 0)
		i := 0
		go func(c chan uintptr) {
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, R.ValueOf(v))
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectRValueChannel(seq chan R.Value, f interface{}) (r chan R.Value, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			for v := range seq {
				c <- f(v.Interface())
			}
			close(c)
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			for v := range seq {
				r <- f(v)
			}
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{}, 0)
		go func(c chan interface{}) {
			i := 0
			for v := range seq {
				c <- f(i, v.Interface())
				i++
			}
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value, 0)
		go func(c chan R.Value) {
			i := 0
			for v := range seq {
				c <- f(i, v)
				i++
			}
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectChannel(seq R.Value, f interface{}) (r R.Value, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRecv(seq, func(v R.Value) {
				c <- f(v.Interface())
			})
			c.Close()
		}(r)
	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c R.Value) {
			eachRecv(seq, func(v R.Value) {
				c <- f(v)
			})
			c.Close()
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			i := 0
			eachRecv(seq, func(v R.Value) {
				c <- f(i, v.Interface())
				i++
			})
			c.Close()
		}(r)
	case func(int, R.Value):
		r = make(chan R.Value)
		go func(c R.Value) {
			i := 0
			eachRecv(seq, func(i int, v R.Value) {
				c <- f(i, v)
				i++
			})
			c.Close()
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
