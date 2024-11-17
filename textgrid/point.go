package textgrid

// Point - A point is a specific time marker with a string value.
type Point struct {
	value float64
	mark  string
}

// GetValue - Returns the time value of the point.
func (point *Point) GetValue() float64 {
	return point.value
}

// GetMark - Returns the label mark of the point.
func (point *Point) GetMark() string {
	return point.mark
}

// SetValue - Sets the value of the point.
func (point *Point) SetValue(value float64) {
	point.value = value
}

// SetMark - Sets the label mark of the point.
func (point *Point) SetMark(mark string) {
	point.mark = mark
}
