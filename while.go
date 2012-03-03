package sequences

import R "reflect"

func whileIndexable(container Indexable, r bool, f interface{}) (count int, e error) {
	var end	int
	switch end, e = Len(container); {
	case e != nil:
		return
	case end > 0:
		switch f := f.(type) {
		case func(interface{}) bool:
			if f(container.AtOffset(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(container.AtOffset(i)) != r {
						break
					}
					count++
				}
			}
		case func(int, interface{}) bool:
			if f(0, container.AtOffset(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, container.AtOffset(i)) != r {
						break
					}
					count++
				}
			}
		case func(interface{}, interface{}) bool:
			if f(0, container.AtOffset(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, container.AtOffset(i)) != r {
						break
					}
					count++
				}
			}
		case func(...interface{}) bool:
			p := make([]interface{}, end, end)
			for i := 0; i < end; i++ {
				p[i] = container.AtOffset(i)
			}
			if f(p...) == r {
				count = 1
			}
		case func(bool, ...interface{}) int:
			p := make([]interface{}, end, end)
			for i := 0; i < end; i++ {
				p[i] = container.AtOffset(i)
			}
			count = f(r, p...)
		case func(R.Value) bool:
			if f(R.ValueOf(container.AtOffset(0))) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(R.ValueOf(container.AtOffset(i))) != r {
						break
					}
					count++
				}
			}
		case func(int, R.Value) bool:
			if f(0, R.ValueOf(container.AtOffset(0))) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, R.ValueOf(container.AtOffset(i))) != r {
						break
					}
					count++
				}
			}
		case func(R.Value, R.Value) bool:
			if f(R.ValueOf(0), R.ValueOf(container.AtOffset(0))) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(R.ValueOf(i), R.ValueOf(container.AtOffset(i))) != r {
						break
					}
					count++
				}
			}
		case func(...R.Value) bool:
			p := make([]R.Value, end, end)
			for i := 0; i < end; i++ {
				p[i] = R.ValueOf(container.AtOffset(i))
			}
			if f(p...) == r {
				count = 1
			}
		case func(bool, ...R.Value) int:
			p := make([]R.Value, end, end)
			for i := 0; i < end; i++ {
				p[i] = R.ValueOf(container.AtOffset(i))
			}
			count = f(r, p...)
		default:
			if f := R.ValueOf(f); f.Kind() == R.Func {
				if t := f.Type(); t.IsVariadic() {
					switch t.NumIn() {
					case 1:				//	f(...v) bool
						p := make([]R.Value, end, end)
						for i := 0; i < end; i++ {
							p[i] = R.ValueOf(container.AtOffset(i))
						}
						if f.Call(p)[0].Bool() == r {
							count = 1
						}
					case 2:				//	f(bool, ...v) int
						p := make([]R.Value, 1, 4)
						p[0] = R.ValueOf(r)
						for i := 0; i < end; i++ {
							p = append(p, R.ValueOf(container.AtOffset(i)))
						}
						count = int(f.Call(p)[0].Int())
					default:
						e = UNKNOWN_ITERATOR
					}
				} else {
					switch t.NumIn() {
					case 1:				//	f(v) bool
						p := make([]R.Value, 1, 1)
						for ; count < end; count++ {
							p[0] = R.ValueOf(container.AtOffset(count))
							if f.Call(p)[0].Bool() != r {
								break
							}
						}
					case 2:				//	f(i, v) bool
						p := make([]R.Value, 2, 2)
						for ; count < end; count++ {
							p[0] = R.ValueOf(count)
							p[1] = R.ValueOf(container.AtOffset(count))
							if f.Call(p)[0].Bool() != r {
								break
							}
						}
					default:
						e = UNKNOWN_ITERATOR
					}
				}
			}
		}
	}
	return
}

func whileSlice(s R.Value, r bool, f interface{}) (count int, e error) {
	var end	int
	switch end, e = Len(s); {
	case e != nil:
		return
	case end > 0:
		switch f := f.(type) {
		case func(interface{}) bool:
			if f(s.Index(0).Interface()) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(s.Index(i).Interface()) != r {
						break
					}
					count++
				}
			}
		case func(int, interface{}) bool:
			if f(0, s.Index(0).Interface()) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, s.Index(i).Interface()) != r {
						break
					}
					count++
				}
			}
		case func(interface{}, interface{}) bool:
			if f(0, s.Index(0).Interface()) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, s.Index(i).Interface()) != r {
						break
					}
					count++
				}
			}
		case func(...interface{}) bool:
			p := make([]interface{}, end, end)
			for i := 0; i < end; i++ {
				p[i] = s.Index(0).Interface()
			}
			if f(p) == r {
				count = 1
			}
		case func(bool, ...interface{}) int:
			p := make([]interface{}, end, end)
			for i := 0; i < end; i++ {
				p[i] = s.Index(0).Interface()
			}
			count = f(r, p...)
		case func(R.Value) bool:
			if f(s.Index(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(s.Index(i)) != r {
						break
					}
					count++
				}
			}
		case func(int, R.Value) bool:
			if f(0, s.Index(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, s.Index(i)) != r {
						break
					}
					count++
				}
			}
		case func(interface{}, R.Value) bool:
			if f(0, s.Index(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(i, s.Index(i)) != r {
						break
					}
					count++
				}
			}
		case func(R.Value, R.Value) bool:
			if f(R.ValueOf(0), s.Index(0)) == r {
				count = 1
				for i := 1; i < end; i++ {
					if f(R.ValueOf(i), s.Index(i)) != r {
						break
					}
					count++
				}
			}
		case func(...R.Value) bool:
			p := make([]R.Value, end, end)
			for i := 0; i < end; i++ {
				p[i] = s.Index(0)
			}
			if f(p...) == r {
				count = 1
			}
		case func(bool, ...R.Value) int:
			p := make([]R.Value, end, end)
			for i := 0; i < end; i++ {
				p[i] = s.Index(0)
			}
			count = f(r, p...)
		default:
			if f := R.ValueOf(f); f.Kind() == R.Func {
				if t := f.Type(); t.IsVariadic() {
					switch t.NumIn() {
					case 1:				//	f(...v) bool
						p := make([]R.Value, end, end)
						for i := 1; i < end; i++ {
							p[i] = s.Index(i)
						}
						if f.Call(p)[0].Bool() == r {
							count = 1
						}
					case 2:				//	f(bool, ...v) int
						p := make([]R.Value, 1, 4)
						p[0] = R.ValueOf(r)
						for i := 0; i < end; i++ {
							p = append(p, s.Index(i))
						}
						count = int(f.Call(p)[0].Int())
					default:
						e = UNKNOWN_ITERATOR
					}
				} else {
					switch t.NumIn() {
					case 1:				//	f(v) bool
						p := make([]R.Value, 1, 1)
						for ; count < end; count++ {
							p[0] = s.Index(count)
							if f.Call(p)[0].Bool() != r {
								break
							}
						}
					case 2:				//	f(i, v) bool
						p := make([]R.Value, 2, 2)
						for ; count < end; count++ {
							p[0] = R.ValueOf(count)
							p[1] = s.Index(count)
							if f.Call(p)[0].Bool() != r {
								break
							}
						}
					default:
						e = UNKNOWN_ITERATOR
					}
				}
			}
		}
	}
	return
}

func whileChannel(c R.Value, r bool, f interface{}) (count int, e error) {
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
	case func(...interface{}) bool:
		p := make([]interface{}, 0, 4)
		for v, open := c.Recv(); open; {
			p = append(p, v.Interface())
			v, open = c.Recv()
		}
		if f(p...) == r {
			count = 1
		}
	case func(bool, ...interface{}) int:
		p := make([]interface{}, 0, 4)
		for v, open := c.Recv(); open; {
			p = append(p, v.Interface())
			v, open = c.Recv()
		}
		count = f(r, p...)
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
	case func(...R.Value) bool:
		p := make([]R.Value, 0, 4)
		for v, open := c.Recv(); open; {
			p = append(p, v)
			v, open = c.Recv()
		}
		if f(p...) == r {
			count = 1
		}
	case func(bool, ...R.Value) int:
		p := make([]R.Value, 0, 4)
		for v, open := c.Recv(); open; {
			p = append(p, v)
			v, open = c.Recv()
		}
		count = f(r, p...)
	default:
		if f := R.ValueOf(f); f.Kind() == R.Func {
			if t := f.Type(); t.IsVariadic() {
				switch t.NumIn() {
				case 1:				//	f(...v) bool
					p := make([]R.Value, 0, 4)
					i := 0
					for v, open := c.Recv(); open; i++ {
						p = append(p, v)
						v, open = c.Recv()
					}
					if f.Call(p)[0].Bool() == r {
						count = 1
					}
				case 2:				//	f(bool, ...v) int
					p := make([]R.Value, 0, 4)
					i := 0
					for v, open := c.Recv(); open; i++ {
						p = append(p, v)
						v, open = c.Recv()
					}
					count = int(f.Call(p)[0].Int())
				default:
					e = UNKNOWN_ITERATOR
				}
			} else {
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
					e = UNKNOWN_ITERATOR
				}
			}
		}
	}
	return
}