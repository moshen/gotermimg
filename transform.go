package gotermimg

import (
	"image"

	"github.com/moshen/gotermimg/vendor/resize"
)

type Transformer func(image.Image) image.Image

func Resize(width, height uint) Transformer {
	return func(img image.Image) image.Image {
		return resize.Resize(width, height, img, resize.NearestNeighbor)
	}
}
