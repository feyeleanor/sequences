package sequences

func Any(container interface{}, f func(interface{}) bool) (b bool, e error) {
	var l	int
	if l, e = Until(container, f); e == nil {
		 b = l > 0
	}
	return 
}

func All(container interface{}, f func(interface{}) bool) (b bool, e error) {
	var l	int
	if l, e = While(container, f); e == nil {
		lc, _ := Len(container)
		b = l == lc
	}
	return
}

func None(container interface{}, f func(interface{}) bool) (b bool, e error) {
	var l	int
	if l, e = Until(container, f); e == nil {
		b = l == 0
	}
	return
}

func One(container interface{}, f func(interface{}) bool) (b bool, e error) {
	var count	int
	_, e = While(container, func(x interface{}) bool {
		if f(x) {
			count++
		}
		if count > 1 {
			return false
		}
		return true
	})
	if e == nil {
		b = count == 1
	}
	return
}

func Density(container interface{}, f func(interface{}) bool) (r float64, e error) {
	var c	int
	if c, e = Count(container, f); e == nil {
		var l	int
		switch l, e = Len(container); {
		case e != nil:
			return
		case l == 0:
			r = float64(1)
		default:
			r = float64(c) / float64(l)
		}
	}
	return
}

func IsDense(container interface{}, threshold float64, f func(interface{}) bool) bool {
	d, e := Density(container, f);
	return e == nil && d > threshold
}

func IsSparse(container interface{}, threshold float64, f func(interface{}) bool) bool {
	d, e := Density(container, f);
	return e == nil && d <= threshold
}