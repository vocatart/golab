package htk

///
/// LAB ACCESSORS AND MUTATORS
///

// Sets the annotations of a lab object
func (lab *Lab) SetAnnotations(annotations []Annotation) {
	lab.annotations = annotations
}

// Gets the annotations of a lab object
func (lab Lab) GetAnnotations() []Annotation {
	return lab.annotations
}

// Pushes an annotation to the end of a lab object
func (lab *Lab) PushAnnotation(annotation Annotation) {
	lab.annotations = append(lab.annotations, annotation)
}

// Appends an annotations object to the end of a lab object
func (lab *Lab) AppendAnnotations(annotations []Annotation) {
	lab.annotations = append(lab.annotations, annotations...)
}

// Removes all annotations from a lab object
func (lab *Lab) ClearAnnotations() {
	lab.annotations = []Annotation{}
}

// Dumps all labels in a lab object to a slice
func (lab *Lab) DumpLabels() []string {
	var result []string

	for _, annotation := range lab.annotations {
		result = append(result, annotation.label)
	}

	return result
}

// Gets the name of a lab object
func (lab Lab) GetName() string {
	return lab.name
}

// Sets the name of a lab object
func (lab *Lab) SetName(name string) {
	lab.name = name
}

// Gets the precision of a lab object
func (lab Lab) GetPrecision() uint8 {
	return lab.precision
}

// Sets the precision of a lab object
func (lab *Lab) SetPrecision(precision uint8) {
	lab.precision = precision
}

// Gets the total duration of a lab by getting the difference in global start and end.
func (lab Lab) GetDuration() (result float64) {
	// calculate using start and end in case lab file doesnt start at 0
	start := lab.annotations[0].start
	end := lab.annotations[len(lab.annotations)-1].end

	return end - start
}

// Gets the total amount of annotations in a lab file.
func (lab Lab) GetLength() (result int) {
	return len(lab.annotations)
}

///
/// ANNOTATION ACCESSORS AND MUTATORS
///

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
