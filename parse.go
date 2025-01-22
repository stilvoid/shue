package shue

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

type parseType int

const (
	hex parseType = iota
	rgb
	hsl
)

type config struct {
	re        *regexp.Regexp
	parseType parseType
	scale     float64
}

var configs []config

func init() {
	configs = []config{
		// RGB
		{regexp.MustCompile(`^#?([0-9a-fA-F])([0-9a-fA-F])([0-9a-fA-F])$`), hex, 15},
		{regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`), hex, 255},
		{regexp.MustCompile(`^rgb\((\d+),(\d+),(\d+)\)$`), rgb, 255},
		{regexp.MustCompile(`^hsl\((\d+),(\d+)%,(\d+)%\)$`), hsl, 100},

		// RGBA
		{regexp.MustCompile(`^#?([0-9a-fA-F])([0-9a-fA-F])([0-9a-fA-F])([0-9a-fA-F])$`), hex, 15},
		{regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`), hex, 255},
		{regexp.MustCompile(`^rgba\((\d+),(\d+),(\d+),([01](?:\.\d+)?)\)$`), rgb, 255},
		{regexp.MustCompile(`^hsla\((\d+),(\d+)%,(\d+)%,([01](?:\.\d+)?)\)$`), hsl, 100},
	}
}

func Parse(input string) (Colour, error) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	for _, config := range configs {
		match := config.re.FindStringSubmatch(input)
		if match != nil {
			parts := make([]float64, 0)

			for i, a := range match[1:] {
				base := 10
				if config.parseType == hex {
					base = 16
				}

				if config.parseType != hex && i == 3 {
					part, err := strconv.ParseFloat(a, 0)
					if err != nil {
						return Colour{}, err
					}

					parts = append(parts, part)
				} else {
					part, err := strconv.ParseInt(a, base, 0)
					if err != nil {
						return Colour{}, err
					}

					scale := config.scale

					if config.parseType == hsl && i == 0 {
						scale = 1
					}

					parts = append(parts, float64(part)/scale)
				}
			}

			if len(parts) == 3 {
				parts = append(parts, 1.0)
			}

			if len(parts) != 4 {
				return Colour{}, errors.New("Invalid input format")
			}

			if config.parseType == hsl {
				col := colorful.Hsl(parts[0], parts[1], parts[2])

				parts[0] = col.R
				parts[1] = col.G
				parts[2] = col.B
			}

			return Colour{parts[0], parts[1], parts[2], parts[3]}, nil
		}
	}

	return Colour{}, errors.New("Invalid input format")
}
