package textgrid

import (
	"fmt"
	"slices"
	"sort"
)

// A tier is an arbitrary type that can hold either intervals or points.
type Tier interface {
	TierType() string
	TierName() string
	TierDuration() float64
	TierXmin() float64
	TierXmax() float64
	SetName(string)
	SetXmin(float64, ...bool) error
	SetXmax(float64, ...bool) error
	GetSize() int
	sort()
}

// An interval tier holds a collection of intervals, along with a tier name and total duration.
type IntervalTier struct {
	name      string
	xmin      float64
	xmax      float64
	intervals []Interval
}

// A point tier holds a collection of points, along with a tier name and total duration.
type PointTier struct {
	name   string
	xmin   float64
	xmax   float64
	points []Point
}

// TODO: make these safer

// Returns tier type for `IntervalTier`
func (iTier IntervalTier) TierType() string {
	return "IntervalTier"
}

// Returns tier type for `PointTier`
func (pTier PointTier) TierType() string {
	return "PointTier"
}

// Returns tier duration for `IntervalTier`
func (iTier IntervalTier) TierDuration() float64 {
	return iTier.xmax - iTier.xmin
}

// Returns tier duration for `PointTier`
func (pTier PointTier) TierDuration() float64 {
	return pTier.xmax - pTier.xmin
}

// Returns tier name for `IntervalTier`
func (iTier IntervalTier) TierName() string {
	return iTier.name
}

// Returns tier name for `PointTier`
func (pTier PointTier) TierName() string {
	return pTier.name
}

// Returns xmin for `IntervalTier`
func (iTier IntervalTier) TierXmin() float64 {
	return iTier.xmin
}

// Returns xmin for `PointTier`
func (pTier PointTier) TierXmin() float64 {
	return pTier.xmin
}

// Returns xmax for `IntervalTier`
func (iTier IntervalTier) TierXmax() float64 {
	return iTier.xmax
}

// Returns the xmax for `PointTier`
func (pTier PointTier) TierXmax() float64 {
	return pTier.xmax
}

// Sets name for tier `IntervalTier`
func (iTier *IntervalTier) SetName(name string) {
	iTier.name = name
}

// Sets name for tier `PointTier`
func (pTier *PointTier) SetName(name string) {
	pTier.name = name
}

// Sets xmin for tier `IntervalTier`
func (iTier *IntervalTier) SetXmin(xmin float64, warn ...bool) error {

	// default value of warn is true
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	// if there is a value in the tier that is smaller than the xmin that is being set, return an error.
	foundXmins := []float64{}
	for _, interval := range iTier.intervals {
		foundXmins = append(foundXmins, interval.xmin)
	}

	smallestXmin := slices.Min(foundXmins)

	if warn[0] && smallestXmin > xmin {
		xminError := fmt.Errorf("warning: you are trying to set xmin %f when xmin %f exists in current interval tier", xmin, smallestXmin)
		return xminError
	}

	iTier.xmin = xmin
	return nil
}

// Sets xmin for tier `PointTier`
func (pTier *PointTier) SetXmin(xmin float64, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	foundValues := []float64{}
	for _, point := range pTier.points {
		foundValues = append(foundValues, point.value)
	}

	smallestValue := slices.Min(foundValues)

	if warn[0] && smallestValue > xmin {
		xminError := fmt.Errorf("warning: you are trying to set xmin of %f when point value %f exists in current point tier", xmin, smallestValue)
		return xminError
	}

	pTier.xmin = xmin
	return nil
}

// Sets xmax for tier `IntervalTier`
func (iTier *IntervalTier) SetXmax(xmax float64, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	foundXmaxs := []float64{}
	for _, interval := range iTier.intervals {
		foundXmaxs = append(foundXmaxs, interval.xmax)
	}

	largestXmax := slices.Max(foundXmaxs)

	if warn[0] && largestXmax < xmax {
		xmaxError := fmt.Errorf("warning: you are trying to set xmax of %f when xmax %f exists in current interval tier", xmax, largestXmax)
		return xmaxError
	}

	iTier.xmax = xmax
	return nil
}

// Sets xmax for tier `PointTier`
func (pTier *PointTier) SetXmax(xmax float64, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	foundValues := []float64{}
	for _, point := range pTier.points {
		foundValues = append(foundValues, point.value)
	}

	largestValue := slices.Max(foundValues)

	if warn[0] && largestValue < xmax {
		return fmt.Errorf("warning: you are trying to set xmax of %f when point value %f exists in current point tier", xmax, largestValue)
	}

	pTier.xmax = xmax
	return nil
}

// Returns the size (number of intervals) of `IntervalTier`
func (iTier IntervalTier) GetSize() int {
	return len(iTier.intervals)
}

// Returns the size (number of points) of `PointTier`
func (pTier PointTier) GetSize() int {
	return len(pTier.points)
}

// Returns slice of intervals from an `IntervalTier“
func (iTier IntervalTier) GetIntervals() []Interval {
	return iTier.intervals
}

// Returns slice of points from a `PointTier`
func (pTier PointTier) GetPoints() []Point {
	return pTier.points
}

// Pushes an interval to the interval tier. Sorts intervals by minimum x value after pushing the interval.
func (iTier *IntervalTier) PushInterval(intervalPush Interval, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	if warn[0] && intervalPush.xmin < iTier.xmin {
		return fmt.Errorf("warning: you are trying to push interval with xmin of %f but tier %s has existing xmin of %f", intervalPush.xmin, iTier.name, iTier.xmin)
	}

	iTier.intervals = append(iTier.intervals, intervalPush)
	iTier.sort()

	return nil
}

// Pushes a slice of intervals to the interval tier. SOrts intervals by minimum x value after pushing the intervals.
func (iTier *IntervalTier) PushIntervals(intervalsPush []Interval, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	if warn[0] {
		for _, interval := range intervalsPush {
			if interval.xmin < iTier.xmin {
				return fmt.Errorf("warning: you are trying to push interval with xmin of %f but tier %s has existing xmin of %f", interval.xmin, iTier.name, iTier.xmin)
			}
		}
	}

	iTier.intervals = append(iTier.intervals, intervalsPush...)
	iTier.sort()

	return nil
}

// Sets all intervals of an interval tier. Checks for xmin and xmax mismatch, and sorts the modified interval tier.
func (iTier *IntervalTier) SetIntervals(newIntervals []Interval, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	if warn[0] {
		for _, interval := range newIntervals {
			if interval.xmin < iTier.xmin {
				return fmt.Errorf("warning: you are trying to push interval with xmin of %f but tier %s has existing xmin of %f", interval.xmin, iTier.name, iTier.xmin)
			} else if interval.xmax > iTier.xmax {
				return fmt.Errorf("warning: you are trying to push interval with xmax of %f but tier %s has existing xmax of %f", interval.xmax, iTier.name, iTier.xmax)
			}
		}
	}

	iTier.intervals = newIntervals
	iTier.sort()

	return nil
}

// Checks for overlaps in an interval tier. Returns a 2d slide of overlapping interval indicies, or nil.
func (iTier IntervalTier) GetOverlapping() [][]int {

	overlaps := [][]int{}

	// iterate over each pair of intervals, comparing the xmax to the next xmin (which should be the same)
	for i, interval := range iTier.intervals {
		nextInterval := iTier.intervals[i+1]

		if interval.xmax != nextInterval.xmin {
			overlaps = append(overlaps, []int{i, i + 1})
		}
	}

	if overlaps != nil {
		return overlaps
	} else {
		return nil
	}
}

func (pTier *PointTier) PushPoint(pointPush Point, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	if pointPush.value < pTier.xmin && warn[0] {
		return fmt.Errorf("warning: you are trying to push a point with value %f when tier %s has xmin of %f", pointPush.value, pTier.name, pTier.xmin)
	}

	pTier.points = append(pTier.points, pointPush)
	pTier.sort()

	return nil
}

func (pTier *PointTier) PushPoints(pointsPush []Point, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	if warn[0] {
		for i, point := range pointsPush {
			if point.value < pTier.xmin {
				return fmt.Errorf("warning: you are trying to push a point with value %f when tier %s has xmin of %f", pointsPush[i].value, pTier.name, pTier.xmin)
			}
		}
	}

	pTier.points = append(pTier.points, pointsPush...)
	pTier.sort()

	return nil
}

// reorder interval tier by xmins
func (iTier *IntervalTier) sort() {
	sort.Slice(iTier.intervals, func(i, j int) bool {
		return iTier.intervals[i].xmin < iTier.intervals[j].xmin
	})
}

// reorder point tier by point values
func (pTier *PointTier) sort() {
	sort.Slice(pTier.points, func(i, j int) bool {
		return pTier.points[i].value < pTier.points[j].value
	})
}
