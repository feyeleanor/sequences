package sequences

import R "reflect"

const(
	_MAXINT_ = int(^uint(0) >> 1)
	DEFERRED_COUNT = -1
)

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
	Step(start, span, steps int, function interface{}) (int, error)
}

type Iterable interface {
	Each(processor interface{}) (int, error)
}

func Step(seq interface{}, start, span, steps int, f interface{}) (i int, e error) {
	enum := &Enumerator{ Sequence: seq, Bounds: Bounds{ Start: start, Span: span, Limit: steps } }
	return enum.Each(f)
}

func Each(seq, f interface{}) (i int, e error) {
	enum := &Enumerator{ Sequence: seq, Bounds: Bounds{ Span: 1, Limit: -1 } }
	return enum.Each(f)
}

func Cycle(seq interface{}, count int, f interface{}) (i int, e error) {
	enum := &Enumerator{ Sequence: seq, Bounds: Bounds{ Span: 1, Limit: -1 } }
	enum.Bind(f)
	for ; i < count && e == nil; i++ {
		_, e = enum.Execute()
	}
	return
}

func CycleForever(container interface{}, count int, f interface{}) {
	for {
		Each(container, f)
	}
}

//	f() should be an interface{} with a call to a generic function call handler
func Count(container interface{}, f func(interface{}) bool) (n int, e error) {
	_, e = Each(container, func(x interface{}) {
		if f(x) {
			n++
		}
	})
	return
}


type PartiallyIterable interface {
	While(condition bool, function interface{}) (count int, e error)
}

func While(container, f interface{}) (count int, e error) {
	switch container := container.(type) {
	case PartiallyIterable:
		count, e = container.While(true, f)
	case Indexable:
		count, e = whileIndexable(container, true, f)
	default:
		switch c := R.ValueOf(container); c.Kind() {
		case R.Slice:
			count, e = whileSlice(c, true, f)
//		case R.Chan:
//			count, e = whileChannel(c, true, f)
//		case R.Func:
//			count, e = whileFunction(c, true, f)
		}
	}
	return
}

func Until(container, f interface{}) (count int, e error) {
	switch container := container.(type) {
	case PartiallyIterable:
		count, e = container.While(false, f)
	case Indexable:
		count, e = whileIndexable(container, false, f)
	default:
		switch c := R.ValueOf(container); c.Kind() {
		case R.Slice:
			count, e = whileSlice(c, false, f)
//		case R.Chan:
//			count, e = whileChannel(c, false, f)
//		case R.Func:
//			count, e = whileFunction(c, false, f)
		}
	}
	return
}


type Reducible interface {
	Reduce(seed, function interface{}) (r interface{}, e error)
}

func Reduce(container, seed interface{}, f interface{}) (r interface{}, e error) {
	switch c := container.(type) {
	case Reducible:
		r, e = c.Reduce(seed, f)
	case Indexable:
		r, e = reduceIndexable(c, seed, f)
	case Mappable:
		r, e = reduceMappable(c, seed, f)
	case Iterable:
		r, e = reduceIterable(c, seed, f)
	default:
		switch c := R.ValueOf(container); c.Kind() {
		case R.Invalid:
			r, e = seed, INVALID_SEQUENCE
		case R.Slice:
			r, e = reduceSlice(c, seed, f)
		case R.Map:
			r, e = reduceMap(c, seed, f)
		case R.Chan:
			r, e = reduceChan(c, seed, f)
		case R.Func:
			r, e = reduceFunction(c, seed, f)
		}
	}
	return
}