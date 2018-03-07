package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"

func printHex(color colorful.Color) {
    fmt.Printf("%s\n", color.Hex())
}

func printH3x(color colorful.Color) {
    r := uint8(color.R*15.0+0.5)
    g := uint8(color.G*15.0+0.5)
    b := uint8(color.B*15.0+0.5)

    fmt.Printf("#%x%x%x\n", r, g, b)
}

func printRgb(color colorful.Color) {
    fmt.Printf("rgb(%f, %f, %f)\n", color.R, color.G, color.B)
}

func printRgb256(color colorful.Color) {
    r, g, b := color.RGB255()

    fmt.Printf("rgb(%d, %d, %d)\n", r, g, b)
}

func printHsv(color colorful.Color) {
    h, s, v := color.Hsv()

    fmt.Printf("hsv(%f, %f, %f)\n", h, s, v)
}

func parseHex(input string) (colorful.Color) {
    color, err := colorful.Hex(input)

    if err != nil {
        fmt.Errorf("Not a valid hex value: %s\n", input)
    }

    return color
}

func parseRgb(input string) (colorful.Color) {
    var color colorful.Color

    fmt.Sscanf(input, "rgb(%f, %f, %f)", &color.R, &color.G, &color.B)

    return color
}

func parseRgb255(input string) (colorful.Color) {
    var color colorful.Color
    var r, g, b int

    fmt.Sscanf(input, "rgb(%d, %d, %d)", &r, &g, &b)

    color.R = float64(r) / 255
    color.G = float64(g) / 255
    color.B = float64(b) / 255

    return color
}

func parseHsv(input string) (colorful.Color) {
    var h, s, v float64

    fmt.Sscanf(input, "hsv(%f, %f, %f)", &h, &s, &v)

    return colorful.Hsv(h, s, v)
}

func main() {
    // #ff804d
    colors := []colorful.Color {
        parseHex("#f85"),
        parseHex("#ff804d"),
        parseRgb("rgb(1.0, 0.5, 0.3)"),
        parseRgb255("rgb(255, 127, 77)"),
        parseHsv("hsv(17.142857, 0.700000, 1.000000)"),
    }

    for _, color := range(colors) {
        fmt.Printf("%v\n", color)
        printHex(color)
        printH3x(color)
        printRgb(color)
        printRgb256(color)
        printHsv(color)
        fmt.Printf("\n")
    }
}
