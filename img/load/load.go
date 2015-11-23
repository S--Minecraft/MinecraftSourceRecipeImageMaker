package load

import (
	"fmt"
	"image"
	"image/png"
	"image/jpeg"
	"image/gif"
	"os"
	"path"
)

var err error

//拡張子を判別してデコード
func decode(filePath *string, file *os.File) (img image.Image) {
	ext := path.Ext(*filePath)

	if ext == ".jpg" {
		if img, err = jpeg.Decode(file); err != nil {
			fmt.Println("Jpeg Image decode error: ", err)
			return
		}
	} else if ext == ".png" {
		if img, err = png.Decode(file); err != nil {
			fmt.Println("Png Image decode error: ", err)
			return
		}
	} else if ext == ".gif" {
		if img, err = gif.Decode(file); err != nil {
			fmt.Println("Gif Image decode error: ", err)
			return
		}
	} else {
		fmt.Println("Not compatible image type error")
		return
	}
	return
}

func Load(filePath string) image.Image {
	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		fmt.Println("File load error: ", err)
	}
	defer file.Close()

	return decode(&filePath, file)
}
