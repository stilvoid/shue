# Shue

A command line tool for modifying and converting colour values for use with CSS etc.

## Installing

`brew install stilvoid/tools/shue`

_or_

Download a binary from the [releases](https://github.com/stilvoid/shue/releases) page.

_or_

Run `go install github.com/stilvoid/shue@latest`

## Usage

```
The following formats can be read and output by shue:

  Name    Example
  ---     ---
  hex3    #123
  hex4    #1234
  hex6    #123456
  hex8    #12345678
  rgb     rgb(12, 34, 56)
  rgba    rgba(12, 45, 56, 0.7)
  hsl     hsl(180, 25%, 50%)
  hsla    hsla(180, 25%, 50%, 0.75)

If the input format contains an alpha value (hex4,hex8,rgba,hsla)
but the output format does not (hex3,hex6,rgb,hsl)
then shue will pre-multiply the colour before outputting.

Shue can lighten a colour by specifying --lighten with a value 100+
or darken a colour by specifying a value below 100.

If you specify --lighten and --invert, inversion takes place _after_ lightening

Usage:
  shue COLOUR [flags]

Flags:
  -f, --format string   Output format. See --help for details (default "all")
  -h, --help            help for shue
  -i, --invert          Invert the colour value - lighten is applied first
  -l, --lighten int     Percentage by which to adjust the brightness (default 100)
  -v, --version         version for shue
```
