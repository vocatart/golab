package textgrid

import "testing"

func TestCreatingInterval(t *testing.T) {
	interval := Interval{xmin: 0.0, xmax: 1.0, text: "test"}

	// testing accessors
	if interval.GetXmin() != 0.0 {
		t.Errorf("Expected xmin 0.0, got %f", interval.xmin)
	}
	if interval.GetXmax() != 1.0 {
		t.Errorf("Expected xmax 1.0, got %f", interval.xmax)
	}
	if interval.GetText() != "test" {
		t.Errorf("Expected interval label \"test\", got %q", interval.text)
	}

	// testing mutators
	interval.SetXmin(2.0)
	interval.SetXmax(3.0)
	interval.SetText("test2")

	if interval.GetXmin() != 2.0 {
		t.Errorf("Expected xmin 2.0, got %f", interval.xmin)
	}
	if interval.GetXmax() != 3.0 {
		t.Errorf("Expected xmax 3.0, got %f", interval.xmax)
	}
	if interval.GetText() != "test2" {
		t.Errorf("Expected interval label \"test2\", got %q", interval.text)
	}

	// testing misc methods
	if interval.GetDuration() != 1.0 {
		t.Errorf("Expected dur of 1.0, got %f", interval.GetDuration())
	}
	if interval.GetMedian() != 2.5 {
		t.Errorf("Expected median of 2.5, got %f", interval.GetMedian())
	}
}
