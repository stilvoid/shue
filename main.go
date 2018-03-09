package main

import "fmt"
import "log"
import "github.com/lucasb-eyer/go-colorful"

func main() {
    var color colorful.Color
    var output string
    var err error

    // #ff804d
    pairs := [][]string {
        {"hex", "#f85"},
        {"hex", "#ff804d"},
        {"rgb", "rgb(1.0, 0.5, 0.3)"},
        {"rgb", "rgb(255, 127, 77)"},
        {"hsv", "hsv(17.142857, 0.700000, 1.000000)"},
        {"hsv", "hsv(17, 0, 1)"},
        {"rgb", "rgb(1, 0, 0)"},
    }

    formats := []string {
        "hex",
        "h3x",
        "rgb",
        "rgb256",
        "hsv",
    }

    for _, pair := range(pairs) {
        name, input := pair[0], pair[1]

        fmt.Printf("%s: %s ->\n", name, input)

        parser, ok := parsers[name]
        if !ok {
            log.Fatal(fmt.Errorf("Parser does not exist: %s\n", name))
        }

        color, err = parser(input)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%#v\n", color)

        for _, format := range(formats) {
            formatter, ok := formatters[format]
            if !ok {
                log.Fatal(fmt.Errorf("Formatter does not exist: %s\n", format))
            }

            output = formatter(color)

            fmt.Printf("  %s: %s\n", format, output)
        }

        fmt.Printf("\n")
    }
}
