package gosequence

import(
	"reflect"
	"testing"
)

func TestWhileSlice(t *testing.T) {
	ConfirmWhileTrue := func(s interface{}, count int) {
		if i := While(s, func(v interface{}) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index int, v interface{}) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index, v interface{}) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}
	}
	ConfirmWhileTrue(partially_iterable_slice{}, 0)
	ConfirmWhileTrue(partially_iterable_slice{false}, 0)
	ConfirmWhileTrue(partially_iterable_slice{false, true}, 0)
	ConfirmWhileTrue(partially_iterable_slice{true, false}, 1)
	ConfirmWhileTrue(partially_iterable_slice{true, true}, 2)
	ConfirmWhileTrue(partially_iterable_slice{true, true, true}, 3)
	ConfirmWhileTrue(partially_iterable_slice{true, true, false, true}, 2)

	ConfirmWhileIndexableTrue := func(s interface{}, count int) {
		ConfirmWhileTrue(s, count)
		if i := While(s, func(v bool) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index int, v bool) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index interface{}, v bool) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}
	}
	ConfirmWhileIndexableTrue(indexable_slice{}, 0)
	ConfirmWhileIndexableTrue(indexable_slice{false}, 0)
	ConfirmWhileIndexableTrue(indexable_slice{false, true}, 0)
	ConfirmWhileIndexableTrue(indexable_slice{true, false}, 1)
	ConfirmWhileIndexableTrue(indexable_slice{true, true}, 2)
	ConfirmWhileIndexableTrue(indexable_slice{true, true, true}, 3)
	ConfirmWhileIndexableTrue(indexable_slice{true, true, false, true}, 2)
	ConfirmWhileIndexableTrue(indexable_function(func(i int) interface{} { return i > 2 }), 0)
	ConfirmWhileIndexableTrue(indexable_function(func(i int) interface{} { return i < 4 }), 4)

	ConfirmWhileSliceTrue := func(s interface{}, count int) {
		ConfirmWhileIndexableTrue(s, count)
		if i := While(s, func(v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index int, v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index interface{}, v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := While(s, func(index, v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}
	}
	ConfirmWhileSliceTrue([]bool{}, 0)
	ConfirmWhileSliceTrue([]bool{false}, 0)
	ConfirmWhileSliceTrue([]bool{false, true}, 0)
	ConfirmWhileSliceTrue([]bool{true, false}, 1)
	ConfirmWhileSliceTrue([]bool{true, true}, 2)
	ConfirmWhileSliceTrue([]bool{true, true, true}, 3)
	ConfirmWhileSliceTrue([]bool{true, true, false, true}, 2)
}

func TestUntilSlice(t *testing.T) {
	ConfirmUntilTrue := func(s interface{}, count int) {
		if i := Until(s, func(v interface{}) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index int, v interface{}) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index, v interface{}) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}
	}
	ConfirmUntilTrue(partially_iterable_slice{}, 0)
	ConfirmUntilTrue(partially_iterable_slice{false}, 1)
	ConfirmUntilTrue(partially_iterable_slice{true, false}, 0)
	ConfirmUntilTrue(partially_iterable_slice{false, true}, 1)
	ConfirmUntilTrue(partially_iterable_slice{false, false}, 2)
	ConfirmUntilTrue(partially_iterable_slice{false, false, false}, 3)
	ConfirmUntilTrue(partially_iterable_slice{false, false, true, false}, 2)

	ConfirmUntilIndexableTrue := func(s interface{}, count int) {
		ConfirmUntilTrue(s, count)
		if i := Until(s, func(v bool) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index int, v bool) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index interface{}, v bool) bool { return v == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}
	}
	ConfirmUntilIndexableTrue(indexable_slice{}, 0)
	ConfirmUntilIndexableTrue(indexable_slice{false}, 1)
	ConfirmUntilIndexableTrue(indexable_slice{true, false}, 0)
	ConfirmUntilIndexableTrue(indexable_slice{false, true}, 1)
	ConfirmUntilIndexableTrue(indexable_slice{false, false}, 2)
	ConfirmUntilIndexableTrue(indexable_slice{false, false, false}, 3)
	ConfirmUntilIndexableTrue(indexable_slice{false, false, true, false}, 2)
	ConfirmUntilIndexableTrue(indexable_function(func(i int) interface{} { return i < 2 }), 0)
	ConfirmUntilIndexableTrue(indexable_function(func(i int) interface{} { return i > 3 }), 4)

	ConfirmUntilSliceTrue := func(s interface{}, count int) {
		ConfirmUntilIndexableTrue(s, count)
		if i := Until(s, func(v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index int, v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index interface{}, v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}

		if i := Until(s, func(index, v reflect.Value) bool { return v.Bool() == true }); i != count {
			t.Fatalf("%v: should have succeeded %v times but instead succeeded %v times", s, count, i)
		}
	}
	ConfirmUntilSliceTrue([]bool{}, 0)
	ConfirmUntilSliceTrue([]bool{false}, 1)
	ConfirmUntilSliceTrue([]bool{true, false}, 0)
	ConfirmUntilSliceTrue([]bool{false, true}, 1)
	ConfirmUntilSliceTrue([]bool{false, false}, 2)
	ConfirmUntilSliceTrue([]bool{false, false, false}, 3)
	ConfirmUntilSliceTrue([]bool{false, false, true, false}, 2)
}