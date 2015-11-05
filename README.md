# Shue

A simple tool for converting rgb colour values into the nearest equivalent bash shell equivalents.

## Usage

    shue [-16|-256] <colour>

    -16     Specify a 16-colour terminal (default)
    -256    Specify a 16-colour terminal

    <colour> must be in one of the following formats:
        rgb
        rrggbb
        #rgb
        #rrggbb

    Note: If you include a hash, you'll need to quote the colour

        e.g. shue -256 "#abc"

The command will output the equivalent terminal colour number which you can then use in vim/mutt/Xresources definitions.
