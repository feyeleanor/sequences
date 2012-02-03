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

type Enumerable interface {
	Step(increment int, function interface{}) bool
}

type Iterable interface {
	Each(function interface{}) bool
}

func Step(container interface{}, increment int, f interface{}) (ok bool) {
	switch container := container.(type) {
	case Enumerable:		ok = container.Step(increment, f)
	case []bool:			ok = stepBool(container, increment, f)
	case []complex64:		ok = stepComplex64(container, increment, f)
	case []complex128:		ok = stepComplex128(container, increment, f)
	case []error:			ok = stepError(container, increment, f)
	case []float32:			ok = stepFloat32(container, increment, f)
	case []float64:			ok = stepFloat64(container, increment, f)
	case []int:				ok = stepInt(container, increment, f)
	case []int8:			ok = stepInt8(container, increment, f)
	case []int16:			ok = stepInt16(container, increment, f)
	case []int32:			ok = stepInt32(container, increment, f)
	case []int64:			ok = stepInt64(container, increment, f)
	case []interface{}:		ok = stepInterface(container, increment, f)
	case []string:			ok = stepString(container, increment, f)
	case []uint:			ok = stepUint(container, increment, f)
	case []uint8:			ok = stepUint8(container, increment, f)
	case []uint16:			ok = stepUint16(container, increment, f)
	case []uint32:			ok = stepUint32(container, increment, f)
	case []uint64:			ok = stepUint64(container, increment, f)
	case []uintptr:			ok = stepUintptr(container, increment, f)
	case []R.Value:			ok = stepRValueSlice(container, increment, f)
	case Indexable:			ok = stepIndexable(container, increment, f)
	case R.Value:			switch container.Kind() {
							case R.Slice:		ok = stepSlice(container, increment, f)
							case R.Chan:		ok = stepChannel(container, increment, f)
//							case R.Func:		ok = stepFunction(container, increment, f)
							}
//	default:				switch c := R.ValueOf(container); c.Kind() {
//							case R.Slice:		ok = stepSlice(c, increment, f)
//							case R.Chan:		ok = stepChannel(c, increment, f)
//							case R.Func:		ok = stepFunction(c, increment, f)
//							}
	}
	return
}

func Each(container, f interface{}) (ok bool) {
	switch container := container.(type) {
	case Iterable:			ok = container.Each(f)
	case Enumerable:		ok = container.Step(1, f)
	case []bool:			ok = stepBool(container, 1, f)
	case []complex64:		ok = stepComplex64(container, 1, f)
	case []complex128:		ok = stepComplex128(container, 1, f)
	case []error:			ok = stepError(container, 1, f)
	case []float32:			ok = stepFloat32(container, 1, f)
	case []float64:			ok = stepFloat64(container, 1, f)
	case []int:				ok = stepInt(container, 1, f)
	case []int8:			ok = stepInt8(container, 1, f)
	case []int16:			ok = stepInt16(container, 1, f)
	case []int32:			ok = stepInt32(container, 1, f)
	case []int64:			ok = stepInt64(container, 1, f)
	case []interface{}:		ok = stepInterface(container, 1, f)
	case []string:			ok = stepString(container, 1, f)
	case []uint:			ok = stepUint(container, 1, f)
	case []uint8:			ok = stepUint8(container, 1, f)
	case []uint16:			ok = stepUint16(container, 1, f)
	case []uint32:			ok = stepUint32(container, 1, f)
	case []uint64:			ok = stepUint64(container, 1, f)
	case []uintptr:			ok = stepUintptr(container, 1, f)
	case []R.Value:			ok = stepRValueSlice(container, 1, f)
	case Indexable:			ok = stepIndexable(container, 1, f)
	case Mappable:			ok = eachMappable(container, f)
	case R.Value:			switch container.Kind() {
							case R.Slice:		ok = stepSlice(container, 1, f)
							case R.Map:			ok = eachMap(container, f)
							case R.Chan:		ok = stepChannel(container, 1, f)
							case R.Func:		ok = stepFunction(container, 1, f)
							}
	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		ok = stepSlice(c, 1, f)
							case R.Map:			ok = eachMap(c, f)
							case R.Chan:		ok = stepChannel(c, 1, f)
							case R.Func:		ok = stepFunction(c, 1, f)
							}
	}
	return
}

func Cycle(container interface{}, count int, f interface{}) (i int, ok bool) {
	for ok = true; i < count && ok; i++ { ok = Each(container, f) }
	return
}

func CycleForever(container interface{}, count int, f interface{}) (i int, ok bool) {
	for ok = true; ok; i++ { ok = Each(container, f) }
	return
}

//	f() should be an interface{} with a call to a generic function call handler
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