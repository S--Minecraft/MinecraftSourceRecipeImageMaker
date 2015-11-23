package main

import (
	"./img/load"
	"./img/edit"
	"./img/output"
	"./cfgReader"
	"./recipeReader"
)

func main() {
	var configs map[string]cfgReader.Config = make(map[string]cfgReader.Config)

	allRecipe := recipeReader.ReadAll("assets")
	for _, recipe := range allRecipe {
		for _, recipeType := range recipe.RecipeType {
			config, ok := configs[recipeType.Type]
			if !ok {
				config = cfgReader.Read(recipeType.Type)
				configs[recipeType.Type] = config
			}
			readRecipe(&recipeType.Recipes, &config)
		}
	}
}

func readRecipe(recipes *[]recipeReader.Recipe, cfg *cfgReader.Config) {
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
