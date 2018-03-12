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

        formatter, ok := formatters[format]
        if !ok {
            t.Errorf("Missing formatter: %s\n", format)
            break
        }

        actual := formatter(color)

        if actual != expected {
            t.Errorf("%s -> unexpected: %v\n", format, actual)
        }
    }
}
