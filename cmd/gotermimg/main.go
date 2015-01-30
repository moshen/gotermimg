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

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: gotermimg [-u] IMAGEFILE
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

	if imgformat == "gif" {
		gifimg, err := gif.DecodeAll(file)
		if err != nil {
			log.Fatal(err)
		}

		if len(gifimg.Image) > 1 {
			timg.PrintAnimation(timg.Gif(gifimg, conv))
		} else {
			timg.PrintImage(gifimg.Image[0], conv)
		}
	} else {
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		timg.PrintImage(img, conv)
	}
}
