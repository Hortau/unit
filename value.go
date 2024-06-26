package unit

type Value struct {
	val  float64
	unit theUnit
}

// NewConverter create a Value to be converted
func NewConverter(v float64) Value {
	return Value{val: v}
}

// From resolve symbol and call a function to returns the type and convert to a base
func (v Value) From(symbol string) Value {
	result, ok := fromMap[symbol]
	if !ok {
		return Value{}
	}
	return result(v.val)
}

// To resolve symbol and call a function to returns a float converted from the base
func (v Value) To(symbol string) (float64, error) {
	if v.unit == 0 {
		return 0, ErrNotFound
	}

	result, ok := toMap[symbol]
	if !ok {
		return 0, ErrNotFound
	}
	return result(v)
}
