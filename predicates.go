package sequences

func Any(container interface{}, f func(interface{}) bool) (b bool) {
	if l := Until(container, f); l > 0 {
		 b = l < Len(container)
	}
	return 
}

func All(container interface{}, f func(interface{}) bool) (b bool) {
	if l := While(container, f); l > 0 {
		b = l == Len(container)
	}
	return
}

func None(container interface{}, f func(interface{}) bool) (b bool) {
	return Until(container, f) == Len(container)
}

func One(container interface{}, f func(interface{}) bool) (b bool) {
	count := 0
	While(container, func(x interface{}) bool {
		if f(x) {
			count++
		}
		if count > 1 {
			return false
		}
		return true
	})
	return count == 1
}