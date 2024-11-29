package htk

import "testing"

func TestCreatingAnnotation(t *testing.T) {
	annotation := Annotation{0.0, 10.0, "test"}

	if annotation.GetStart() != 0.0 {
		t.Errorf("expected 0.0, got %f", annotation.GetStart())
	}
	if annotation.GetEnd() != 10.0 {
		t.Errorf("expected 10.0, got %f", annotation.GetEnd())
	}
	if annotation.GetLabel() != "test" {
		t.Errorf("expected test, got %s", annotation.GetLabel())
	}
	if annotation.GetDuration() != 10.0 {
		t.Errorf("expected 10.0, got %f", annotation.GetDuration())
	}

	annotation.SetStart(15.0)
	annotation.SetEnd(20.0)
	annotation.SetLabel("test2")

	if annotation.GetStart() != 15.0 {
		t.Errorf("expected 15.0, got %f", annotation.GetStart())
	}
	if annotation.GetEnd() != 20.0 {
		t.Errorf("expected 20.0, got %f", annotation.GetEnd())
	}
	if annotation.GetLabel() != "test2" {
		t.Errorf("expected test2, got %s", annotation.GetLabel())
	}
}
