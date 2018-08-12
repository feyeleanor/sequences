package sequences

func LessThanZero(i interface{}) (r bool) {
	switch i := i.(type) {
	case int:
		r = i < 0
	case int8:
		r = i < 0
	case int16:
		r = i < 0
	case int32:
		r = i < 0
	case int64:
		r = i < 0
	}
	return
}

func EqualZero(i interface{}) (r bool) {
	switch i := i.(type) {
	case int:
		r = i == 0
	case int8:
		r = i == 0
	case int16:
		r = i == 0
	case int32:
		r = i == 0
	case int64:
		r = i == 0
	case uint:
		r = i == 0
	case uint8:
		r = i == 0
	case uint16:
		r = i == 0
	case uint32:
		r = i == 0
	case uint64:
		r = i == 0
	}
	return
}

func GreaterThanZero(i interface{}) (r bool) {
	switch i := i.(type) {
	case int:
		r = i > 0
	case int8:
		r = i > 0
	case int16:
		r = i > 0
	case int32:
		r = i > 0
	case int64:
		r = i > 0
	case uint:
		r = i > 0
	case uint8:
		r = i > 0
	case uint16:
		r = i > 0
	case uint32:
		r = i > 0
	case uint64:
		r = i > 0
	}
	return
}
