package edit

import (
	"image"
)

func Copy(img image.Image) image.Image {
	bounds := img.Bounds()
	size := bounds.Size()
	img2 := image.NewNRGBA(bounds)
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := img.At(x, y)
			img2.Set(x, y, color)
		}
	}
	return img2
}
