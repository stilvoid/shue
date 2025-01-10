# Shue

A command line tool for modifying and converting colour values for use with CSS etc.

## Installing

`brew install stilvoid/tools/shue`

_or_

Download a binary from the [releases](https://github.com/stilvoid/shue/releases) page.

_or_

Run `go install github.com/stilvoid/shue@latest`

## Supported formats

* Six-digit hex, e.g. `#ff8000`
* Three-digit hex, e.g. `#f80`
* Three-part RGB, e.g. `rgb(255, 128, 0)`

## Supported operations

* Lighten n%
* Set value
* Invert

## Usage

    Usage: shue [OPTIONS] [COLOUR] FORMAT

      Converts COLOUR to FORMAT. COLOUR must be in one of the supported formats.

    Supported formats:
      hex:      #rrggbb
      h3x:      #rgb
      rgb:      rgb(red, green, blue)

    Options:
      -l PERCENT    Lighten COLOUR by PERCENT% before converting to FORMAT
      -i            Invert COLOUR before converting to FORMAT

## Examples

Converting a colour to a specific format:

    $ shue -f rgb ff8800
    rgb(255, 136, 0)

Inverting a colour:

    $ shue -i ff8800
    #07f
    #0077ff
    rgb(0, 119, 255)

Lightening a colour:

    $ shue -l 200 884400
    #f80
    #ff8800
    rgb(255, 136, 0)

Darkening a colour:

    $ shue -l 50 884400
    #420
    #442200
    rgb(68, 34, 0)
