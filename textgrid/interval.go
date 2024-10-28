package textgrid

// Interval - annotation that holds start and end values (in seconds), and a string value.
type Interval struct {
	xmin float64
	xmax float64
	text string
}

// GetDuration - Returns the duration of the interval.
func (interval *Interval) GetDuration() float64 {
	return interval.xmax - interval.xmin
}

// GetMedian - Returns the midpoint of the interval.
func (interval *Interval) GetMedian() float64 {
	return (interval.xmin + interval.xmax) / 2.0
}

// GetXmin - Returns xmin of the interval.
func (interval *Interval) GetXmin() float64 {
	return interval.xmin
}

// GetXmax - Returns xmax of the interval.
func (interval *Interval) GetXmax() float64 {
	return interval.xmax
}

func (interval *Interval) GetText() string {
	return interval.text
}

// SetXmin - Sets the xmin value of the interval.
func (interval *Interval) SetXmin(xmin float64) {
	interval.xmin = xmin
}

// SetXmax - Sets the xmax value of the inerval.
func (interval *Interval) SetXmax(xmax float64) {
	interval.xmax = xmax
}

// SetText - Sets the text of the interval.
func (interval *Interval) SetText(text string) {
	interval.text = text
}
