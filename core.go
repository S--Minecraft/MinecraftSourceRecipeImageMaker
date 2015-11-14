package main

import (
	"./img/load"
	"./img/edit"
	"./img/output"
	"./cfgReader"
	"./recipeReader"
	//"fmt"
)

func main() {
	config := cfgReader.Read()

	recipe := recipeReader.Read("assets/amt.json")

	for _, recipeType := range recipe.RecipeType {
		if recipeType.Type == "craftingTable" {
			craftingTable(&recipeType.Recipes, &config)
		}
	}

	/*
	layerImg := load.Load("cfg/" + config.CraftingTable.Gui)
	coverImg := load.Load("assets/2.png")

	place := config.CraftingTable.Place
	edit.PasteArr(&layerImg, place[0], &coverImg)
	edit.PasteArr(&layerImg, place[2], &coverImg)
	edit.PasteArr(&layerImg, place[4], &coverImg)
	edit.PasteArr(&layerImg, place[9], &coverImg)

	layerImg = edit.TrimArr(&layerImg, config.CraftingTable.Trim)

	output.Output("output/3.png", &layerImg)
	*/
}

func craftingTable(recipes *[]recipeReader.Recipe, config *cfgReader.Config) {
	layerImg := load.Load("cfg/" + config.CraftingTable.Gui)
	place := config.CraftingTable.Place

	rs := *recipes
	for _, recipe := range rs {
		lImg := layerImg

		for i, item := range recipe.Shape {
			if item != "" {
				imgPath := "assets/" + recipe.Img[item] + ".png"
				img := load.Load(imgPath)
				edit.PasteArr(&lImg, place[i], &img)
			}
		}

		lImg = edit.TrimArr(&lImg, config.CraftingTable.Trim)
		output.Output("output/" + recipe.Img[recipe.Shape[9]] + ".png", &lImg)
	}
}
