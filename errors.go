package sequences

type SequenceError int

const (
	UNKNOWN_ITERATOR = SequenceError(iota)
	NOT_AN_ITERATOR
	INVALID_SEQUENCE
	ZERO_STEP
	INVALID_START
	ASCENDING_SEQUENCE
)

func (s SequenceError) Error() (r string) {
	switch s {
	case UNKNOWN_ITERATOR:
		r = "Unhandled iterator function"
	case NOT_AN_ITERATOR:
		r = "Iterator function or channel required"
	case INVALID_SEQUENCE:
		r = "Invalid Sequence"
	case ZERO_STEP:
		r = "Step integral 0 invalidates iteration"
	case INVALID_START:
		r = "Start index outside valid bounds."
	case ASCENDING_SEQUENCE:
		r = "This sequence can only be iterated in ascending order"
	}
	return
}