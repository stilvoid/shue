package shue

import (
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

type Colour struct {
	R, G, B, A float64
}

func scale(in, factor float64) int {
	return int(math.Min(in*factor+0.5, factor))
}

func (c *Colour) toColor() color.Color {
	return color.NRGBA{
		R: uint8(scale(c.R, 255)),
		G: uint8(scale(c.G, 255)),
		B: uint8(scale(c.B, 255)),
		A: uint8(scale(c.A, 255)),
	}
}

// HSL returns pre-multiplied HSL values 0..1
func (c *Colour) HSL() (float64, float64, float64) {
	r, g, b := c.RGB()

	col := colorful.Color{
		R: r,
		G: g,
		B: b,
	}

	h, s, l := col.Hsl()

	h /= 360

	return h, s, l
}

// HSL returns non-pre-multiplied HSL values 0..1
func (c *Colour) HSLA() (float64, float64, float64, float64) {
	r, g, b, a := c.RGBA()

	col := colorful.Color{
		R: r,
		G: g,
		B: b,
	}

	h, s, l := col.Hsl()

	h /= 360

	return h, s, l, a
}

// RGB returns pre-multiplied RGB values 0..1
func (c *Colour) RGB() (float64, float64, float64) {
	r, g, b, _ := c.toColor().RGBA()

	return float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff
}

// RGBA returns non-pre-multiplied RGBA values 0..1
func (c *Colour) RGBA() (float64, float64, float64, float64) {
	return c.R, c.G, c.B, c.A
}

func (c *Colour) Invert() {
	c.R = 1 - c.R
	c.G = 1 - c.G
	c.B = 1 - c.B
}

func (c *Colour) Lighten(value int) {
	col := colorful.Color{
		R: c.R,
		G: c.G,
		B: c.B,
	}

	h, s, l := col.Hsl()

	l = l * (float64(value) / 100)

	l = math.Min(math.Max(l, 0), 1.0)

	col = colorful.Hsl(h, s, l)

	c.R = col.R
	c.G = col.G
	c.B = col.B
}

func (c *Colour) Equal(other Colour) bool {
	a, _ := colorful.MakeColor(c.toColor())
	b, _ := colorful.MakeColor(other.toColor())

	return a.AlmostEqualRgb(b)
}
