# golab

[![Go Reference](https://pkg.go.dev/badge/github.com/vocatart/golab.svg)](https://pkg.go.dev/github.com/vocatart/golab)

simple HTK & TextGrid handling for go.

## example

allows you to read and write annotation data from lab files that are in seconds. will automatically detect floating point preicision of a read file. can also view the data of individual annotations.

```go
package main

import "github.com/vocatart/golab/htk"

func main() {
    lab = golab.ReadLab("path/to/file.lab")

    fmt.Printf("lab file %s, duration of %f, has %d labels.", lab.GetName(), lab.GetDuration(), lab.GetLength()) 
    fmt.Printf(lab.ToString())

    lab.WriteLab("path/to/new/lab.lab")

    annotations = lab.GetAnnotations()

    fmt.Printf("annotation 0 has start time of %f, end time of %f, and label %s", annotations[0].GetStart(), annotations[0].GetEnd(), annotations[0].GetLabel())
}
```

for textgrid, implementation is mostly there, but cannot write to files yet.

```go
package main

import "github.com/vocatart/golab/textgrid"

func main() {
    tg = golab.ReadTextgrid("path/to/file.TextGrid")

    fmt.Printf("textgrid file %s, duration of %f, has %d tiers.", tg.GetName(), tg.GetDuration(), tg.GetTiers().GetSize()) 
}
```
some .lab examples taken from [kiritan_singing](https://github.com/mmorise/kiritan_singing).
