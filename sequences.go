package sequences

import R "reflect"

const (
	_MAXINT_           = int(^uint(0) >> 1)
	INDEX_OUT_OF_RANGE = "runtime error: index out of range"
)

var (
	NOT_ENUMERABLE         error = errors.New("Value is not enumerable")
	NO_ENUMERATOR_PROVIDED error = errors.New("Not an enumerator")
)

func PanicWithIndex(x int) {
	_ = []int{}[x]
}

func IgnoreIndexOutOfRange(f func()) {
	defer func() {
		if x := recover(); x != INDEX_OUT_OF_RANGE {
			panic(x)
		}
	}()
	f()
}
