package golab

func (lab *Lab) SetDenomination(denomination int8) {
	lab.Denomination = &Denomination{Denomination: denomination}
}

func (lab *Lab) ClearDenomination() {
	lab.Denomination = nil
}

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

func (lab Lab) GetAnnotations() []Annotation {
	return lab.Annotations
}

func (lab *Lab) SetAnnotations(annotations []Annotation) {
	lab.Annotations = annotations
}

func (lab *Lab) PushAnnotation(annotation Annotation) {
	lab.Annotations = append(lab.Annotations, annotation)
}

func (lab *Lab) AppendAnnotations(annotations []Annotation) {
	lab.Annotations = append(lab.Annotations, annotations...)
}

func (lab *Lab) InsertAnnotation(index int, annotation Annotation) {
	lab.Annotations = append(lab.Annotations[:index], append([]Annotation{annotation}, lab.Annotations[index:]...)...)
}

func (lab *Lab) RemoveAnnotation(index int) {
	lab.Annotations = append(lab.Annotations[:index], lab.Annotations[index+1:]...)
}

func (lab Lab) GetName() string {
	return lab.Name
}

func (lab *Lab) SetName(name string) {
	lab.Name = name
}

func (lab Lab) GetPrecision() uint8 {
	return lab.Precision
}

func (lab *Lab) SetPrecision(precision uint8) {
	lab.Precision = precision
}
