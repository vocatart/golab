package htk

import "testing"

func TestReadingLab(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	if lab.name != "short.lab" {
		t.Fatalf("wanted 'short', recieved %s", lab.name)
	} else if lab.annotations == nil {
		t.Fatalf("wanted annotations, recieved nil")
	} else if lab.precision != 7 {
		t.Fatalf("wanted precision of 7, recieved %d", lab.precision)
	}

	t.Log("lab reading successful!")
}

func TestReadingDifferentPrecision(t *testing.T) {
	lab, err := ReadLab("examples/02.lab")
	if err != nil {
		t.Fatal(err)
	}

	if lab.precision != 6 {
		t.Fatalf("wanted precision of 6, recieved %d", lab.precision)
	}

	t.Log("different precisions reading successful!")
}

func TestWritingLab(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	err = lab.WriteLab("examples/output.lab", true)
	if err != nil {
		return
	}

	t.Log("lab writing successful!")
}

func TestWritingDifferentPrecision(t *testing.T) {
	lab, err := ReadLab("examples/02.lab")
	if err != nil {
		t.Fatal(err)
	}

	err = lab.WriteLab("examples/output2.lab", true)
	if err != nil {
		return
	}

	t.Log("different pricision writing successful!")
}

func TestPrintingLabString(t *testing.T) {
	lab, err := ReadLab("examples/one_line.lab")
	if err != nil {
		t.Fatal(err)
	}

	if lab.ToString() != "0.0000000 10.0000000 test\n" {
		t.Fatal("malformed lab string!")
	}

	t.Log("printing lab to string successful!")
}

func TestSettingAnnotations(t *testing.T) {
	lab := Lab{}
	lab.SetAnnotations([]Annotation{
		{start: 0, end: 1, label: "test1"},
		{start: 1, end: 2, label: "test2"},
		{start: 2, end: 3, label: "test3"},
	})

	if lab.ToString() != "0 1 test1\n1 2 test2\n2 3 test3\n" {
		t.Fatal("annotations not set correctly!")
	}

	t.Log("setting annotations successful!")
}

func TestGettingAnnotations(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	if lab.annotations[0].start != 0 {
		t.Fatalf("incorrect annotation start time! wanted 0, recieved %f", lab.annotations[0].start)
	} else if lab.annotations[0].end != 10 {
		t.Fatalf("incorrect annotation end time! wanted 10, recieved %f", lab.annotations[0].end)
	} else if lab.annotations[0].label != "test" {
		t.Fatalf("incorrect annotation label!, wanted string \"test\", recieved string \"%s\"", lab.annotations[0].label)
	}

	t.Log("getting annotations successful!")
}

func TestPushingAnnotation(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	lab.PushAnnotation(Annotation{start: 20, end: 30, label: "new_annotation"})

	if lab.GetLength() != 3 {
		t.Fatalf("incorrect annotation index! wanted index length 3, recieved %d", lab.GetLength())
	} else if lab.annotations[2].start != 20 {
		t.Fatalf("incorrect annotation start time! wanted start time 20, recieved %f", lab.annotations[2].start)
	} else if lab.annotations[2].end != 30 {
		t.Fatalf("incorrect annotation end time! wanted end time 30, recieved %f", lab.annotations[2].end)
	}

	t.Log("pushing annotations successful!")
}

func TestAppendingAnnotations(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	lab.AppendAnnotations([]Annotation{
		{start: 20, end: 30, label: "new_annotation"},
		{start: 30, end: 40, label: "new_annotation2"},
	})

	if lab.GetLength() != 4 {
		t.Fatalf("incorrect annotation index! wanted index length of 4, recieved %d", lab.GetLength())
	} else if lab.annotations[2].start != 20 {
		t.Fatalf("incorrect annotation start time!")
	} else if lab.annotations[2].end != 30 {
		t.Fatalf("incorrect annotation end time!")
	} else if lab.annotations[2].label != "new_annotation" {
		t.Fatalf("incorrect annotation label!")
	} else if lab.annotations[3].start != 30 {
		t.Fatalf("incorrect annotation start time!")
	} else if lab.annotations[3].end != 40 {
		t.Fatalf("incorrect annotation end time!")
	} else if lab.annotations[3].label != "new_annotation2" {
		t.Fatalf("incorrect annotation label!")
	}

	t.Log("appending annotations successful!")
}

func TestClearingAnnotations(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	lab.ClearAnnotations()

	if lab.ToString() != "" {
		t.Fatal("lab not properly cleared!")
	}

	t.Log("clearing annotations successful!")
}

func TestDumpingLabels(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}
	labSlice := lab.DumpLabels()

	groundTruthSlice := []string{"test", "test2"}

	if isEqualSlice(labSlice, groundTruthSlice) == false {
		t.Fatal("returned slices not identical!")
	}

	t.Log("annotation dumping successful!")
}

func TestGettingLabName(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	if lab.GetName() != lab.name {
		t.Fatalf("wanted lab name %s, recieved %s", lab.name, lab.GetName())
	}

	t.Log("getting lab name successful!")
}

func TestSettingLabName(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	lab.SetName("newName")

	if lab.name != "newName" {
		t.Fatalf("wanted lab name of \"newName\", recieved \"%s\"", lab.name)
	}

	t.Log("setting lab name successful!")
}

func TestGettingPrecision(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	if lab.GetPrecision() != lab.precision {
		t.Fatalf("wanted precision %d, recieved %d", lab.precision, lab.GetPrecision())
	}

	t.Log("getting precision successful!")
}

func TestSettingPrecision(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	lab.SetPrecision(6)

	if lab.precision != 6 {
		t.Fatalf("wanted precision 6, recieved %d", lab.precision)
	}

	err = lab.WriteLab("examples/output3.lab", true)
	if err != nil {
		return
	}

	t.Log("setting precision successful!")
}

func TestGettingLabDuration(t *testing.T) {
	lab, err := ReadLab("examples/short.lab")
	if err != nil {
		t.Fatal(err)
	}

	trueDuration := lab.annotations[len(lab.annotations)-1].end

	if lab.GetDuration() != trueDuration {
		t.Fatalf("wanted duration of %f, recieved %f", trueDuration, lab.GetDuration())
	}

	t.Log("getting lab duration successful!")
}

func TestGettingLabLength(t *testing.T) {
	lab, err := ReadLab("examples/01.lab")
	if err != nil {
		t.Fatal(err)
	}

	trueLength := len(lab.annotations)

	if lab.GetLength() != trueLength {
		t.Fatalf("wanted length of %d, recieved %d", trueLength, lab.GetLength())
	}

	t.Log("getting lab length successful!")
}
