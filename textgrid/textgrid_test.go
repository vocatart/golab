package textgrid

import "testing"

func TestCreatingTextGrid(t *testing.T) {
	tg := TextGrid{
		xmin: 0,
		xmax: 10.0,
		tiers: []Tier{&IntervalTier{
			name: "IntervalTier",
			xmin: 0,
			xmax: 10.0,
			intervals: []Interval{{
				xmin: 0,
				xmax: 10.0,
				text: "Interval",
			}},
		}, &PointTier{
			name: "PointTier",
			xmin: 0,
			xmax: 10.0,
			points: []Point{{
				value: 5.0,
				mark:  "Point",
			}},
		}},
		name: "TextGrid",
	}

	newTiers := []Tier{&PointTier{
		name: "NewPointTier",
		xmin: 10.0,
		xmax: 20.0,
		points: []Point{{
			value: 15.0,
			mark:  "NewPoint",
		}},
	}, &IntervalTier{
		name: "NewIntervalTier",
		xmin: 10.0,
		xmax: 20.0,
		intervals: []Interval{{
			xmin: 10.0,
			xmax: 20.0,
			text: "NewInterval",
		}},
	}}

	// testing accessors
	if tg.GetXmin() != 0 {
		t.Errorf("expected xmin 0, got %f", tg.GetXmin())
	}
	if tg.GetXmax() != 10.0 {
		t.Errorf("expected xmax 10.0, got %f", tg.GetXmax())
	}
	if tg.GetName() != "TextGrid" {
		t.Errorf("expected name \"TextGrid\", got %q", tg.GetName())
	}
	if tg.GetTiers()[0].GetType() != "IntervalTier" {
		t.Errorf("expected type \"IntervalTier\", got %q", tg.GetTiers()[0].GetType())
	}
	if tg.GetTiers()[1].GetType() != "TextTier" {
		t.Errorf("expected type \"TextTier\", got %q", tg.GetTiers()[1].GetType())
	}

	// testing misc methods
	if !tg.HasIntervalTier() {
		t.Errorf("expected HasIntervalTier to be true")
	}
	if !tg.HasPointTier() {
		t.Errorf("expected HasPointTier to be true")
	}

	if tg.GetTier("IntervalTier").GetType() != "IntervalTier" {
		t.Errorf("\"IntervalTier\" incorrectly retrieved")
	}
	if tg.GetTier("PointTier").GetType() != "TextTier" {
		t.Errorf("\"TextTier\" incorrectly retrieved")
	}
	if tg.TierAtIndex(0).GetType() != "IntervalTier" {
		t.Errorf("\"IntervalTier\" incorrectly retrieved")
	}
	if tg.TierAtIndex(1).GetType() != "TextTier" {
		t.Errorf("\"TextTier\" incorrectly retrieved")
	}

	if tg.GetSize() != 2 {
		t.Errorf("expected size 2, got %d", tg.GetSize())
	}

	// testing mutators
	tg.SetXmin(10.0)
	tg.SetXmax(20.0)
	tg.SetName("NewTextGrid")
	tg.SetTiers(newTiers)

	if tg.GetXmin() != 10.0 {
		t.Errorf("expected xmin 10.0, got %f", tg.GetXmin())
	}
	if tg.GetXmax() != 20.0 {
		t.Errorf("expected xmax 20.0, got %f", tg.GetXmax())
	}
	if tg.GetName() != "NewTextGrid" {
		t.Errorf("expected name NewTextGrid, got %q", tg.GetName())
	}
	if tg.GetTiers()[0].GetType() != "TextTier" {
		t.Errorf("expected type \"TextTier\", got %q", tg.GetTiers()[0].GetType())
	}
	if tg.GetTiers()[1].GetType() != "IntervalTier" {
		t.Errorf("expected type \"IntervalTier\", got %q", tg.GetTiers()[1].GetType())
	}
}

func TestReadingTextgridASCIILong(t *testing.T) {
	_, err := ReadTextgrid("examples/long.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestReadingTextgridASCIIShort(t *testing.T) {
	_, err := ReadTextgrid("examples/short.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestReadingTextgridUTF16(t *testing.T) {
	_, err := ReadTextgrid("examples/polish64.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestReadingTextgridUTF8(t *testing.T) {
	_, err := ReadTextgrid("examples/polish65.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestWritingLongTextgrid(t *testing.T) {
	tg, err := ReadTextgrid("examples/long.TextGrid")
	if err != nil {
		t.Error(err)
	}

	err = tg.WriteLong("examples/long_output.TextGrid", true)
	if err != nil {
		t.Error(err)
	}
}

func TestWritingShortTextgrid(t *testing.T) {
	tg, err := ReadTextgrid("examples/short.TextGrid")
	if err != nil {
		t.Error(err)
	}

	err = tg.WriteShort("examples/short_output.TextGrid", true)
	if err != nil {
		t.Error(err)
	}
}
