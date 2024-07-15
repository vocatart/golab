# golab

Simple TextGrid reading and writing for go.

## example

very scuffed and written in one day and also my first real open source project. expect it to be unstable

```go
package main

import "github.com/vocatart/golab"

func main() {
    golab.readLab("path/to/file") // read lab into memory
    golab.writeLab("path/to/file") // write lab into plaintext file

    lab = golab.readLab("path/to/file")

    // get the three values inside annotations
    lab.annotations[0].start
    lab.annotations[0].end
    lab.annotations[0].label

    lab.getDuration() // get the duration of the entire file

    lab.annotations[0].getDuration() // get the duration of an exact annotation
}
```

examples directory shows an input .lab file from [kiritan_singing](https://github.com/mmorise/kiritan_singing) being read into memory by golab and exported into output.lab
