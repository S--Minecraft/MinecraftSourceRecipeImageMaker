package main

import (
	"./img/load"
	"./img/edit"
	"./img/output"
	"./cfgReader"
)

func main() {
	config := cfgReader.Read()

	layerImg := load.Load("cfg/" + config.CraftingTable.Gui)
	coverImg := load.Load("cfg/2.png")

	place := config.CraftingTable.Place
	edit.PasteArr(&layerImg, place[0], &coverImg)
	edit.PasteArr(&layerImg, place[2], &coverImg)
	edit.PasteArr(&layerImg, place[4], &coverImg)

	output.Output("cfg/3.png", &layerImg)
}

