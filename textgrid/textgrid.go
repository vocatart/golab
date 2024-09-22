package textgrid

import (
	"log"
	"os"
)

type TextGrid struct {
	xmin  float64
	xmax  float64
	tiers []Tier
	name  string
}

// Returns xmin of TextGrid
func (tg TextGrid) GetXmin() float64 {
	return tg.xmin
}

// Sets xmin of TextGrid
func (tg *TextGrid) SetXmin(xmin float64) {
	tg.xmin = xmin
}

// Returns xmax of TextGrid
func (tg TextGrid) GetXmax() float64 {
	return tg.xmax
}

// Sets xmax of TextGrid
func (tg *TextGrid) SetXmax(xmax float64) {
	tg.xmax = xmax
}

// Returns name of TextGrid
func (tg TextGrid) GetName() string {
	return tg.name
}

// Sets name of TextGrid
func (tg *TextGrid) SetName(name string) {
	tg.name = name
}

func (tg TextGrid) GetTiers() []Tier {
	return tg.tiers
}

func ReadTextgrid(path string) TextGrid {
	tg := TextGrid{}

	tgData, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer tgData.Close()

}
