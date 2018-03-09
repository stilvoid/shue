package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"

type formatter func(colorful.Color)string
var formatters map[string]formatter

func init() {
    formatters = map[string]formatter {
        "hex": formatHex,
        "h3x": formatH3x,
        "rgb": formatRgb,
        "rgb256": formatRgb256,
        "hsv": formatHsv,
    }
}

func formatHex(color colorful.Color) string {
    return fmt.Sprintf("%s", color.Hex())
}

func formatH3x(color colorful.Color) string {
    r := uint8(color.R*15.0+0.5)
    g := uint8(color.G*15.0+0.5)
    b := uint8(color.B*15.0+0.5)

    return fmt.Sprintf("#%x%x%x", r, g, b)
}

func formatRgb(color colorful.Color) string {
    return fmt.Sprintf("rgb(%f, %f, %f)", color.R, color.G, color.B)
}

func formatRgb256(color colorful.Color) string {
    r, g, b := color.RGB255()

    return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

func formatHsv(color colorful.Color) string {
    h, s, v := color.Hsv()

    return fmt.Sprintf("hsv(%f, %f, %f)", h, s, v)
}
