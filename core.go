package main

import (
	"image"
	"./img/load"
	"./img/edit"
	"./img/output"
	"./cfgReader"
	"./recipeReader"
)

func main() {
	var configs map[string]cfgReader.Config = make(map[string]cfgReader.Config)
	var layerImgs map[string]image.Image = make(map[string]image.Image)

	allRecipe := recipeReader.ReadAll("assets")
	for _, recipe := range allRecipe {
		for _, recipeType := range recipe.RecipeType {
			// 対応config読み込み(読み込み済みでない場合ファイルロード)
			config, ok := configs[recipeType.Type]
			if !ok {
				config = cfgReader.Read(recipeType.Type)
				configs[recipeType.Type] = config
			}
			// 対応画像読み込み(読み込み済みでない場合ファイルロード)
			layer, ok := layerImgs[recipeType.Type]
			if !ok {
				layer = load.Load("cfg/" + config.Gui)
				layer = edit.TrimArr(&layer, config.Trim)
				layerImgs[recipeType.Type] = layer
			}

			readRecipe(&recipeType.Recipes, &config, layer)
		}
	}
}

func readRecipe(recipes *[]recipeReader.Recipe, cfg *cfgReader.Config, layerImg image.Image) {
	place := cfg.Place

	rs := *recipes
	for _, recipe := range rs {
		for i, item := range recipe.Shape {
			if item != "" {
				imgPath := "assets/" + recipe.Img[item] + ".png"
				img := load.Load(imgPath)
				edit.PasteArrOffset(&layerImg, place[i], cfg.Trim, &img)
			}
		}

		shape := recipe.Shape
		output.Output("output/" + recipe.Img[shape[len(shape) - 1]] + ".png", &layerImg)
	}
}
