package textgrid

import "testing"

func TestReadingTextgridASCIILong(t *testing.T) {
	ReadTextgrid("examples/long.TextGrid")
}

func TestReadingTextgridASCIIShort(t *testing.T) {
	ReadTextgrid("examples/short.TextGrid")
}

func TestReadingTextgridUTF16(t *testing.T) {
	ReadTextgrid("examples/polish64.TextGrid")
}

func TestReadingTextgridUTF8(t *testing.T) {
	ReadTextgrid("examples/polish65.TextGrid")
}
