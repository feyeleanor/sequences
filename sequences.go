package sequences

import R "reflect"

const(
	_MAXINT_ = int(^uint(0) >> 1)
	DEFERRED_COUNT = -1
	INDEX_OUT_OF_RANGE = "runtime error: index out of range"
	NOT_A_SEQUENCE = "sequences: not a valid sequence"
	NOT_AN_ITERATOR = "sequences: not a valid iterator"
	NEGATIVE_SPAN = "sequences: span must be a non-negative number"
	UNHANDLED_ITERATOR = "sequences: type cannot handle iterator"
	ASCENDING_SEQUENCE = "sequences: sequence cannot be iterated in reverse order"
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
	Each(enum *Enumerator) (count int)
}

var(
	_MEASURABLE = R.TypeOf(new(Measurable)).Elem()
	_CONFINED = R.TypeOf(new(Confined)).Elem()
	_MAPPABLE = R.TypeOf(new(Mappable)).Elem()
	_ENUMERABLE = R.TypeOf(new(Enumerable)).Elem()
)

func PanicWithIndex(x int) {
	_ = []int{}[x]
}

func IgnoreIndexOutOfRange() {
	if x := recover(); x != nil {
		if x, ok := x.(error); !ok || x.Error() != INDEX_OUT_OF_RANGE {
			panic(x)
		}
	} 
}

/*
	Determines whether or not its parameter is a function which can be used as a sequence.
*/
func isFunctionSequence(f R.Value) (b bool) {
	defer func() {
		recover()
	}()
	if t := f.Type(); t.NumIn() == 1 && t.NumOut() == 1 {
		switch t.In(0).Kind() {
		case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
			fallthrough
		case R.Uint, R.Uint8, R.Uint16, R.Uint32, R.Uint64:
			fallthrough
		case R.Uintptr:
			b = true
		}
	}
	return
}

func Each(seq, f interface{}) (count int) {
	return EachBy(seq, 1, f)
}

func EachBy(seq interface{}, span int, f interface{}) int {
	enum := &Enumerator{ Sequence: seq, Span: span, Limit: -1 }
	return enum.Each(f)
}

func Cycle(seq interface{}, count int, f interface{}) (i int) {
	enum := &Enumerator{ Sequence: seq, Span: 1, Limit: -1 }
	for ; i < count; i++ {
		enum.Each(f)
	}
	return
}

func CycleForever(container interface{}, f interface{}) {
	for {
		Each(container, f)
	}
}

//	f() should be an interface{} with a call to a generic function call handler
func Count(container interface{}, f func(interface{}) bool) (n, l int) {
	return CountBy(container, 1, f)
//	l = Each(container, func(x interface{}) {
//		if f(x) {
//			n++
//		}
//	})
	return
}

func CountBy(container interface{}, span int, f func(interface{}) bool) (n, l int) {
	enum := &Enumerator{ Sequence: container, Span: span }
	l = enum.Each(func(x interface{}) {
		if f(x) {
			n++
		}
	})
	return
}

type PartiallyEnumerable interface {
	While(condition bool, function interface{}) (count int)
}

func While(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyEnumerable:
		count = container.While(true, f)
	case Indexable:
		count = whileIndexable(container, true, f)
	default:
		switch c := R.ValueOf(container); c.Kind() {
		case R.Slice:
			count = whileSlice(c, true, f)
//		case R.Chan:
//			count = whileChannel(c, true, f)
//		case R.Func:
//			count = whileFunction(c, true, f)
		}
	}
	return
}

func Until(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyEnumerable:
		count = container.While(false, f)
	case Indexable:
		count = whileIndexable(container, false, f)
	default:
		switch c := R.ValueOf(container); c.Kind() {
		case R.Slice:
			count = whileSlice(c, false, f)
//		case R.Chan:
//			count = whileChannel(c, false, f)
//		case R.Func:
//			count = whileFunction(c, false, f)
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
	case Enumerable:
		r, e = reduceEnumerable(c, seed, f)
	default:
		switch c := R.ValueOf(container); c.Kind() {
		case R.Invalid:
			r = seed
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