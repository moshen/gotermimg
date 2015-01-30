// Package gotermimg provides functions to convert image.Image structs
// to 256 color terminal compatible []string.
// It also provides functions to convert and "play" gif animations.
package gotermimg

import (
	"fmt"
	"image"

	"github.com/moshen/gotermimg/terminal"
)

type Converter func(image.Image, Transformer) []string

const (
	space  = " "
	spaces = "  "
	top    = "▀"
	bottom = "▄"
)

// Converts image.Image img to a []string of 256 color terminal compatbile
// ANSI using 2 spaces as a "pixel".  Applies Transformer trans to img
// before conversion
func ANSI(img image.Image, trans Transformer) []string {
	if trans != nil {
		img = trans(img)
	}

	bounds := img.Bounds()
	termimg := make([]string, bounds.Dy(), bounds.Dy())

	for i := 0; i < bounds.Dy(); i++ {
		row := ""
		local := ""
		curcolor := terminal.ColorCode(0)

		for j := 0; j < bounds.Dx(); j++ {
			newcolor := terminal.FindColorCode(img.At(j, i))

			if newcolor != curcolor {
				if curcolor > 0 {
					row += terminal.Bg(curcolor, local)
				} else {
					row += terminal.Reset + local
				}
				curcolor = newcolor
				local = ""
			}

			local += spaces
		}

		if curcolor > 0 {
			row += terminal.Bg(curcolor, local)
		} else {
			row += terminal.Reset + local
		}
		termimg[i] = row + terminal.Reset
	}

	return termimg
}

// Converts image.Image img to a []string of 256 color terminal compatbile
// UTF8 using UTF8 1/2 blocks as a "pixel".  Applies Transformer trans to img
// before conversion
func UTF8(img image.Image, trans Transformer) []string {
	if trans != nil {
		img = trans(img)
	}

	bounds := img.Bounds()
	termimg := make([]string, (bounds.Dy()/2)+1, (bounds.Dy()/2)+1)
	rownum := 0

	for i := 0; i < bounds.Dy(); i += 2 {
		row := ""

		for j := 0; j < bounds.Dx(); j++ {
			colortop := terminal.FindColorCode(img.At(j, i))
			colorbottom := terminal.FindColorCode(img.At(j, i+1))

			if colortop == colorbottom {
				if colortop > 0 {
					row += terminal.Bg(colortop, space)
				} else {
					row += terminal.Reset + space
				}
			} else {
				if colortop > 0 && colorbottom > 0 {
					row += terminal.Bg(colorbottom, terminal.Fg(colortop, top))
				} else if colortop > 0 && colorbottom == 0 {
					row += terminal.Reset + terminal.Fg(colortop, top)
				} else if colortop == 0 && colorbottom > 0 {
					row += terminal.Reset + terminal.Fg(colorbottom, bottom)
				}
			}
		}

		termimg[rownum] = row + terminal.Reset
		rownum++
	}

	return termimg
}

// Prints image.Image img to os.Stdout using Converter with Transformer trans
func PrintImage(img image.Image, conv Converter, trans Transformer) {
	for _, v := range conv(img, trans) {
		fmt.Println(v)
	}
}
