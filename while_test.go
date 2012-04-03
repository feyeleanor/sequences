package sequences

import(
	"reflect"
	"testing"
)

func TestWhileSlice(t *testing.T) {
	var iterators	[]interface{}
	reportIterationError := func(o interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("While(%v, f) iteration failed with error %v", o, e)
		}
	}

	ConfirmWhileTrue := func(s interface{}, count int) {
		defer reportIterationError(s)
		for n, f := range iterators {
			if i := While(s, f); i != count {
				t.Fatalf("While(%v, iterators[%v]): should have succeeded %v times but instead succeeded %v times", s, n, count, i)
			}
		}
	}

	iterators = []interface{}{
		func(v interface{}) bool { return v.(bool) },
		func(index int, v interface{}) bool { return v.(bool) },
		func(index, v interface{}) bool { return v.(bool) },
	}
	ConfirmWhileTrue(partially_enumerable_slice{}, 0)
	ConfirmWhileTrue(partially_enumerable_slice{false}, 0)
	ConfirmWhileTrue(partially_enumerable_slice{false, true}, 0)
	ConfirmWhileTrue(partially_enumerable_slice{true, false}, 1)
	ConfirmWhileTrue(partially_enumerable_slice{true, true}, 2)
	ConfirmWhileTrue(partially_enumerable_slice{true, true, true}, 3)
	ConfirmWhileTrue(partially_enumerable_slice{true, true, false, true}, 2)

	iterators = append(
		iterators,
		func(v bool) bool { return v },
		func(index int, v bool) bool { return v },
		func(index interface{}, v bool) bool { return v },
	)
	ConfirmWhileTrue(indexable_slice{}, 0)
	ConfirmWhileTrue(indexable_slice{false}, 0)
	ConfirmWhileTrue(indexable_slice{false, true}, 0)
	ConfirmWhileTrue(indexable_slice{true, false}, 1)
	ConfirmWhileTrue(indexable_slice{true, true}, 2)
	ConfirmWhileTrue(indexable_slice{true, true, true}, 3)
	ConfirmWhileTrue(indexable_slice{true, true, false, true}, 2)
	ConfirmWhileTrue(indexable_function(func(i int) interface{} { return i > 2 }), 0)
	ConfirmWhileTrue(indexable_function(func(i int) interface{} { return i < 4 }), 4)

	iterators = append(
		iterators,
		func(v reflect.Value) bool { return v.Bool() },
		func(index int, v reflect.Value) bool { return v.Bool() },
		func(index interface{}, v reflect.Value) bool { return v.Bool() },
		func(index, v reflect.Value) bool { return v.Bool() },
	)

	ConfirmWhileTrue([]bool{}, 0)
	ConfirmWhileTrue([]bool{false}, 0)
	ConfirmWhileTrue([]bool{false, true}, 0)
	ConfirmWhileTrue([]bool{true, false}, 1)
	ConfirmWhileTrue([]bool{true, true}, 2)
	ConfirmWhileTrue([]bool{true, true, true}, 3)
	ConfirmWhileTrue([]bool{true, true, false, true}, 2)
}

func TestUntilSlice(t *testing.T) {
	var iterators	[]interface{}
	reportIterationError := func(o interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("Until(%v, f) iteration failed with error %v", o, e)
		}
	}

	ConfirmUntilTrue := func(s interface{}, count int) {
		defer reportIterationError(s)
		for n, f := range iterators {
			if i := Until(s, f); i != count {
				t.Fatalf("Until(%v, iterators[%v]): should have succeeded %v times but instead succeeded %v times", s, n, count, i)
			}
		}
	}

	iterators = []interface{}{
		func(v interface{}) bool { return v == true },
		func(index int, v interface{}) bool { return v == true },
		func(index, v interface{}) bool { return v == true },
	}
	ConfirmUntilTrue(partially_enumerable_slice{}, 0)
	ConfirmUntilTrue(partially_enumerable_slice{false}, 1)
	ConfirmUntilTrue(partially_enumerable_slice{true, false}, 0)
	ConfirmUntilTrue(partially_enumerable_slice{false, true}, 1)
	ConfirmUntilTrue(partially_enumerable_slice{false, false}, 2)
	ConfirmUntilTrue(partially_enumerable_slice{false, false, false}, 3)
	ConfirmUntilTrue(partially_enumerable_slice{false, false, true, false}, 2)

	iterators = append(
		iterators,
		func(v bool) bool { return v == true },
		func(index int, v bool) bool { return v == true },
		func(index interface{}, v bool) bool { return v == true },
	)
	ConfirmUntilTrue(indexable_slice{}, 0)
	ConfirmUntilTrue(indexable_slice{false}, 1)
	ConfirmUntilTrue(indexable_slice{true, false}, 0)
	ConfirmUntilTrue(indexable_slice{false, true}, 1)
	ConfirmUntilTrue(indexable_slice{false, false}, 2)
	ConfirmUntilTrue(indexable_slice{false, false, false}, 3)
	ConfirmUntilTrue(indexable_slice{false, false, true, false}, 2)
	ConfirmUntilTrue(indexable_function(func(i int) interface{} { return i < 2 }), 0)
	ConfirmUntilTrue(indexable_function(func(i int) interface{} { return i > 3 }), 4)

	iterators = append(
		iterators,
		func(v reflect.Value) bool { return v.Bool() == true },
		func(index int, v reflect.Value) bool { return v.Bool() == true },
		func(index interface{}, v reflect.Value) bool { return v.Bool() == true },
		func(index, v reflect.Value) bool { return v.Bool() == true },
	)
	ConfirmUntilTrue([]bool{}, 0)
	ConfirmUntilTrue([]bool{false}, 1)
	ConfirmUntilTrue([]bool{true, false}, 0)
	ConfirmUntilTrue([]bool{false, true}, 1)
	ConfirmUntilTrue([]bool{false, false}, 2)
	ConfirmUntilTrue([]bool{false, false, false}, 3)
	ConfirmUntilTrue([]bool{false, false, true, false}, 2)
}