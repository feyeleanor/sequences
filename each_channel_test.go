package sequences

import (
	R "reflect"
	"testing"
)

func spannedLength(length, span int) (r int) {
	if span < 0 {
		r = length / -span
	} else {
		r = length / span
	}
	if length%span != 0 {
		r++
	}
	return
}

func TestEnumerateChannelBool(t *testing.T) {
	var count int
	var sequence func() chan bool
	var elements []interface{}

	s := []bool{true, false, false, true, true}
	sequence = func() (r chan bool) {
		r = make(chan bool)
		go func() {
			for _, v := range s {
				r <- v
			}
			close(r)
		}()
		return
	}

	spannedLen := func(span int) (r int) {
		return spannedLength(len(s), span)
	}

	ConfirmEachBy := func(s interface{}, span int, f interface{}) {
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("%T span %v: iteration failed with error %v", s, span, e)
			}
		}()

		count = 0
		if EachBy(s, span, f); count != spannedLen(span) {
			t.Fatalf("%T span %v: total iterations should be %v but are %v", s, span, spannedLen(span), count)
		}
	}

	ConfirmChannelEachBy := func(seq chan bool, span int, f interface{}) {
		elements = make([]interface{}, 0, 0)
		l := spannedLen(span)
		done := make(chan bool)
		switch r := f.(type) {
		case chan bool:
			go func() {
				for p := l; p > 0; p-- {
					v := <-r
					elements = append(elements, v)
				}
				close(done)
			}()
		case chan interface{}:
			go func() {
				for p := l; p > 0; p-- {
					v := <-r
					elements = append(elements, v)
				}
				close(done)
			}()
		case chan R.Value:
			go func() {
				for p := l; p > 0; p-- {
					v := <-r
					elements = append(elements, v)
				}
				close(done)
			}()
			/*		case []chan bool:
								go func() {
									for p := l; p > 0; p-- {
										x := len(r) -1
										v := <- r[x]
										elements = append(elements, v)
										y := len(elements) - 1
										for ; x > -1; x-- {
											w := <- r[x]
					println("v =", v, "w =", w)
											if w != v {
												elements[y] = w
											}
										}
									}
									close(done)
								}()
							case []chan interface{}:
								go func() {
									for p := l; p > 0; p-- {
										v := <-r[0]
										elements = append(elements, v)
									}
									close(done)
								}()
							case []chan R.Value:
								go func() {
									for p := l; p > 0; p-- {
										v := <-r[0]
										elements = append(elements, v)
									}
									close(done)
								}()
			*/
		}
		defer func() {
			if e := recover(); e != nil {
				t.Fatalf("%T span %v: iteration failed with error %v", seq, span, e)
			}
		}()
		count = 0
		EachBy(seq, span, f)
		for range done {
		}
		switch {
		case count != l:
			t.Fatalf("%T span %v: total iterations should be %v but are %v", seq, span, l, count)
		case len(elements) != l:
			t.Fatalf("%T span %v: total iterations should be %v but are %v", seq, span, l, len(elements))
		}
	}

	Confirm := func(seq func() chan bool, span int) {
		//	The channel produces a value only once so AtOffset can't be used to confirm that value is correct
		ConfirmEachBy(seq(), span, func(v bool) {
			count++
		})

		ConfirmEachBy(seq(), span, func(v interface{}) {
			count++
		})

		ConfirmEachBy(seq(), span, func(i int, v bool) {
			if offset := span * count; i != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, i)
			}
			count++
		})

		ConfirmEachBy(seq(), span, func(i int, v interface{}) {
			if offset := span * count; i != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, i)
			}
			count++
		})

		ConfirmEachBy(seq(), span, func(k, v interface{}) {
			if offset := span * count; k != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, k)
			}
			count++
		})

		ConfirmEachBy(seq(), span, func(v R.Value) {
			count++
		})

		ConfirmEachBy(seq(), span, func(i int, v R.Value) {
			if offset := span * count; i != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, i)
			}
			count++
		})

		ConfirmEachBy(seq(), span, func(k interface{}, v R.Value) {
			if offset := span * count; k != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, k)
			}
			count++
		})

		ConfirmEachBy(seq(), span, func(k, v R.Value) {
			if offset := span * count; k.Interface() != offset {
				t.Fatalf("%T span %v: index %v erroneously reported as %v", seq, span, offset, k)
			}
			count++
		})

		ConfirmChannelEachBy(seq(), span, make(chan bool))
		ConfirmChannelEachBy(seq(), span, make(chan interface{}))
		ConfirmChannelEachBy(seq(), span, make(chan R.Value))
		/*		ConfirmChannelEachBy(seq(), span, []chan bool{make(chan bool, 6), make(chan bool, 6), make(chan bool, 6)})
				ConfirmChannelEachBy(seq(), span, []chan interface{}{make(chan interface{}, 6), make(chan interface{}, 6), make(chan interface{}, 6)})
				ConfirmChannelEachBy(seq(), span, []chan R.Value{make(chan R.Value, 6), make(chan R.Value, 6), make(chan R.Value, 6)})
		*/
	}

	t.Logf("Write failing test case for span < 0")
	t.Logf("Write test case for span 0")
	Confirm(sequence, 1)
	Confirm(sequence, 2)
	Confirm(sequence, 3)
	Confirm(sequence, 4)
	Confirm(sequence, 5)
	Confirm(sequence, 6)
}
