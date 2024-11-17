package textgrid

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/TomOnTime/utfutil"
	"github.com/gammazero/deque"
	"github.com/saintfish/chardet"
)

type TextGrid struct {
	xmin  float64
	xmax  float64
	tiers []Tier
	name  string
}

// GetXmin - Returns xmin of TextGrid
func (tg *TextGrid) GetXmin() float64 {
	return tg.xmin
}

// SetXmin - Sets xmin of TextGrid
func (tg *TextGrid) SetXmin(xmin float64) {
	tg.xmin = xmin
}

// GetXmax - Returns xmax of TextGrid
func (tg *TextGrid) GetXmax() float64 {
	return tg.xmax
}

// SetXmax - Sets xmax of TextGrid
func (tg *TextGrid) SetXmax(xmax float64) {
	tg.xmax = xmax
}

// GetName - Returns name of TextGrid
func (tg *TextGrid) GetName() string {
	return tg.name
}

// SetName - Sets name of TextGrid
func (tg *TextGrid) SetName(name string) {
	tg.name = name
}

// GetTiers - Gets tiers of TextGrid
func (tg *TextGrid) GetTiers() []Tier {
	return tg.tiers
}

// HasIntervalTier - Returns true if TextGrid has IntervalTier
func (tg *TextGrid) HasIntervalTier() bool {
	for _, tier := range tg.tiers {
		if tier.TierType() == "IntervalTier" {
			return true
		}
	}
	return false
}

// HasPointTier - Returns true if TextGrid has PointTier
func (tg *TextGrid) HasPointTier() bool {
	for _, tier := range tg.tiers {
		if tier.TierType() == "PointTier" {
			return true
		}
	}
	return false
}

// GetTier - Returns tier with given name, if it exists
func (tg *TextGrid) GetTier(name string) Tier {
	for _, tier := range tg.tiers {
		if tier.TierName() == name {
			return tier
		}
	}
	return nil
}

// GetSize returns the amount of Tier entries in a TextGrid.
func (tg *TextGrid) GetSize() int {
	return len(tg.tiers)
}

// ReadTextgrid - Takes a path to a .TextGrid file and reads its contents into a TextGrid.
func ReadTextgrid(path string) (TextGrid, error) {
	var tg = TextGrid{}
	tgDeque := deque.New[string]()

	// grab the name element from the path
	tg.name = filepath.Base(path)

	// check if the file exists
	// TextGrid files are USUALLY UTF-8, UTF-16, or ASCII.
	tgData, err := utfutil.ReadFile(path, utfutil.UTF8)
	if err != nil {
		return tg, err
	}

	// if somehow utf-util finds something else, error out
	// chardet will sometimes read ASCII as ISO-8859-1, which go will always interpret correctly when casting its byte slice to a string.
	retrievedEncoding := getEncoding(tgData)
	if retrievedEncoding != "UTF-8" && retrievedEncoding != "ISO-8859-1" {
		return tg, fmt.Errorf("error: encoding out of scope, recieved %q encoding for %q", retrievedEncoding, tg.name)
	}

	// convert string slice into deque
	tgContent := processContent(tgData)
	for _, str := range tgContent {
		tgDeque.PushBack(str)
	}

	// verify the first two entries in the deque
	err = verifyHead(tgDeque)
	if err != nil {
		return tg, fmt.Errorf("error: textgrid %s has malformed header\n %s", tg.name, err.Error())
	}

	// pop the next two values, which should be xmin and xmax respectively.
	globalXmin := pullFloat(tgDeque.PopFront())
	globalXmax := pullFloat(tgDeque.PopFront())

	fmt.Println(globalXmax)

	// set the xmin and xmax preemptively in case the status is <absent>
	tg.xmin = globalXmin
	tg.xmax = globalXmax

	// the next value is the status, which should either be <absent> or <exists>
	tierStatus := pullBracketedValue(tgDeque.PopFront())
	if tierStatus == "absent" {
		log.Println("warning: tierStatus is <absent>, a textgrid with 0 tiers will be returned")
		return tg, nil
	} else if tierStatus != "exists" {
		return tg, fmt.Errorf("error: expected tier status <exists> or <absent>, recieved <%s>", tierStatus)
	}

	// get the number of tiers that exist in this textgrid
	numTiers := pullInt(tgDeque.PopFront())

	tiers, err := parseTiers(globalXmin, globalXmax, tgDeque, numTiers)
	if err != nil {
		return tg, err
	}
	tg.tiers = tiers

	return tg, nil
}

// WriteLong writes to a .TextGrid file in long format.
// Will overwrite existing files unless otherwise specified.
func (tg *TextGrid) WriteLong(path string, overwrite ...bool) error {
	// default to false
	if len(overwrite) == 0 {
		overwrite = append(overwrite, false)
	}

	// replace backslashes with forward slashes
	path = strings.Replace(path, "\\", "/", -1)

	var fileName string
	pathInfo, err := os.Stat(path)
	if err != nil {
		if filepath.Ext(path) != "" {
			// if the path is a file, make the directory it is in
			pathSplit := strings.Split(path, "/")
			err := os.MkdirAll(strings.Join(pathSplit[0:len(pathSplit)-1], "/"), os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// if the path is a directory, make it
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	// if path is a directory, construct the desired filename. If path is a file, make it.
	if filepath.Ext(path) == "" || (pathInfo != nil && pathInfo.IsDir()) {
		fileName = filepath.Join(path, tg.name+".TextGrid")
	} else {
		// if the path is a file, check if it already exists and if overwrite is false
		if pathInfo != nil && !overwrite[0] {
			return fmt.Errorf("error writing textgrid %q: file %s already exists", tg.name, path)
		}
		fileName = path
	}

	// create the file with the filename defined above
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		closingError := file.Close()
		if err == nil {
			err = closingError
		}
	}(file)

	// create the header of the textgrid file
	_, err = fmt.Fprintf(file, "File type = \"ooTextFile\"\nObject class = \"TextGrid\"\n\n")
	if err != nil {
		return err
	}

	// create the xmin and xmax of the textgrid file
	_, err = fmt.Fprintf(file, "xmin = %s\nxmax = %s\n", f2s(tg.xmin), f2s(tg.xmax))
	if err != nil {
		return err
	}

	// create the tier flag
	// is usually <exists> if you have tiers, but can also be <absent> if you somehow have a tier-less textgrid
	if tg.tiers == nil {
		_, err = fmt.Fprintf(file, "tiers? <absent>")
		return nil
	} else {
		_, err = fmt.Fprintf(file, "tiers? <exists>")
	}

	// create the size field, which is the number of tiers in the textgrid.
	_, err = fmt.Fprintf(file, "\nsize = %d\n", tg.GetSize())
	if err != nil {
		return err
	}

	// begin writing tiers, which starts with the blank item [] field
	_, err = fmt.Fprintf(file, "item []:\n")
	if err != nil {
		return err
	}

	for tierNum, tier := range tg.tiers {
		// write tier info
		_, err = fmt.Fprintf(file, "\titem [%d]:\n", tierNum+1)
		if err != nil {
			return err
		}

		// tier class
		_, err = fmt.Fprintf(file, "\t\tclass = \"%s\"\n", tier.TierType())
		if err != nil {
			return err
		}

		// tier name
		_, err = fmt.Fprintf(file, "\t\tname = \"%s\"\n", tier.TierName())
		if err != nil {
			return err
		}

		// xmin and xmax
		_, err = fmt.Fprintf(file, "\t\txmin = %s\n\t\txmax = %s\n", f2s(tier.TierXmin()), f2s(tier.TierXmax()))

		// write content info and contents of tier
		if tier.TierType() == "IntervalTier" {
			// if tier is interval tier
			_, err = fmt.Fprintf(file, "\t\tintervals: size = %d\n", tier.GetSize())
			if err != nil {
				return err
			}

			for intervalNum, interval := range tier.GetIntervals() {
				// write interval number
				_, err = fmt.Fprintf(file, "\t\tintervals [%d]:\n", intervalNum)
				if err != nil {
					return err
				}

				// xmin and xmax
				_, err = fmt.Fprintf(file, "\t\t\txmin = %f\n\t\t\txmax = %f\n", tg.xmin, tg.xmax)
				if err != nil {
					return err
				}

				// text
				_, err = fmt.Fprintf(file, "\t\t\ttext = \"%s\"\n", interval.text)
				if err != nil {
					return err
				}
			}
		} else {
			// if tier is point tier
			_, err = fmt.Fprintf(file, "\t\tpoints: size = %d\n", tier.GetSize())
			if err != nil {
				return err
			}

			for pointNum, point := range tier.GetPoints() {
				// write point number
				_, err = fmt.Fprintf(file, "\t\tpoints [%d]:\n", pointNum)
				if err != nil {
					return err
				}

				// value
				_, err = fmt.Fprintf(file, "\t\t\tnumber = %f\n", point.value)
				if err != nil {
					return err
				}

				// mark
				_, err = fmt.Fprintf(file, "\t\t\tmark = \"%s\"\n", point.mark)
			}
		}
	}

	return nil
}

func parseTiers(globalXmin float64, globalXmax float64, content *deque.Deque[string], numTiers int) ([]Tier, error) {
	var tiers []Tier
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
			return nil, fmt.Errorf("error: %s %s has xmin %f, when TextGrid xmin is %f", tierType, tierName, tierXmin, globalXmin)
		}
		if tierXmax > globalXmax {
			return nil, fmt.Errorf("error: %s %s has xmax %f, when TextGrid xmax is %f", tierType, tierName, tierXmax, globalXmax)
		}

		// the last value before the intervals/points begin should be the number of intervals/points in the unique tier
		tierContentCount := pullInt(content.PopFront())
		contentCounter := 0

		// loop for each tier type
		if tierType == "IntervalTier" {
			var intervals []Interval

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
			var points []Point

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
			return nil, fmt.Errorf("error: unexpected tier type %s", tierType)
		}
		tierCounter++
	}

	return tiers, nil
}

// turns textgrid file content into a slice of usable strings.
// internally, any textgrid given is converted into a "short" type textgrid with any empty lines removed
func processContent(data []byte) []string {

	// remove all empty lines
	tgString := strings.ReplaceAll(string(data), "\n\n", "\n")

	bracketRegex := regexp.MustCompile(`\[\d+]`)

	// a short textgrid is basically a textgrid that is only labels, numbers, and flags.
	// we will use regex to remove everything that isn't needed by praat to recognize a textgrid.
	// `\d+(\.\d+)?` matches all floats and integers
	// `\"[^\"]*\` matches all content in between double quotes
	// `\<[^>]*>` matches all content in between angle brackets
	textgridRegex := regexp.MustCompile(`(\d+(\.\d+)?|"[^"]*"|<[^>]*)`)

	tgSanitized := strings.Join(textgridRegex.FindAllString(bracketRegex.ReplaceAllString(tgString, ""), -1), "\n")

	return strings.Split(tgSanitized, "\n")
}

// checks the necessary `File Type` and `Object Class` fields of a TextGrid file
func verifyHead(tgContent *deque.Deque[string]) error {
	fileType := tgContent.PopFront()
	objectClass := tgContent.PopFront()

	if pullQuotedValue(fileType) != "ooTextFile" {
		return fmt.Errorf("error: wanted fileType ooTextFile, recieved %s", fileType)
	}

	if pullQuotedValue(objectClass) != "TextGrid" {
		return fmt.Errorf("error: wanted objectClass TextGrid, recieved %s", objectClass)
	}

	return nil
}

// takes a value contained in quotes and returns it without quotes.
// quotes inside quotes will be preserved.
func pullQuotedValue(str string) string {
	stringRegex := regexp.MustCompile(`^"(.*)"$`)
	return stringRegex.ReplaceAllString(str, `$1`)
}

// takes a value that has angle brackets and removes them.
// this is only used for flags in TextGrids, so we don't care about internal brackets.
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

// get encoding type of binary data slice
func getEncoding(data []byte) string {
	detector := chardet.NewTextDetector()

	detectedEncoding, err := detector.DetectBest(data)
	if err != nil {
		log.Fatal(err)
	}

	return detectedEncoding.Charset
}

// convert float to string
func f2s(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
