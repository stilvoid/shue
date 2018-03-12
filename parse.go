package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"

type parseFunc func(string)(colorful.Color, error)
var parsers map[string]parseFunc

func init() {
    parsers = map[string]parseFunc {
        "hex": parseHex,
        "h3x": parseHex,
        "rgb": parseRgb,
    }
}

func parseHex(input string) (colorful.Color, error) {
	var color colorful.Color
    var format string
    var factor float64

    switch len(input) {
    case 7:
        format = "#%02x%02x%02x"
        factor = 1.0/255.0
    case 6:
        format = "%02x%02x%02x"
        factor = 1.0/255.0
    case 4:
        format = "#%1x%1x%1x"
        factor = 1.0/15.0
    case 3:
        format = "%1x%1x%1x"
        factor = 1.0/15.0
    default:
        return color, fmt.Errorf("color: %v is not a hex-color", input)
    }

    var r, g, b uint8

    n, err := fmt.Sscanf(input, format, &r, &g, &b)
    if err != nil {
        return color, err
    }
    if n != 3 {
        return color, fmt.Errorf("color: %v is not a hex-color", input)
    }

	color.R = float64(r) * factor
	color.G = float64(g) * factor
	color.B = float64(b) * factor

	return color, nil
}

func parseRgb(input string) (colorful.Color, error) {
    var color colorful.Color
    var err error
    var n, r, g, b int

    formats := []string {
        "%d %d %d",
        "%d,%d,%d",
        "rgb(%d %d %d)",
        "rgb(%d,%d,%d)",
    }

    for _, format := range(formats) {
        n, err = fmt.Sscanf(input, format, &r, &g, &b)
        if err != nil {
            continue
        }
        if n != 3 {
            continue
        }

        color.R = float64(r) / 255
        color.G = float64(g) / 255
        color.B = float64(b) / 255

        return color, nil
    }

    return color, fmt.Errorf("Badly formatted rgb value: %s", input)
}
