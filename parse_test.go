package main

import (
    "math"
	"testing"
    "github.com/lucasb-eyer/go-colorful"
)

func TestParse(t *testing.T) {
    expected := colorful.Color{1, 0.5, 0}

    pairs := [][]string {
        {"h3x", "f80"},
        {"h3x", "#f80"},
        {"hex", "ff8800"},
        {"hex", "#ff8800"},
        {"rgb", "255 128 0"},
        {"rgb", "255, 128, 0"},
        {"rgb", "rgb(255 128 0)"},
        {"rgb", "rgb(255, 128, 0)"},
        {"rgb", "rgb(255,128,0)"},
    }

    for _, pair := range(pairs) {
        format, input := pair[0], pair[1]

        parser, ok := parsers[format]
        if !ok {
            t.Errorf("Missing parser: %s\n", format)
            break
        }

        actual, err := parser(input)
        if err != nil {
            t.Error(err)
            break
        }

        rDist := math.Abs(actual.R - expected.R)
        gDist := math.Abs(actual.G - expected.G)
        bDist := math.Abs(actual.B - expected.B)

        dist := (rDist + gDist + bDist) / 3.0

        if dist > 1 {
            t.Errorf("%s -> unexpected: %v\n", input, actual)
        }
    }
}
