package edit

import (
	"image"
	"../../cfgReader"
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
	return
}

func PasteOffset(back *image.Image, xx int, yy int, offsetX int, offsetY int, front *image.Image) {
	Paste(back, xx - offsetX, yy - offsetY, front)
	return
}

func PasteArr(back *image.Image, arr []int, front *image.Image) {
	Paste(back, arr[0], arr[1], front)
	return
}

func PasteArrOffset(back *image.Image, arr []int, offset []uint, front *image.Image) {
	PasteOffset(back, arr[0], arr[1], int(offset[0]), int(offset[1]), front)
	return
}

func Override(img *image.Image, x1 int, y1 int, x2 int, y2 int, toX int, toY int) {
	coverSizeX := x2 - x1
	coverSizeY := y2 - y1

	for x := 0; x <= coverSizeX; x++ {
		for y := 0; y <= coverSizeY; y++ {
			color := (*img).At(x1 + x, y1 + y)
			r, g, b, _ :=color.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				(*img).(*image.NRGBA).Set(toX + x, toY + y, color)
			}
		}
	}
	return
}

func OverrideCfg(img *image.Image, override *cfgReader.Override) {
	b := override.Before
	a := override.After
	Override(img, int(b[0]), int(b[1]), int(b[2]), int(b[3]), int(a[0]), int(a[1]))
	return
}
