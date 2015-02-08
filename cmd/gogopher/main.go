package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"image/gif"
	"log"
	"os"
	"regexp"

	timg "github.com/moshen/gotermimg"
)

const (
	gophergif = `
    H4sIAO7zx1QAA3P3dLOwTBRgEGD4zMCQceOHhoXDjx8/NlQkCHBwSCgoSEgIZ
    Kw4kLFgw////xkwgOJ/bj/XkGBnxwBXIz0DZkaQ0E8WTiMGTgYdkDzIWAaWtQ
    ZLpwR5aQKJJhDBodgiwOSjuVTECaREgUGxRYSBAchmFGBiEHBicBKIYBBoEhF
    w4nJwcORicJrA5OTFyKTAYsDGwejgYcAK1JMMxA2KS4GkcwcjA4vCEqABDDIs
    CgwMLVOYQIYCHdIwIQjI4AC7IoiLwQHiYJBREBbUKAa4USAANIqDch+0cmgwa
    LhwAH2wwEGJBchlJM4HC5ip5ANrADjNEJTVAQAA`
)

var (
	cleanbase64 = regexp.MustCompile(`\s+`)
)

func main() {
	isUTF8 := flag.Bool("u", false, "Enable UTF8 output")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: gogopher [-u]
Prints a looping animation of the Go gopher looking shifty!
`)
		flag.PrintDefaults()
	}

	flag.Parse()

	data, err := base64.StdEncoding.DecodeString(cleanbase64.ReplaceAllString(gophergif, ""))
	if err != nil {
		log.Fatal(err)
	}

	read, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	gifimg, err := gif.DecodeAll(read)
	if err != nil {
		log.Fatal(err)
	}

	var conv timg.Converter
	if *isUTF8 {
		conv = timg.UTF8
	} else {
		conv = timg.ANSI
	}

	timg.PrintAnimation(timg.Gif(gifimg, conv, nil), nil)
}
