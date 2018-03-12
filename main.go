package main

import (
    "os"
    "strings"
    "flag"
    "github.com/lucasb-eyer/go-colorful"
    "log"
    "fmt"
    "io/ioutil"
)

var lighten *int
var invert  *bool

func printUsage() {
    fmt.Println("Usage: shue [OPTIONS] FORMAT")
    fmt.Println()
    flag.PrintDefaults()
    fmt.Println()
}

func init() {
    flag.Usage = printUsage
    lighten = flag.Int("l", 100, "Percentage to lighten by")
    invert = flag.Bool("i", false, "Invert the colour")
}

func main() {
    flag.Parse()

    if flag.NArg() != 1 {
        printUsage()
        os.Exit(1)
    }

    formatter, ok := formatters[flag.Arg(0)]
    if !ok {
        log.Fatalf("Unknown formatter: %s\n", flag.Arg(0))
    }

    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }

    input := strings.TrimSpace(string(bytes))

    for _, parser := range(parsers) {
        color, err := parser(input)
        if err != nil {
            continue
        }

        if *invert {
            h, s, v := color.Hsv()
            h = h + 180 % 360

            color = colorful.Hsv(h, s, v)
        }

        if *lighten != 100 {
            h, s, v := color.Hsv()
            v = v * (float64(*lighten)/100)

            color = colorful.Hsv(h, s, v)
        }

        fmt.Println(formatter(color))
        os.Exit(0)
    }

    fmt.Println("DEAD")
    os.Exit(1)
}
