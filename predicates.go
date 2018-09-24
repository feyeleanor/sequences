package sequences

func None(seq interface{}, f func(interface{}) bool) bool {
	var count int
	UntilTrue(seq, func(x interface{}) bool {
		if f(x) {
			count++
		}
		return count > 0
	})
	return count == 0
}

func One(seq interface{}, f func(interface{}) bool) bool {
	var count int
	WhileTrue(seq, func(x interface{}) bool {
		if f(x) {
			count++
		}
		return count < 2
	})
	return count == 1
}

func Any(seq interface{}, f func(interface{}) bool) bool {
	return UntilTrue(seq, f) > 0
}

func All(seq interface{}, f func(interface{}) bool) bool {
	l := Len(seq)
	return l > 0 && l == WhileTrue(seq, f)
}

/*
	Density determines the proportion of a sequence which matches a given predicate
*/
func Density(seq interface{}, f func(interface{}) bool) (r float64) {
	var l, n int

	Each(seq, func(x interface{}) {
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

func IsDense(seq interface{}, threshold float64, f func(interface{}) bool) bool {
	return Density(seq, f) > threshold
}

func IsSparse(seq interface{}, threshold float64, f func(interface{}) bool) bool {
	return Density(seq, f) <= threshold
}
