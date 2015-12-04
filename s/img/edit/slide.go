package edit

import (
	"image"
)

// imgの片方の辺をずらして変形する
// (option=trueで縦/falseで横に伸ばす/option2=trueで左上固定/falseで右下固定)
// ＿　　　 ＿
// | |　→ / /
// ￣　　　￣
//
func Slide(img *image.Image, length int, option bool, option2 bool) image.Image {
	size := (*img).Bounds().Size()
	//横を伸ばす場合は横の長さがx+x/lenになる
	img2 := image.NewNRGBA(image.Rect(0, 0, size.X+int(float64(size.X)/float64(length)), size.Y))
	if option {
		//縦を伸ばす場合は縦の長さがx+x/lenになる
		img2 = image.NewNRGBA(image.Rect(0, 0, size.X, size.Y+int(float64(size.Y)/float64(length))))
	}
	size2 := (*img2).Bounds().Size()
	//統一で扱うために長い方を取得
	sizeB := size2.Y
	if size2.X > size2.Y {
		sizeB = size2.X
	}

	offset := 0
	for x, x2 := 0, 0; x < sizeB; x++ {
		//逆スタートの場合は反転
		x2 = x
		if option2 {
			if option {
				x2 = size2.X - x
			} else {
				x2 = size2.Y - x
			}
		}

		for y, y2 := 0, 0; y < sizeB; y++ {
			//逆スタートの場合は反転
			y2 = y
			if option2 {
				if option {
					y2 = size2.Y - y
				} else {
					y2 = size2.X - y
				}
			}

			c := (*img).At(x2, y2)
			r, g, b, _ := c.RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				//縦横の入れ替えの場合
				if option {
					img2.Set(x2, y2 + offset, c)
				} else {
					img2.Set(y2 + offset, x2, c)
				}
			}
		}

		//ずらす分
		if (x2 % length == 0) {
			offset++
		}
	}
	return img2
}
