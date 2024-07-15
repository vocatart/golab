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
	testDuration := lab.annotations[0].getDuration()

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

	testLength := lab.getDuration()

	if testLength != 10.5 {
		test.Fatalf("wanted lab length of 10.5 sec, receieved %g", testLength)
	}
}

func TestReadWritingLab(test *testing.T) {
	lab := readLab("examples/01.lab")

	writeLab(lab, "examples/output.lab")
}
