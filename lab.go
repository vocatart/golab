package golab

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
}

// Gets the total duration of a lab by getting the difference in global start and end.
func (lab Lab) getDuration() (result float64) {
	// calculate using start and end in case lab file doesnt start at 0
	start := lab.annotations[0].start
	end := lab.annotations[len(lab.annotations)-1].end

	return end - start
}

// Gets the total duration of an annotation
func (annotation Annotation) getDuration() (result float64) {
	return annotation.end - annotation.start
}

// Takes a path to a .lab file and reads its contents into a Lab. If the file doesn't exist or the start and end arent valid floats, an empty Lab will be returned.
func readLab(path string) (result Lab) {
	label := Lab{}

	labData, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return Lab{}
	}
	defer labData.Close()

	line := bufio.NewScanner(labData)
	for line.Scan() {
		labLine := strings.Split(line.Text(), " ")

		start, err := strconv.ParseFloat(labLine[0], 64)
		if err != nil {
			log.Fatal(err)
			return Lab{}
		}
		end, err := strconv.ParseFloat(labLine[1], 64)
		if err != nil {
			log.Fatal(err)
			return Lab{}
		}

		label.annotations = append(label.annotations, Annotation{start: start, end: end, label: labLine[2]})
	}

	return label
}

func writeLab(lab Lab, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, labEntry := range lab.annotations {
		fmt.Fprintln(file, strconv.FormatFloat(labEntry.start, 'f', 7, 64), strconv.FormatFloat(labEntry.end, 'f', 7, 64), labEntry.label)
	}
}
