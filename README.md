# golab

[![Go Reference](https://pkg.go.dev/badge/github.com/vocatart/golab.svg "Go Reference")](https://pkg.go.dev/github.com/vocatart/golab)
![go test](https://github.com/vocatart/golab/actions/workflows/go.yml/badge.svg "Go Test Status")

simple HTK & TextGrid handling for go.

## installation

```plaintext
$ go get github.com/vocatart/golab
```

## about HTK and TextGrid

[HTK](http://www.seas.ucla.edu/spapl/weichu/htkbook/node113_mn.html) and [TextGrid](https://www.fon.hum.uva.nl/praat/manual/TextGrid_file_formats.html) are two popular formats used for annotating audio, mainly for use in linguistic analysis.

### HTK

while HTK is officially defined as `[start  [end] ] name [score] { auxname [auxscore] } [comment]` golab only 
supports HTK labels in the format `[start] [end] [name]`. Due to the non-standardization of the HTK format, it is 
hard to guarantee compatibility. For most linguistic applications, this should be sufficient.

### TextGrid

TextGrid files are internally stored as short format TextGrids. all other information in between the relevant data 
will be ignored.

#### tiers flag

TextGrid files have a flag that designate whether they contain tiers. For example, a TextGrid that contains tiers 
would look like the following:

```TextGrid
File type = "ooTextFile"
Object class = "TextGrid"

xmin = 0 
xmax = 2.3510204081632655 
tiers? <exists> 
size = 3 
item []: 
    item [1]:
        class = "IntervalTier" 
        name = "Mary" 
        xmin = 0 
        xmax = 2.3510204081632655 
        intervals: size = 3 
        intervals [1]:
            xmin = 0 
            xmax = 0.7427342752056899 
            text = "1_label1" 
        intervals [2]:
            xmin = 0.7427342752056899 
            xmax = 1.7447703580322245 
            text = "1_label2"
        intervals [3]:
            xmin = 1.7447703580322245 
            xmax = 2.3510204081632655 
            text = "1_label3" 
    item [2]:
        class = "IntervalTier" 
        name = "John" 
        xmin = 0 
        xmax = 2.3510204081632655 
        intervals: size = 2 
        intervals [1]:
            xmin = 0 
            xmax = 1.2402970197816243 
            text = "2_label1" 
        intervals [2]:
            xmin = 1.2402970197816243 
            xmax = 2.3510204081632655 
            text = "2_label2" 
    item [3]:
        class = "TextTier" 
        name = "Bell" 
        xmin = 0 
        xmax = 2.3510204081632655 
        points: size = 3 
        points [1]:
            number = 0.40238753672840144 
            mark = "point1" 
        points [2]:
            number = 1.1677357861976339 
            mark = "point2" 
        points [3]:
            number = 1.8950757704562047 
            mark = "point3" 
```

while a TextGrid with no tiers would have an `<absent>` tag instead.

```TextGrid
File type = "ooTextFile"
Object class = "TextGrid"

xmin = 0 
xmax = 2.3510204081632655 
tiers? <absent>
```

## example

```go
package main

import (
	"github.com/vocatart/golab/htk"
	"github.com/vocatart/golab/textgrid"
)

func main() {
	lab, err := htk.ReadLab("examples/short.lab")
	if err != nil {
		panic(err)
	}

	lab.GetName()        // returns "short"
	lab.GetAnnotations() // returns []Annotation
	lab.GetPrecision()   // returns 7 (floating point precision of file, parsed when read in)
	// etc

	tg, err := textgrid.ReadTextgrid("examples/long.TextGrid")
	if err != nil {
		panic(err)
    }

	tg.GetXmin() // returns 0.0
	tg.GetXmax() // returns 2.3510204081632655
	tg.GetTiers() // returns []Tier (can contain IntervalTier or PointTier)
	tg.GetName() // returns "long"
}
```

some .lab examples taken from [kiritan_singing](https://github.com/mmorise/kiritan_singing).
