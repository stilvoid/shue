package main

import (
    "os"
    "strings"
    "flag"
    "github.com/lucasb-eyer/go-colorful"
    "log"
    "fmt"
)

func printUsage() {
    fmt.Println("Usage: shue [OPTIONS] COLOUR")
    fmt.Println()
    flag.PrintDefaults()
    fmt.Println()
}

func shue(input string, format string, lighten int, invert bool) {
    formatter, ok := formatters[format]
    if !ok {
        log.Fatalf("Unknown formatter: %s\n", format)
    }

    color, err := parse(input)
    if err != nil {
        log.Fatal(err)
    }

    if invert {
        color.R = 1 - color.R
        color.G = 1 - color.G
        color.B = 1 - color.B
    }

    if lighten != 100 {
        h, s, v := color.Hsv()
        v = v * (float64(lighten)/100)

        color = colorful.Hsv(h, s, v)
    }

    fmt.Println(formatter(color))
}

func main() {
    format := flag.String("f", "all", fmt.Sprintf(
        "Format to output: %s",
        strings.Join(formats, ", "),
    ))
    lighten := flag.Int("l", 100, "Percentage by which to adjust the brightness.")
    invert := flag.Bool("i", false, "Invert the colour")

    flag.Usage = printUsage
    flag.Parse()

    if flag.NArg() != 1 {
        printUsage()
        os.Exit(1)
    }

    input := flag.Arg(0)

    shue(input, *format, *lighten, *invert)
}
