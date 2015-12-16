package edit

import (
	"../../cfgReader"
	"image"
)

// backの(xx, yy)にfrontを貼り付け
func Paste(back *image.Image, xx int, yy int, front *image.Image) {
	size := (*front).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			color := (*front).At(x, y)
			r, g, b, _ := color.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				(*back).(*image.NRGBA).Set(x+xx, y+yy, color)
			}
		}
	}
	return
}

// backの(xx+offsetX, yy+offsetY)にfrontを貼り付け
func PasteOffset(back *image.Image, xx int, yy int, offsetX int, offsetY int, front *image.Image) {
	Paste(back, xx-offsetX, yy-offsetY, front)
	return
}

// backの[x, y]にfrontを貼り付け
func PasteArr(back *image.Image, arr image.Point, front *image.Image) {
	Paste(back, arr.X, arr.Y, front)
	return
}

// backの[x, y]にオフセットをつけてfrontを貼り付け
func PasteArrOffset(back *image.Image, arr image.Point, offset image.Point, front *image.Image) {
	PasteOffset(back, arr.X, arr.Y, offset.X, offset.Y, front)
	return
}

// frontの(x1,y1),(x2,y2)の長方形をbackの(toX,toY)に貼り付け
func OverrideMulti(front *image.Image, x1 int, y1 int, x2 int, y2 int, back *image.Image, toX int, toY int) {
	coverSizeX := x2 - x1
	coverSizeY := y2 - y1

	for x := 0; x <= coverSizeX; x++ {
		for y := 0; y <= coverSizeY; y++ {
			color := (*front).At(x1+x, y1+y)
			r, g, b, _ := color.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				(*back).(*image.NRGBA).Set(toX+x, toY+y, color)
			}
		}
	}
	return
}

func OverrideMultiArr(front *image.Image, from []image.Point, back *image.Image, to image.Point) {
	OverrideMulti(front, from[0].X, from[0].Y, from[1].X, from[1].Y, back, to.X, to.Y)
	return
}

// imgの(x1,y1),(x2,y2)の長方形を(toX,toY)に貼り付け
func Override(img *image.Image, x1 int, y1 int, x2 int, y2 int, toX int, toY int) {
	OverrideMulti(img, x1, y1, x2, y2, img, toX, toY)
	return
}

// imgの(b[0], b[1]),(b[2],b[3])の長方形を(a[0],a[1])に貼り付け
func OverrideCfg(img *image.Image, override *cfgReader.Override) {
	b := override.Before
	a := override.After
	Override(img, int(b[0]), int(b[1]), int(b[2]), int(b[3]), int(a[0]), int(a[1]))
	return
}
