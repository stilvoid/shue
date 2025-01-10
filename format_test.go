package shue_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stilvoid/shue"
)

func TestFormatAll(t *testing.T) {
	color := shue.Colour{1, 0.5, 0.3, 0.25}

	expected := `#421
#f854
#402013
#ff804d40
hsl(17, 54%, 16%)
hsla(17, 100%, 65%, 0.25)
rgb(64, 32, 19)
rgba(255, 128, 77, 0.25)`

	actual, err := shue.Format("all", color)
	if err != nil {
		t.Error(err)
	}

	if d := cmp.Diff(expected, actual); d != "" {
		t.Error(d)
	}
}

func TestFormat(t *testing.T) {
	color := shue.Colour{1, 0.5, 0.3, 1}

	cases := [][]string{
		{"hex3", "#f85"},
		{"hex6", "#ff804d"},
		{"hsl", "hsl(17, 100%, 65%)"},
		{"rgb", "rgb(255, 128, 77)"},
	}

	for _, pair := range cases {
		format, expected := pair[0], pair[1]

		actual, err := shue.Format(format, color)
		if err != nil {
			t.Error(err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}

func TestFormatMultiplyAlpha(t *testing.T) {
	color := shue.Colour{1, 0.5, 0.3, 0.5}

	cases := [][]string{
		{"hex3", "#842"},
		{"hex6", "#804027"},
		{"hsl", "hsl(17, 54%, 33%)"},
		{"rgb", "rgb(128, 64, 39)"},
	}

	for _, pair := range cases {
		format, expected := pair[0], pair[1]

		actual, err := shue.Format(format, color)
		if err != nil {
			t.Error(err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}

func TestFormatRetainAlpha(t *testing.T) {
	color := shue.Colour{1, 0.5, 0.3, 0.5}

	cases := [][]string{
		{"hex4", "#f858"},
		{"hex8", "#ff804d80"},
		{"hsla", "hsla(17, 100%, 65%, 0.5)"},
		{"rgba", "rgba(255, 128, 77, 0.5)"},
	}

	for _, pair := range cases {
		format, expected := pair[0], pair[1]

		actual, err := shue.Format(format, color)
		if err != nil {
			t.Error(err)
		}

		if d := cmp.Diff(expected, actual); d != "" {
			t.Error(d)
		}
	}
}

func TestOverflow(t *testing.T) {
	colors := []shue.Colour{
		shue.Colour{1, 1, 1, 1},
		shue.Colour{2, 1, 1, 1},
		shue.Colour{1, 2, 1, 1},
		shue.Colour{2, 2, 1, 1},
		shue.Colour{1, 1, 2, 1},
		shue.Colour{2, 1, 2, 1},
		shue.Colour{1, 2, 2, 1},
		shue.Colour{2, 2, 2, 1},
	}

	cases := [][]string{
		{"hex3", "#fff"},
		{"hex6", "#ffffff"},
		{"rgb", "rgb(255, 255, 255)"},
	}

	for _, pair := range cases {
		format, expected := pair[0], pair[1]

		for _, color := range colors {
			actual, err := shue.Format(format, color)
			if err != nil {
				t.Error(err)
			}

			if actual != expected {
				t.Errorf("%s -> unexpected: %v\n", format, actual)
			}
		}
	}
}
