package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"

type parseFunc func(string)(colorful.Color, error)
var parsers map[string]parseFunc

func init() {
    parsers = map[string]parseFunc {
        "hex": parseHex,
        "rgb": parseRgb,
        "hsv": parseHsv,
    }
}

func parseHex(input string) (colorful.Color, error) {
    return colorful.Hex(input)
}

func parseRgb(input string) (colorful.Color, error) {
    var color colorful.Color
    var n int
    var err error

    n, err = fmt.Sscanf(input, "rgb(%f, %f, %f)", &color.R, &color.G, &color.B)
    if err == nil && n == 3 {
        return color, nil
    }

    var r, g, b int

    n, err = fmt.Sscanf(input, "rgb(%d, %d, %d)", &r, &g, &b)

    if err != nil {
        return color, err
    }

    if n != 3 {
        return color, fmt.Errorf("Badly formatted rgb value: %s", input)
    }

    color.R = float64(r) / 255
    color.G = float64(g) / 255
    color.B = float64(b) / 255

    return color, nil
}

func parseHsv(input string) (colorful.Color, error) {
    var h, s, v float64

    n, err := fmt.Sscanf(input, "hsv(%f, %f, %f)", &h, &s, &v)

    if err != nil {
        return colorful.Color{}, err
    }

    if n != 3 {
        return colorful.Color{}, fmt.Errorf("Badly formatted hsv value: %s", input)
    }

    return colorful.Hsv(h, s, v), nil
}
