package sequences

import "testing"

func TestAny(t *testing.T) {
	reportIterationError := func(o interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("Any(%v, f) iteration failed with error %v", o, e)
		}
	}

	ConfirmAny := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if !Any(o, f) {
			t.Fatalf("Any(%v, f) should be true but is false", o)
		}
	}

	RefuteAny := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if Any(o, f) {
			t.Fatalf("Any(%v, f) should be false but is true", o)
		}
	}

	RefuteAny(nil, isPositive)
	RefuteAny([]int{}, isPositive)
	RefuteAny([]int{0}, isPositive)
	ConfirmAny([]int{0, 1}, isPositive)
	ConfirmAny([]int{0, 0, 1}, isPositive)
}

func TestAll(t *testing.T) {
	reportIterationError := func(o interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("All(%v, f) iteration failed with error %v", o, e)
		}
	}

	ConfirmAll := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if !All(o, f) {
			t.Fatalf("All(%v, f) should be true but is false", o)
		}
	}

	RefuteAll := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if All(o, f) {
			t.Fatalf("Any(%v, f) should be false but is true", o)
		}
	}

	RefuteAll(nil, isPositive)
	RefuteAll([]int{}, isPositive)
	RefuteAll([]int{0}, isPositive)
	RefuteAll([]int{0, 1}, isPositive)
	RefuteAll([]int{0, 0, 1}, isPositive)

	ConfirmAll([]int{1}, isPositive)
	ConfirmAll([]int{1, 1}, isPositive)
	ConfirmAll([]int{1, 1, 1}, isPositive)

	RefuteAll([]interface{}{}, isPositive)
	RefuteAll([]interface{}{0}, isPositive)
	RefuteAll([]interface{}{0, 1}, isPositive)
	RefuteAll([]interface{}{0, 0, 1}, isPositive)

	ConfirmAll([]interface{}{1}, isPositive)
	ConfirmAll([]interface{}{1, 1}, isPositive)
	ConfirmAll([]interface{}{1, 1, 1}, isPositive)
}

func TestNone(t *testing.T) {
	reportIterationError := func(o interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("None(%v, f) iteration failed with error %v", o, e)
		}
	}

	ConfirmNone := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if !None(o, f) {
			t.Fatalf("None(%v, f) should be true but is false", o)
		}
	}

	RefuteNone := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if None(o, f) {
			t.Fatalf("None(%v, f) should be false but is true", o)
		}
	}

	ConfirmNone(nil, isPositive)
	ConfirmNone([]int{}, isPositive)
	ConfirmNone([]int{0}, isPositive)
	RefuteNone([]int{0, 1}, isPositive)
	RefuteNone([]int{0, 0, 1}, isPositive)

	RefuteNone([]int{1}, isPositive)
	RefuteNone([]int{1, 1}, isPositive)
	RefuteNone([]int{1, 1, 1}, isPositive)

	ConfirmNone([]interface{}{}, isPositive)
	ConfirmNone([]interface{}{0}, isPositive)
	RefuteNone([]interface{}{0, 1}, isPositive)
	RefuteNone([]interface{}{0, 0, 1}, isPositive)

	RefuteNone([]interface{}{1}, isPositive)
	RefuteNone([]interface{}{1, 1}, isPositive)
	RefuteNone([]interface{}{1, 1, 1}, isPositive)
}

func TestOne(t *testing.T) {
	reportIterationError := func(o interface{}) {
		if e := recover(); e != nil {
			t.Fatalf("One(%v, f) iteration failed with error %v", o, e)
		}
	}

	ConfirmOne := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if !One(o, f) {
			t.Fatalf("One(%v, f) should be true but is false", o)
		}
	}

	RefuteOne := func(o interface{}, f func(i interface{}) bool) {
		defer reportIterationError(o)
		if One(o, f) {
			t.Fatalf("One(%v, f) should be false but is true", o)
		}
	}

	RefuteOne(nil, isPositive)
	RefuteOne([]int{}, isPositive)
	RefuteOne([]int{0}, isPositive)
	ConfirmOne([]int{0, 1}, isPositive)
	ConfirmOne([]int{0, 0, 1}, isPositive)
	RefuteOne([]int{0, 0, 1, 1}, isPositive)

	ConfirmOne([]int{1}, isPositive)
	RefuteOne([]int{1, 1}, isPositive)
	RefuteOne([]int{1, 1, 1}, isPositive)

	RefuteOne([]interface{}{}, isPositive)
	RefuteOne([]interface{}{0}, isPositive)
	ConfirmOne([]interface{}{0, 1}, isPositive)
	ConfirmOne([]interface{}{0, 0, 1}, isPositive)
	RefuteOne([]interface{}{0, 0, 1, 1}, isPositive)

	ConfirmOne([]interface{}{1}, isPositive)
	RefuteOne([]interface{}{1, 1}, isPositive)
	RefuteOne([]interface{}{1, 1, 1}, isPositive)
}

func TestDensity(t *testing.T) {
	t.Fatalf("Fix Tests")
	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmDensity := func(o interface{}, r float64) {
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("Density(%v, f) iteration failed with error %v", o, e)
			}
		}()

		tol := 0.0001
		if d := Density(o, IsPositive); d-r > tol && r-d < tol {
			t.Fatalf("Density(%v, f) should be true but is false", o)
		}
	}
	ConfirmDensity(nil, 0.0)
	ConfirmDensity([]int{}, 0.0)
	ConfirmDensity([]int{0}, 0.0)
	ConfirmDensity([]int{1}, 1.0)

	ConfirmDensity([]int{0, 1}, 0.5)
	ConfirmDensity([]int{1, 0}, 0.5)

	ConfirmDensity([]int{0, 0, 1}, 0.3333)
	ConfirmDensity([]int{0, 1, 0}, 0.3333)
	ConfirmDensity([]int{1, 0, 0}, 0.3333)

	ConfirmDensity([]int{1, 0, 1}, 0.6666)
	ConfirmDensity([]int{1, 1, 0}, 0.6666)
	ConfirmDensity([]int{0, 1, 1}, 0.6666)

	ConfirmDensity([]int{0, 0, 0, 1}, 0.25)
	ConfirmDensity([]int{0, 0, 1, 0}, 0.25)
	ConfirmDensity([]int{0, 1, 0, 0}, 0.25)
	ConfirmDensity([]int{1, 0, 0, 0}, 0.25)

	ConfirmDensity([]int{1, 1, 0, 1}, 0.75)
	ConfirmDensity([]int{1, 0, 1, 1}, 0.75)
	ConfirmDensity([]int{0, 1, 1, 1}, 0.75)
	ConfirmDensity([]int{1, 1, 1, 0}, 0.75)
}

func TestIsDense(t *testing.T) {
	t.Fatalf("Fix Tests")
	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmIsDense := func(o interface{}, d float64) {
		if !IsDense(o, d, IsPositive) {
			t.Fatalf("Dense(%v, %v, f) should be true but is false", o, d)
		}
	}

	RefuteIsDense := func(o interface{}, d float64) {
		if IsDense(o, d, IsPositive) {
			t.Fatalf("Dense(%v, %v, f) should be false but is true", o, d)
		}
	}
	RefuteIsDense(nil, 0.0)
	RefuteIsDense(nil, 0.5)
	RefuteIsDense(nil, 1.0)

	RefuteIsDense([]int{}, 0.0)
	RefuteIsDense([]int{}, 0.5)
	RefuteIsDense([]int{}, 1.0)

	RefuteIsDense([]int{0}, 0.0)
	RefuteIsDense([]int{0}, 0.5)
	RefuteIsDense([]int{0}, 1.0)

	ConfirmIsDense([]int{0, 1}, 0.0)
	ConfirmIsDense([]int{0, 1}, 0.45)
	RefuteIsDense([]int{0, 1}, 0.55)
	RefuteIsDense([]int{0, 1}, 1.0)

	ConfirmIsDense([]int{0, 0, 1}, 0.0)
	RefuteIsDense([]int{0, 0, 1}, 0.5)
	RefuteIsDense([]int{0, 0, 1}, 1.0)

	ConfirmIsDense([]int{0, 1, 1}, 0.0)
	ConfirmIsDense([]int{0, 1, 1}, 0.5)
	RefuteIsDense([]int{0, 1, 1}, 1.0)

	ConfirmIsDense([]int{0, 0, 0, 1}, 0.0)
	RefuteIsDense([]int{0, 0, 0, 1}, 0.5)
	RefuteIsDense([]int{0, 0, 0, 1}, 1.0)

	ConfirmIsDense([]int{0, 0, 1, 1}, 0.0)
	ConfirmIsDense([]int{0, 0, 1, 1}, 0.45)
	RefuteIsDense([]int{0, 0, 1, 1}, 0.55)
	RefuteIsDense([]int{0, 0, 1, 1}, 1.0)

	ConfirmIsDense([]int{0, 1, 1, 1}, 0.0)
	ConfirmIsDense([]int{0, 1, 1, 1}, 0.5)
	RefuteIsDense([]int{0, 1, 1, 1}, 1.0)

	ConfirmIsDense([]int{1, 1, 1, 1}, 0.0)
	ConfirmIsDense([]int{1, 1, 1, 1}, 0.5)
	ConfirmIsDense([]int{1, 1, 1, 1}, 0.99)
	RefuteIsDense([]int{1, 1, 1, 1}, 1.0)
}

func TestIsSparse(t *testing.T) {
	t.Fatalf("Fix Tests")
	IsPositive := func(i interface{}) bool {
		if i, ok := i.(int); ok {
			return i > 0
		}
		return false
	}

	ConfirmIsSparse := func(o interface{}, d float64) {
		if !IsSparse(o, d, IsPositive) {
			t.Fatalf("Dense(%v, %v, f) should be true but is false", o, d)
		}
	}

	RefuteIsSparse := func(o interface{}, d float64) {
		if IsSparse(o, d, IsPositive) {
			t.Fatalf("Dense(%v, %v, f) should be false but is true", o, d)
		}
	}
	ConfirmIsSparse(nil, 0.0)
	ConfirmIsSparse(nil, 0.5)
	ConfirmIsSparse(nil, 1.0)

	ConfirmIsSparse([]int{}, 0.0)
	ConfirmIsSparse([]int{}, 0.5)
	ConfirmIsSparse([]int{}, 1.0)

	ConfirmIsSparse([]int{0}, 0.0)
	ConfirmIsSparse([]int{0}, 0.5)
	ConfirmIsSparse([]int{0}, 1.0)

	RefuteIsSparse([]int{0, 1}, 0.0)
	RefuteIsSparse([]int{0, 1}, 0.45)
	ConfirmIsSparse([]int{0, 1}, 0.55)
	ConfirmIsSparse([]int{0, 1}, 1.0)

	RefuteIsSparse([]int{0, 0, 1}, 0.0)
	ConfirmIsSparse([]int{0, 0, 1}, 0.5)
	ConfirmIsSparse([]int{0, 0, 1}, 1.0)

	RefuteIsSparse([]int{0, 1, 1}, 0.0)
	RefuteIsSparse([]int{0, 1, 1}, 0.5)
	ConfirmIsSparse([]int{0, 1, 1}, 1.0)

	RefuteIsSparse([]int{0, 0, 0, 1}, 0.0)
	ConfirmIsSparse([]int{0, 0, 0, 1}, 0.5)
	ConfirmIsSparse([]int{0, 0, 0, 1}, 1.0)

	RefuteIsSparse([]int{0, 0, 1, 1}, 0.0)
	RefuteIsSparse([]int{0, 0, 1, 1}, 0.45)
	ConfirmIsSparse([]int{0, 0, 1, 1}, 0.55)
	ConfirmIsSparse([]int{0, 0, 1, 1}, 1.0)

	RefuteIsSparse([]int{0, 1, 1, 1}, 0.0)
	RefuteIsSparse([]int{0, 1, 1, 1}, 0.5)
	ConfirmIsSparse([]int{0, 1, 1, 1}, 1.0)

	RefuteIsSparse([]int{1, 1, 1, 1}, 0.0)
	RefuteIsSparse([]int{1, 1, 1, 1}, 0.5)
	RefuteIsSparse([]int{1, 1, 1, 1}, 0.99)
	ConfirmIsSparse([]int{1, 1, 1, 1}, 1.0)
}
