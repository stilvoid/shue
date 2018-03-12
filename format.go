package main

import (
    "math"
    "fmt"
    "sort"
    "strings"
    "github.com/lucasb-eyer/go-colorful"
)

type formatter func(colorful.Color)string
var formatters map[string]formatter
var formats []string

func init() {
    formatters = map[string]formatter {
        "hex": formatHex,
        "h3x": formatH3x,
        "rgb": formatRgb,
    }

    formats = make([]string, 0, len(formatters))
    for format := range(formatters) {
        formats = append(formats, format)
    }
    sort.Strings(formats)

    formatters["all"] = func(color colorful.Color) string {
        outputs := make([]string, 0, len(formats))

        for _, format := range(formats) {
            outputs = append(outputs, formatters[format](color))
        }

        return strings.Join(outputs, "\n")
    }
}

func convert(in, factor float64) int {
    return int(math.Max(math.Min(in * factor + 0.5, factor), 0))
}

func formatH3x(color colorful.Color) string {
    r := convert(color.R, 15)
    g := convert(color.G, 15)
    b := convert(color.B, 15)

    return fmt.Sprintf("#%x%x%x", r, g, b)
}

func formatHex(color colorful.Color) string {
    r := convert(color.R, 255)
    g := convert(color.G, 255)
    b := convert(color.B, 255)

    return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func formatRgb(color colorful.Color) string {
    r := convert(color.R, 255)
    g := convert(color.G, 255)
    b := convert(color.B, 255)

    return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}
