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
			readTypes(&configs, &layerImgs, &recipeType)
		}
	}
	return
}

func readTypes(configs *map[string]cfgReader.Config, layerImgs *map[string]image.Image, recipeType *recipeReader.RecipeType) {
	// 対応config読み込み(読み込み済みでない場合ファイルロード)
	cfgs := *configs
	config, ok := cfgs[recipeType.Type]
	if !ok {
		config = cfgReader.Read(recipeType.Type)
		cfgs[recipeType.Type] = config
	}
	// 対応画像読み込み(読み込み済みでない場合ファイルロード)
	lImgs := *layerImgs
	layer, ok := lImgs[recipeType.Type]
	if !ok {
		layer = load.Load("cfg/" + config.Gui)
		if config.Override != nil {
			for _, override := range config.Override {
				edit.OverrideCfg(&layer, &override)
			}
		}
		layer = edit.TrimArr(&layer, config.Trim)
		lImgs[recipeType.Type] = layer
	}
	readRecipe(&recipeType.Recipes, &config, &layer)
	return
}

func readRecipe(recipes *[]recipeReader.Recipe, cfg *cfgReader.Config, layer *image.Image) {
	place := cfg.Place

	rs := *recipes
	for _, recipe := range rs {
		makeImg(&recipe, cfg, *layer, &place)
	}
	return
}

func makeImg(recipe *recipeReader.Recipe, cfg *cfgReader.Config, layerImg image.Image, placeP *[][]int) {
	place := *placeP

	for i, item := range recipe.Shape {
		if item != "" {
			imgPath := "assets/" + recipe.Img[item] + ".png"
			img := load.Load(imgPath)
			edit.PasteArrOffset(&layerImg, place[i], cfg.Trim, &img)
		}
	}
	shape := recipe.Shape
	output.Output("output/" + recipe.Img[shape[len(shape) - 1]] + ".png", &layerImg)
	return
}
