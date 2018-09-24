package sequences

import R "reflect"

func readSlice(seq R.Value, i int) (r R.Value) {
	return R.Index(seq, R.ValueOf(i))
}

func readSliceInterface(seq R.Value, i int) (r interface{}) {
	return readSlice(seq, i).Interface()
}
