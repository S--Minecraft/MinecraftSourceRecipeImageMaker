package edit

import (
	"image"
)

// backの(a, b)にfrontを貼り付け
func Paste(back *image.Image, xx int, yy int, front *image.Image) {
	size := (*front).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := (*front).At(x,y)
			r, g, b, _ :=color.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				(*back).(*image.NRGBA).Set(x+xx, y+yy, color)
			}
		}
	}
}

func PasteArr(back *image.Image, arr []int, front *image.Image) {
	Paste(back, arr[0], arr[1], front)
}
