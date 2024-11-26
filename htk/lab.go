package htk

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Lab structs are a collection of annotations.
// The HTK Label format is defined at http://www.seas.ucla.edu/spapl/weichu/htkbook/node113_mn.html
type Lab struct {
	annotations []Annotation
	name        string
	precision   uint8
}

// SetAnnotations sets the annotations field in a Lab.
func (lab *Lab) SetAnnotations(annotations []Annotation) {
	lab.annotations = annotations
}

// GetAnnotations gets the annotations field in a Lab.
func (lab *Lab) GetAnnotations() []Annotation {
	return lab.annotations
}

// PushAnnotation pushes a single Annotation into the annotations field of a Lab.
func (lab *Lab) PushAnnotation(annotation Annotation) {
	lab.annotations = append(lab.annotations, annotation)
}

// AppendAnnotations appends an Annotation slice to the annotations field in a Lab.
func (lab *Lab) AppendAnnotations(annotations []Annotation) {
	lab.annotations = append(lab.annotations, annotations...)
}

// ClearAnnotations removes all annotations in a Lab.
func (lab *Lab) ClearAnnotations() {
	lab.annotations = nil
}

// GetLabels returns the annotations field in a Lab as a slice of strings.
func (lab *Lab) GetLabels() []string {
	var result []string

	for _, annotation := range lab.annotations {
		result = append(result, annotation.label)
	}

	return result
}

// GetName gets the name of a Lab.
func (lab *Lab) GetName() string {
	return lab.name
}

// SetName sets the name of a Lab.
func (lab *Lab) SetName(name string) {
	lab.name = name
}

// GetPrecision gets the precision of a Lab.
func (lab *Lab) GetPrecision() uint8 {
	return lab.precision
}

// SetPrecision sets the precision of a Lab.
func (lab *Lab) SetPrecision(precision uint8) {
	lab.precision = precision
}

// GetDuration gets the total duration of a Lab by getting the difference in global start and end.
func (lab *Lab) GetDuration() (result float64) {
	// calculate using start and end in case lab file doesn't start at 0
	start := lab.annotations[0].start
	end := lab.annotations[len(lab.annotations)-1].end

	return end - start
}

// GetLength gets the total amount of annotations in a Lab.
func (lab *Lab) GetLength() (result int) {
	return len(lab.annotations)
}

// ReadLab takes a path to a .lab file and reads its contents into a Lab.
func ReadLab(path string) (Lab, error) {
	lab := Lab{}
	parsedPrecision := false

	// check if the file exists
	labData, err := os.Open(path)
	if err != nil {
		return lab, err
	}
	defer func() {
		closingError := labData.Close()
		if err == nil {
			err = closingError
		}
	}()

	// iterate through the file and write to lab object
	line := bufio.NewScanner(labData)
	for line.Scan() {
		// split by whitespace
		labLine := strings.Split(line.Text(), " ")

		// the lab is malformed if there are less than 3 elements per line
		if len(labLine) < 3 {
			return lab, fmt.Errorf("error: malformed lab file %s", path)
		}

		// parse the precision if it hasn't been parsed yet
		if !parsedPrecision {
			lab.parsePrecision(labLine[1])
			parsedPrecision = true
		}

		// parse the start and end times
		start, err := strconv.ParseFloat(labLine[0], 64)
		if err != nil {
			return lab, err
		}
		end, err := strconv.ParseFloat(labLine[1], 64)
		if err != nil {
			return lab, err
		}

		// join the rest of the line into a single string and add it to the annotations
		label := strings.Join(labLine[2:], " ")
		lab.annotations = append(lab.annotations, Annotation{start: start, end: end, label: label})
	}

	lab.name = filepath.Base(path)

	return lab, err
}

// WriteLab writes a Lab to a file from a given path. If the file already exists, it will be overwritten unless overwrite is set to false.
// If the path is a directory, the contents will be written to a file with the same name as the Lab, in the directory.
func (lab *Lab) WriteLab(path string, overwrite ...bool) error {
	// if no overwrite is specified, default to false
	if len(overwrite) == 0 {
		overwrite = append(overwrite, false)
	}

	// replace backslashes with forward slashes
	path = strings.Replace(path, "\\", "/", -1)

	// initialize the filename
	var fileName string

	// check if path exists, if it doesn't, make it.
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

	// if path is a directory, construct the desired filename. if path is a file, make it.
	if filepath.Ext(path) == "" || (pathInfo != nil && pathInfo.IsDir()) {
		// create the filename if the destination is a directory
		fileName = filepath.Join(path, lab.name+".lab")
	} else {
		// if the path is a file, check if it already exists and if overwrite is false
		if pathInfo != nil && !overwrite[0] {
			log.Fatalf("error: file %s already exists", path)
		}
		fileName = path
	}

	// create the file with the filename defined above
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		closingError := file.Close()
		if err == nil {
			err = closingError
		}
	}(file)

	// iterate through the annotations and write them to the file
	precision := lab.precision

	for _, labEntry := range lab.annotations {
		start := labEntry.start
		end := labEntry.end

		_, err := fmt.Fprintln(file, strconv.FormatFloat(start, 'f', int(precision), 64), strconv.FormatFloat(end, 'f', int(precision), 64), labEntry.label)
		if err != nil {
			return err
		}
	}
	return nil
}

// ToString converts a lab to a string.
func (lab *Lab) ToString() string {
	var result string

	for _, annotation := range lab.annotations {
		result += fmt.Sprintf("%s %s %s\n", strconv.FormatFloat(annotation.start, 'f', int(lab.precision), 64), strconv.FormatFloat(annotation.end, 'f', int(lab.precision), 64), annotation.label)
	}

	return result
}

// parsePrecision retrieves the precision field of a Lab based on the context of the time durations.
func (lab *Lab) parsePrecision(secondTime string) {
	periodIndex := strings.Index(secondTime, ".")

	if periodIndex == -1 {
		lab.precision = 7
		return
	}

	lab.precision = uint8(len(secondTime) - periodIndex - 1)
}

// isEqualSlice compares two slices, returning true if they are identical.
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
