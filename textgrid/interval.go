package textgrid

// Interval structs are annotations that hold a start time, end time, and text label.
type Interval struct {
	xmin float64
	xmax float64
	text string
}

// GetDuration returns the duration of an Interval.
func (interval *Interval) GetDuration() float64 {
	return interval.xmax - interval.xmin
}

// GetMedian returns the midpoint of an Interval.
func (interval *Interval) GetMedian() float64 {
	return (interval.xmin + interval.xmax) / 2.0
}

// GetXmin returns xmin of an Interval.
func (interval *Interval) GetXmin() float64 {
	return interval.xmin
}

// GetXmax returns xmax of an Interval.
func (interval *Interval) GetXmax() float64 {
	return interval.xmax
}

func (interval *Interval) GetText() string {
	return interval.text
}

// SetXmin sets xmin of an Interval.
func (interval *Interval) SetXmin(xmin float64) {
	interval.xmin = xmin
}

// SetXmax sets xmax of an Interval.
func (interval *Interval) SetXmax(xmax float64) {
	interval.xmax = xmax
}

// SetText sets the text label of an Interval.
func (interval *Interval) SetText(text string) {
	interval.text = text
}
