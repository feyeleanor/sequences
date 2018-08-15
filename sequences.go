package sequences

import R "reflect"

const (
	_MAXINT_           = int(^uint(0) >> 1)
	DEFERRED_COUNT     = -1
	INDEX_OUT_OF_RANGE = "runtime error: index out of range"
	NOT_A_SEQUENCE     = "sequences: not a valid sequence"
	NOT_AN_ITERATOR    = "sequences: not a valid iterator"
	NEGATIVE_SPAN      = "sequences: span must be a non-negative number"
	UNHANDLED_ITERATOR = "sequences: type cannot handle iterator"
	ASCENDING_SEQUENCE = "sequences: sequence cannot be iterated in reverse order"
)

type Measurable interface {
	Len() int
}

var _MEASURABLE = R.TypeOf(new(Measurable)).Elem()

type Confined interface {
	Cap() int
}

var _CONFINED = R.TypeOf(new(Confined)).Elem()

type Mappable interface {
	StoredAs(key interface{}) interface{}
	Keys() interface{}
}

var _MAPPABLE = R.TypeOf(new(Mappable)).Elem()

type Enumerable interface {
	Each(enum *Enumerator) (count int)
}

var _ENUMERABLE = R.TypeOf(new(Enumerable)).Elem()

type Reducible interface {
	Reduce(enum *Enumerator) interface{}
}

var _REDUCIBLE = R.TypeOf(new(Reducible)).Elem()

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

func Each(seq, f interface{}) {
	Enumerator{Sequence: seq, Span: 1}.Each(f)
}

func EachBy(seq interface{}, span int, f interface{}) {
	Enumerator{Sequence: seq, Span: span}.Each(f)
}

func Cycle(seq interface{}, count int, f interface{}) {
	enum := Enumerator{Sequence: seq, Span: 1}
	for i := 0; i < count; i++ {
		enum.Each(f)
	}
}

func CycleForever(container interface{}, f interface{}) {
	for {
		Each(container, f)
	}
}

//	f() should be an interface{} with a call to a generic function call handler
func Count(container interface{}, f func(interface{}) bool) (n int) {
	return CountBy(container, 1, f)
}

func CountBy(container interface{}, span int, f func(interface{}) bool) (n int) {
	Enumerator{Sequence: container, Span: span}.Each(func(x interface{}) {
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

func Reduce(seq, seed, f interface{}) interface{} {
	return ReduceBy(seq, seed, 1, f)
}

func ReduceBy(seq, seed interface{}, span int, f interface{}) interface{} {
	return Enumerator{Sequence: seq, Span: span}.Reduce(seed, f)
}
