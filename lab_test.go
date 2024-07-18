package golab

import (
	"testing"
)

func TestAnnotations(test *testing.T) {
	annotations := []Annotation{
		{Start: 0.0, End: 1.0, Label: "Test Annotation"},
	}

	lab := Lab{Annotations: annotations}

	testAnnotation := lab.Annotations[0].Label
	testStart := lab.Annotations[0].Start
	testEnd := lab.Annotations[0].End
	testDuration := lab.Annotations[0].GetDuration()

	if testAnnotation != "Test Annotation" {
		test.Fatalf("wanted 'Test Annotation', recieved %s", testAnnotation)
	} else if testStart != 0.0 {
		test.Fatalf("wanted start time of 0 sec, received %g", testStart)
	} else if testEnd != 1.0 {
		test.Fatalf("wanted end time of 1.0 sec, recieved %g", testEnd)
	} else if testDuration != 1.0 {
		test.Fatalf("wanted duration of 1.0 sec, recieved %g", testDuration)
	}

}

func TestLab(test *testing.T) {
	annotations := []Annotation{
		{Start: 0.0, End: 1.0, Label: "Test Annotation"},
		{Start: 1.0, End: 10.5, Label: "Test Annotation 2"},
	}

	lab := Lab{Annotations: annotations}

	testLength := lab.GetDuration()

	if testLength != 10.5 {
		test.Fatalf("wanted lab length of 10.5 sec, receieved %g", testLength)
	}
}

func TestReadWritingLab(test *testing.T) {
	lab1 := ReadLab("examples/01.lab")
	lab1.WriteLab("examples/output.lab", true)

	lab2 := ReadLab("examples/02.lab")
	lab2.WriteLab("examples/output2.lab", true)
}
