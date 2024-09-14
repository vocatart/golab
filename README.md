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
    lab = golab.ReadLab("path/to/file.lab")

    fmt.Print("lab file %s, duration of %f, has %d labels.", lab.GetName(), lab.GetDuration(), lab.GetLength()) 
    fmt.Print(lab.ToString())
}
```

some .lab examples taken from [kiritan_singing](https://github.com/mmorise/kiritan_singing).
