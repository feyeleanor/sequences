package sequences

import R "reflect"

func reduceBoolChannel(enum *Enumerator, seq chan bool) (r bool) {
	var cursor int

	forEachReceived := forEachBoolReceived(enum)
	switch f := enum.f.(type) {
	case func(bool, bool) bool:
		r = enum.seed.(bool)
		forEachReceived(func(v bool) {
			r = f(r, v)
		})
	case func(bool, int, bool) bool:
		r = enum.seed.(bool)
		forEachReceived(func(v bool) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(bool, interface{}, bool) bool:
		r = enum.seed.(bool)
		forEachReceived(func(v bool) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v bool) {
			ri = f(ri, v)
		})
		r = ri.(bool)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v bool) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(bool)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v bool) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(bool)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Bool()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Bool()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Bool()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v bool) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Bool()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceComplex64Channel(enum *Enumerator, seq chan complex64) (r complex64) {
	var cursor int

	forEachReceived := forEachComplex64Received(enum)
	switch f := enum.f.(type) {
	case func(complex64, complex64) complex64:
		r = enum.seed.(complex64)
		forEachReceived(func(v complex64) {
			r = f(r, v)
		})
	case func(complex64, int, complex64) complex64:
		r = enum.seed.(complex64)
		forEachReceived(func(v complex64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(complex64, interface{}, complex64) complex64:
		r = enum.seed.(complex64)
		forEachReceived(func(v complex64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex64) {
			ri = f(ri, v)
		})
		r = ri.(complex64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(complex64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(complex64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = complex64(rv.Complex())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = complex64(rv.Complex())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = complex64(rv.Complex())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex64) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = complex64(rv.Complex())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceComplex128Channel(enum *Enumerator, seq chan complex128) (r complex128) {
	var cursor int

	forEachReceived := forEachComplex128Received(enum)
	switch f := enum.f.(type) {
	case func(complex128, complex128) complex128:
		r = enum.seed.(complex128)
		forEachReceived(func(v complex128) {
			r = f(r, v)
		})
	case func(complex128, int, complex128) complex128:
		r = enum.seed.(complex128)
		forEachReceived(func(v complex128) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(complex128, interface{}, complex128) complex128:
		r = enum.seed.(complex128)
		forEachReceived(func(v complex128) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex128) {
			ri = f(ri, v)
		})
		r = ri.(complex128)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex128) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(complex128)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v complex128) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(complex128)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Complex()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Complex()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Complex()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v complex128) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Complex()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceErrorChannel(enum *Enumerator, seq chan error) (r error) {
	var cursor int

	forEachReceived := forEachErrorReceived(enum)
	switch f := enum.f.(type) {
	case func(error, error) error:
		r = enum.seed.(error)
		forEachReceived(func(v error) {
			r = f(r, v)
		})
	case func(error, int, error) error:
		r = enum.seed.(error)
		forEachReceived(func(v error) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(error, interface{}, error) error:
		r = enum.seed.(error)
		forEachReceived(func(v error) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v error) {
			ri = f(ri, v)
		})
		r = ri.(error)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v error) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(error)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v error) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(error)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Interface().(error)
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Interface().(error)
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Interface().(error)
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v error) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Interface().(error)
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceFloat32Channel(enum *Enumerator, seq chan float32) (r float32) {
	var cursor int

	forEachReceived := forEachFloat32Received(enum)
	switch f := enum.f.(type) {
	case func(float32, float32) float32:
		r = enum.seed.(float32)
		forEachReceived(func(v float32) {
			r = f(r, v)
		})
	case func(float32, int, float32) float32:
		r = enum.seed.(float32)
		forEachReceived(func(v float32) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(float32, interface{}, float32) float32:
		r = enum.seed.(float32)
		forEachReceived(func(v float32) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float32) {
			ri = f(ri, v)
		})
		r = ri.(float32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float32) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(float32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float32) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(float32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, R.ValueOf(v))
		})
		r = float32(rv.Float())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = float32(rv.Float())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = float32(rv.Float())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float32) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = float32(rv.Float())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceFloat64Channel(enum *Enumerator, seq chan float64) (r float64) {
	var cursor int

	forEachReceived := forEachFloat64Received(enum)
	switch f := enum.f.(type) {
	case func(float64, float64) float64:
		r = enum.seed.(float64)
		forEachReceived(func(v float64) {
			r = f(r, v)
		})
	case func(float64, int, float64) float64:
		r = enum.seed.(float64)
		forEachReceived(func(v float64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(float64, interface{}, float64) float64:
		r = enum.seed.(float64)
		forEachReceived(func(v float64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float64) {
			ri = f(ri, v)
		})
		r = ri.(float64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(float64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v float64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(float64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Float()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Float()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Float()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v float64) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Float()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceIntChannel(enum *Enumerator, seq chan int) (r int) {
	var cursor int

	forEachReceived := forEachIntReceived(enum)
	switch f := enum.f.(type) {
	case func(int, int) int:
		r = enum.seed.(int)
		forEachReceived(func(v int) {
			r = f(r, v)
		})
	case func(int, int, int) int:
		r = enum.seed.(int)
		forEachReceived(func(v int) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int, interface{}, int) int:
		r = enum.seed.(int)
		forEachReceived(func(v int) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int) {
			ri = f(ri, v)
		})
		r = ri.(int)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = int(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt8Channel(enum *Enumerator, seq chan int8) (r int8) {
	var cursor int

	forEachReceived := forEachInt8Received(enum)
	switch f := enum.f.(type) {
	case func(int8, int8) int8:
		r = enum.seed.(int8)
		forEachReceived(func(v int8) {
			r = f(r, v)
		})
	case func(int8, int, int8) int8:
		r = enum.seed.(int8)
		forEachReceived(func(v int8) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int8, interface{}, int8) int8:
		r = enum.seed.(int8)
		forEachReceived(func(v int8) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int, int) int:
		ri := int(enum.seed.(int8))
		forEachReceived(func(v int8) {
			ri = f(ri, int(v))
		})
		r = int8(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int8))
		forEachReceived(func(v int8) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int8(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int8))
		forEachReceived(func(v int8) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int8(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int8) {
			ri = f(ri, v)
		})
		r = ri.(int8)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int8) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int8)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int8) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int8)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int8(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int8(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int8(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int8) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = int8(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt16Channel(enum *Enumerator, seq chan int16) (r int16) {
	var cursor int

	forEachReceived := forEachInt16Received(enum)
	switch f := enum.f.(type) {
	case func(int16, int16) int16:
		r = enum.seed.(int16)
		forEachReceived(func(v int16) {
			r = f(r, v)
		})
	case func(int16, int, int16) int16:
		r = enum.seed.(int16)
		forEachReceived(func(v int16) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int16, interface{}, int16) int16:
		r = enum.seed.(int16)
		forEachReceived(func(v int16) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int, int) int:
		ri := int(enum.seed.(int16))
		forEachReceived(func(v int16) {
			ri = f(ri, int(v))
		})
		r = int16(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int16))
		forEachReceived(func(v int16) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int16(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int16))
		forEachReceived(func(v int16) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int16(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int16) {
			ri = f(ri, v)
		})
		r = ri.(int16)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int16) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int16)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int16) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int16)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int16(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int16(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int16(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int16) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = int16(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt32Channel(enum *Enumerator, seq chan int32) (r int32) {
	var cursor int

	forEachReceived := forEachInt32Received(enum)
	switch f := enum.f.(type) {
	case func(int32, int32) int32:
		r = enum.seed.(int32)
		forEachReceived(func(v int32) {
			r = f(r, v)
		})
	case func(int32, int, int32) int32:
		r = enum.seed.(int32)
		forEachReceived(func(v int32) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int32, interface{}, int32) int32:
		r = enum.seed.(int32)
		forEachReceived(func(v int32) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int, int) int:
		ri := int(enum.seed.(int32))
		forEachReceived(func(v int32) {
			ri = f(ri, int(v))
		})
		r = int32(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int32))
		forEachReceived(func(v int32) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int32(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int32))
		forEachReceived(func(v int32) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int32(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int32) {
			ri = f(ri, v)
		})
		r = ri.(int32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int32) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int32) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, R.ValueOf(v))
		})
		r = int32(rv.Int())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int32(rv.Int())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = int32(rv.Int())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int32) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = int32(rv.Int())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInt64Channel(enum *Enumerator, seq chan int64) (r int64) {
	var cursor int

	forEachReceived := forEachInt64Received(enum)
	switch f := enum.f.(type) {
	case func(int64, int64) int64:
		r = enum.seed.(int64)
		forEachReceived(func(v int64) {
			r = f(r, v)
		})
	case func(int64, int, int64) int64:
		r = enum.seed.(int64)
		forEachReceived(func(v int64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int64, interface{}, int64) int64:
		r = enum.seed.(int64)
		forEachReceived(func(v int64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(int, int) int:
		ri := int(enum.seed.(int64))
		forEachReceived(func(v int64) {
			ri = f(ri, int(v))
		})
		r = int64(ri)
	case func(int, int, int) int:
		ri := int(enum.seed.(int64))
		forEachReceived(func(v int64) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int64(ri)
	case func(int, interface{}, int) int:
		ri := int(enum.seed.(int64))
		forEachReceived(func(v int64) {
			ri = f(ri, cursor, int(v))
			cursor++
		})
		r = int64(ri)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int64) {
			ri = f(ri, v)
		})
		r = ri.(int64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v int64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(int64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Int()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Int()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Int()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v int64) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Int()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceInterfaceChannel(enum *Enumerator, seq chan interface{}) (r interface{}) {
	var cursor int

	forEachReceived := forEachInterfaceReceived(enum)
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		r = enum.seed
		forEachReceived(func(v interface{}) {
			r = f(r, v)
		})
	case func(interface{}, int, interface{}) interface{}:
		r = enum.seed
		forEachReceived(func(v interface{}) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}, interface{}) interface{}:
		r = enum.seed
		forEachReceived(func(v interface{}) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Interface()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Interface()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Interface()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v interface{}) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Interface()
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			switch f.Type().NumIn() {
			case 2:
				p := make([]R.Value, 2, 2)
				p[0] = R.ValueOf(enum.seed)
				forEachReceived(func(v interface{}) {
					p[1] = R.ValueOf(v)
					p[0] = f.Call(p)[0]
				})
				r = p[0].Interface()
			case 3:
				p := make([]R.Value, 3, 3)
				p[0] = R.ValueOf(enum.seed)
				forEachReceived(func(v interface{}) {
					p[1], p[2] = R.ValueOf(cursor), R.ValueOf(v)
					p[0] = f.Call(p)[0]
					cursor++
				})
				r = p[0].Interface()
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceStringChannel(enum *Enumerator, seq chan string) (r string) {
	var cursor int

	forEachReceived := forEachStringReceived(enum)
	switch f := enum.f.(type) {
	case func(string, string) string:
		r = enum.seed.(string)
		forEachReceived(func(v string) {
			r = f(r, v)
		})
	case func(string, int, string) string:
		r = enum.seed.(string)
		forEachReceived(func(v string) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(string, interface{}, string) string:
		r = enum.seed.(string)
		forEachReceived(func(v string) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v string) {
			ri = f(ri, v)
		})
		r = ri.(string)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v string) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(string)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v string) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(string)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.String()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.String()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.String()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v string) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.String()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUintChannel(enum *Enumerator, seq chan uint) (r uint) {
	var cursor int

	forEachReceived := forEachUintReceived(enum)
	switch f := enum.f.(type) {
	case func(uint, uint) uint:
		r = enum.seed.(uint)
		forEachReceived(func(v uint) {
			r = f(r, v)
		})
	case func(uint, int, uint) uint:
		r = enum.seed.(uint)
		forEachReceived(func(v uint) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint, interface{}, uint) uint:
		r = enum.seed.(uint)
		forEachReceived(func(v uint) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint) {
			ri = f(ri, v)
		})
		r = ri.(uint)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = uint(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint8Channel(enum *Enumerator, seq chan uint8) (r uint8) {
	var cursor int

	forEachReceived := forEachUint8Received(enum)
	switch f := enum.f.(type) {
	case func(uint8, uint8) uint8:
		r = enum.seed.(uint8)
		forEachReceived(func(v uint8) {
			r = f(r, v)
		})
	case func(uint8, int, uint8) uint8:
		r = enum.seed.(uint8)
		forEachReceived(func(v uint8) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint8, interface{}, uint8) uint8:
		r = enum.seed.(uint8)
		forEachReceived(func(v uint8) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint8))
		forEachReceived(func(v uint8) {
			ru = f(ru, uint(v))
		})
		r = uint8(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint8))
		forEachReceived(func(v uint8) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint8(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint8))
		forEachReceived(func(v uint8) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint8(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint8) {
			ri = f(ri, v)
		})
		r = ri.(uint8)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint8) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint8)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint8) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint8)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint8(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint8(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint8(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint8) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = uint8(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint16Channel(enum *Enumerator, seq chan uint16) (r uint16) {
	var cursor int

	forEachReceived := forEachUint16Received(enum)
	switch f := enum.f.(type) {
	case func(uint16, uint16) uint16:
		r = enum.seed.(uint16)
		forEachReceived(func(v uint16) {
			r = f(r, v)
		})
	case func(uint16, int, uint16) uint16:
		r = enum.seed.(uint16)
		forEachReceived(func(v uint16) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint16, interface{}, uint16) uint16:
		r = enum.seed.(uint16)
		forEachReceived(func(v uint16) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint16))
		forEachReceived(func(v uint16) {
			ru = f(ru, uint(v))
		})
		r = uint16(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint16))
		forEachReceived(func(v uint16) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint16(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint16))
		forEachReceived(func(v uint16) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint16(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint16) {
			ri = f(ri, v)
		})
		r = ri.(uint16)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint16) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint16)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint16) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint16)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint16(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint16(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint16(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint16) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = uint16(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint32Channel(enum *Enumerator, seq chan uint32) (r uint32) {
	var cursor int

	forEachReceived := forEachUint32Received(enum)
	switch f := enum.f.(type) {
	case func(uint32, uint32) uint32:
		r = enum.seed.(uint32)
		forEachReceived(func(v uint32) {
			r = f(r, v)
		})
	case func(uint32, int, uint32) uint32:
		r = enum.seed.(uint32)
		forEachReceived(func(v uint32) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint32, interface{}, uint32) uint32:
		r = enum.seed.(uint32)
		forEachReceived(func(v uint32) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint32))
		forEachReceived(func(v uint32) {
			ru = f(ru, uint(v))
		})
		r = uint32(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint32))
		forEachReceived(func(v uint32) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint32(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint32))
		forEachReceived(func(v uint32) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint32(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint32) {
			ri = f(ri, v)
		})
		r = ri.(uint32)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint32) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint32)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint32) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint32)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uint32(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint32(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uint32(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint32) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = uint32(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUint64Channel(enum *Enumerator, seq chan uint64) (r uint64) {
	var cursor int

	forEachReceived := forEachUint64Received(enum)
	switch f := enum.f.(type) {
	case func(uint64, uint64) uint64:
		r = enum.seed.(uint64)
		forEachReceived(func(v uint64) {
			r = f(r, v)
		})
	case func(uint64, int, uint64) uint64:
		r = enum.seed.(uint64)
		forEachReceived(func(v uint64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint64, interface{}, uint64) uint64:
		r = enum.seed.(uint64)
		forEachReceived(func(v uint64) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uint64))
		forEachReceived(func(v uint64) {
			ru = f(ru, uint(v))
		})
		r = uint64(ru)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uint64))
		forEachReceived(func(v uint64) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint64(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uint64))
		forEachReceived(func(v uint64) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uint64(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint64) {
			ri = f(ri, v)
		})
		r = ri.(uint64)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint64)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uint64) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uint64)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, R.ValueOf(v))
		})
		r = rv.Uint()
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Uint()
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = rv.Uint()
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uint64) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = rv.Uint()
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceUintptrChannel(enum *Enumerator, seq chan uintptr) (r uintptr) {
	var cursor int

	forEachReceived := forEachUintptrReceived(enum)
	switch f := enum.f.(type) {
	case func(uintptr, uintptr) uintptr:
		r = enum.seed.(uintptr)
		forEachReceived(func(v uintptr) {
			r = f(r, v)
		})
	case func(uintptr, int, uintptr) uintptr:
		r = enum.seed.(uintptr)
		forEachReceived(func(v uintptr) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uintptr, interface{}, uintptr) uintptr:
		r = enum.seed.(uintptr)
		forEachReceived(func(v uintptr) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(uint, uint) uint:
		ru := uint(enum.seed.(uintptr))
		forEachReceived(func(v uintptr) {
			ru = f(ru, uint(v))
		})
		r = uintptr(r)
	case func(uint, int, uint) uint:
		ru := uint(enum.seed.(uintptr))
		forEachReceived(func(v uintptr) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uintptr(ru)
	case func(uint, interface{}, uint) uint:
		ru := uint(enum.seed.(uintptr))
		forEachReceived(func(v uintptr) {
			ru = f(ru, cursor, uint(v))
			cursor++
		})
		r = uintptr(ru)
	case func(interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uintptr) {
			ri = f(ri, v)
		})
		r = ri.(uintptr)
	case func(interface{}, int, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uintptr) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uintptr)
	case func(interface{}, interface{}, interface{}) interface{}:
		ri := enum.seed
		forEachReceived(func(v uintptr) {
			ri = f(ri, cursor, v)
			cursor++
		})
		r = ri.(uintptr)
	case func(R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, R.ValueOf(v))
		})
		r = uintptr(rv.Uint())
	case func(R.Value, int, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uintptr(rv.Uint())
	case func(R.Value, interface{}, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, cursor, R.ValueOf(v))
			cursor++
		})
		r = uintptr(rv.Uint())
	case func(R.Value, R.Value, R.Value) R.Value:
		rv := R.ValueOf(enum.seed)
		forEachReceived(func(v uintptr) {
			rv = f(rv, R.ValueOf(cursor), R.ValueOf(v))
			cursor++
		})
		r = uintptr(rv.Uint())
	default:
		panic(UNHANDLED_ITERATOR)
	}
	return
}

func reduceRValueChannel(enum *Enumerator, seq chan R.Value) (r R.Value) {
	var cursor int

	forEachReceived := forEachRValueReceived(enum)
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		var ri interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, v.Interface())
		})
		r = R.ValueOf(ri)
	case func(interface{}, int, interface{}) interface{}:
		var ri interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, cursor, v.Interface())
			cursor++
		})
		r = R.ValueOf(ri)
	case func(interface{}, interface{}, interface{}) interface{}:
		var ri interface{}
		if r, ok := enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, cursor, v.Interface())
			cursor++
		})
		r = R.ValueOf(ri)
	case func(R.Value, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, v)
		})
	case func(R.Value, int, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(R.Value, interface{}, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(R.Value, R.Value, R.Value) R.Value:
		if rv, ok := enum.seed.(R.Value); ok {
			r = rv
		} else {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, R.ValueOf(cursor), v)
			cursor++
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
				forEachReceived(func(v R.Value) {
					p[1] = v
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
				forEachReceived(func(v R.Value) {
					p[1], p[2] = R.ValueOf(cursor), v
					p[0] = f.Call(p)[0]
					cursor++
				})
				r = p[0]
			default:
				panic(UNHANDLED_ITERATOR)
			}
		}
	}
	return
}

func reduceChannel(enum *Enumerator, s R.Value) (r R.Value) {
	var cursor int

	forEachReceived := forEachChannelReceived(enum)
	var ok bool
	switch f := enum.f.(type) {
	case func(interface{}, interface{}) interface{}:
		var ri interface{}
		if r, ok = enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, v.Interface())
		})
		r = R.ValueOf(ri)
	case func(interface{}, int, interface{}) interface{}:
		var ri interface{}
		if r, ok = enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, cursor, v.Interface())
			cursor++
		})
		r = R.ValueOf(ri)
	case func(interface{}, interface{}, interface{}) interface{}:
		var ri interface{}
		if r, ok = enum.seed.(R.Value); ok {
			ri = r.Interface()
		} else {
			ri = enum.seed
		}
		forEachReceived(func(v R.Value) {
			ri = f(ri, cursor, v.Interface())
			cursor++
		})
		r = R.ValueOf(ri)
	case func(R.Value, R.Value) R.Value:
		if r, ok = enum.seed.(R.Value); !ok {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, v)
		})
	case func(R.Value, int, R.Value) R.Value:
		if r, ok = enum.seed.(R.Value); !ok {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(R.Value, interface{}, R.Value) R.Value:
		if r, ok = enum.seed.(R.Value); !ok {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, cursor, v)
			cursor++
		})
	case func(R.Value, R.Value, R.Value) R.Value:
		if r, ok = enum.seed.(R.Value); !ok {
			r = R.ValueOf(enum.seed)
		}
		forEachReceived(func(v R.Value) {
			r = f(r, R.ValueOf(cursor), v)
			cursor++
		})
	default:
		switch f := R.ValueOf(f); f.Kind() {
		case R.Func:
			if t := f.Type(); !t.IsVariadic() {
				switch t.NumIn() {
				case 2:
					p := make([]R.Value, 2, 2)
					if p[0], ok = enum.seed.(R.Value); !ok {
						p[0] = R.ValueOf(enum.seed)
					}
					forEachReceived(func(v R.Value) {
						p[1] = v
						p[0] = f.Call(p)[0]
					})
					r = p[0]
				case 3:
					p := make([]R.Value, 3, 3)
					if p[0], ok = enum.seed.(R.Value); !ok {
						p[0] = R.ValueOf(enum.seed)
					}
					forEachReceived(func(v R.Value) {
						p[1], p[2] = R.ValueOf(cursor), v
						p[0] = f.Call(p)[0]
						cursor++
					})
					r = p[0]
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		default:
			panic(UNHANDLED_ITERATOR)
		}
	}
	return
}
