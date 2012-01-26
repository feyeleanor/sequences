package sequences

import "testing"

func TestAny(t *testing.T) {
	ConfirmAny := func(o interface{}, f func(i interface{}) bool) {
		if !Any(o, f) {
			t.Fatalf("Any(%v, f) should be true but is false", o)
		}
	}

	RefuteAny := func(o interface{}, f func(i interface{}) bool) {
		if Any(o, f) {
			t.Fatalf("Any(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteAny(nil, IsPositive)
	RefuteAny([]int{}, IsPositive)
	RefuteAny([]int{0}, IsPositive)
	ConfirmAny([]int{0, 1}, IsPositive)
	ConfirmAny([]int{0, 0, 1}, IsPositive)
}

func TestAll(t *testing.T) {
	ConfirmAll := func(o interface{}, f func(i interface{}) bool) {
		if !All(o, f) {
			t.Fatalf("All(%v, f) should be true but is false", o)
		}
	}

	RefuteAll := func(o interface{}, f func(i interface{}) bool) {
		if All(o, f) {
			t.Fatalf("All(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteAll(nil, IsPositive)
	RefuteAll([]int{}, IsPositive)
	RefuteAll([]int{0}, IsPositive)
	RefuteAll([]int{0, 1}, IsPositive)
	RefuteAll([]int{0, 0, 1}, IsPositive)

	ConfirmAll([]int{1}, IsPositive)
	ConfirmAll([]int{1, 1}, IsPositive)
	ConfirmAll([]int{1, 1, 1}, IsPositive)

	RefuteAll([]interface{}{}, IsPositive)
	RefuteAll([]interface{}{0}, IsPositive)
	RefuteAll([]interface{}{0, 1}, IsPositive)
	RefuteAll([]interface{}{0, 0, 1}, IsPositive)

	ConfirmAll([]interface{}{1}, IsPositive)
	ConfirmAll([]interface{}{1, 1}, IsPositive)
	ConfirmAll([]interface{}{1, 1, 1}, IsPositive)
}

func TestNone(t *testing.T) {
	ConfirmNone := func(o interface{}, f func(i interface{}) bool) {
		if !None(o, f) {
			t.Fatalf("None(%v, f) should be true but is false", o)
		}
	}

	RefuteNone := func(o interface{}, f func(i interface{}) bool) {
		if None(o, f) {
			t.Fatalf("None(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmNone(nil, IsPositive)
	ConfirmNone([]int{}, IsPositive)
	ConfirmNone([]int{0}, IsPositive)
	RefuteNone([]int{0, 1}, IsPositive)
	RefuteNone([]int{0, 0, 1}, IsPositive)

	RefuteNone([]int{1}, IsPositive)
	RefuteNone([]int{1, 1}, IsPositive)
	RefuteNone([]int{1, 1, 1}, IsPositive)

	ConfirmNone([]interface{}{}, IsPositive)
	ConfirmNone([]interface{}{0}, IsPositive)
	RefuteNone([]interface{}{0, 1}, IsPositive)
	RefuteNone([]interface{}{0, 0, 1}, IsPositive)

	RefuteNone([]interface{}{1}, IsPositive)
	RefuteNone([]interface{}{1, 1}, IsPositive)
	RefuteNone([]interface{}{1, 1, 1}, IsPositive)
}

func TestOne(t *testing.T) {
	ConfirmOne := func(o interface{}, f func(i interface{}) bool) {
		if !One(o, f) {
			t.Fatalf("One(%v, f) should be true but is false", o)
		}
	}

	RefuteOne := func(o interface{}, f func(i interface{}) bool) {
		if One(o, f) {
			t.Fatalf("One(%v, f) should be false but is true", o)
		}
	}

	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	RefuteOne(nil, IsPositive)
	RefuteOne([]int{}, IsPositive)
	RefuteOne([]int{0}, IsPositive)
	ConfirmOne([]int{0, 1}, IsPositive)
	ConfirmOne([]int{0, 0, 1}, IsPositive)
	RefuteOne([]int{0, 0, 1, 1}, IsPositive)

	ConfirmOne([]int{1}, IsPositive)
	RefuteOne([]int{1, 1}, IsPositive)
	RefuteOne([]int{1, 1, 1}, IsPositive)

	RefuteOne([]interface{}{}, IsPositive)
	RefuteOne([]interface{}{0}, IsPositive)
	ConfirmOne([]interface{}{0, 1}, IsPositive)
	ConfirmOne([]interface{}{0, 0, 1}, IsPositive)
	RefuteOne([]interface{}{0, 0, 1, 1}, IsPositive)

	ConfirmOne([]interface{}{1}, IsPositive)
	RefuteOne([]interface{}{1, 1}, IsPositive)
	RefuteOne([]interface{}{1, 1, 1}, IsPositive)
}