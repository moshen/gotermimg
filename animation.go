package gotermimg

import (
	"fmt"
	"time"

	// Vendor
	"gif"

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
	for i, v := range g.Explode() {
		ani[i] = Frame{conv(v, trans), time.Duration(g.Delay[i]) * 10 * time.Millisecond}
	}

	return ani
}

type KeepLooping func() bool

// Returns true until n calls have been made
func LoopTimes(n uint) KeepLooping {
	loopCount := uint(0)
	return func() bool {
		switch {
		case n == 0:
			return true
		case loopCount+1 >= n:
			return false
		default:
			loopCount++
			return true
		}
	}
}

// Returns true until s seconds have elapsed
func LoopSeconds(s uint) KeepLooping {
	var c <-chan time.Time
	return func() bool {
		if c == nil {
			c = time.After(time.Duration(s) * time.Second)
		}
		select {
		case <-c:
			return false
		default:
			return true
		}
	}
}

// Prints Animation ani to os.Stdout after clearing the terminal
// Does not return unless loop() returns false
func PrintAnimation(ani Animation, loop KeepLooping) {
	fmt.Print(terminal.Clear)
	for {
		for _, f := range ani {
			fmt.Print(terminal.Origin)
			for _, v := range f.ANSI {
				fmt.Println(v)
			}
			time.Sleep(f.Delay)
		}
		if loop != nil && !loop() {
			break
		}
	}
}
