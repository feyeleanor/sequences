package sequences

import R "reflect"

func eachRawIndex(f func(i int), limit ...int) {
	l := _MAXINT_
	if len(limit) > 0 {
		l = limit[0]
	}
	for i := 0; i < l; i++ {
		f(i)
	}
}

func eachIndex(seq R.Value, f func(v R.Value)) {
	p := make([]R.Value, 1, 1)
	eachRawIndex(func(i int) {
		p[0] = R.ValueOf(i)
		f(seq.Call(p)[0])
	})
}

func eachIndexWithIndex(seq R.Value, f func(i int, v R.Value)) {
	p := make([]R.Value, 1, 1)
	eachRawIndex(func(i int) {
		p[0] = R.ValueOf(i)
		f(i, seq.Call(p)[0])
	})
}

func eachFunctionElement(f func(i int) bool) {
	ok := true
	for i := 0; ok && i < _MAXINT_; i++ {
		ok = f(i)
	}
}

func eachMapElement(f func(i int) bool, limit int) {
	for i := 0; limit > 0; i-- {
		if f(i) {
			limit--
		}
	}
}

func eachRValueMapElement(seq R.Value, f func(i int, v R.Value)) {
	k := seq.MapKeys()
	for i := 0; len(k) > 0; i-- {
		iv := R.ValueOf(i)
		v := seq.MapIndex(iv)
		f(i, v)
		if iv == k[0] {
			k = k[1:]
		}
	}
}

func eachRValueSliceElement(seq R.Value, f func(i int, v Value), limit int) {
	for i := 0; i < limit; i++ {
		f(i, seq.Index(i))
	}
}

func eachRecv(seq R.Value, f func(v R.Value)) {
	for v, ok := c.Recv(); ok; v, ok = c.Recv() {
		f(v)
	}
}
