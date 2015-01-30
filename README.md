
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

### gogotermimg

    Usage: gotermimg [-u] IMAGEFILE
      IMAGEFILE - png, gif or jpg.  gif will auto-play
      -u=false: Enable UTF8 output

[![gotermimg on a png with transparency](http://fat.gfycat.com/AbsoluteShockingHerring.gif)](http://gfycat.com/AbsoluteShockingHerring)

[![gotermimg on an animated gif with transparency](http://zippy.gfycat.com/IcyBlindBlesbok.gif)](http://gfycat.com/IcyBlindBlesbok)

### gogopher

    Usage: gogopher [-u]
    Prints a looping animation of the Go gopher looking shifty!
      -u=false: Enable UTF8 output

[![gogopher printing a shify go gopher](http://zippy.gfycat.com/ConsciousTimelyHuman.gif)](http://gfycat.com/ConsciousTimelyHuman)

## Author

[Colin Kennedy](https://github.com/moshen)

## License

[MIT](http://colken.mit-license.org/)
, See LICENSE file

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
The design is licensed under the Creative Commons 3.0 Attributions license.
Read this article for more details: http://blog.golang.org/gopher

