package shue_test

import (
	"fmt"

	"github.com/stilvoid/shue"
)

func exec(input string, format string, lighten int, invert bool) {
	c, _ := shue.Parse(input)

	c.Lighten(lighten)

	if invert {
		c.Invert()
	}

	out, _ := shue.Format(format, c)

	fmt.Println(out)
}

func ExampleColour_Invert() {
	exec("fff", "hex6", 100, true)
	exec("000", "hex6", 100, true)
	exec("0180f0", "hex6", 100, true)

	// Output:
	// #000000
	// #ffffff
	// #fe7f0f
}

func ExampleColour_Lighten_lighter() {
	exec("000", "hex6", 150, false)
	exec("fff", "hex6", 150, false)
	exec("808080", "hex6", 150, false)
	exec("804020", "hex6", 150, false)

	// Output:
	// #000000
	// #ffffff
	// #c0c0c0
	// #c06030
}

func ExampleColour_Lighten_darker() {
	exec("000", "hex6", 50, false)
	exec("fff", "hex6", 50, false)
	exec("808080", "hex6", 50, false)
	exec("ff8060", "hex6", 50, false)

	// Output:
	// #000000
	// #808080
	// #404040
	// #b02300
}

func ExampleColour_Lighten_and_invert() {
	exec("888", "hex6", 150, true)
	exec("444", "hex6", 200, true)
	exec("fff", "hex6", 20, true)

	// Output:
	// #333333
	// #777777
	// #cccccc
}

func Example_no_change() {
	exec("123", "hex3", 100, false)
	exec("1234", "hex4", 100, false)
	exec("123456", "hex6", 100, false)
	exec("12345678", "hex8", 100, false)
	exec("rgb(32,64,128)", "rgb", 100, false)
	exec("rgba(32,64,128,0.125)", "rgba", 100, false)
	exec("hsl(120,50%,66%)", "hsl", 100, false)
	exec("hsla(120,50%,66%,0.1)", "hsla", 100, false)

	// Output:
	// #123
	// #1234
	// #123456
	// #12345678
	// rgb(32, 64, 128)
	// rgba(32, 64, 128, 0.125)
	// hsl(120, 50%, 66%)
	// hsla(120, 50%, 66%, 0.1)
}
