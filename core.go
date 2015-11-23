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
			if recipeType.Type == "furnace" {
				furnace(&recipeType.Recipes, &config)
			}
		}
	}
}

func craftingTable(recipes *[]recipeReader.Recipe, config *cfgReader.Config) {
	cfg := config.CraftingTable

	layerImg := load.Load("cfg/" + cfg.Gui)
	place := cfg.Place

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

		lImg = edit.TrimArr(&lImg, cfg.Trim)
		output.Output("output/" + recipe.Img[recipe.Shape[len(shape) - 1]] + ".png", &lImg)
	}
}

func furnace(recipes *[]recipeReader.Recipe, config *cfgReader.Config) {
	cfg := config.Furnace

	layerImg := load.Load("cfg/" + cfg.Gui)
	place := cfg.Place

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

		lImg = edit.TrimArr(&lImg, cfg.Trim)
		shape := recipe.Shape
		output.Output("output/" + recipe.Img[shape[len(shape) - 1]] + ".png", &lImg)
	}
}
