package gotermimg

import (
	"fmt"
	"image/gif"
	"time"

	"github.com/moshen/gotermimg/terminal"
)

// Defines a frame of an Animation
type Frame struct {
	ANSI  []string
	Delay time.Duration
}

type Animation []Frame

// Converts *git.GIF g to an Animation using Converter conv
// Currently assumes all gif frames are the same size
func Gif(g *gif.GIF, conv Converter, trans Transformer) Animation {
	giflen := len(g.Image)
	ani := make(Animation, giflen, giflen)
	for i, v := range g.Image {
		ani[i] = Frame{conv(v, trans), time.Duration(g.Delay[i]) * 10 * time.Millisecond}
	}

	return ani
}

// Prints Animation ani to os.Stdout after clearing the terminal
// Does not return!  Loops indefinately
func PrintAnimation(ani Animation) {
	fmt.Print(terminal.Clear)
	for {
		for _, f := range ani {
			fmt.Print(terminal.Origin)
			for _, v := range f.ANSI {
				fmt.Println(v)
			}
			time.Sleep(f.Delay)
		}
	}
}
