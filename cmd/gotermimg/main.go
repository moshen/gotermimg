package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	timg "github.com/moshen/gotermimg"
)

func main() {
	isUTF8 := flag.Bool("u", false, "Enable UTF8 output")
	width := flag.Uint("x", 0, `Scale to n*2 columns wide in ANSI mode, n columns wide in UTF8 mode.
        When -x=0 (the default), aspect ratio is maintained.
        For example if -x is provided without -y, height is scaled to
        maintain aspect ratio`)
	height := flag.Uint("y", 0, `Scale to n rows high in ANSI mode, n/2 rows high in UTF8 mode.
        When -y=0 (the default), aspect ratio is maintained.
        For example if -y is provided without -x, width is scaled to
        maintain aspect ratio`)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: gotermimg [-u|-x=n|-y=n] IMAGEFILE
  IMAGEFILE - png, gif or jpg.  gif will auto-play
`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, imgformat, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Seek(0, 0)

	var conv timg.Converter
	if *isUTF8 {
		conv = timg.UTF8
	} else {
		conv = timg.ANSI
	}

	var trans timg.Transformer
	if *width != 0 || *height != 0 {
		trans = timg.Resize(*width, *height)
	}

	if imgformat == "gif" {
		gifimg, err := gif.DecodeAll(file)
		if err != nil {
			log.Fatal(err)
		}

		if len(gifimg.Image) > 1 {
			timg.PrintAnimation(timg.Gif(gifimg, conv, trans))
		} else {
			timg.PrintImage(gifimg.Image[0], conv, trans)
		}
	} else {
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		timg.PrintImage(img, conv, trans)
	}
}
