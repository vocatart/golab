package textgrid

import "testing"

func TestCreatingPoint(t *testing.T) {
	point := Point{mark: "point", value: 10.0}

	if point.value != 10.0 {
		t.Errorf("Expected 10.0, got %f", point.value)
	}
	if point.mark != "point" {
		t.Errorf("Expected 'mark', got %s", point.mark)
	}

	if !t.Failed() {
		t.Logf("point created successfully")
	}
}
