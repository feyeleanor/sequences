package sequences

import (
	"fmt"
	R "reflect"
)

type UDT int

func isPositive(i interface{}) bool {
	if i, ok := i.(int); ok {
		return i > 0
	}
	return false
}

type Error int

func (e Error) Error() string {
	return fmt.Sprintf("Error Code: %v\n", e)
}

type enumerable_slice []interface{}

func (s enumerable_slice) Len() int {
	return len(s)
}

func (s enumerable_slice) Each(enum *Enumerator) (i int) {
	switch f := enum.f().(type) {
	case func(interface{}):
		eachRawIndex(func(i int) {
			f(s[i])
		}, len(s))
	case func(int, interface{}):
		eachRawIndex(func(i int) {
			f(i, s[i])
		}, len(s))
	case func(interface{}, interface{}):
		eachRawIndex(func(i int) {
			f(i, s[i])
		}, len(s))
	case func(R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(s[i]))
		}, len(s))
	case func(int, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(s[i]))
		}, len(s))
	case func(interface{}, R.Value):
		eachRawIndex(func(i int) {
			f(i, R.ValueOf(s[i]))
		}, len(s))
	case func(R.Value, R.Value):
		eachRawIndex(func(i int) {
			f(R.ValueOf(i), R.ValueOf(s[i]))
		}, len(s))
	}
	return
}

type indexable_slice []interface{}

func (i indexable_slice) Len() int {
	return len(i)
}

func (i indexable_slice) AtOffset(x int) interface{} {
	return i[x]
}

type indexable_function func(i int) interface{}

func (f indexable_function) Len() int {
	return 10
}

func (f indexable_function) AtOffset(x int) (r interface{}) {
	if x < f.Len() {
		r = f(x)
	} else {
		PanicWithIndex(x)
	}
	return
}
