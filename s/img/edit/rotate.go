package edit

import (
	"image"
	"math"
)

// imgをangle°だけ回転させた画像を返す
func Rotate(img *image.Image, angle int) image.Image {
	r := math.Pi * 2 * float64(angle) / 360//ラジアン
	size := (*img).Bounds().Size()
	dw := int(math.Abs(float64(size.X)*math.Cos(r))+math.Abs(float64(size.Y)*math.Sin(r)))
	dh := int(math.Abs(float64(size.X)*math.Sin(r))+math.Abs(float64(size.Y)*math.Cos(r)))
	img2 := image.NewNRGBA(image.Rect(0, 0, dw, dh))
	size2 := img2.Bounds().Size()
	// (a, b)を中心としてangle(度数法)だけ回転
	a := dw/2
	b := dh/2
	for i := 0/*-size2.X*/; i < size2.X; i++ {
		for j := 0/*-size2.Y*/; j < size2.Y; j++ {
			x := int(float64(i-a)*math.Cos(r) - float64(j-b)*math.Sin(r))+a
			y := int(float64(i-a)*math.Sin(r) + float64(j-b)*math.Cos(r))+b
			c := (*img).At(x, y)
			//r, g, b, _ := c.RGBA()
			//if !(r == 0 && g == 0 && b == 0) {
			img2.Set(i, j, c)
			//}
		}
	}
	return img2
}
