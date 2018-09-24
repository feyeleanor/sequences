package sequences

import R "reflect"

func collectBoolFunction(seq func(int) bool, f interface{}) (r chan bool, e error) {
	switch f := f.(type) {
	case func(bool) bool:
		r = make(chan bool)
		go func(c chan bool) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, bool) bool:
		r = make(chan bool)
		go func(c chan bool) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectComplex64Function(seq func(int) complex64, f interface{}) (r []complex64, e error) {
	switch f := f.(type) {
	case func(complex64) complex64:
		r = make([]complex64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, complex64) complex64:
		r = make([]complex64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectComplex128Function(seq func(int) complex128, f interface{}) (r []complex128, e error) {
	switch f := f.(type) {
	case func(complex128) complex128:
		r = make([]complex128, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, complex128) complex128:
		r = make([]complex128, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectErrorFunction(seq func(int) error, f interface{}) (r []error, e error) {
	switch f := f.(type) {
	case func(error) error:
		r = make([]error, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, error) error:
		r = make([]error, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFloat32Function(seq func(int) float32, f interface{}) (r []float32, e error) {
	switch f := f.(type) {
	case func(float32) float32:
		r = make([]float32, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, float32) float32:
		r = make([]float32, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFloat64Function(seq func(int) float64, f interface{}) (r []float64, e error) {
	switch f := f.(type) {
	case func(float64) float64:
		r = make([]float64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, float64) float64:
		r = make([]float64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectIntFunction(seq func(int) int, f interface{}) (r []int, e error) {
	switch f := f.(type) {
	case func(int) int:
		r = make([]int, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, int) int:
		r = make([]int, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt8Function(seq func(int) int8, f interface{}) (r []int8, e error) {
	switch f := f.(type) {
	case func(int8) int8:
		r = make([]int8, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, int8) int8:
		r = make([]int8, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt16Function(seq func(int) int16, f interface{}) (r []int16, e error) {
	switch f := f.(type) {
	case func(int16) int16:
		r = make([]int16, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, int16) int16:
		r = make([]int16, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt32Function(seq func(int) int32, f interface{}) (r []int32, e error) {
	switch f := f.(type) {
	case func(int32) int32:
		r = make([]int32, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, int32) int32:
		r = make([]int32, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt64Function(seq func(int) int64, f interface{}) (r []int64, e error) {
	switch f := f.(type) {
	case func(int64) int64:
		r = make([]int64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, int64) int64:
		r = make([]int64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectStringFunction(seq func(int) string, f interface{}) (r []string, e error) {
	switch f := f.(type) {
	case func(string) string:
		r = make([]string, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, string) string:
		r = make([]string, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUintFunction(seq func(int) uint, f interface{}) (r []uint, e error) {
	switch f := f.(type) {
	case func(uint) uint:
		r = make([]uint, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, uint) uint:
		r = make([]uint, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint8Function(seq func(int) uint8, f interface{}) (r []uint8, e error) {
	switch f := f.(type) {
	case func(uint8) uint8:
		r = make([]uint8, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, uint8) uint8:
		r = make([]uint8, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint16Function(seq func(int) uint16, f interface{}) (r []uint16, e error) {
	switch f := f.(type) {
	case func(uint16) uint16:
		r = make([]uint16, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, uint16) uint16:
		r = make([]uint16, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint32Function(g func(int) uint32, f interface{}) (r []uint32, e error) {
	switch f := f.(type) {
	case func(uint32) uint32:
		r = make([]uint32, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, uint32) uint32:
		r = make([]uint32, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint64Function(g func(int) uint64, f interface{}) (r []uint64, e error) {
	switch f := f.(type) {
	case func(uint64) uint64:
		r = make([]uint64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, uint64) uint64:
		r = make([]uint64, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUintptrFunction(g func(int) uintptr, f interface{}) (r []uintptr, e error) {
	switch f := f.(type) {
	case func(uintptr) uintptr:
		r = make([]uintptr, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, uintptr) uintptr:
		r = make([]uintptr, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(seq(i))
			})
			close(c)
		}(r)
	case func(int, interface{}) interface{}:
		r = make(chan interface{})
		go func(c chan interface{}) {
			eachRawIndex(func(i int) {
				c <- f(i, seq(i))
			})
			close(c)
		}(r)

	case func(R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	case func(int, R.Value) R.Value:
		r = make(chan R.Value)
		go func(c chan R.Value) {
			eachRawIndex(func(i int) {
				c <- f(i, R.ValueOf(seq(i)))
			})
			close(c)
		}(r)
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInterfaceFunction(seq func(int) interface{}, f interface{}) (r []interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = make([]interface{}, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, interface{}) interface{}:
		r = make([]interface{}, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})

	case func(R.Value) R.Value:
		r = make([]R.Value, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(R.ValueOf(seq(i))))
		})
	case func(int, R.Value) R.Value:
		r = make([]R.Value, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, R.ValueOf(seq(i))))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectRValueFunction(seq func(int) R.Value, f interface{}) (r []R.Value, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = make([]interface{}, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i).Interface()))
		})
	case func(int, interface{}) interface{}:
		r = make([]interface{}, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i).Interface()))
		})

	case func(R.Value) R.Value:
		r = make([]R.Value, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(seq(i)))
		})
	case func(int, R.Value) R.Value:
		r = make([]R.Value, 0)
		eachRawIndex(func(i int) {
			r = append(r, f(i, seq(i)))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFunction(seq R.Value, f interface{}) (r R.Value, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = R.MakeSlice(R.TypeOf([]interface{}{}), 0, 1)
		eachIndex(seq, func(v R.Value) {
			r = R.Append(r, f(v.Interface()))
		})
	case func(int, interface{}):
		r = R.MakeSlice(R.TypeOf([]interface{}{}), 0, 1)
		eachIndexWithIndex(seq, func(i int, v R.Value) {
			r = R.Append(r, f(i, v.Interface()))
		})

	case func(R.Value):
		r = R.MakeSlice(seq.Type(), 0, 1)
		eachIndex(seq, func(v R.Value) {
			r = R.Append(r, f(v))
		})
	case func(int, R.Value):
		r = R.MakeSlice(seq.Type(), 0, 1)
		eachIndexWithIndex(seq, func(i int, v R.Value) {
			r = R.Append(r, f(i, v))
		})
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
