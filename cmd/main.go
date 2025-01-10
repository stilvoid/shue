package main

import (
	"fmt"
	"strings"

	"github.com/stilvoid/shue"

	"github.com/spf13/cobra"
)

var version = "git"

var format string
var lighten int
var invert bool

func init() {
	rootCmd.Flags().StringVarP(&format, "format", "f", "all", fmt.Sprintf("Format to output: %s", strings.Join(shue.Formats, ", ")))
	rootCmd.Flags().IntVarP(&lighten, "lighten", "l", 100, "Percentage by which to adjust the brightness")
	rootCmd.Flags().BoolVarP(&invert, "invert", "i", false, "Invert the colour value - lighten is applied first")

	rootCmd.Version = version
}

var rootCmd = &cobra.Command{
	Use:   "shue [COLOUR]",
	Short: "Shue converts and modifies colour values for use in CSS etc.",
	Args:  cobra.ExactArgs(1),
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
