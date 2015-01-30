
# gotermimg

Something hacked together to display images and play gifs in 256 color
terminals.  I have
[implemented something similar before](https://github.com/moshen/Image-Term256Color)
, as have others...

Wrote this largely as an exercise in Go.

## Installation

Install using `go get`:

    go get github.com/moshen/gotermimg/...

Installs the `gotermimg` and `gogopher` command line applications.

## Usage

### gotermimg

    Usage: gotermimg [-u|-x=n|-y=n] IMAGEFILE
      IMAGEFILE - png, gif or jpg.  gif will auto-play
      -u=false: Enable UTF8 output
      -x=0: Scale to n*2 columns wide in ANSI mode, n columns wide in UTF8 mode.
            When -x=0 (the default), aspect ratio is maintained.
            For example if -x is provided without -y, height is scaled to
            maintain aspect ratio
      -y=0: Scale to n rows high in ANSI mode, n/2 rows high in UTF8 mode.
            When -y=0 (the default), aspect ratio is maintained.
            For example if -y is provided without -x, width is scaled to
            maintain aspect ratio

[![gotermimg on a png with transparency](http://fat.gfycat.com/AbsoluteShockingHerring.gif)](http://gfycat.com/AbsoluteShockingHerring)

[![gotermimg on an animated gif with transparency](http://zippy.gfycat.com/IcyBlindBlesbok.gif)](http://gfycat.com/IcyBlindBlesbok)

### gogopher

    Usage: gogopher [-u]
    Prints a looping animation of the Go gopher looking shifty!
      -u=false: Enable UTF8 output

[![gogopher printing a shify go gopher](http://zippy.gfycat.com/ConsciousTimelyHuman.gif)](http://gfycat.com/ConsciousTimelyHuman)

## Author

[Colin Kennedy](https://github.com/moshen)

### vendor/resize

[Jan Schlicht](https://github.com/nfnt/resize)

## License

[MIT](http://colken.mit-license.org/)
(unless otherwise noted), See LICENSE file

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
The design is licensed under the Creative Commons 3.0 Attributions license.
Read this article for more details: http://blog.golang.org/gopher

### vendor/resize

[MIT Style](https://github.com/moshen/gotermimg/blob/master/vendor/resize/LICENSE)
