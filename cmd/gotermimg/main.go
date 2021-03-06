package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	// Vendor
	"gif"
	"termutil"

	timg "github.com/moshen/gotermimg"
	"github.com/moshen/gotermimg/terminal"
)

func main() {
	isUTF8 := flag.Bool("u", false, "Enable UTF8 output")
	width := flag.Uint("x", 0, `Scale to n*2 columns wide in ANSI mode, n columns wide in UTF8 mode.
        When -x=0 (the default), aspect ratio is maintained.
        For example if -y is provided without -x, width is scaled to
        maintain aspect ratio`)
	height := flag.Uint("y", 0, `Scale to n rows high in ANSI mode, n/2 rows high in UTF8 mode.
        When -y=0 (the default), aspect ratio is maintained.
        For example if -x is provided without -y, height is scaled to
        maintain aspect ratio`)
	loopTimes := flag.Uint("l", 0, `Loop animation n times
        When -l=0 (the default), animation is looped indefinitely. Supersedes -s
        Only applies to multi-frame gifs`)
	loopSeconds := flag.Uint("s", 0, `Loop animation n seconds
        When -s=0 (the default), this option is ignored.
        Only applies to multi-frame gifs`)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: gotermimg [-u] [-x=n] [-y=n] [-l=n|-s=n] [IMAGEFILE]
  IMAGEFILE - png, gif or jpg.  gif will auto-play.
  Image data can be piped to stdin instead of providing IMAGEFILE.

  If neither -x or -y are provided, and the image is larger than your current
  terminal, it will be automatically scaled to fit.

`)
		flag.PrintDefaults()
	}

	flag.Parse()

	var buf *bytes.Reader
	switch {
	case !termutil.Isatty(os.Stdin.Fd()):
		bufData, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		buf = bytes.NewReader(bufData)
	case len(flag.Args()) < 1:
		flag.Usage()
		os.Exit(1)
	default:
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		bufData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		buf = bytes.NewReader(bufData)
	}

	conf, imgformat, err := image.DecodeConfig(buf)
	if err != nil {
		log.Fatal(err)
	}
	buf.Seek(0, 0)

	var conv timg.Converter
	if *isUTF8 {
		conv = timg.UTF8
	} else {
		conv = timg.ANSI
	}

	var trans timg.Transformer
	if *width != 0 || *height != 0 {
		trans = timg.Resize(*width, *height)
	} else if termutil.Isatty(os.Stdout.Fd()) {
		x, y, err := terminal.Size(os.Stdout.Fd())
		if err != nil {
			log.Fatal(err)
		}

		y = y - 1

		// Convert the actual terminal dimensions into effective dimensions
		switch {
		case *isUTF8:
			y = y * 2
		case x%2 == 0:
			x = x / 2
		default:
			x = (x - 1) / 2
		}

		if uint(conf.Width) > x || uint(conf.Height) > y {
			aspectTerm := float32(x) / float32(y)
			aspectImg := float32(conf.Width) / float32(conf.Height)

			if aspectImg > aspectTerm {
				trans = timg.Resize(x, 0)
			} else {
				trans = timg.Resize(0, y)
			}
		}
	}

	if imgformat == "gif" {
		gifimg, err := gif.DecodeAll(buf)
		if err != nil {
			log.Fatal(err)
		}

		if len(gifimg.Image) > 1 {
			var loop timg.KeepLooping
			switch {
			// Don't bother looping if we're not outputting to a tty
			case !termutil.Isatty(os.Stdout.Fd()):
				loop = timg.LoopTimes(1)
			case *loopTimes > 0:
				loop = timg.LoopTimes(*loopTimes)
			case *loopSeconds > 0:
				loop = timg.LoopSeconds(*loopSeconds)
			}

			timg.PrintAnimation(timg.Gif(gifimg, conv, trans), loop)
		} else {
			timg.PrintImage(gifimg.Image[0], conv, trans)
		}
	} else {
		img, _, err := image.Decode(buf)
		if err != nil {
			log.Fatal(err)
		}
		timg.PrintImage(img, conv, trans)
	}
}
