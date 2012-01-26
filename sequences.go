package sequences

import "reflect"

type Measurable interface {
	Len() int
}

type Confined interface {
	Cap() int
}

type Indexable interface {
	At(index int) interface{}
}

type Mappable interface {
	Lookup(key interface{}) interface{}
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
	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		ok = eachSlice(c, f)
							case reflect.Map:		ok = eachMap(c, f)
							case reflect.Chan:		ok = eachChannel(c, f)
							case reflect.Func:		ok = eachFunction(c, f)
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
	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		count = whileSlice(c, true, f)
//							case reflect.Chan:		count = whileChannel(c, true, f)
//							case reflect.Func:		count = whileFunction(c, true, f)
							}
	}
	return
}

func Until(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyIterable:	count = container.While(false, f)
	case Indexable:			count = whileIndexable(container, false, f)
	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		count = whileSlice(c, false, f)
//							case reflect.Chan:		count = whileChannel(c, false, f)
//							case reflect.Func:		count = whileFunction(c, false, f)
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
	default:					switch c := reflect.ValueOf(container); c.Kind() {
								case reflect.Invalid:	r, ok = seed, false
								case reflect.Slice:		r, ok = reduceSlice(c, seed, f)
								case reflect.Map:		r, ok = reduceMap(c, seed, f)
								case reflect.Chan:		r, ok = reduceChan(c, seed, f)
								case reflect.Func:		r, ok = reduceFunction(c, seed, f)
								}
	}
	return
}