package shue

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	inputs := []string{
		"f60",
		"#f90",
		"339966",
		"#663399",
		"rgb(255, 128, 0)",
		"rgb(16,64,128)",
		"hsl(0,100%,50%)",
	}

	expecteds := []Colour{
		{1, 0.4, 0, 1},
		{1, 0.6, 0, 1},
		{0.2, 0.6, 0.4, 1},
		{0.4, 0.2, 0.6, 1},
		{1, 0.5, 0, 1},
		{0.0625, 0.25, 0.5, 1},
		{1, 0, 0, 1},
	}

	for i, input := range inputs {
		expected := expecteds[i]

		actual, err := Parse(input)
		if err != nil {
			t.Error(err)
		}

		if !expected.Equal(actual) {
			t.Error(fmt.Errorf("%#v != %#v", actual, expected))
		}
	}
}

func TestParseWithAlpha(t *testing.T) {
	inputs := []string{
		"f603",
		"#f906",
		"33996600",
		"#66339999",
		"rgba(255, 128, 0, 0.25)",
		"rgba(16,64,128, 0.9)",
		"hsla(0,100%,50%,0.75)",
	}

	expecteds := []Colour{
		{1, 0.4, 0, 0.2},
		{1, 0.6, 0, 0.4},
		{0.2, 0.6, 0.4, 0},
		{0.4, 0.2, 0.6, 0.6},
		{1, 0.5, 0, 0.25},
		{0.0625, 0.25, 0.5, 0.9},
		{1, 0, 0, 0.75},
	}

	for i, input := range inputs {
		expected := expecteds[i]

		actual, err := Parse(input)
		if err != nil {
			t.Error(err)
		}

		if !expected.Equal(actual) {
			t.Error(fmt.Errorf("%#v != %#v", actual, expected))
		}
	}
}
