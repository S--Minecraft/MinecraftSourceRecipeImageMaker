package edit

import (
	"image"
)

// backの(a, b)にfrontを貼り付け
func Paste(back *image.Image, a int, b int, front *image.Image) {
	size := (*front).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			(*back).(*image.NRGBA).Set(x+a, y+b, (*front).At(x,y))
		}
	}
}

func PasteArr(back *image.Image, arr []int, front *image.Image) {
	Paste(back, arr[0], arr[1], front)
}
