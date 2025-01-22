package shue

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
)

type formatter func(Colour) string

var Formats []string

var formatters map[string]formatter

func init() {
	formatters = map[string]formatter{
		// RGB
		"hex3": formatHex3,
		"hex6": formatHex6,
		"rgb":  formatRgb,
		"hsl":  formatHsl,

		// RGBA
		"hex4": formatHex4,
		"hex8": formatHex8,
		"rgba": formatRgba,
		"hsla": formatHsla,

		// All
		"all": formatAll,
	}

	Formats = slices.Collect(maps.Keys(formatters))
	sort.Strings(Formats)
}

func Format(format string, colour Colour) (string, error) {
	formatter, ok := formatters[format]
	if !ok {
		return "", fmt.Errorf("Unknown formatter: %s", format)
	}

	return formatter(colour), nil
}

func formatAll(color Colour) string {
	parts := make([]string, 0)

	for _, format := range Formats {
		if format != "all" {
			s, _ := Format(format, color)

			parts = append(parts, s)
		}
	}

	return strings.Join(parts, "\n")
}

func formatHex3(color Colour) string {
	r, g, b := color.RGB()

	return fmt.Sprintf("#%x%x%x", scale(r, 15), scale(g, 15), scale(b, 15))
}

func formatHex4(color Colour) string {
	r := scale(color.R, 15)
	g := scale(color.G, 15)
	b := scale(color.B, 15)
	a := scale(color.A, 15)

	return fmt.Sprintf("#%x%x%x%x", r, g, b, a)
}

func formatHex6(color Colour) string {
	r, g, b := color.RGB()

	return fmt.Sprintf("#%02x%02x%02x", scale(r, 255), scale(g, 255), scale(b, 255))
}

func formatHex8(color Colour) string {
	r := scale(color.R, 255)
	g := scale(color.G, 255)
	b := scale(color.B, 255)
	a := scale(color.A, 255)

	return fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
}

func formatRgb(color Colour) string {
	r, g, b := color.RGB()

	return fmt.Sprintf("rgb(%d, %d, %d)", scale(r, 255), scale(g, 255), scale(b, 255))
}

func formatRgba(color Colour) string {
	r, g, b, a := color.RGBA()

	return fmt.Sprintf("rgba(%d, %d, %d, %g)", scale(r, 255), scale(g, 255), scale(b, 255), a)
}

func formatHsl(color Colour) string {
	h, s, l := color.HSL()

	return fmt.Sprintf("hsl(%d, %d%%, %d%%)", scale(h, 360), scale(s, 100), scale(l, 100))
}

func formatHsla(color Colour) string {
	h, s, l, a := color.HSLA()

	return fmt.Sprintf("hsla(%d, %d%%, %d%%, %g)", scale(h, 360), scale(s, 100), scale(l, 100), a)
}
