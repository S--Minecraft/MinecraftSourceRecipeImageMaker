package main

import (
	"./cfgReader"
	"./img/edit"
	"./img/load"
	"./img/output"
	"./recipeReader"
	"image"
	"sync"
	"regexp"
)

var waitGroup sync.WaitGroup //平行処理用

func main() {
	//ロードされたコンフィグと背景画像
	var configs map[string]cfgReader.Config = make(map[string]cfgReader.Config)
	var layerImgs map[string]image.Image = make(map[string]image.Image)
	edit.InitNumber()

	//すべてのファイルのレシピ一覧
	allRecipe := recipeReader.ReadAll("assets")

	for _, recipe := range allRecipe {
		for _, recipeType := range recipe.RecipeType {
			readTypes(&configs, &layerImgs, &recipeType)
		}
	}

	waitGroup.Wait() //完了まで待つ
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
	for _, recipe := range *recipes {
		waitGroup.Add(1) //処理をカウント
		go makeImg(recipe, cfg, *layer)
	}
	return
}

func makeImg(recipe recipeReader.Recipe, cfg *cfgReader.Config, layer image.Image) {
	layerImg := edit.Copy(layer)
	place := cfg.Place

	for i, item := range recipe.Shape {
		if item != "" {
			//アイテムのロード
			imgPath := "assets/" + recipe.Img[item] + ".png"
			img := load.Load(imgPath)
			if isBlock(imgPath) {
				img = edit.Cube(&img)
				img = edit.Resize(&img, 16, 16)
			}
			//アイテムを貼り付け
			edit.PasteArrOffset(&layerImg, place[i], cfg.Trim, &img)
		}
	}
	//出力個数を出力
	shape := recipe.Shape
	if recipe.Number != 0 && recipe.Number != 1 {
		edit.PasteNumberArrOffset(&layerImg, recipe.Number, place[len(shape)-1], cfg.Trim)
	}

	//作成されるものの名前.pngで出力
	output.Output("output/"+recipe.Img[shape[len(shape)-1]]+".png", &layerImg)

	defer waitGroup.Done() //完了をwaitGroupに知らせる
	return
}

func isBlock(path string) bool {
	reg := regexp.MustCompile(`block/.*`)
	match := reg.FindString(path)
	return (match != "")
}
