package sequences

import R "reflect"

type Measurable interface {
	Len() int
}

type Confined interface {
	Cap() int
}

var _MEASURABLE = R.TypeOf(new(Measurable)).Elem()

var _CONFINED = R.TypeOf(new(Confined)).Elem()

/*
	Determines whether or not its parameter is a function which can be used as a sequence.
*/
func isFunctionSequence(f R.Value) bool {
	t := f.Type()
	return f.Kind() == R.Func && (isInfiniteFunction(t) || isBoundedFunction(t))
}

func isInfiniteFunction(t R.Type) bool {
	return t.NumIn() == 1 && t.In(0).Kind() == R.Int && t.NumOut() == 1
}

func isBoundedFunction(t R.Type) bool {
	return t.NumIn() == 1 && t.In(0).Kind() == R.Int && t.NumOut() == 2 && t.Out(1).Kind() == R.Bool
}

/*
	Determines the length of a sequence. For non-sequence values the length will be one, with the
	exception of nil which is defined to have length zero.
*/
func Len(seq interface{}) (l int) {
	switch seq := seq.(type) {
	case R.Value:
		defer func() {
			recover()
		}()
		if seq.Type().Implements(_MEASURABLE) {
			l = seq.Interface().(Measurable).Len()
		} else {
			switch seq.Kind() {
			case R.Slice:
				l = seq.Len()
			case R.Chan:
				l = _MAXINT_
			case R.Func:
				if isFunctionSequence(seq) {
					l = _MAXINT_
				} else {
					l = 1
				}
			default:
				l = 1
			}
		}
	case Measurable:
		l = seq.Len()
	case []bool:
		l = len(seq)
	case []complex64:
		l = len(seq)
	case []complex128:
		l = len(seq)
	case []error:
		l = len(seq)
	case []float32:
		l = len(seq)
	case []float64:
		l = len(seq)
	case []int:
		l = len(seq)
	case []int8:
		l = len(seq)
	case []int16:
		l = len(seq)
	case []int32:
		l = len(seq)
	case []int64:
		l = len(seq)
	case []interface{}:
		l = len(seq)
	case []string:
		l = len(seq)
	case []uint:
		l = len(seq)
	case []uint8:
		l = len(seq)
	case []uint16:
		l = len(seq)
	case []uint32:
		l = len(seq)
	case []uint64:
		l = len(seq)
	case []uintptr:
		l = len(seq)
	case []R.Value:
		l = len(seq)
	case chan bool, chan complex64, chan complex128, chan error,
		chan float32, chan float64, chan int, chan int8,
		chan int16, chan int32, chan int64, chan interface{},
		chan string, chan uint, chan uint8, chan uint16,
		chan uint32, chan uint64, chan uintptr, chan R.Value:
		l = _MAXINT_
	default:
		l = Len(R.ValueOf(seq))
	}
	return
}

/*
	Determines the capacity of a sequence. For non-sequence values the capacity will be one, with the
	exception of nil which is defined to have capacity zero.
	NOTE: The length of a native map is zero as defined by the runtime.
*/
func Cap(seq interface{}) (c int) {
	switch seq := seq.(type) {
	case R.Value:
		defer func() {
			recover()
		}()
		if seq.Type().Implements(_CONFINED) {
			c = seq.Interface().(Confined).Cap()
		} else {
			switch seq.Kind() {
			case R.Slice:
				c = seq.Cap()
			case R.Chan:
				c = _MAXINT_
			case R.Func:
				if isFunctionSequence(seq) {
					c = _MAXINT_
				} else {
					c = 1
				}
			default:
				c = 1
			}
		}
	case Confined:
		c = seq.Cap()
	case []bool:
		c = cap(seq)
	case []int:
		c = cap(seq)
	case []int8:
		c = cap(seq)
	case []int16:
		c = cap(seq)
	case []int32:
		c = cap(seq)
	case []int64:
		c = cap(seq)
	case []uint:
		c = cap(seq)
	case []uint8:
		c = cap(seq)
	case []uint16:
		c = cap(seq)
	case []uint32:
		c = cap(seq)
	case []uint64:
		c = cap(seq)
	case []uintptr:
		c = cap(seq)
	case []complex64:
		c = cap(seq)
	case []complex128:
		c = cap(seq)
	case []float32:
		c = cap(seq)
	case []float64:
		c = cap(seq)
	case []string:
		c = cap(seq)
	case []interface{}:
		c = cap(seq)
	case chan bool, chan complex64, chan complex128, chan error,
		chan float32, chan float64, chan int, chan int8,
		chan int16, chan int32, chan int64, chan interface{},
		chan string, chan uint, chan uint8, chan uint16,
		chan uint32, chan uint64, chan uintptr, chan R.Value:
		c = _MAXINT_
	default:
		c = Cap(R.ValueOf(seq))
	}
	return
}
