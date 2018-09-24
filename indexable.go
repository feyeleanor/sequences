package sequences

import R "reflect"

type Indexable interface {
	Len() int
	AtOffset(int) interface{}
}

func AtOffset(seq interface{}, i int) (r interface{}) {
	switch seq := seq.(type) {
	case Indexable:
		r = seq.AtOffset(i)

	case []bool:
		r = seq[i]
	case map[int]bool:
		r = seq[i]
	case func(int) bool:
		r = seq(i)

	case []complex64:
		r = seq[i]
	case map[int]complex64:
		r = seq[i]
	case func(int) complex64:
		r = seq(i)

	case []complex128:
		r = seq[i]
	case map[int]complex128:
		r = seq[i]
	case func(int) complex128:
		r = seq(i)

	case []error:
		r = seq[i]
	case map[int]error:
		r = seq[i]
	case func(int) error:
		r = seq(i)

	case []float32:
		r = seq[i]
	case map[int]float32:
		r = seq[i]
	case func(int) float32:
		r = seq(i)

	case []float64:
		r = seq[i]
	case map[int]float64:
		r = seq[i]
	case func(int) float64:
		r = seq(i)

	case []int:
		r = seq[i]
	case map[int]int:
		r = seq[i]
	case func(int) int:
		r = seq(i)

	case []int8:
		r = seq[i]
	case map[int]int8:
		r = seq[i]
	case func(int) int8:
		r = seq(i)

	case []int16:
		r = seq[i]
	case map[int]int16:
		r = seq[i]
	case func(int) int16:
		r = seq(i)

	case []int32:
		r = seq[i]
	case map[int]int32:
		r = seq[i]
	case func(int) int32:
		r = seq(i)

	case []int64:
		r = seq[i]
	case map[int]int64:
		r = seq[i]
	case func(int) int64:
		r = seq(i)

	case []string:
		r = seq[i]
	case map[int]string:
		r = seq[i]
	case func(int) string:
		r = seq(i)

	case []uint:
		r = seq[i]
	case map[int]uint:
		r = seq[i]
	case func(int) uint:
		r = seq(i)

	case []uint8:
		r = seq[i]
	case map[int]uint8:
		r = seq[i]
	case func(int) uint8:
		r = seq(i)

	case []uint16:
		r = seq[i]
	case map[int]uint16:
		r = seq[i]
	case func(int) uint16:
		r = seq(i)

	case []uint32:
		r = seq[i]
	case map[int]uint32:
		r = seq[i]
	case func(int) uint32:
		r = seq(i)

	case []uint64:
		r = seq[i]
	case map[int]uint64:
		r = seq[i]
	case func(int) uint64:
		r = seq(i)

	case []uintptr:
		r = seq[i]
	case map[int]uintptr:
		r = seq[i]
	case func(int) uintptr:
		r = seq(i)

	case []interface{}:
		r = seq[i]
	case map[int]interface{}:
		r = seq[i]
	case func(int) interface{}:
		r = seq(i)

	case []R.Value:
		r = seq[i]
	case map[int]R.Value:
		r = seq[i]
	case func(int) R.Value:
		r = seq(i)

	case R.Value:
		switch seq.Kind() {
		case R.Slice:
			r = seq.Index(i).Interface()
		case R.Map:
			r = seq.MapIndex(R.ValueOf(i)).Interface()
		case R.Func:
			r = readFunctionInterface(seq, i)
		}

	default:
		AtOffset(R.ValueOf(seq), i)
	}
	return
}

func AtOffsets(seq interface{}, offsets ...int) (r interface{}) {
	switch seq := seq.(type) {
	case Indexable:
		n := make([]interface{}, len(offsets))
		for i, _ := range n {
			n[i] = seq.AtOffset(offsets[i])
		}
		r = n
	case []bool:
		n := make([]bool, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []interface{}:
		n := make([]interface{}, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []int:
		n := make([]int, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []int8:
		n := make([]int8, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []int16:
		n := make([]int16, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []int32:
		n := make([]int32, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []int64:
		n := make([]int64, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []uint:
		n := make([]uint, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []uint8:
		n := make([]uint8, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []uint16:
		n := make([]uint16, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []uint32:
		n := make([]uint32, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []uint64:
		n := make([]uint64, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []uintptr:
		n := make([]uintptr, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []float32:
		n := make([]float32, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []float64:
		n := make([]float64, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []complex64:
		n := make([]complex64, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []complex128:
		n := make([]complex128, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []string:
		n := make([]string, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []error:
		n := make([]error, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case []R.Value:
		n := make([]R.Value, len(offsets))
		for i, _ := range n {
			n[i] = seq[offsets[i]]
		}
		r = n
	case R.Value:
		switch seq.Kind() {
		case R.Slice:
			panic("implement for arbitrary slices")
		case R.Map:
			panic("implement for arbitrary maps")
		case R.Func:
			panic("implement for arbitrary functions")
		case R.Chan:
			panic("Channels not yet supported")
		default:
			panic("this object cannot be indexed")
		}
	default:
		AtOffsets(R.ValueOf(seq), offsets...)
	}
	return
}
