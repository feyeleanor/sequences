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
	Step(increment int, function interface{}) int
}

type Iterable interface {
	Each(function interface{}) int
}

func Step(container interface{}, increment int, f interface{}) (i int) {
	switch container := container.(type) {
	case Enumerable:		i = container.Step(increment, f)
	case []bool:			i = stepBool(container, increment, f)
	case []complex64:		i = stepComplex64(container, increment, f)
	case []complex128:		i = stepComplex128(container, increment, f)
	case []error:			i = stepError(container, increment, f)
	case []float32:			i = stepFloat32(container, increment, f)
	case []float64:			i = stepFloat64(container, increment, f)
	case []int:				i = stepInt(container, increment, f)
	case []int8:			i = stepInt8(container, increment, f)
	case []int16:			i = stepInt16(container, increment, f)
	case []int32:			i = stepInt32(container, increment, f)
	case []int64:			i = stepInt64(container, increment, f)
	case []interface{}:		i = stepInterface(container, increment, f)
	case []string:			i = stepString(container, increment, f)
	case []uint:			i = stepUint(container, increment, f)
	case []uint8:			i = stepUint8(container, increment, f)
	case []uint16:			i = stepUint16(container, increment, f)
	case []uint32:			i = stepUint32(container, increment, f)
	case []uint64:			i = stepUint64(container, increment, f)
	case []uintptr:			i = stepUintptr(container, increment, f)
	case []R.Value:			i = stepRValueSlice(container, increment, f)
	case Indexable:			i = stepIndexable(container, increment, f)
	case R.Value:			switch container.Kind() {
							case R.Slice:		i = stepSlice(container, increment, f)
							case R.Chan:		i = stepChannel(container, increment, f)
							case R.Func:		i = stepFunction(container, increment, f)
							}
	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		i = stepSlice(c, increment, f)
							case R.Chan:		i = stepChannel(c, increment, f)
							case R.Func:		i = stepFunction(c, increment, f)
							}
	}
	return
}

func Each(container, f interface{}) (i int) {
	switch container := container.(type) {
	case Iterable:			i = container.Each(f)
	case Enumerable:		i = container.Step(1, f)
	case []bool:			i = stepBool(container, 1, f)
	case []complex64:		i = stepComplex64(container, 1, f)
	case []complex128:		i = stepComplex128(container, 1, f)
	case []error:			i = stepError(container, 1, f)
	case []float32:			i = stepFloat32(container, 1, f)
	case []float64:			i = stepFloat64(container, 1, f)
	case []int:				i = stepInt(container, 1, f)
	case []int8:			i = stepInt8(container, 1, f)
	case []int16:			i = stepInt16(container, 1, f)
	case []int32:			i = stepInt32(container, 1, f)
	case []int64:			i = stepInt64(container, 1, f)
	case []interface{}:		i = stepInterface(container, 1, f)
	case []string:			i = stepString(container, 1, f)
	case []uint:			i = stepUint(container, 1, f)
	case []uint8:			i = stepUint8(container, 1, f)
	case []uint16:			i = stepUint16(container, 1, f)
	case []uint32:			i = stepUint32(container, 1, f)
	case []uint64:			i = stepUint64(container, 1, f)
	case []uintptr:			i = stepUintptr(container, 1, f)
	case []R.Value:			i = stepRValueSlice(container, 1, f)
	case Indexable:			i = stepIndexable(container, 1, f)
	case Mappable:			i = eachMappable(container, f)
	case R.Value:			switch container.Kind() {
							case R.Slice:		i = stepSlice(container, 1, f)
							case R.Map:			i = eachMap(container, f)
							case R.Chan:		i = stepChannel(container, 1, f)
							case R.Func:		i = stepFunction(container, 1, f)
							}
	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		i = stepSlice(c, 1, f)
							case R.Map:			i = eachMap(c, f)
							case R.Chan:		i = stepChannel(c, 1, f)
							case R.Func:		i = stepFunction(c, 1, f)
							}
	}
	return
}

func Cycle(container interface{}, count int, f interface{}) (i int) {
	for ; i < count; i++ { Each(container, f) }
	return
}

func CycleForever(container interface{}, count int, f interface{}) {
	for {
		Each(container, f)
	}
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