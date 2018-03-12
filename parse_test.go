package main

import (
    "math"
	"testing"
    "github.com/lucasb-eyer/go-colorful"
)

func TestParse(t *testing.T) {
    expected := colorful.Color{1, 0.5, 0}

    pairs := []string {
        "f80",
        "#f80",
        "ff8800",
        "#ff8800",
        "255 128 0",
        "255, 128, 0",
        "rgb(255 128 0)",
        "rgb(255, 128, 0)",
        "rgb(255,128,0)",
    }

    for _, input := range(pairs) {
        actual, err := parse(input)
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
