package edit

import (
	"image"
	"math"
)

// imgをangle°だけ回転させた画像を返す
// http://homepage2.nifty.com/tsugu/sotuken/rotation/
func Rotate(img *image.Image, angle int) image.Image {
	r := math.Pi * 2 * float64(angle) / 360 //ラジアン
	size := (*img).Bounds().Size()
	dw := int(math.Abs(float64(size.X)*math.Cos(r)) + math.Abs(float64(size.Y)*math.Sin(r)) + 0.5)
	dh := int(math.Abs(float64(size.X)*math.Sin(r)) + math.Abs(float64(size.Y)*math.Cos(r)) + 0.5)
	img2 := image.NewNRGBA(image.Rect(0, 0, dw, dh))
	// (scx, scy)が入力、(dcx, dcy)が出力の中心座標
	scx, scy := size.X/2, size.Y/2
	dcx, dcy := dw/2, dh/2
	sinR, cosR := int(math.Sin(r)*1024), int(math.Cos(r)*1024)
	for x2 := 0; x2 < dh; x2++ {
		for y2 := 0; y2 < dw; y2++ {
			x1 := ((x2-dcx)*cosR-(y2-dcy)*sinR)/1024 + scx
			y1 := ((x2-dcx)*sinR+(y2-dcy)*cosR)/1024 + scy
			if 0 <= x1 && x1 < size.X && 0 <= y1 && y1 < size.Y {
				c := (*img).At(x1, y1)
				r, g, b, _ := c.RGBA()
				if !(r == 0 && g == 0 && b == 0) {
					img2.Set(x2, y2, c)
				}
			}
		}
	}
	return img2
}
