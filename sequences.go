package sequences

import R "reflect"

type Measurable interface {
	Len() int
}

type Confined interface {
	Cap() int
}

type Mappable interface {
	StoredAs(key interface{}) interface{}
	Keys() interface{}
}

type Iterable interface {
	Each(function interface{}) bool
}

func Each(container, f interface{}) (ok bool) {
	switch container := container.(type) {
	case Iterable:			ok = container.Each(f)
	case Indexable:			ok = eachIndexable(container, f)
	case Mappable:			ok = eachMappable(container, f)
	case []bool:			ok = eachBool(container, f)
	case []complex64:		ok = eachComplex64(container, f)
	case []complex128:		ok = eachComplex128(container, f)
	case []error:			ok = eachError(container, f)
	case []float32:			ok = eachFloat32(container, f)
	case []float64:			ok = eachFloat64(container, f)
	case []int:				ok = eachInt(container, f)
	case []int8:			ok = eachInt8(container, f)
	case []int16:			ok = eachInt16(container, f)
	case []int32:			ok = eachInt32(container, f)
	case []int64:			ok = eachInt64(container, f)
	case []interface{}:		ok = eachInterface(container, f)
	case []string:			ok = eachString(container, f)
	case []uint:			ok = eachUint(container, f)
	case []uint8:			ok = eachUint8(container, f)
	case []uint16:			ok = eachUint16(container, f)
	case []uint32:			ok = eachUint32(container, f)
	case []uint64:			ok = eachUint64(container, f)
	case []uintptr:			ok = eachUintptr(container, f)
	case []R.Value:			ok = eachRValueSlice(container, f)
	case R.Value:			switch container.Kind() {
							case R.Slice:		ok = eachSlice(container, f)
							case R.Map:			ok = eachMap(container, f)
							case R.Chan:		ok = eachChannel(container, f)
							case R.Func:		ok = eachFunction(container, f)
							}
	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		ok = eachSlice(c, f)
							case R.Map:			ok = eachMap(c, f)
							case R.Chan:		ok = eachChannel(c, f)
							case R.Func:		ok = eachFunction(c, f)
							}
	}
	return
}

func Cycle(container interface{}, count int, f func(interface{})) (i int, ok bool) {
	for ok = true; i < count && ok; i++ { ok = Each(container, f) }
	return
}

func CycleForever(container interface{}, count int, f func(interface{})) (i int, ok bool) {
	for ok = true; ok; i++ { ok = Each(container, f) }
	return
}

func Count(container interface{}, f func(interface{}) bool) (n int) {
	Each(container, func(x interface{}) {
		if f(x) { n++ }
	})
	return
}


type PartiallyIterable interface {
	While(condition bool, function interface{}) (count int)
}

func While(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyIterable:	count = container.While(true, f)
	case Indexable:			count = whileIndexable(container, true, f)
	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		count = whileSlice(c, true, f)
//							case R.Chan:		count = whileChannel(c, true, f)
//							case R.Func:		count = whileFunction(c, true, f)
							}
	}
	return
}

func Until(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyIterable:	count = container.While(false, f)
	case Indexable:			count = whileIndexable(container, false, f)
	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		count = whileSlice(c, false, f)
//							case R.Chan:		count = whileChannel(c, false, f)
//							case R.Func:		count = whileFunction(c, false, f)
							}
	}
	return
}


type Reducible interface {
	Reduce(seed, function interface{}) (r interface{}, ok bool)
}

func Reduce(container, seed interface{}, f interface{}) (r interface{}, ok bool) {
	switch c := container.(type) {
	case Reducible:				r, ok = c.Reduce(seed, f)
	case Indexable:				r, ok = reduceIndexable(c, seed, f)
	case Mappable:				r, ok = reduceMappable(c, seed, f)
	case Iterable:				r, ok = reduceIterable(c, seed, f)
	default:					switch c := R.ValueOf(container); c.Kind() {
								case R.Invalid:	r, ok = seed, false
								case R.Slice:		r, ok = reduceSlice(c, seed, f)
								case R.Map:			r, ok = reduceMap(c, seed, f)
								case R.Chan:		r, ok = reduceChan(c, seed, f)
								case R.Func:		r, ok = reduceFunction(c, seed, f)
								}
	}
	return
}