package textgrid

import (
	"log"
	"os"
	"strings"
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

// Gets tiers of TextGrid
func (tg TextGrid) GetTiers() []Tier {
	return tg.tiers
}

// Takes a path to a .TextGrid file and reads its contents into a TextGrid.
func ReadTextgrid(path string) TextGrid {
	tg := TextGrid{}

	// check if the file exists
	// reading into memory is alot easier than line-by-line with textgrids
	tgData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	tgContent := processContent(tgData)

	if !verifyHead(tgContent) {
		log.Fatal("malformed textgrid head!")
	}

	return tg

}

// turns textgrid file content into a slice of useable strings
func processContent(data []byte) []string {
	// remove all empty lines
	tgString := strings.ReplaceAll(string(data), "\n\n", "\n")

	var sanitized strings.Builder
	inQuotes := false

	// remove all spaces unless the space is between quotations (part of a label)
	for i := 0; i < len(tgString); i++ {
		char := tgString[i]

		if char == '"' {
			inQuotes = !inQuotes
		}

		if inQuotes || char != ' ' {
			sanitized.WriteByte(char)
		}
	}

	return strings.Split(sanitized.String(), "\n")
}

// verifies the existence and content of FileType and ObjectClass
func verifyHead(tgContent []string) bool {
	fileType := tgContent[0]
	objectClass := tgContent[1]

	if !strings.Contains(fileType, "ooTextFile") {
		log.Println("malformed fileType")
		return false
	}

	if !strings.Contains(objectClass, "TextGrid") {
		log.Println("malformed objectClass")
		return false
	}

	return true
}
