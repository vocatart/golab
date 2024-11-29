package textgrid

import (
	"fmt"
	"slices"
	"sort"
)

// Tier structs are an arbitrary type that can represent either IntervalTier or PointTier structs.
type Tier interface {
	GetType() string
	GetName() string
	GetDuration() float64
	GetXmin() float64
	GetXmax() float64
	SetName(string)
	SetXmin(float64, ...bool) error
	SetXmax(float64, ...bool) error
	SetIntervals([]Interval, ...bool) error
	GetIntervals() []Interval
	PushInterval(Interval, ...bool) error
	PushIntervals([]Interval, ...bool) error
	GetPoints() []Point
	SetPoints([]Point, ...bool) error
	PushPoint(Point, ...bool) error
	PushPoints([]Point, ...bool) error
	GetSize() int
	GetOverlapping() [][]int
	sort()
}

// IntervalTier is a Tier type that has a name, starting point, ending point, and Interval slice.
type IntervalTier struct {
	name      string
	xmin      float64
	xmax      float64
	intervals []Interval
}

// PointTier is a Tier type that has a name, starting point, ending point, and Point slice.
type PointTier struct {
	name   string
	xmin   float64
	xmax   float64
	points []Point
}

// GetType returns tier type for IntervalTier.
func (iTier *IntervalTier) GetType() string {
	return "IntervalTier"
}

// GetType returns tier type for IntervalTier.
func (pTier *PointTier) GetType() string {
	return "TextTier"
}

// GetDuration returns duration for IntervalTier.
func (iTier *IntervalTier) GetDuration() float64 {
	return iTier.xmax - iTier.xmin
}

// GetDuration returns duration for PointTier.
func (pTier *PointTier) GetDuration() float64 {
	return pTier.xmax - pTier.xmin
}

// GetName returns name for IntervalTier.
func (iTier *IntervalTier) GetName() string {
	return iTier.name
}

// GetName returns name for PointTier.
func (pTier *PointTier) GetName() string {
	return pTier.name
}

// GetXmin returns xmin for IntervalTier.
func (iTier *IntervalTier) GetXmin() float64 {
	return iTier.xmin
}

// GetXmin returns xmin for PointTier.
func (pTier *PointTier) GetXmin() float64 {
	return pTier.xmin
}

// GetXmax returns xmax for IntervalTier.
func (iTier *IntervalTier) GetXmax() float64 {
	return iTier.xmax
}

// GetXmax returns xmax for PointTier.
func (pTier *PointTier) GetXmax() float64 {
	return pTier.xmax
}

// SetName sets name for IntervalTier.
func (iTier *IntervalTier) SetName(name string) {
	iTier.name = name
}

// SetName sets name for PointTier.
func (pTier *PointTier) SetName(name string) {
	pTier.name = name
}

// SetXmin sets xmin for IntervalTier.
// By default, will return an error if you attempt to set an xmin that is smaller than the smallest xmin in IntervalTier.
func (iTier *IntervalTier) SetXmin(xmin float64, warn ...bool) error {

	// default value of warn is true
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	// if there is a value in the tier that is smaller than the xmin that is being set, return an error.
	var foundXmins []float64
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

// SetXmin sets xmin for PointTier.
// By default, will return an error if you attempt to set an xmin that is smaller than the smallest xmin in PointTier.
func (pTier *PointTier) SetXmin(xmin float64, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	var foundValues []float64
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

// SetXmax sets xmax for IntervalTier.
// By default, will return an error if you attempt to set an xmax that is larger than the largest xmax in IntervalTier.
func (iTier *IntervalTier) SetXmax(xmax float64, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	var foundXmaxs []float64
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

// SetXmax sets xmax for PointTier.
// By default, will return an error if you attempt to set an xmax that is larger than the largest xmax in PointTier.
func (pTier *PointTier) SetXmax(xmax float64, warn ...bool) error {

	if len(warn) == 0 {
		warn = append(warn, true)
	}

	var foundValues []float64
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

// GetSize returns the number of intervals in IntervalTier.
func (iTier *IntervalTier) GetSize() int {
	return len(iTier.intervals)
}

// GetSize returns the number of points in PointTier.
func (pTier *PointTier) GetSize() int {
	return len(pTier.points)
}

// GetIntervals returns Interval slice from IntervalTier.
func (iTier *IntervalTier) GetIntervals() []Interval {
	return iTier.intervals
}

// GetPoints implements Tier.
func (iTier *IntervalTier) GetPoints() []Point {
	panic("error: cannot return points from IntervalTier (type mismatch)")
}

// PushPoint implements Tier.
func (iTier *IntervalTier) PushPoint(point Point, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	return fmt.Errorf("error: cannot push point [%s, %s] to IntervalTier %q (type mismatch)", point.GetMark(), f2s(point.GetValue()), iTier.name)
}

// PushPoints implements Tier.
func (iTier *IntervalTier) PushPoints(points []Point, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	return fmt.Errorf("error: cannot push points starting with point [%s, %s] to IntervalTier %q (type mismatch)", points[0].GetMark(), f2s(points[0].GetValue()), iTier.name)
}

// GetPoints returns Point slice from PointTier.
func (pTier *PointTier) GetPoints() []Point {
	return pTier.points
}

// GetIntervals implements Tier.
func (pTier *PointTier) GetIntervals() []Interval {
	panic("error: cannot return intervals from PointTier (type mismatch)")
}

// PushInterval implements Tier.
func (pTier *PointTier) PushInterval(intervalTier Interval, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	return fmt.Errorf("error: cannot push interval [%s, %s, %s] to PointTier %q (type mismatch)", f2s(intervalTier.xmin), f2s(intervalTier.xmax), intervalTier.text, pTier.name)
}

// PushIntervals implements Tier.
func (pTier *PointTier) PushIntervals(intervals []Interval, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	return fmt.Errorf("error: cannot push intervals starting with [%s, %s, %s] to PointTier %q (type mismatch)", f2s(intervals[0].xmin), f2s(intervals[0].xmax), intervals[0].text, pTier.name)
}

// PushInterval appends an Interval to IntervalTier. Sorts Interval slice by all xmin values after pushing the Interval.
// By default, will return an error if you attempt to append an Interval that has a smaller xmin than the IntervalTier xmin.
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

// SetPoints sets Point slice field of PointTier.
func (pTier *PointTier) SetPoints(newPoints []Point, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	for _, point := range newPoints {
		if point.value < pTier.xmin && warn[0] || point.value > pTier.xmax && warn[0] {
			return fmt.Errorf("warning: you are trying to push point with value of %f but tier %s has existing xmin and xmax [%f, %f]", point.value, pTier.name, pTier.xmin, pTier.xmax)
		}
	}

	pTier.points = newPoints
	pTier.sort()

	return nil
}

// SetPoints implements Tier.
func (iTier *IntervalTier) SetPoints(newPoints []Point, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	return fmt.Errorf("error: cannot set points starting with [%s, %s] of IntervalTier %s (type mismatch)", f2s(newPoints[0].value), newPoints[0].mark, iTier.name)
}

// PushIntervals appends an Interval slice to IntervalTier. Sorts Interval slice by all xmin values after pushing the Interval slice.
// By default, will return an error if you attempt to append an Interval that has a smaller xmin than the IntervalTier xmin.
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

// SetIntervals sets intervals field of IntervalTier.
// By default, will return an error if there is an Interval that has an xmin or xmax not inside the TextGrid range.
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

// SetIntervals implements Tier.
func (pTier *PointTier) SetIntervals(intervals []Interval, warn ...bool) error {
	if len(warn) == 0 {
		warn = append(warn, true)
	}

	return fmt.Errorf("error: cannot set intervals starting with [%s, %s, %s] to PointTier %q (type mismatch)", f2s(intervals[0].xmin), f2s(intervals[0].xmax), intervals[0].text, pTier.name)
}

// GetOverlapping checks for overlaps in PointTier. Returns 2d slice of overlapping Point indices, or nil.
func (pTier *PointTier) GetOverlapping() [][]int {
	var overlaps [][]int

	// iterate over each pair of points, comparing the value to the next value (which should not be the same)
	for i, point := range pTier.points {
		if len(pTier.points) > i+1 {
			nextPoint := pTier.points[i+1]

			if point.value == nextPoint.value {
				overlaps = append(overlaps, []int{i, i + 1})
			}
		}
	}

	if overlaps != nil {
		return overlaps
	} else {
		return nil
	}
}

// GetOverlapping checks for overlaps in IntervalTier. Returns a 2d slide of overlapping interval indices, or nil.
func (iTier *IntervalTier) GetOverlapping() [][]int {
	var overlaps [][]int

	// iterate over each pair of intervals, comparing the xmax to the next xmin (which should be the same)
	for i, interval := range iTier.intervals {
		if len(iTier.intervals) > i+1 {
			nextInterval := iTier.intervals[i+1]

			if interval.xmax != nextInterval.xmin {
				overlaps = append(overlaps, []int{i, i + 1})
			}
		}
	}

	if overlaps != nil {
		return overlaps
	} else {
		return nil
	}
}

// PushPoint appends a Point to PointTier.
// By default, will return an error if you attempt to append a Point that has a value smaller than the PointTier xmin.
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

// PushPoints appends a Point slice to PointTier.
// By default, will return an error if you attempt to append a Point that has a value smaller than the PointTier xmin.
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

// sort reorders IntervalTier Interval slice by Interval xmin field.
func (iTier *IntervalTier) sort() {
	sort.Slice(iTier.intervals, func(i, j int) bool {
		return iTier.intervals[i].xmin < iTier.intervals[j].xmin
	})
}

// sort reorders PointTier Point slice by Point value field.
func (pTier *PointTier) sort() {
	sort.Slice(pTier.points, func(i, j int) bool {
		return pTier.points[i].value < pTier.points[j].value
	})
}
