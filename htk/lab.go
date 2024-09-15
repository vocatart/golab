package golab

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// An annotation has a starting and ending time in seconds, with a text label.
type Annotation struct {
	start float64
	end   float64
	label string
}

// A lab contains a collection of annotations.
type Lab struct {
	annotations []Annotation
	name        string
	precision   uint8
}

// Takes a path to a .lab file and reads its contents into a Lab.
func ReadLab(path string) (result Lab) {
	lab := Lab{}
	parsedPrecision := false

	// check if the file exists
	labData, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer labData.Close()

	// iterate through the file and write to lab object
	line := bufio.NewScanner(labData)
	for line.Scan() {
		// split by whitespace
		labLine := strings.Split(line.Text(), " ")

		// the lab is malformed if there are less than 3 elements per line
		if len(labLine) < 3 {
			log.Fatal("malformed lab file")
		}

		// parse the precision if it hasnt been parsed yet
		if !parsedPrecision {
			lab.parsePrecision(labLine[1])
			parsedPrecision = true
		}

		// parse the start and end times
		start, err := strconv.ParseFloat(labLine[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.ParseFloat(labLine[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// join the rest of the line into a single string and add it to the annotations
		label := strings.Join(labLine[2:], " ")
		lab.annotations = append(lab.annotations, Annotation{start: start, end: end, label: label})
	}

	lab.name = filepath.Base(path)

	return lab
}

// Writes a lab to a file from a given path. If the file already exists, it will be overwritten unless overwrite is set to false.
// If the path is a directory, the lab will be written to a file with the same name as the lab, in the directory.
func (lab Lab) WriteLab(path string, overwrite ...bool) {
	// if no overwrite is specified, default to false
	if len(overwrite) == 0 {
		overwrite = append(overwrite, false)
	}

	// replace backslashes with forward slashes
	path = strings.Replace(path, "\\", "/", -1)

	// initialize the filename
	var fileName string

	// check if path exists, if it doesnt, make it.
	pathInfo, err := os.Stat(path)
	if err != nil {
		if filepath.Ext(path) != "" {
			// if the path is a file, make the directory it is in
			pathSplit := strings.Split(path, "/")
			os.MkdirAll(strings.Join(pathSplit[0:len(pathSplit)-1], "/"), os.ModePerm)
		} else {
			// if the path is a directory, make it
			os.MkdirAll(path, os.ModePerm)
		}
	}

	// if path is a directory, construct the desired filename. if path is a file, make it.
	if filepath.Ext(path) == "" || (pathInfo != nil && pathInfo.IsDir()) {
		// create the filename if the destination is a directory
		fileName = filepath.Join(path, (lab.name + ".lab"))
	} else {
		// if the path is a file, check if it already exists and if overwrite is false
		if pathInfo != nil && !overwrite[0] {
			log.Fatalf("file %s already exists", path)
		}
		fileName = path
	}

	// create the file with the filename defined above
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// iterate through the annotations and write them to the file
	precision := lab.precision

	for _, labEntry := range lab.annotations {
		start := labEntry.start
		end := labEntry.end

		fmt.Fprintln(file, strconv.FormatFloat(start, 'f', int(precision), 64), strconv.FormatFloat(end, 'f', int(precision), 64), labEntry.label)
	}
}

// Converts a lab to a string
func (lab Lab) ToString() string {
	var result string

	for _, annotation := range lab.annotations {
		result += fmt.Sprintf("%s %s %s\n", strconv.FormatFloat(annotation.start, 'f', int(lab.precision), 64), strconv.FormatFloat(annotation.end, 'f', int(lab.precision), 64), annotation.label)
	}

	return result
}

// Parses the precision of a lab file based on the context of the time durations
func (lab *Lab) parsePrecision(secondTime string) {
	periodIndex := strings.Index(secondTime, ".")

	if periodIndex == -1 {
		lab.precision = 7
		return
	}

	lab.precision = uint8(len(secondTime) - periodIndex - 1)
}

// Compare two slices to see if they are identical
func isEqualSlice(slice1, slice2 []string) bool {
	// if they aren't the same length, we return false right away
	if len(slice1) != len(slice2) {
		return false
	}

	// check and compare each element
	for index, element := range slice1 {
		if element != slice2[index] {
			return false
		}
	}
	return true
}
