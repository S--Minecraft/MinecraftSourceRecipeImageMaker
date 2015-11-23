package edit

import (
	"image"
)

//(x1,y1),(x2,y2)を切り出し
func Trim(img *image.Image, x1 uint, y1 uint, x2 uint, y2 uint) image.Image {
	img2 := image.NewNRGBA(image.Rect(0, 0, int(x2-x1), int(y2-y1)))
	size := (*img2).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			(*img2).Set(x, y, (*img).At(int(x1)+x, int(y1)+y))
		}
	}
	return img2
}

func TrimArr(img *image.Image, arr []uint) image.Image {
	return Trim(img, arr[0], arr[1], arr[2], arr[3])
}
