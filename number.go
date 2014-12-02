package jason

type Number struct {
	Float64 float64
	Valid   bool
}

func (j *Jason) number() *Number {

	var valid bool

	// Check the type of this data
	switch j.data.(type) {
	case float64:
		valid = true
		break
	}

	n := new(Number)
	n.Valid = valid

	if valid {
		n.Float64 = j.data.(float64)
	}

	return n
}

func (j *Jason) Number() float64 {
	n := j.number()
	return n.Float64
}

// Conveniece method
func (j *Jason) Float64() float64 {
	return j.Number()
}

// Conveniece method
func (j *Jason) Int64() int64 {
	return int64(j.Number())
}

// Returns true if the object is actually an object
func (j *Jason) IsNumber() bool {
	n := j.number()
	return n.Valid
}
