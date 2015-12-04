package edit

import(
	"image"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"strconv"
	"../load"
)

type Number struct {
	Numbers [][]int `json:"numbers"`
}

var numberJson Number;
var number [][]int
var numberImg image.Image;

func loadJson() {
	file, err := ioutil.ReadFile("assets/number/number.json")
	if err != nil {
		fmt.Println("Number json read error: ", err)
	}
	json.Unmarshal(file, &numberJson)
	number = numberJson.Numbers
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
			width += number[i][2] - number[i][0] + 1
			if number[i][3] - number[i][1] > height {
				height = number[i][3] - number[i][1]
			}
		}
	}

	offsetX := 0
	for _, n := range strconv.Itoa(num) {
		if i, err := strconv.Atoi(string(n)); err != nil {
			fmt.Println("Number String Convert Error: ", err)
		} else {
			OverrideMultiArr(&numberImg, number[i], back, []int{xx - width + 1 + offsetX, yy - height + 1})

			offsetX += number[i][3] - number[i][1] + 1
		}
	}
	return
}

func PasteNumberOffset(back *image.Image, num int, xx int, yy int, offsetX int, offsetY int) {
	PasteNumber(back, num, xx-offsetX, yy-offsetY)
	return
}

func PasteNumberArrOffset(back *image.Image, num int, place []int, offset []uint) {
	PasteNumber(back, num, place[0]+16-int(offset[0]), place[1]+16-int(offset[1]))
	return
}
