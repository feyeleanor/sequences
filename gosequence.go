package gosequence

import "reflect"

type Indexable interface {
	Len() int
	At(index int) interface{}
}

type Mappable interface {
	Len() int
	At(key interface{}) interface{}
	Keys() []interface{}
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
	if count == 0 {
		for ok = true; ok; i++ { ok = Each(container, f) }
	} else {
		for ok = true; i < count && ok; i++ { ok = Each(container, f) }
	}
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