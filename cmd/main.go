package main

import (
	"fmt"

	"github.com/stilvoid/shue"

	"github.com/spf13/cobra"
)

var version = "git"

var format string
var lighten int
var invert bool

func init() {
	rootCmd.Flags().StringVarP(&format, "format", "f", "all", "Output format. See --help for details")
	rootCmd.Flags().IntVarP(&lighten, "lighten", "l", 100, "Percentage by which to adjust the brightness")
	rootCmd.Flags().BoolVarP(&invert, "invert", "i", false, "Invert the colour value - lighten is applied first")

	rootCmd.Version = version
}

var rootCmd = &cobra.Command{
	Use:   "shue COLOUR",
	Short: "Shue converts and modifies colour values for use in CSS etc.",
	Long: `The following formats can be read and output by shue:

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

If you specify --lighten and --invert, inversion takes place _after_ lightening`,
	Example: `  shue -f rgb 112233
  shue -f rgb "rgba(12,34,45,0.67)"
  shue -i "#ff0000"
  shue -l 150 "hsl(120, 100%, 50%)"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		color, err := shue.Parse(input)
		cobra.CheckErr(err)

		if lighten != 100 {
			color.Lighten(lighten)
		}

		if invert {
			color.Invert()
		}

		out, err := shue.Format(format, color)
		cobra.CheckErr(err)

		fmt.Println(out)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}
