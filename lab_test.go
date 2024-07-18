package golab

import (
	"testing"
)

func TestAnnotations(test *testing.T) {
	annotations := []Annotation{
		{start: 0.0, end: 1.0, label: "Test Annotation"},
	}

	lab := Lab{annotations: annotations}

	testAnnotation := lab.annotations[0].label
	testStart := lab.annotations[0].start
	testEnd := lab.annotations[0].end
	testDuration := lab.annotations[0].GetDuration()

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
		{start: 0.0, end: 1.0, label: "Test Annotation"},
		{start: 1.0, end: 10.5, label: "Test Annotation 2"},
	}

	lab := Lab{annotations: annotations}

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
