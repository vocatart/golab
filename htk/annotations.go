package htk

// Annotation contains a starting and ending time in seconds, with a text label.
type Annotation struct {
	start float64
	end   float64
	label string
}

// GetDuration gets the total duration of an annotation.
func (annotation *Annotation) GetDuration() (result float64) {
	return annotation.end - annotation.start
}

// GetStart gets the start time of an annotation.
func (annotation *Annotation) GetStart() float64 {
	return annotation.start
}

// SetStart sets the start time of an annotation.
func (annotation *Annotation) SetStart(start float64) {
	annotation.start = start
}

// GetEnd gets the end time of an annotation.
func (annotation *Annotation) GetEnd() float64 {
	return annotation.end
}

// SetEnd sets the end time of an annotation.
func (annotation *Annotation) SetEnd(end float64) {
	annotation.end = end
}

// GetLabel gets the label of an annotation.
func (annotation *Annotation) GetLabel() string {
	return annotation.label
}

// SetLabel sets the label of an annotation
func (annotation *Annotation) SetLabel(label string) {
	annotation.label = label
}
