package textgrid

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
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
	tgDeque := deque.New[string]()

	// grab the name element from the path
	tg.name = filepath.Base(path)

	// check if the file exists
	// reading into memory is alot easier than line-by-line with textgrids
	tgData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// process the textgrid data into a slice of machine-friendly strings
	tgContent := processContent(tgData)

	// convert string slice into deque
	for _, str := range tgContent {
		tgDeque.PushBack(str)
	}

	// verify the first two entries in the deque
	if !verifyHead(tgDeque) {
		log.Fatal("malformed textgrid at head!")
	}

	// pop the next two values, which should be xmin and xmax respectively.
	globalXmin := pullFloat(tgDeque.PopFront())
	globalXmax := pullFloat(tgDeque.PopFront())

	// set the xmin and xmax preemptivley incase the status is <absent>
	tg.xmin = globalXmin
	tg.xmax = globalXmax

	// the next value is the status, which should either be <absent> or <exists>
	tierStatus := pullBracketedValue(tgDeque.PopFront())
	if tierStatus == "absent" {
		log.Println("warning: tierStatus is <absent>, a textgrid with 0 tiers will be returned")
		return tg
	} else if tierStatus != "exists" {
		log.Fatalf("error: expected tier status <exists> or <absent>, recieved <%s>", tierStatus)
	}

	// get the number of tiers that exist in this textgrid
	numTiers := pullInt(tgDeque.PopFront())

	tg.tiers = parseTiers(globalXmin, globalXmax, tgDeque, numTiers)

	return tg
}

// TODO: Implement
// func (tg TextGrid) WriteLong(path string, overwrite ...bool) {
// }

func parseTiers(globalXmin float64, globalXmax float64, content *deque.Deque[string], numTiers int) []Tier {
	tiers := []Tier{}
	tierCounter := 0

	for tierCounter < numTiers {
		// at the start of a tier, the first two values will be tierType and tierName
		tierType := pullQuotedValue(content.PopFront())
		tierName := pullQuotedValue(content.PopFront())

		// the next two values should be the xmin and xmax of the unique tier
		tierXmin := pullFloat(content.PopFront())
		tierXmax := pullFloat(content.PopFront())

		// check to see if any boundaries are inconsistent
		if tierXmin < globalXmin {
			log.Fatalf("error: %s %s has xmin %f, when TextGrid xmin is %f", tierType, tierName, tierXmin, globalXmin)
		}
		if tierXmax > globalXmax {
			log.Fatalf("error: %s %s has xmax %f, when TextGrid xmax is %f", tierType, tierName, tierXmax, globalXmax)
		}

		// the last value before the intervals/points begin should be the number of intervals/points in the unique tier
		tierContentCount := pullInt(content.PopFront())
		contentCounter := 0

		// loop for each tier type
		if tierType == "IntervalTier" {
			intervals := []Interval{}

			for contentCounter != tierContentCount {
				// the next three values in an interval are the xmin, xmax, and text
				intervalXmin := pullFloat(content.PopFront())
				intervalXmax := pullFloat(content.PopFront())
				intervalText := pullQuotedValue(content.PopFront())

				// create the new interval
				newInterval := Interval{xmin: intervalXmin, xmax: intervalXmax, text: intervalText}
				intervals = append(intervals, newInterval)

				contentCounter++
			}

			// once we've gotten all the intervals, create the new tier and push it to the tiers slice.
			newIntervalTier := IntervalTier{name: tierName, xmin: tierXmin, xmax: tierXmax, intervals: intervals}
			tiers = append(tiers, &newIntervalTier)

		} else if tierType == "TextTier" { // point tier
			points := []Point{}

			for contentCounter != tierContentCount {
				// the next two values in a point are the value and the mark
				pointValue := pullFloat(content.PopFront())
				pointMark := pullQuotedValue(content.PopFront())

				// create the new point
				newPoint := Point{value: pointValue, mark: pointMark}
				points = append(points, newPoint)

				contentCounter++
			}

			// once we've gotten all the points, create the new tier and push it to the tiers slice.
			newPointTier := PointTier{name: tierName, xmin: tierXmin, xmax: tierXmax, points: points}
			tiers = append(tiers, &newPointTier)
		} else {
			log.Fatalf("error: unexpected tier type %s", tierType)
		}
		tierCounter++
	}

	return tiers
}

// turns textgrid file content into a slice of useable strings.
// internally, any textgrid given is converted into a "short" type textgrid with any empty lines removed
func processContent(data []byte) []string {
	// remove all empty lines
	tgString := strings.ReplaceAll(string(data), "\n\n", "\n")

	bracketRegex := regexp.MustCompile(`\[\d+\]`)

	// a short textgrid is basically a textgrid that is only labels, numbers, and flags.
	// we will use regex to remove everything that isnt needed by praat to recognize a textgrid.
	// `\d+(\.\d+)?` matches all floats and integers
	// `\"[^\"]*\` matches all content inbetween double quotes
	// `\<[^>]*>` matches all content inbetween angle brackets
	textgridRegex := regexp.MustCompile(`(\d+(\.\d+)?|\"[^\"]*\"|<[^>]*>)`)

	tgSanitized := strings.Join(textgridRegex.FindAllString(bracketRegex.ReplaceAllString(tgString, ""), -1), "\n")

	return strings.Split(tgSanitized, "\n")
}

// checks the nessecary `File Type` and `Object Class` fields of a TextGrid file
func verifyHead(tgContent *deque.Deque[string]) bool {
	fileType := tgContent.PopFront()
	objectClass := tgContent.PopFront()

	if pullQuotedValue(fileType) != "ooTextFile" {
		log.Printf("error: wanted fileType ooTextFile, recieved %s", fileType)
		return false
	}

	if pullQuotedValue(objectClass) != "TextGrid" {
		log.Printf("error: wanted objectClass TextGrid, recieved %s", objectClass)
		return false
	}

	return true
}

// takes a value contained in quotes and returns it without quotes.
// quotes inside quotes will be preserved.
func pullQuotedValue(str string) string {
	stringRegex := regexp.MustCompile(`^"(.*)"$`)
	return stringRegex.ReplaceAllString(str, `$1`)
}

// takes a value that has angle brackets and removes them.
// this is only used for flags in textgrids, so we don't care about internal brackets.
func pullBracketedValue(str string) string {
	stringRegex := regexp.MustCompile(`[<>]`)
	return stringRegex.ReplaceAllString(str, "")
}

// string into int
func pullInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// string into float
func pullFloat(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
