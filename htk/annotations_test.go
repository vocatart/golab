package htk

import "testing"

func TestGettingAnnotationDuration(t *testing.T) {
	// loading the first annotation from the first example as a dummy annotation
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	annotation := lab.annotations[0]
	trueDuration := annotation.end - annotation.start

	if annotation.GetDuration() != trueDuration {
		t.Fatalf("wanted duration %f, received %f", trueDuration, annotation.GetDuration())
	}

	t.Log("Getting annotation duration successful!")
}

func TestGettingAnnotationStart(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	annotation := lab.annotations[0]
	trueStart := annotation.start

	if annotation.GetStart() != trueStart {
		t.Fatalf("wanted start time of %f, received %f", trueStart, annotation.GetStart())
	}

	t.Log("getting annotation start successful!")
}

func TestGettingAnnotationEnd(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	annotation := lab.annotations[0]
	trueEnd := annotation.end

	if annotation.GetEnd() != trueEnd {
		t.Fatalf("wanted end time of %f, received %f", trueEnd, annotation.GetEnd())
	}

	t.Log("getting annotation end successful!")
}

func TestGettingAnnotationLabel(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	annotation := lab.annotations[0]
	trueLabel := annotation.label

	if annotation.GetLabel() != trueLabel {
		t.Fatalf("wanted label %s, received %s", trueLabel, annotation.GetLabel())
	}

	t.Log("test setting annotation label successful!")
}

func TestSettingAnnotationStartEnd(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	annotation := lab.annotations[0]
	annotation.SetStart(10)
	annotation.SetEnd(20)

	if annotation.start != 10 {
		t.Fatalf("wanted start time of 10, received %f", annotation.start)
	} else if annotation.end != 20 {
		t.Fatalf("wanted end time of 20, received %f", annotation.end)
	}

	t.Log("setting annotation start and end successful!")
}

func TestSettingAnnotationLabel(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	annotation := lab.annotations[0]
	annotation.SetLabel("label")

	if annotation.label != "label" {
		t.Fatalf("wanted label string of \"label\", recieved %s", annotation.label)
	}
}
