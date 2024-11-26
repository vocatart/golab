package textgrid

import "testing"

func TestReadingTextgridASCIILong(t *testing.T) {
	_, err := ReadTextgrid("examples/long.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestReadingTextgridASCIIShort(t *testing.T) {
	_, err := ReadTextgrid("examples/short.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestReadingTextgridUTF16(t *testing.T) {
	_, err := ReadTextgrid("examples/polish64.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestReadingTextgridUTF8(t *testing.T) {
	_, err := ReadTextgrid("examples/polish65.TextGrid")
	if err != nil {
		t.Error(err)
	}
}

func TestWritingLongTextgrid(t *testing.T) {
	tg, err := ReadTextgrid("examples/long.TextGrid")
	if err != nil {
		t.Error(err)
	}

	err = tg.WriteLong("examples/long_output.TextGrid", true)
	if err != nil {
		t.Error(err)
	}
}

func TestWritingShortTextgrid(t *testing.T) {
	tg, err := ReadTextgrid("examples/short.TextGrid")
	if err != nil {
		t.Error(err)
	}

	err = tg.WriteShort("examples/short_output.TextGrid", true)
	if err != nil {
		t.Error(err)
	}
}
