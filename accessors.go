package golab

func (lab *Lab) SetDenomination(denomination int8) {
	lab.denomination = &Denomination{denomination: denomination}
}

func (lab *Lab) ClearDenomination() {
	lab.denomination = nil
}

func (lab *Lab) GetDenomination() *Denomination {
	return lab.denomination
}

// Gets the total duration of a lab by getting the difference in global start and end.
func (lab Lab) GetDuration() (result float64) {
	// calculate using start and end in case lab file doesnt start at 0
	start := lab.annotations[0].start
	end := lab.annotations[len(lab.annotations)-1].end

	return end - start
}

// Gets the total duration of an annotation
func (annotation Annotation) GetDuration() (result float64) {
	return annotation.end - annotation.start
}

func (lab Lab) GetAnnotations() []Annotation {
	return lab.annotations
}

func (lab *Lab) SetAnnotations(annotations []Annotation) {
	lab.annotations = annotations
}

func (lab *Lab) PushAnnotation(annotation Annotation) {
	lab.annotations = append(lab.annotations, annotation)
}

func (lab *Lab) AppendAnnotations(annotations []Annotation) {
	lab.annotations = append(lab.annotations, annotations...)
}

func (lab *Lab) InsertAnnotation(index int, annotation Annotation) {
	lab.annotations = append(lab.annotations[:index], append([]Annotation{annotation}, lab.annotations[index:]...)...)
}

func (lab *Lab) RemoveAnnotation(index int) {
	lab.annotations = append(lab.annotations[:index], lab.annotations[index+1:]...)
}

func (lab Lab) GetName() string {
	return lab.name
}

func (lab *Lab) SetName(name string) {
	lab.name = name
}

func (lab Lab) GetPrecision() uint8 {
	return lab.precision
}

func (lab *Lab) SetPrecision(precision uint8) {
	lab.precision = precision
}
