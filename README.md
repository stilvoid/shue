# Shue

A command line tool for modifying and converting colour values for use with CSS etc.

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
