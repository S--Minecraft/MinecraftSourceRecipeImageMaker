package output

import (
	"fmt"
	"image"
	"image/png"
	"image/jpeg"
	"image/gif"
	"regexp"
	"strconv"
	"os"
	"path"
)

var err error

func encode(filePath *string, img *image.Image, file *os.File) {
	ext := path.Ext(*filePath)

	if ext == ".jpg" {
		if err = jpeg.Encode(file, *img, nil); err != nil {
			fmt.Println("Jpeg Image encode error: ", err)
		}
	} else if ext == ".png" {
		if err = png.Encode(file, *img); err != nil {
			fmt.Println("Png Image encode error: ", err)
		}
	} else if ext == ".gif" {
		if err = gif.Encode(file, *img, nil); err != nil {
			fmt.Println("Gif Image encode error: ", err)
		}
	} else {
		fmt.Println("Not compatible image type error")
	}
}

func isExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func Output(filePath string, img *image.Image) {
	var file *os.File
	if isExist(filePath) {
		reg := regexp.MustCompile(`(.*)\((\d)\)\.(.*?)$`)
		match := reg.FindAllStringSubmatch(filePath, -1)
		if len(match) > 0 {
			if n, errr := strconv.Atoi(match[0][2]); errr != nil {
				fmt.Println("File name number string convert error: ", err)
			} else {
				nStr := strconv.Itoa(n + 1)
				Output(match[0][1] + "(" + nStr + ")." + match[0][3], img)
			}
		} else {
			reg = regexp.MustCompile(`(.*)\.(.*?)$`)
			match = reg.FindAllStringSubmatch(filePath, -1)
			Output(match[0][1] + "(1)." + match[0][2], img)
		}
	}
	if file, err = os.Create(filePath); err != nil {
		fmt.Println("Output create error: ", err)
		return
	}
	encode(&filePath, img, file)
	defer file.Close()
}