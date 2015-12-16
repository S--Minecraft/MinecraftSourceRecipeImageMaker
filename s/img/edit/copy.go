package edit

import (
	"image"
	"image/draw"
)

func Copy(img *image.Image) image.Image {
	img2 := image.NewNRGBA((*img).Bounds())
	draw.Draw(img2, img2.Bounds(), *img, image.ZP, draw.Src)
	return img2
}
