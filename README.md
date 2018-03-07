# Shue

A command line tool for modifying and converting colour values in various formats.

## Supported formats

* Hexadecimal
    * `hex`: e.g. `#ff8000`
    * `h3x`: e.g. `#f80`

* Decimal
    * `rgb`: e.g. `rgb(255, 128, 0)`
    * `hsv`: e.g. `hsv(255, 128, 0)`

* Bash
    * `bash16`: e.g. `3`
    * `bash256`: e.g. `208`

## Supported operations

* Lighten n%
* Set value
* Invert

## Usage

    Usage: shue [COLOUR] [OPTIONS] [-to] FORMAT

      Converts COLOUR to FORMAT. COLOUR must be in one of the supported formats.

    Supported formats:
      hex:      #rrggbb
      h3x:      #rgb
      rgb:      rgb(red, green, blue)
      hsv:      hsv(hue, saturation, value)
      bash16:   colour
      bash256:  colour

    Options:
      -l PERCENT    Lighten COLOUR by PERCENT% before converting to FORMAT
      -v VALUE      Set COLOUR's hsv value to VALUE before converting to FORMAT
      -i            Invert COLOUR before converting to FORMAT
      -r            Output raw values without the formatting, i.e.
                      hex:      rrggbb
                      h3x:      rgb
                      rgb:      red green blue
                      hsv:      hue saturation value
                      bash16:   colour
                      bash256:  colour
