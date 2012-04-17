package sequences

import R "reflect"

func reduceBoolSlice(enum *Enumerator, seq []bool) (r bool) {
	switch f := enum.f.(type) {
	case func(bool, bool) bool:
		r = enum.seed.(bool)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(bool, int, bool) bool:
		r = enum.seed.(bool)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(bool, interface{}, bool) bool:
		r = enum.seed.(bool)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(bool)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(bool)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(bool)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = rv.Bool()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}


func reduceComplex64Slice(enum *Enumerator, seq []complex64) (r complex64) {
	switch f := enum.f.(type) {
	case func(complex64, complex64) complex64:
		r = enum.seed.(complex64)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(complex64, int, complex64) complex64:
		r = enum.seed.(complex64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(complex64, interface{}, complex64) complex64:
		r = enum.seed.(complex64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(complex64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(complex64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(complex64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = complex64(rv.Complex())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = complex64(rv.Complex())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = complex64(rv.Complex())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = complex64(rv.Complex())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceComplex128Slice(enum *Enumerator, seq []complex128) (r complex128) {
	switch f := enum.f.(type) {
	case func(complex128, complex128) complex128:
		r = enum.seed.(complex128)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(complex128, int, complex128) complex128:
		r = enum.seed.(complex128)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(complex128, interface{}, complex128) complex128:
		r = enum.seed.(complex128)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(complex128)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(complex128)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(complex128)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = complex128(rv.Complex())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = complex128(rv.Complex())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = complex128(rv.Complex())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = complex128(rv.Complex())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceErrorSlice(enum *Enumerator, seq []error) (r error) {
	switch f := enum.f.(type) {
	case func(error, error) error:
		r = enum.seed.(error)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(error, int, error) error:
		r = enum.seed.(error)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(error, interface{}, error) error:
		r = enum.seed.(error)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(error)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(error)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(error)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = rv.Interface().(error)
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Interface().(error)
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Interface().(error)
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = rv.Interface().(error)
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceFloat32Slice(enum *Enumerator, seq []float32) (r float32) {
	switch f := enum.f.(type) {
	case func(float32, float32) float32:
		r = enum.seed.(float32)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(float32, int, float32) float32:
		r = enum.seed.(float32)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(float32, interface{}, float32) float32:
		r = enum.seed.(float32)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(float32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(float32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(float32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = float32(rv.Float())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = float32(rv.Float())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = float32(rv.Float())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = float32(rv.Float())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceFloat64Slice(enum *Enumerator, seq []float64) (r float64) {
	switch f := enum.f.(type) {
	case func(float64, float64) float64:
		r = enum.seed.(float64)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(float64, int, float64) float64:
		r = enum.seed.(float64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(float64, interface{}, float64) float64:
		r = enum.seed.(float64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(float64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(float64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(float64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = float64(rv.Float())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = float64(rv.Float())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = float64(rv.Float())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = float64(rv.Float())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceIntSlice(enum *Enumerator, seq []int) (r int) {
	switch f := enum.f.(type) {
	case func(int, int) int:
		r = enum.seed.(int)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(int, int, int) int:
		r = enum.seed.(int)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int, interface{}, int) int:
		r = enum.seed.(int)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(int)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = int(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = int(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt8Slice(enum *Enumerator, seq []int8) (r int8) {
	switch f := enum.f.(type) {
	case func(int8, int8) int8:
		r = enum.seed.(int8)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(int8, int, int8) int8:
		r = enum.seed.(int8)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int8, interface{}, int8) int8:
		r = enum.seed.(int8)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int, int) int:
		ri := int(enum.seed.(int8))
		enum.reduce(func(cursor int) {
			ri = f(ri, int(seq[cursor]))
		})
		r = int8(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int8))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int8(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int8))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int8(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(int8)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int8)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int8)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = int8(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int8(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int8(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = int8(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt16Slice(enum *Enumerator, seq []int16) (r int16) {
	switch f := enum.f.(type) {
	case func(int16, int16) int16:
		r = enum.seed.(int16)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(int16, int, int16) int16:
		r = enum.seed.(int16)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int16, interface{}, int16) int16:
		r = enum.seed.(int16)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int, int) int:
		ri := int(enum.seed.(int16))
		enum.reduce(func(cursor int) {
			ri = f(ri, int(seq[cursor]))
		})
		r = int16(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int16))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int16(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int16))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int16(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(int16)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int16)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int16)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = int16(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int16(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int16(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = int16(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt32Slice(enum *Enumerator, seq []int32) (r int32) {
	switch f := enum.f.(type) {
	case func(int32, int32) int32:
		r = enum.seed.(int32)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(int32, int, int32) int32:
		r = enum.seed.(int32)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int32, interface{}, int32) int32:
		r = enum.seed.(int32)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int, int) int:
		ri := int(enum.seed.(int32))
		enum.reduce(func(cursor int) {
			ri = f(ri, int(seq[cursor]))
		})
		r = int32(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int32))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int32(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int32))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int32(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(int32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = int32(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int32(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int32(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = int32(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt64Slice(enum *Enumerator, seq []int64) (r int64) {
	switch f := enum.f.(type) {
	case func(int64, int64) int64:
		r = enum.seed.(int64)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(int64, int, int64) int64:
		r = enum.seed.(int64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int64, interface{}, int64) int64:
		r = enum.seed.(int64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(int, int) int:
		ri := int(enum.seed.(int64))
		enum.reduce(func(cursor int) {
			ri = f(ri, int(seq[cursor]))
		})
		r = int64(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int64))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int64(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int64))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, int(seq[cursor]))
		})
		r = int64(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(int64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(int64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = int64(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int64(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = int64(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = int64(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInterfaceSlice(enum *Enumerator, seq []interface{}) (r interface{}) {
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		r = enum.seed
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(interface{}, int, interface{}) interface{}:
		r = enum.seed
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}, interface{}) interface{}:
		r = enum.seed
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = rv.Interface()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Interface()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.Interface()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = rv.Interface()
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			switch f.Type().NumIn() {
			case 2:
				p := make([]R.Value, 2, 2)
				p[0] = R.ValueOf(enum.seed)
				enum.reduce(func(cursor int) {
					p[1] = R.ValueOf(seq[cursor])
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			case 3:
				p := make([]R.Value, 3, 3)
				p[0] = R.ValueOf(enum.seed)
				enum.reduce(func(cursor int) {
					p[1], p[2] = R.ValueOf(cursor), R.ValueOf(seq[cursor])
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceStringSlice(enum *Enumerator, seq []string) (r string) {
	switch f := enum.f.(type) {
	case func(string, string) string:
		r = enum.seed.(string)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(string, int, string) string:
		r = enum.seed.(string)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(string, interface{}, string) string:
		r = enum.seed.(string)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(string)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(string)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(string)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = rv.String()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.String()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = rv.String()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = rv.String()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUintSlice(enum *Enumerator, seq []uint) (r uint) {
	switch f := enum.f.(type) {
	case func(uint, uint) uint:
		r = enum.seed.(uint)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(uint, int, uint) uint:
		r = enum.seed.(uint)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint, interface{}, uint) uint:
		r = enum.seed.(uint)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(uint)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = uint(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = uint(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint8Slice(enum *Enumerator, seq []uint8) (r uint8) {
	switch f := enum.f.(type) {
	case func(uint8, uint8) uint8:
		r = enum.seed.(uint8)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(uint8, int, uint8) uint8:
		r = enum.seed.(uint8)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint8, interface{}, uint8) uint8:
		r = enum.seed.(uint8)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint, uint) uint:
		ri := uint(enum.seed.(uint8))
		enum.reduce(func(cursor int) {
			ri = f(ri, uint(seq[cursor]))
		})
		r = uint8(ri)
	case func(uint, int, uint) uint:
		ri := uint(enum.seed.(uint8))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint8(ri)
	case func(uint, interface{}, uint) uint:
		ri := uint(enum.seed.(uint8))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint8(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(uint8)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint8)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint8)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = uint8(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint8(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint8(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = uint8(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint16Slice(enum *Enumerator, seq []uint16) (r uint16) {
	switch f := enum.f.(type) {
	case func(uint16, uint16) uint16:
		r = enum.seed.(uint16)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(uint16, int, uint16) uint16:
		r = enum.seed.(uint16)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint16, interface{}, uint16) uint16:
		r = enum.seed.(uint16)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint, uint) uint:
		ri := uint(enum.seed.(uint16))
		enum.reduce(func(cursor int) {
			ri = f(ri, uint(seq[cursor]))
		})
		r = uint16(ri)
	case func(uint, int, uint) uint:
		ri := uint(enum.seed.(uint16))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint16(ri)
	case func(uint, interface{}, uint) uint:
		ri := uint(enum.seed.(uint16))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint16(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(uint16)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint16)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint16)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = uint16(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint16(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint16(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = uint16(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint32Slice(enum *Enumerator, seq []uint32) (r uint32) {
	switch f := enum.f.(type) {
	case func(uint32, uint32) uint32:
		r = enum.seed.(uint32)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(uint32, int, uint32) uint32:
		r = enum.seed.(uint32)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint32, interface{}, uint32) uint32:
		r = enum.seed.(uint32)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint, uint) uint:
		ri := uint(enum.seed.(uint32))
		enum.reduce(func(cursor int) {
			ri = f(ri, uint(seq[cursor]))
		})
		r = uint32(ri)
	case func(uint, int, uint) uint:
		ri := uint(enum.seed.(uint32))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint32(ri)
	case func(uint, interface{}, uint) uint:
		ri := uint(enum.seed.(uint32))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint32(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(uint32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = uint32(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint32(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint32(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = uint32(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint64Slice(enum *Enumerator, seq []uint64) (r uint64) {
	switch f := enum.f.(type) {
	case func(uint64, uint64) uint64:
		r = enum.seed.(uint64)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(uint64, int, uint64) uint64:
		r = enum.seed.(uint64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint64, interface{}, uint64) uint64:
		r = enum.seed.(uint64)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint, uint) uint:
		ri := uint(enum.seed.(uint64))
		enum.reduce(func(cursor int) {
			ri = f(ri, uint(seq[cursor]))
		})
		r = uint64(ri)
	case func(uint, int, uint) uint:
		ri := uint(enum.seed.(uint64))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint64(ri)
	case func(uint, interface{}, uint) uint:
		ri := uint(enum.seed.(uint64))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uint64(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(uint64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uint64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = uint64(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint64(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uint64(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = uint64(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUintptrSlice(enum *Enumerator, seq []uintptr) (r uintptr) {
	switch f := enum.f.(type) {
	case func(uintptr, uintptr) uintptr:
		r = enum.seed.(uintptr)
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(uintptr, int, uintptr) uintptr:
		r = enum.seed.(uintptr)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uintptr, interface{}, uintptr) uintptr:
		r = enum.seed.(uintptr)
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(uint, uint) uint:
		ri := uint(enum.seed.(uintptr))
		enum.reduce(func(cursor int) {
			ri = f(ri, uint(seq[cursor]))
		})
		r = uintptr(ri)
	case func(uint, int, uint) uint:
		ri := uint(enum.seed.(uintptr))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uintptr(ri)
	case func(uint, interface{}, uint) uint:
		ri := uint(enum.seed.(uintptr))
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, uint(seq[cursor]))
		})
		r = uintptr(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor])
		})
		r = ri.(uintptr)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uintptr)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor])
		})
		r = ri.(uintptr)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
 		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(seq[cursor]))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, cursor, R.ValueOf(seq[cursor]))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		enum.reduce(func(cursor int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(seq[cursor]))
		})
		r = uintptr(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceRValueSlice(enum *Enumerator, seq []R.Value) (r R.Value) {
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		var ri	interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		enum.reduce(func(cursor int) {
			ri = f(ri, seq[cursor].Interface())
		})
		r = R.ValueOf(ri)
	case func(interface{}, int, interface{}) interface{}:
		var ri	interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor].Interface())
		})
		r = R.ValueOf(ri)
	case func(interface{}, interface{}, interface{}) interface{}:
		var ri	interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		enum.reduce(func(cursor int) {
			ri = f(ri, cursor, seq[cursor].Interface())
		})
		r = R.ValueOf(ri)
	case func(R.Value, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			r = f(r, seq[cursor])
		})
	case func(R.Value, int, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(R.Value, interface{}, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			r = f(r, cursor, seq[cursor])
		})
	case func(R.Value, R.Value, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		enum.reduce(func(cursor int) {
			r = f(r, R.ValueOf(cursor), seq[cursor])
		})
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			switch f.Type().NumIn() {
			case 2:
				p := make([]R.Value, 2, 2)
				if rv, ok := enum.seed.(R.Value); ok {
					p[0] = rv
				} else {
					p[0] = R.ValueOf(enum.seed)
				}
				enum.reduce(func(cursor int) {
					p[1] = R.ValueOf(seq[cursor])
					p[0] = f.Call(p)[0]
				})
				r = p[0]
			case 3:
				p := make([]R.Value, 3, 3)
				if rv, ok := enum.seed.(R.Value); ok {
					p[0] = rv
				} else {
					p[0] = R.ValueOf(enum.seed)
				}
				enum.reduce(func(cursor int) {
					p[1], p[2] = R.ValueOf(cursor), R.ValueOf(seq[cursor])
					p[0] = f.Call(p)[0]
				})
				r = p[0]
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

/*
func reduceSlice(enum *Enumerator, s R.Value) (i int) {
	switch f := enum.f.(type) {
	case func(interface{}):
		i = enum.reduce(func(cursor int) {
			f(s.Index(cursor).Interface())
		})
	case func(interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(s.Index(cursor).Interface())
		})
	case func(int, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor).Interface())
		})
	case func(int, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor).Interface())
		})
	case func(interface{}, interface{}):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor).Interface())
		})
	case func(interface{}, interface{}) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor).Interface())
		})
	case func(R.Value):
		i = enum.reduce(func(cursor int) {
			f(s.Index(cursor))
		})
	case func(R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(s.Index(cursor))
		})
	case func(int, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor))
		})
	case func(int, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor))
		})
	case func(interface{}, R.Value):
		i = enum.reduce(func(cursor int) {
			f(cursor, s.Index(cursor))
		})
	case func(interface{}, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(cursor, s.Index(cursor))
		})
	case func(R.Value, R.Value):
		i = enum.reduce(func(cursor int) {
			f(R.ValueOf(cursor), s.Index(cursor))
		})
	case func(R.Value, R.Value) bool:
		i = enum.reduce(func(cursor int) bool {
			return f(R.ValueOf(cursor), s.Index(cursor))
		})
	default:
		switch f := R.ValueOf(f); f.Kind() {
		case R.Func:
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 1:				//	f(v)
					if t.NumOut() == 0 {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) {
							p[0] = s.Index(cursor)
							f.Call(p)
						})
					} else {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) bool {
							p[0] = s.Index(cursor)
							return f.Call(p)[0].Bool()
						})
					}
				case 2:				//	f(i, v)
					if t.NumOut() == 0 {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) {
							p[0], p[1] = R.ValueOf(cursor), s.Index(cursor)
							f.Call(p)
						})
					} else {
						p := []R.Value{}
						i = enum.reduce(func(cursor int) bool {
							p[0], p[1] = R.ValueOf(cursor), s.Index(cursor)
							return f.Call(p)[0].Bool()
						})
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		case R.Chan:
			i = enum.reduce(func(cursor int) {
				f.Send(s.Index(cursor))
			})
		case R.Slice:
			if f.Type().Elem().Kind() == R.Chan {
				n := f.Len()
				i = enum.reduce(func(cursor int) {
					for i := n; i > 0; i-- {
						f.Index(i).Send(s.Index(cursor))
					}
				})
			} else {
				panic(UNHANDLED_ITERATOR)
			}
		default:
			panic(UNHANDLED_ITERATOR)
		}
	}
	return
}
*/