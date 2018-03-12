package main

import (
)

func ExampleHexFormat() {
    shue("ff8001", "hex", 100, false)
    // Output: #ff8001
}

func ExampleH3xFormat() {
    shue("ff8001", "h3x", 100, false)
    // Output: #f80
}

func ExampleRgbFormat() {
    shue("ff8001", "rgb", 100, false)
    // Output: rgb(255, 128, 1)
}

func ExampleAllFormat() {
    shue("ff8001", "all", 100, false)
    // Output:
    // #f80
    // #ff8001
    // rgb(255, 128, 1)
}

func ExampleInvert() {
    // Output:

    shue("fff", "hex", 100, true)
    // #000000

    shue("000", "hex", 100, true)
    // #ffffff

    shue("0180f0", "hex", 100, true)
    // #fe7f0f
}

func ExampleLighten() {
    // Output:

    shue("000", "hex", 200, false)
    // #000000

    shue("fff", "hex", 200, false)
    // #ffffff

    shue("808080", "hex", 200, false)
    // #ffffff

    shue("804020", "hex", 200, false)
    // #ff8040
}

func ExampleDarken() {
    // Output:

    shue("000", "hex", 50, false)
    // #000000

    shue("fff", "hex", 50, false)
    // #808080

    shue("808080", "hex", 50, false)
    // #404040

    shue("ff8060", "hex", 50, false)
    // #804030
}
