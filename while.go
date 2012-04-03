package sequences

import R "reflect"

func whileIndexable(container Indexable, r bool, f interface{}) (count int) {
	defer IgnoreIndexOutOfRange()
	switch f := f.(type) {
	case func(interface{}) bool:
		for ; f(container.AtOffset(count)) == r; count++ {}
	case func(int, interface{}) bool:
		for ; f(count, container.AtOffset(count)) == r; count++ {}
	case func(interface{}, interface{}) bool:
		for ; f(count, container.AtOffset(count)) == r; count++ {}
	case func(R.Value) bool:
		for ; f(R.ValueOf(container.AtOffset(count))) == r; count++ {}
	case func(int, R.Value) bool:
		for ; f(count, R.ValueOf(container.AtOffset(count))) == r; count++ {}
	case func(R.Value, R.Value) bool:
		for ; f(R.ValueOf(count), R.ValueOf(container.AtOffset(count))) == r; count++ {}
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); t.NumOut() == 1 {
				switch t.NumIn() {
				case 1:				//	f(v) bool
					p := []R.Value{ R.ValueOf(container.AtOffset(count)) }
					for f.Call(p)[0].Bool() == r {
						count++
						p[0] = R.ValueOf(container.AtOffset(count))
					}
				case 2:				//	f(i, v) bool
					p := []R.Value{ R.ValueOf(count), R.ValueOf(container.AtOffset(count)) }
					for f.Call(p)[0].Bool() == r {
						count++
						p[0] = R.ValueOf(count)
						p[1] = R.ValueOf(container.AtOffset(count))
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		}
	}
	return
}

func whileSlice(s R.Value, r bool, f interface{}) (count int) {
	defer IgnoreIndexOutOfRange()
	switch f := f.(type) {
	case func(interface{}) bool:
		for ; f(s.Index(count).Interface()) == r; count++ {}
	case func(int, interface{}) bool:
		for ; f(count, s.Index(count).Interface()) == r; count++ {}
	case func(interface{}, interface{}) bool:
		for ; f(count, s.Index(count).Interface()) == r; count++ {}
	case func(R.Value) bool:
		for ; f(s.Index(count)) == r; count++ {}
	case func(int, R.Value) bool:
		for ; f(count, s.Index(count)) == r; count++ {}
	case func(interface{}, R.Value) bool:
		for ; f(count, s.Index(count)) == r; count++ {}
	case func(R.Value, R.Value) bool:
		for ; f(R.ValueOf(count), s.Index(count)) == r; count++ {}
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); t.NumOut() == 1 {
				switch t.NumIn() {
				case 1:				//	f(v) bool
					p := []R.Value{ s.Index(count) }
					for f.Call(p)[0].Bool() == r {
						count++
						p[0] = s.Index(count)
					}
				case 2:				//	f(i, v) bool
					p := []R.Value{ R.ValueOf(count), s.Index(count) }
					for f.Call(p)[0].Bool() == r {
						count++
						p[0] = R.ValueOf(count)
						p[1] = s.Index(count)
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		}
	}
	return
}

func whileChannel(c R.Value, r bool, f interface{}) (count int) {
	switch f := f.(type) {
	case func(interface{}) bool:
		for v, open := c.Recv(); open && f(v.Interface()) == r; count++ {
			v, open = c.Recv()
		}
	case func(int, interface{}) bool:
		for v, open := c.Recv(); open && f(count, v.Interface()) == r; count++ {
			v, open = c.Recv()
		}
	case func(interface{}, interface{}) bool:
		for v, open := c.Recv(); open && f(count, v.Interface()) == r; count++ {
			v, open = c.Recv()
		}
	case func(R.Value) bool:
		for v, open := c.Recv(); open && f(v) == r; count++ {
			v, open = c.Recv()
		}
	case func(int, R.Value) bool:
		for v, open := c.Recv(); open && f(count, v) == r; count++ {
			v, open = c.Recv()
		}
	case func(interface{}, R.Value) bool:
		for v, open := c.Recv(); open && f(count, v) == r; count++ {
			v, open = c.Recv()
		}
	case func(R.Value, R.Value) bool:
		for v, open := c.Recv(); open && f(R.ValueOf(count), v) == r; count++ {
			v, open = c.Recv()
		}
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); t.NumOut() == 1 {
				switch t.NumIn() {
				case 1:				//	f(v) bool
					open := false
					p := make([]R.Value, 1, 1)
					for p[0], open = c.Recv(); open && f.Call(p)[0].Bool() == r; count++ {
						p[0], open = c.Recv()
					}
				case 2:				//	f(i, v) bool
					open := false
					p := make([]R.Value, 2, 2)
					p[0] = R.ValueOf(0)
					for p[1], open = c.Recv(); open && f.Call(p)[0].Bool() == r; count++ {
						p[0] = R.ValueOf(count)
						p[1], open = c.Recv()
					}
				default:
					panic(UNHANDLED_ITERATOR)
				}
			}
		}
	}
	return
}