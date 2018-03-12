package main

import (
	"testing"
    "github.com/lucasb-eyer/go-colorful"
)

func TestFormat(t *testing.T) {
    color := colorful.Color{1, 0.5, 0.3}

    cases := [][]string {
        {"h3x", "#f85"},
        {"hex", "#ff804d"},
        {"rgb", "rgb(255, 128, 77)"},
    }

    for _, pair := range(cases) {
        format, expected := pair[0], pair[1]

        actual := formatters[format](color)

        if actual != expected {
            t.Errorf("%s -> unexpected: %v\n", format, actual)
        }
    }
}

func TestOverflow(t *testing.T) {
    colors := []colorful.Color{
        {1, 1, 1},
        {2, 1, 1},
        {1, 2, 1},
        {2, 2, 1},
        {1, 1, 2},
        {2, 1, 2},
        {1, 2, 2},
        {2, 2, 2},
    }

    cases := [][]string {
        {"h3x", "#fff"},
        {"hex", "#ffffff"},
        {"rgb", "rgb(255, 255, 255)"},
    }

    for _, pair := range(cases) {
        format, expected := pair[0], pair[1]

        for _, color := range(colors) {
            actual := formatters[format](color)

            if actual != expected {
                t.Errorf("%s -> unexpected: %v\n", format, actual)
            }
        }
    }
}
