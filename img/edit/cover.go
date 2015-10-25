package edit

import (
	"image"
)

// backの(a, b)にfrontを貼り付け
func Paste(back *image.Image, a int, b int, front *image.Image) {
	size := (*front).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			(*back).(*image.RGBA).Set(x+a, y+b, (*front).At(x,y))
		}
	}
}
