package textgrid

import "testing"

func TestCreatingIntervalTier(t *testing.T) {
	intervals := []Interval{
		{0.0, 1.0, "test1"},
		{1.0, 2.0, "test2"},
		{2.0, 3.0, "test3"},
	}

	tier := IntervalTier{name: "TestTier", xmin: 0.0, xmax: 3.0, intervals: intervals}

	if tier.name != "TestTier" {
		t.Errorf("Expected 'TestTier', got %s", tier.name)
	}
	if tier.xmin != 0.0 {
		t.Errorf("Expected 0.0, got %f", tier.xmin)
	}
	if tier.xmax != 3.0 {
		t.Errorf("Expected 3.0, got %f", tier.xmax)
	}
	if len(tier.intervals) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(tier.intervals))
	}

	for _, interval := range tier.intervals {
		if interval.text != "test1" && interval.text != "test2" && interval.text != "test3" {
			t.Errorf("Expected 'test1', 'test2', or 'test3', got %s", interval.text)
		}

		if interval.xmin != 0.0 && interval.xmin != 1.0 && interval.xmin != 2.0 {
			t.Errorf("Expected 0.0, 1.0, or 2.0, got %f", interval.xmin)
		}

		if interval.xmax != 1.0 && interval.xmax != 2.0 && interval.xmax != 3.0 {
			t.Errorf("Expected 1.0, 2.0, or 3.0, got %f", interval.xmax)
		}
	}

	if !t.Failed() {
		t.Logf("interval tier created successfully")
	}
}

func TestSettingIntervals(t *testing.T) {
	intervals := []Interval{
		{0.0, 1.0, "test1"},
		{1.0, 2.0, "test2"},
		{2.0, 3.0, "test3"},
	}

	tier := IntervalTier{name: "TestTier", xmin: 0.0, xmax: 3.0}

	err := tier.SetIntervals(intervals)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(tier.intervals) != 3 {
		t.Errorf("Expected 3 intervals, got %d", len(tier.intervals))
	}

	for _, interval := range tier.intervals {
		if interval.text != "test1" && interval.text != "test2" && interval.text != "test3" {
			t.Errorf("Expected 'test1', 'test2', or 'test3', got %s", interval.text)
		}

		if interval.xmin != 0.0 && interval.xmin != 1.0 && interval.xmin != 2.0 {
			t.Errorf("Expected 0.0, 1.0, or 2.0, got %f", interval.xmin)
		}

		if interval.xmax != 1.0 && interval.xmax != 2.0 && interval.xmax != 3.0 {
			t.Errorf("Expected 1.0, 2.0, or 3.0, got %f", interval.xmax)
		}
	}

	if !t.Failed() {
		t.Logf("intervals set successfully")
	}
}
