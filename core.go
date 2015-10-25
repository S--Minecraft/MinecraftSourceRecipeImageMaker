package main

import (
	"./img/load"
	"./img/edit"
	"./img/output"
)

func main() {
	layerImg := load.Load("1.png")
	coverImg := load.Load("2.png")

	edit.Paste(&layerImg, 50, 0, &coverImg)

	output.Output("3.gif", &layerImg)
}

