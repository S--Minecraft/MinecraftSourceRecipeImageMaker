package edit

import (
	"image"
)

// imgを32x32のブロックの画像に変換する
func Cube(img *image.Image) image.Image {
	layer := image.Image(image.NewNRGBA(image.Rect(0, 0, 32, 32)))

	top := Rotate(img, 45)
	top = Resize(&top, 28, 15)
	left := Slide(img, 2, true, false)
	left = Resize(&left, 14, 24)
	right := Slide(img, 2, true, true)
	right = Resize(&right, 14, 24)

	Paste(&layer, 2, 0, &top)
	Paste(&layer, 2, 8, &left)
	Paste(&layer, 16, 6, &right)

	return layer
}
