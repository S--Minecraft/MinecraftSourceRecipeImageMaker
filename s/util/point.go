package util

import (
	"fmt"
	"image"
)

func IToPoint(i []int) (p image.Point) {
	if len(i) != 2 {
		fmt.Println("failed to convert []int to image.Point")
	} else {
		p = image.Pt(i[0], i[1])
	}
	return
}

func IToPoints(i []int) []image.Point {
	p := make([]image.Point, len(i))
	if !(2 < len(i) && len(i)%2 == 0) {
		fmt.Println("failed to convert []int to []image.Point")
	} else {
		for j := 0; j < len(i); j = +2 {
			p[j] = image.Pt(i[j], i[j+1])
		}
	}
	return p
}

func IArrToPoints(i [][]int) []image.Point {
	p := make([]image.Point, len(i))
	for j := 0; j < len(i); j++ {
		if len(i[j]) != 2 {
			fmt.Println("failed to convert []int to [][]image.Point")
		} else {
			p[j] = image.Pt(i[j][0], i[j][1])
		}
	}
	return p
}

func UiToPoint(i []uint) (p image.Point) {
	if len(i) != 2 {
		fmt.Println("failed to convert []int to image.Point")
	} else {
		p = image.Pt(int(i[0]), int(i[1]))
	}
	return
}

func UiToPoints(i []uint) []image.Point {
	p := make([]image.Point, len(i))
	if !(2 < len(i) && len(i)%2 == 0) {
		fmt.Println("failed to convert []uint to []image.Point")
	} else {
		k := 0
		for j := 0; j < len(i); j++ {
			p[j] = image.Pt(int(i[k]), int(i[k+1]))
			k = +2
		}
	}
	return p
}

func UiArrToPoints(i [][]uint) []image.Point {
	p := make([]image.Point, len(i))
	for j := 0; j < len(i); j++ {
		if len(i[j]) != 2 {
			fmt.Println("failed to convert []uint to [][]image.Point")
		} else {
			p[j] = image.Pt(int(i[j][0]), int(i[j][1]))
		}
	}
	return p
}
