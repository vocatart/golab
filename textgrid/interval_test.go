package textgrid

import "testing"

func TestCreatingInterval(t *testing.T) {
	interval := Interval{xmin: 0.0, xmax: 1.0, text: "test"}

	if interval.xmin != 0.0 {
		t.Errorf("Expected 0.0, got %f", interval.xmin)
	}
	if interval.xmax != 1.0 {
		t.Errorf("Expected 1.0, got %f", interval.xmax)
	}
	if interval.text != "test" {
		t.Errorf("Expected 'test', got %s", interval.text)
	}

	if !t.Failed() {
		t.Logf("interval created successfully")
	}
}
