package sequences

import R "reflect"

func collectBoolSlice(seq []bool, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(bool) bool:
		s := make([]bool, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, bool) bool:
		s := make([]bool, 0, len(seq))
		for i, v := range seq {
			s = append(r, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectComplex64Slice(seq []complex64, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(complex64) complex64:
		s := make([]complex64, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, complex64) complex64:
		s := make([]complex64, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectComplex128Slice(seq []complex128, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(complex128) complex128:
		s := make([]complex128, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, complex128) complex128:
		s := make([]complex128, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectErrorSlice(seq []error, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(error) error:
		s := make([]error, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, error) error:
		s := make([]error, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFloat32Slice(seq []float32, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(float32) float32:
		s := make([]float32, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, float32) float32:
		s := make([]float32, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectFloat64Slice(seq []float64, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(float64) float64:
		s := make([]float64, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, float64) float64:
		s := make([]float64, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectIntSlice(seq []int, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(int) int:
		s := make([]int, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, int) int:
		s := make([]int, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt8Slice(seq []int8, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(int8) int8:
		s := make([]int8, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, int8) int8:
		s := make([]int8, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt16Slice(seq []int16, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(int16):
		s := make([]int16, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, int16):
		s := make([]int16, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt32Slice(seq []int32, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(int32):
		s := make([]int32, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, int32):
		s := make([]int32, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInt64Slice(seq []int64, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(int64):
		s := make([]int64, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, int64):
		s := make([]int64, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectStringSlice(seq []string, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(string) string:
		s := make([]string, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, string) string:
		s := make([]string, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUintSlice(seq []uint, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(uint) uint:
		s := make([]uint, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, uint) uint:
		s := make([]uint, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint8Slice(seq []uint8, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(uint8) uint8:
		s := make([]uint8, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, uint8) uint8:
		s := make([]uint8, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint16Slice(seq []uint16, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(uint16) uint16:
		s := make([]uint16, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, uint16) uint16:
		s := make([]uint16, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint32Slice(seq []uint32, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(uint32) uint32:
		s := make([]uint32, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, uint32) uint32:
		s := make([]uint32, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUint64Slice(seq []uint64, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(uint64) uint64:
		s := make([]uint64, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, uint64) uint64:
		s := make([]uint64, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectUintptrSlice(seq []uintptr, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(uintptr) uintptr:
		s := make([]uintptr, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, uintptr) uintptr:
		s := make([]uintptr, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectInterfaceSlice(seq []interface{}, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(r, f(R.ValueOf(v)))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, R.ValueOf(v)))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectRValueSlice(seq []R.Value, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v.Interface()))
		}
		r = s
	case func(int, interface{}) interface{}:
		s := make([]interface{}, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v.Interface()))
		}
		r = s

	case func(R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for _, v := range seq {
			s = append(s, f(v))
		}
		r = s
	case func(int, R.Value) R.Value:
		s := make([]R.Value, 0, len(seq))
		for i, v := range seq {
			s = append(s, f(i, v))
		}
		r = s
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}

func collectSlice(s R.Value, f interface{}) (r R.Value, e error) {
	switch f := f.(type) {
	case func(interface{}) interface{}:
		r = MakeSlice(R.TypeOf([]interface{}{}), 0, s.Len())
		eachRValueSliceElement(func(i int, v R.Value) {
			r = Append(r, R.ValueOf(f(v.Interface())))
		}, s.Len())
	case func(R.Value) R.Value:
		r = MakeSlice(s.Type(), 0, s.Len())
		eachRValueSliceElement(func(i int, v R.Value) {
			r = Append(r, f(v))
		}, s.Len())
	case func(int, interface{}) interface{}:
		r = MakeSlice(R.TypeOf([]interface{}{}), 0, s.Len())
		eachRValueSliceElement(func(i int, v R.Value) {
			r = Append(r, f(i, v.Interface()))
		}, s.Len())
	case func(int, R.Value) R.Value:
		r = MakeSlice(s.Type(), 0, s.Len())
		eachRValueSliceElement(func(i int, v R.Value) {
			r = Append(r, f(i, v))
		}, s.Len())
	default:
		e = NO_ENUMERATOR_PROVIDED
	}
	return
}
