package edit

import (
	"image"
	"math"
)

// imgをangle°だけ回転させた画像を返す
func Rotate(img *image.Image, angle int) image.Image {
	img2 := image.NewNRGBA((*img).Bounds())
	img2Size := (*img2).Bounds().Size()
	r := math.Pi * 2 * float64(angle) / 360
	x := img2Size.X / 2
	y := img2Size.Y / 2
	for i := 0; i < img2Size.Y; i++ {
		for j := 0; j < img2Size.X; j++ {
			xx := int(float64(j-x)*math.Cos(-r)-float64(i-y)*math.Sin(-r))+x
			yy := int(float64(j-x)*math.Sin(-r)-float64(i-y)*math.Cos(-r))+y
			c := (*img).At(xx, yy)
			r, g, b, _ := c.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				img2.Set(j, i, c)
			}
		}
	}
	return img2
}
