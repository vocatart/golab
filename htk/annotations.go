package htk

// An annotation has a starting and ending time in seconds, with a text label.
type Annotation struct {
	start float64
	end   float64
	label string
}

// Gets the total duration of an annotation
func (annotation Annotation) GetDuration() (result float64) {
	return annotation.end - annotation.start
}

// Gets the start time of an annotation
func (annotation Annotation) GetStart() float64 {
	return annotation.start
}

// Sets the start time of an annotation
func (annotation *Annotation) SetStart(start float64) {
	annotation.start = start
}

// Gets the end time of an annotation
func (annotation Annotation) GetEnd() float64 {
	return annotation.end
}

// Sets the end time of an annotation
func (annotation *Annotation) SetEnd(end float64) {
	annotation.end = end
}

// Gets the label of an annotation
func (annotation Annotation) GetLabel() string {
	return annotation.label
}

// Sets the label of an annotation
func (annotation *Annotation) SetLabel(label string) {
	annotation.label = label
}
