package textgrid

import (
	"testing"
)

func TestCreatingIntervalTier(t *testing.T) {
	intervals := []Interval{
		{0.0, 1.0, "test1"},
		{1.0, 2.0, "test2"},
		{2.0, 3.0, "test3"},
	}

	singlePushInterval := Interval{3.0, 4.0, "test4"}

	multiplePushIntervals := []Interval{
		{4.0, 5.0, "test5"},
		{6.0, 7.0, "test7"},
		{5.0, 6.0, "test6"},
	}

	tier := IntervalTier{name: "TestTier", xmin: 0.0, xmax: 3.0, intervals: intervals}

	// testing accessors
	if tier.GetName() != "TestTier" {
		t.Errorf("Expected 'TestTier', got %s", tier.name)
	}
	if tier.GetXmin() != 0.0 {
		t.Errorf("Expected 0.0, got %f", tier.xmin)
	}
	if tier.GetXmax() != 3.0 {
		t.Errorf("Expected 3.0, got %f", tier.xmax)
	}
	if tier.GetSize() != 3 {
		t.Errorf("Expected 3 intervals, got %d", tier.GetSize())
	}
	if tier.GetType() != "IntervalTier" {
		t.Errorf("Expected 'IntervalTier', got %s", tier.GetType())
	}
	if tier.GetDuration() != 3.0 {
		t.Errorf("Expected 3.0, got %f", tier.GetDuration())
	}

	for _, interval := range tier.GetIntervals() {
		if interval.GetText() != "test1" && interval.GetText() != "test2" && interval.GetText() != "test3" {
			t.Errorf("Expected 'test1', 'test2', or 'test3', got %s", interval.GetText())
		}

		if interval.GetXmin() != 0.0 && interval.GetXmin() != 1.0 && interval.GetXmin() != 2.0 {
			t.Errorf("Expected 0.0, 1.0, or 2.0, got %f", interval.GetXmin())
		}

		if interval.GetXmax() != 1.0 && interval.GetXmax() != 2.0 && interval.GetXmax() != 3.0 {
			t.Errorf("Expected 1.0, 2.0, or 3.0, got %f", interval.GetXmax())
		}
	}

	err := tier.PushInterval(singlePushInterval)
	if err != nil {
		t.Error(err)
	}

	err = tier.PushIntervals(multiplePushIntervals)
	if err != nil {
		t.Error(err)
	}

	// should be fully sorted
	t.Logf("Interval Tier: %v\n", tier)
}

func TestOverlapping(t *testing.T) {
	overlappingIntervalTier := IntervalTier{
		name: "OverlappingIntervalTier",
		xmin: 0,
		xmax: 10.0,
		intervals: []Interval{
			{0.0, 1.0, "test1"},
			{0.5, 1.0, "test2"},
		},
	}

	overlappingPointTier := PointTier{
		name: "OverlappingPointTier",
		xmin: 0.0,
		xmax: 10.0,
		points: []Point{
			{5.0, "test3"},
			{5.0, "test4"},
		},
	}

	// should be [0, 1]
	t.Logf("Overlapping Interval Indicies: %v\n", overlappingIntervalTier.GetOverlapping())
	t.Logf("Overlapping Point Indicies: %v\n", overlappingPointTier.GetOverlapping())
}

func TestCreatingPointTier(t *testing.T) {
	points := []Point{
		{0.0, "test1"},
		{1.0, "test2"},
		{2.0, "test3"},
	}

	singlePushPoint := Point{3.0, "test4"}

	multiplePushPoints := []Point{
		{4.0, "test5"},
		{6.0, "test7"},
		{5.0, "test6"},
	}

	tier := PointTier{name: "TestTier", xmin: 0.0, xmax: 3.0, points: points}

	if tier.GetName() != "TestTier" {
		t.Errorf("Expected 'TestTier', got %s", tier.name)
	}
	if tier.GetXmin() != 0.0 {
		t.Errorf("Expected 0.0, got %f", tier.xmin)
	}
	if tier.GetXmax() != 3.0 {
		t.Errorf("Expected 3.0, got %f", tier.xmax)
	}
	if tier.GetSize() != 3 {
		t.Errorf("Expected 3 intervals, got %d", tier.GetSize())
	}

	for _, point := range tier.GetPoints() {
		if point.GetMark() != "test1" && point.GetMark() != "test2" && point.GetMark() != "test3" {
			t.Errorf("Expected 'test1', 'test2', or 'test3', got %s", point.GetMark())
		}

		if point.GetValue() != 0.0 && point.GetValue() != 1.0 && point.GetValue() != 2.0 {
			t.Errorf("Expected 0.0, 1.0, or 2.0, got %f", point.GetValue())
		}
	}

	err := tier.PushPoint(singlePushPoint)
	if err != nil {
		t.Error(err)
	}

	err = tier.PushPoints(multiplePushPoints)
	if err != nil {
		t.Error(err)
	}

	// should be fully sorted
	t.Logf("Point Tier: %v\n", tier)
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
		t.Errorf(err.Error())
	}

	if tier.GetSize() != 3 {
		t.Errorf("Expected 3 intervals, got %d", tier.GetSize())
	}

	for _, interval := range tier.GetIntervals() {
		if interval.GetText() != "test1" && interval.GetText() != "test2" && interval.GetText() != "test3" {
			t.Errorf("Expected 'test1', 'test2', or 'test3', got %s", interval.GetText())
		}

		if interval.GetXmin() != 0.0 && interval.GetXmin() != 1.0 && interval.GetXmin() != 2.0 {
			t.Errorf("Expected 0.0, 1.0, or 2.0, got %f", interval.GetXmin())
		}

		if interval.GetXmax() != 1.0 && interval.GetXmax() != 2.0 && interval.GetXmax() != 3.0 {
			t.Errorf("Expected 1.0, 2.0, or 3.0, got %f", interval.GetXmax())
		}
	}
}

func TestSettingPoints(t *testing.T) {
	points := []Point{
		{0.0, "test1"},
		{1.0, "test2"},
		{2.0, "test3"},
	}

	tier := PointTier{name: "TestTier", xmin: 0.0, xmax: 3.0}

	err := tier.SetPoints(points)
	if err != nil {
		t.Errorf(err.Error())
	}

	if tier.GetSize() != 3 {
		t.Errorf("Expected 3 intervals, got %d", tier.GetSize())
	}

	for _, point := range tier.GetPoints() {
		if point.GetMark() != "test1" && point.GetMark() != "test2" && point.GetMark() != "test3" {
			t.Errorf("Expected 'test1', 'test2', or 'test3', got %s", point.GetMark())
		}

		if point.GetValue() != 0.0 && point.GetValue() != 1.0 && point.GetValue() != 2.0 {
			t.Errorf("Expected 0.0, 1.0, or 2.0, got %f", point.GetValue())
		}
	}
}
