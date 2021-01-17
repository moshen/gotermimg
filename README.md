
# gotermimg
[![Build Status](https://travis-ci.org/moshen/gotermimg.svg?branch=master)](https://travis-ci.org/moshen/gotermimg)

Something hacked together to display images and play gifs in 256 color
terminals.  I have
[implemented something similar before](https://github.com/moshen/Image-Term256Color)
, as have others...

Wrote this largely as an exercise in Go.

## Installation

Install using `go get`:

    go get github.com/moshen/gotermimg/...

Installs the `gotermimg` command line application.

## Usage

### gotermimg

    Usage: gotermimg [-u] [-x=n] [-y=n] [-l=n|-s=n] [IMAGEFILE]
      IMAGEFILE - png, gif or jpg.  gif will auto-play.
      Image data can be piped to stdin instead of providing IMAGEFILE.

      If neither -x or -y are provided, and the image is larger than your current
      terminal, it will be automatically scaled to fit.

      -l=0: Loop animation n times
            When -l=0 (the default), animation is looped indefinitely. Supersedes -s
            Only applies to multi-frame gifs
      -s=0: Loop animation n seconds
            When -s=0 (the default), this option is ignored.
            Only applies to multi-frame gifs
      -u=false: Enable UTF8 output
      -x=0: Scale to n*2 columns wide in ANSI mode, n columns wide in UTF8 mode.
            When -x=0 (the default), aspect ratio is maintained.
            For example if -y is provided without -x, width is scaled to
            maintain aspect ratio
      -y=0: Scale to n rows high in ANSI mode, n/2 rows high in UTF8 mode.
            When -y=0 (the default), aspect ratio is maintained.
            For example if -x is provided without -y, height is scaled to
            maintain aspect ratio

[![gotermimg on a png with transparency](https://media.giphy.com/media/vpYeVwn2cRxstBp5hS/giphy.gif)](https://media.giphy.com/media/vpYeVwn2cRxstBp5hS/giphy.gif)

[![gotermimg on an animated gif with transparency](https://media.giphy.com/media/b9sXmD1dWBUvbgr87r/giphy.gif)](https://media.giphy.com/media/b9sXmD1dWBUvbgr87r/giphy.gif)

While the render speed on some slower terminals might not look very good, urxvt
looks amazing (click through for HQ).

[![gotermimg on urxvt](https://media.giphy.com/media/Jsg9KArYyntBPgoH4o/giphy.gif)](https://media.giphy.com/media/Jsg9KArYyntBPgoH4o/giphy.gif)

## Author

[Colin Kennedy](https://github.com/moshen)

### vendor/gif

[Go AUTHORS](https://golang.org/AUTHORS)

### vendor/resize

[Jan Schlicht](https://github.com/nfnt/resize)

### vendor/termutil

[Andrew](https://github.com/andrew-d/go-termutil)

## License

[MIT](http://colken.mit-license.org/)
(unless otherwise noted), See LICENSE file

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
The design is licensed under the Creative Commons 3.0 Attributions license.
Read this article for more details: http://blog.golang.org/gopher

### vendor/gif

[BSD](https://github.com/moshen/gotermimg/blob/master/vendor/gif/LICENSE)

### vendor/resize

[MIT Style](https://github.com/moshen/gotermimg/blob/master/vendor/resize/LICENSE)

### vendor/termutil

[MIT](https://github.com/moshen/gotermimg/blob/master/vendor/termutil/LICENSE)
