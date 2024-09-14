# golab

[![Go Reference](https://pkg.go.dev/badge/github.com/vocatart/golab.svg)](https://pkg.go.dev/github.com/vocatart/golab)

simple HTK label reading and writing for go.

## example

allows you to read and write annotation data from lab files that are in seconds. will automatically detect floating point preicision of a read file. can also view the data of individual annotations.

```go
package main

import "github.com/vocatart/golab"

func main() {
    lab = golab.ReadLab("path/to/file.lab")

    fmt.Printf("lab file %s, duration of %f, has %d labels.", lab.GetName(), lab.GetDuration(), lab.GetLength()) 
    fmt.Printf(lab.ToString())

    lab.WriteLab("path/to/new/lab.lab")

    annotations = lab.GetAnnotations()

    fmt.Printf("annotation 0 has start time of %f, end time of %f, and label %s", annotations[0].GetStart(), annotations[0].GetEnd(), annotations[0].GetLabel())
}
```

some .lab examples taken from [kiritan_singing](https://github.com/mmorise/kiritan_singing).
