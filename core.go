package main

import (
	"./img/load"
	"./img/edit"
	"./img/output"
	"./cfgReader"
	"./recipeReader"
)

func main() {
	config := cfgReader.Read()

	allRecipe := recipeReader.ReadAll("assets")
	for _, recipe := range allRecipe {
		for _, recipeType := range recipe.RecipeType {
			if recipeType.Type == "craftingTable" {
				craftingTable(&recipeType.Recipes, &config)
			}
		}
	}
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
