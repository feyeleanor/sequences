package sequences

import R "reflect"

func sortedMapKeys(seq R.Value) (r []int) {
	l := seq.Len()
	r = make([]int, l, l)
	for i, v := range seq.MapKeys() {
		r[i] = int(v.Int())
	}
	sort.Ints(r)
	return
}

func readMap(seq R.Value, i int) (r R.Value) {
	return R.MapIndex(seq, R.ValueOf(i))
}

func readMapInterface(seq R.Value, i int) (r interface{}) {
	return readMap(seq, i).Interface()
}
