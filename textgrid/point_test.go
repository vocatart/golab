package textgrid

import "testing"

func TestCreatingPoint(t *testing.T) {
	point := Point{mark: "point", value: 10.0}

	// testing accessors
	if point.GetValue() != 10.0 {
		t.Errorf("Expected 10.0, got %f", point.value)
	}
	if point.GetMark() != "point" {
		t.Errorf("Expected 'mark', got %s", point.mark)
	}

	// testing mutators
	point.SetValue(15.0)
	point.SetMark("test")

	if point.GetValue() != 15.0 {
		t.Errorf("Expected 15.0, got %f", point.value)
	}
	if point.GetMark() != "test" {
		t.Errorf("Expected 'mark', got %s", point.mark)
	}
}
