package edit

import (
	"image"
)

//(x1,y1),(x2,y2)を切り出し
func Trim(img *image.Image, x1 int, y1 int, x2 int, y2 int) image.Image {
	img2 := image.NewNRGBA(image.Rect(0, 0, x2-x1, y2-y1))
	size := (*img2).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			(*img2).Set(x, y, (*img).At(x1+x, y1+y))
		}
	}
	return img2
}

func TrimArr(img *image.Image, arr []image.Point) image.Image {
	return Trim(img, arr[0].X, arr[0].Y, arr[1].X, arr[1].Y)
}
