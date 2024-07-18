package golab

// Sets the denomination of a lab object
func (lab *Lab) SetDenomination(denomination int8) {
	lab.Denomination = &Denomination{Denomination: denomination}
}

// Sets the denomination of a lab object to nil
func (lab *Lab) ClearDenomination() {
	lab.Denomination = nil
}

// Gets the denomination of a lab object
func (lab *Lab) GetDenomination() *Denomination {
	return lab.Denomination
}

// Gets the total duration of a lab by getting the difference in global start and end.
func (lab Lab) GetDuration() (result float64) {
	// calculate using start and end in case lab file doesnt start at 0
	start := lab.Annotations[0].Start
	end := lab.Annotations[len(lab.Annotations)-1].End

	return end - start
}

// Gets the total duration of an annotation
func (annotation Annotation) GetDuration() (result float64) {
	return annotation.End - annotation.Start
}

// Gets the annotations of a lab object
func (lab Lab) GetAnnotations() []Annotation {
	return lab.Annotations
}

// Sets the annotations of a lab object
func (lab *Lab) SetAnnotations(annotations []Annotation) {
	lab.Annotations = annotations
}

// Pushes an annotation to the end of a lab object
func (lab *Lab) PushAnnotation(annotation Annotation) {
	lab.Annotations = append(lab.Annotations, annotation)
}

// Appends an annotations object to the end of a lab object
func (lab *Lab) AppendAnnotations(annotations []Annotation) {
	lab.Annotations = append(lab.Annotations, annotations...)
}

// Inserts an annotation at a specific index
func (lab *Lab) InsertAnnotation(index int, annotation Annotation) {
	lab.Annotations = append(lab.Annotations[:index], append([]Annotation{annotation}, lab.Annotations[index:]...)...)
}

// Removes an annotation at a specific index
func (lab *Lab) RemoveAnnotation(index int) {
	lab.Annotations = append(lab.Annotations[:index], lab.Annotations[index+1:]...)
}

// Gets the name of a lab object
func (lab Lab) GetName() string {
	return lab.Name
}

// Sets the name of a lab object
func (lab *Lab) SetName(name string) {
	lab.Name = name
}

// Gets the precision of a lab object
func (lab Lab) GetPrecision() uint8 {
	return lab.Precision
}

// Sets the precision of a lab object
func (lab *Lab) SetPrecision(precision uint8) {
	lab.Precision = precision
}
