# golab

[![Go Reference](https://pkg.go.dev/badge/github.com/vocatart/golab.svg)](https://pkg.go.dev/github.com/vocatart/golab)

Simple HTK label reading and writing for go.

## example

very scuffed and written in one day and also my first real open source project. expect it to be unstable

todo: add better testing and error handling

```go
package main

import "github.com/vocatart/golab"

func main() {
    lab = golab.ReadLab("path/to/file.lab") // read a lab file into memory
    lab.WriteLab("path/to/file.lab") // write directly to a lab file
    lab.WriteLab("path/to/directory") // or write to a directory, using the name field of the lab object as the filename

    lab.ToString() // convert the entire lab into a string

    lab.SetDenomination(10) // set the denomination of the start and end times of a lab file

    lab.GetDenomination() // get the denomination of the start and end times of a lab file

    lab.GetDuration() // get the duration of the entire file

    lab.annotations[0].GetDuration() // get the duration of an exact annotation

    lab.GetAnnotations() // get the annotations of a lab file

    lab.SetAnnotations(Annotations[]{Start: 0.0, End: 1.0, Label: "Annotation"}) // set the annotations of a lab file

    lab.PushAnnotation(Annotation{Start: 0.0, End: 1.0, Label: "Annotation"}) // push an annotation to the end of the annotations slice in a lab file

    lab.AppendAnnotations(Annotations[]{Start: 0.0, End: 1.0, Label: "Annotation", Start: 1.0, End: 2.0, Label: "Annotation 2"}) // append annotations to the end of the annotations slice in a lab file

    lab.InsertAnnotation(index: 0, Annotation{Start: 0.0, End: 1.0, Label: "Annotation"}) // insert an annotation at a specific index in the annotations slice in a lab file

    lab.RemoveAnnotation(index: 0) // remove an annotation at a specific index in the annotations slice in a lab file

    lab.GetName() // get the name of a lab file

    lab.SetName("label") // set the name of a lab file

    lab.GetPrecision() // get the precision of a lab file

    lab.SetPrecision(7) // set the precision of a lab file

    // get the three values inside annotations
    lab.Annotations[0].Start // float64
    lab.Annotations[0].End // float64
    lab.Annotations[0].Label // string

    // get the four values inside a lab
    lab.Annotations // []Annotation
    lab.Name // string
    lab.Denomination // *Denomination
    lab.Precision // uint8
}
```

examples directory shows an input .lab file from [kiritan_singing](https://github.com/mmorise/kiritan_singing) being read into memory by golab and exported into output.lab
