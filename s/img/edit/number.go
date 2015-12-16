package edit

import (
	"../load"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
)

type Number struct {
	Numbers [][]int `json:"numbers"`
}

var numberJson Number
var number [][]image.Point
var numberImg image.Image

func loadJson() {
	file, err := ioutil.ReadFile("assets/number/number.json")
	if err != nil {
		fmt.Println("Number json read error: ", err)
	}
	json.Unmarshal(file, &numberJson)
	for _, n := range numberJson.Numbers {
		number = append(number, []image.Point{image.Pt(n[0], n[1]), image.Pt(n[2], n[3])})
	}
	return
}
func loadImg() {
	numberImg = load.Load("assets/number/ascii.png")
}

func InitNumber() {
	loadImg()
	loadJson()
	return
}

// 数字の右下の座標(xx,yy)になるように数字を入れる
func PasteNumber(back *image.Image, num int, xx int, yy int) {
	height := 0
	width := 0

	for _, n := range strconv.Itoa(num) {
		if i, err := strconv.Atoi(string(n)); err != nil {
			fmt.Println("Number String Convert Error: ", err)
		} else {
			width += number[i][1].X - number[i][0].X + 1
			if number[i][1].Y-number[i][0].Y > height {
				height = number[i][1].Y - number[i][0].Y
			}
		}
	}

	offsetX := 0
	for _, n := range strconv.Itoa(num) {
		if i, err := strconv.Atoi(string(n)); err != nil {
			fmt.Println("Number String Convert Error: ", err)
		} else {
			OverrideMultiArr(&numberImg, number[i], back, image.Pt(xx-width+1+offsetX, yy-height+1))

			offsetX += number[i][1].Y - number[i][0].Y + 1
		}
	}
	return
}

func PasteNumberOffset(back *image.Image, num int, xx int, yy int, offsetX int, offsetY int) {
	PasteNumber(back, num, xx-offsetX, yy-offsetY)
	return
}

func PasteNumberArrOffset(back *image.Image, num int, place image.Point, offset image.Point) {
	PasteNumber(back, num, place.X+16-offset.X, place.Y+16-offset.Y)
	return
}
