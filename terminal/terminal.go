// Package terminal provides the functions and constants required
// for 256 color terminal output and the conversion of color.Color structs
// to ColorCode
package terminal

import (
	"fmt"
	"image/color"
)

type ColorCode uint8

const (
	// Reset terminal colors
	Reset = "\033[0m"
	// Clear terminal
	Clear = "\033[2J"
	// Move the cursor to 0,0
	Origin = "\033[0;0H"
	// Foreground
	fgCode = uint8(38)
	// Background
	bgCode = uint8(48)
)

func colorize(fb uint8, i ColorCode, s string) string {
	return fmt.Sprintf("\033[%d;5;%dm%s", fb, i, s)
}

// Returns string s with ColorCode i applied to the foreground
func Fg(i ColorCode, s string) string {
	return colorize(fgCode, i, s)
}

// Returns string s with ColorCode i applied to the background
func Bg(i ColorCode, s string) string {
	return colorize(bgCode, i, s)
}

// Finds the closest ColorCode to color.Color c and returns it.
// If color.Color s is transparent, returns 0
func FindColorCode(c color.Color) ColorCode {
	_, _, _, alpha := c.RGBA()
	if alpha == 0 {
		return 0
	} else {
		return ColorLookup[TermPalette.Index(c)]
	}
}
