package sequences

func (enum Enumerator) ifTrue(cond func(int) bool, f func(int)) (count int) {
	count = enum.each(func(cursor int) (ok bool) {
		if ok = cond(cursor); ok {
			f(cursor)
		}
		return
	})
	return
}

func (enum Enumerator) ifFalse(cond func(int) bool, f func(int)) (count int) {
	count = enum.each(func(cursor int) (ok bool) {
		if ok = !cond(cursor); ok {
			f(cursor)
		}
		return
	})
	return
}

func None(container interface{}, f func(interface{}) bool) bool {
	var count int
	Until(container, func(x interface{}) bool {
		if f(x) {
			count++
		}
		return count > 0
	})
	return count == 0
}

func One(container interface{}, f func(interface{}) bool) bool {
	var count int
	While(container, func(x interface{}) bool {
		if f(x) {
			count++
		}
		return count < 2
	})
	return count == 1
}

func Any(container interface{}, f func(interface{}) bool) bool {
	return Until(container, f) > 0
}

func All(container interface{}, f func(interface{}) bool) bool {
	l := Len(container)
	return l > 0 && l == While(container, f)
}

/*
	Density determines the proportion of a sequence which matches a given predicate
*/
func Density(container interface{}, f func(interface{}) bool) (r float64) {
	var l, n int

	Each(container, func(x interface{}) {
		l++
		if f(x) {
			n++
		}
	})
	if l > 0 {
		r = float64(n) / float64(l)
	}
	return
}

func IsDense(container interface{}, threshold float64, f func(interface{}) bool) bool {
	return Density(container, f) > threshold
}

func IsSparse(container interface{}, threshold float64, f func(interface{}) bool) bool {
	return Density(container, f) <= threshold
}
