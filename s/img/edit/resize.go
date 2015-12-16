package edit

import (
	"image"
)

// imgを0,0-x,yにリサイズする
func Resize(img *image.Image, x int, y int) image.Image {
	img2 := image.NewNRGBA(image.Rect(0, 0, x, y))
	size := (*img).Bounds().Size()
	size2 := (*img2).Bounds().Size()
	for xx := 0; xx < size2.X; xx++ {
		for yy := 0; yy < size2.Y; yy++ {
			atX := int(float64(xx) / float64(size2.X) * float64(size.X))
			atY := int(float64(yy) / float64(size2.Y) * float64(size.Y))
			c := (*img).At(atX, atY)
			r, g, b, _ := c.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				img2.Set(xx, yy, c)
			}
		}
	}
	return img2
}
