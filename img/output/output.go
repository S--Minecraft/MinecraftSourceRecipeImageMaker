package output

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

func Output(filePath string, img *image.Image) {
	var file *os.File
	if file, err = os.Create(filePath); err != nil {
		fmt.Println("Output create error: ", err)
		return
	}
	encode(&filePath, img, file)
	defer file.Close()
}
