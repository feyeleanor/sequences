package sequences

import R "reflect"

type Reducible interface {
	Reduce(seed, f interface{}) interface{}
}

var _REDUCIBLE = R.TypeOf(new(Reducible)).Elem()

func AsRValue(v interface{}) (r R.Value) {
	if x, ok := seed.(R.Value); ok {
		r = v
	} else {
		r = R.ValueOf(seed)
	}
	return
}

func Reduce(seq, seed, f interface{}) (r interface{}, e error) {
	switch f := f.(type) {
	case func(bool, bool) bool:
		rv := seed.(bool)
		e = Each(seq, func(v bool) {
			rv = f(rv, v)
		})
		r = rv

	case func(bool, int, bool) bool:
		rv := seed.(bool)
		e = Each(seq, func(i int, v bool) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(complex64, complex64) complex64:
		rv := seed.(complex64)
		e = Each(seq, func(v complex64) {
			rv = f(rv, v)
		})
		r = rv

	case func(complex64, int, complex64) complex64:
		rv := seed.(complex64)
		e = Each(seq, func(i int, v complex64) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(complex128, complex128) complex128:
		rv := seed.(complex128)
		e = Each(seq, func(v complex128) {
			rv = f(rv, v)
		})
		r = rv

	case func(complex128, int, complex128) complex128:
		rv := seed.(complex128)
		e = Each(seq, func(i int, v complex128) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(error, error) error:
		rv := seed.(error)
		e = Each(seq, func(v error) {
			rv = f(rv, v)
		})
		r = rv

	case func(error, int, error) error:
		rv := seed.(error)
		e = Each(seq, func(i int, v error) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(float32, float32) float32:
		rv := seed.(float32)
		e = Each(seq, func(v float32) {
			rv = f(rv, v)
		})
		r = rv

	case func(float32, int, float32) float32:
		rv := seed.(float32)
		e = Each(seq, func(i int, v float32) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(float64, float64) float64:
		rv := seed.(float64)
		e = Each(seq, func(v float64) {
			rv = f(rv, v)
		})
		r = rv

	case func(float64, int, float64) float64:
		rv := seed.(float64)
		e = Each(seq, func(i int, v float64) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(int, int) int:
		rv := seed.(int)
		e = Each(seq, func(v int) {
			rv = f(rv, v)
		})
		r = rv

	case func(int, int, int) int:
		rv := seed.(int)
		e = Each(seq, func(i, v int) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(int8, int8) int8:
		rv := seed.(int8)
		e = Each(seq, func(v int8) {
			rv = f(rv, v)
		})
		r = rv

	case func(int8, int, int8) int8:
		rv := seed.(int8)
		e = Each(seq, func(i int, v int8) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(int16, int16) int16:
		rv := seed.(int16)
		e = Each(seq, func(v int16) {
			rv = f(rv, v)
		})
		r = rv

	case func(int16, int, int16) int16:
		rv := seed.(int16)
		e = Each(seq, func(i int, v int16) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(int32, int32) int32:
		rv := seed.(int32)
		e = Each(seq, func(v int32) {
			rv = f(rv, v)
		})
		r = rv

	case func(int32, int, int32) int32:
		rv := seed.(int32)
		e = Each(seq, func(i int, v int32) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(int64, int64) int64:
		rv := seed.(int64)
		e = Each(seq, func(v int64) {
			rv = f(rv, v)
		})
		r = rv

	case func(int64, int, int64) int64:
		rv := seed.(int64)
		e = Each(seq, func(i int, v int64) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(string, string) string:
		rv := seed.(string)
		e = Each(seq, func(v string) {
			rv = f(rv, v)
		})
		r = rv

	case func(string, int, string) string:
		rv := seed.(string)
		e = Each(seq, func(i int, v string) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(uint, uint) uint:
		rv := seed.(uint)
		e = Each(seq, func(v uint) {
			rv = f(rv, v)
		})
		r = rv

	case func(uint, int, uint) uint:
		rv := seed.(uint)
		e = Each(seq, func(i int, v uint) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(uint8, uint8) uint8:
		rv := seed.(uint8)
		e = Each(seq, func(v uint8) {
			rv = f(rv, v)
		})
		r = rv

	case func(uint8, int, uint8) uint8:
		rv := seed.(uint8)
		e = Each(seq, func(i int, v uint8) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(uint16, uint16) uint16:
		rv := seed.(uint16)
		e = Each(seq, func(v uint16) {
			rv = f(rv, v)
		})
		r = rv

	case func(uint16, int, uint16) uint16:
		rv := seed.(uint16)
		e = Each(seq, func(i int, v uint16) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(uint32, uint32) uint32:
		rv := seed.(uint32)
		e = Each(seq, func(v uint32) {
			rv = f(rv, v)
		})
		r = rv

	case func(uint32, int, uint32) uint32:
		rv := seed.(uint32)
		e = Each(seq, func(i int, v uint32) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(uint64, uint64) uint64:
		rv := seed.(uint64)
		e = Each(seq, func(v uint64) {
			rv = f(rv, v)
		})
		r = rv

	case func(uint64, int, uint64) uint64:
		rv := seed.(uint64)
		e = Each(seq, func(i int, v uint64) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(uintptr, uintptr) uintptr:
		rv := seed.(uintptr)
		e = Each(seq, func(v uintptr) {
			rv = f(rv, v)
		})
		r = rv

	case func(uintptr, int, uintptr) uintptr:
		rv := seed.(uintptr)
		e = Each(seq, func(i int, v uintptr) {
			rv = f(rv, i, v)
		})
		r = rv

	case func(R.Value, R.Value) R.Value:
		rv := AsRValue(seed)
		e = Each(seq, func(v R.Value) {
			rv = f(rv, v)
		})
		r = rv.Interface()

	case func(R.Value, int, R.Value) R.Value:
		rv := AsRValue(seed)
		e = Each(seq, func(i int, v R.Value) {
			rv = f(rv, i, v)
		})
		r = rv.Interface()

	case func(interface{}, int, interface{}) interface{}:
		r = seed
		e = Each(seq, func(i int, v interface{}) {
			r = f(r, i, v)
		})

	case func(interface{}, interface{}) interface{}:
		r = seed
		e = Each(seq, func(v interface{}) {
			r = f(r, v)
		})

	default:
		e = NO_ENUMERATOR_PROVIDED
	}

	if e != nil {
		r = nil
	}
	return
}
