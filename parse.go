package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"

type config struct {
    format string
    factor float64
}

var configs []config

func init() {
    configs = []config {
        {"#%02x%02x%02x", 255},
        {"%02x%02x%02x", 255},
        {"#%1x%1x%1x", 15},
        {"%1x%1x%1x", 15},
        {"rgb(%d,%d,%d)", 255},
        {"rgb(%d %d %d)", 255},
        {"%d,%d,%d", 255},
        {"%d %d %d", 255},
    }
}

func parse(input string) (colorful.Color, error) {
    var r, g, b uint8

    for _, config := range(configs) {
        n, err := fmt.Sscanf(input, config.format, &r, &g, &b)
        if err != nil {
            continue
        }
        if n != 3 {
            continue
        }

        return colorful.Color {
            float64(r) / config.factor,
            float64(g) / config.factor,
            float64(b) / config.factor,
        }, nil
    }

    return colorful.Color{}, fmt.Errorf("not a valid color: %s", input)
}
