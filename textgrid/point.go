package textgrid

// A point is a specifc time marker with a string value.
type Point struct {
	value float64
	mark  string
}

// Returns the time value of the point.
func (point Point) GetValue() float64 {
	return point.value
}

// Returns the label mark of the point.
func (point Point) GetMark() string {
	return point.mark
}

// Sets the value of the point.
func (point *Point) SetValue(value float64) {
	point.value = value
}

// Sets the label mark of the point.
func (point *Point) SetMark(mark string) {
	point.mark = mark
}
