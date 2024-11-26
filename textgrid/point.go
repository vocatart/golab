package textgrid

// Point structs are annotations that represent a specific point in time, with a text label.
type Point struct {
	value float64
	mark  string
}

// GetValue returns value of a Point.
func (point *Point) GetValue() float64 {
	return point.value
}

// GetMark returns the text label of a Point.
func (point *Point) GetMark() string {
	return point.mark
}

// SetValue sets the value of a Point.
func (point *Point) SetValue(value float64) {
	point.value = value
}

// SetMark sets the mark of a Point.
func (point *Point) SetMark(mark string) {
	point.mark = mark
}
