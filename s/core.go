package main

import (
	"fmt"
	"github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/s/cfgReader"
	"github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/s/img/edit"
	"github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/s/img/load"
	"github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/s/img/output"
	"github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/s/recipeReader"
	"image"
	"regexp"
	"sync"
)

var waitGroup sync.WaitGroup //平行処理用

//ロードされたコンフィグと背景画像
var configs map[string]cfgReader.Config = make(map[string]cfgReader.Config)
var layerImgs map[string]image.Image = make(map[string]image.Image)

func main() {
	edit.InitNumber()

	//すべてのファイルのレシピ一覧
	allRecipe := recipeReader.ReadAll("assets")

	for _, recipe := range allRecipe {
		for _, recipeType := range recipe.RecipeType {
			readTypes(&recipeType)
		}
	}

	waitGroup.Wait() //完了まで待つ
	return
}

// クラフト別処理
func readTypes(recipeType *recipeReader.RecipeType) {
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
		if config.Override != nil {
			for _, override := range config.Override {
				edit.OverrideCfg(&layer, &override)
			}
		}
		layer = edit.TrimArr(&layer, config.Trim)
		layerImgs[recipeType.Type] = layer
	}

	for _, recipe := range recipeType.Recipes {
		waitGroup.Add(1) //処理をカウント
		go makeImg(recipe, &config, layer)
	}
	return
}

// レシピのそれぞれの画像出力
func makeImg(recipe recipeReader.Recipe, cfg *cfgReader.Config, layer image.Image) {
	layerImg := edit.Copy(&layer)
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
			edit.PasteArrOffset(&layerImg, place[i], cfg.Trim[0], &img)
		}
	}
	//出力個数を出力
	shape := recipe.Shape
	if recipe.Number != 0 && recipe.Number != 1 {
		edit.PasteNumberArrOffset(&layerImg, recipe.Number, place[len(shape)-1], cfg.Trim[0])
	}

	//作成されるものの名前.pngで出力
	output.Output("output/"+recipe.Img[shape[len(shape)-1]]+".png", &layerImg)

	fmt.Println("Outputed:", recipe.Img[shape[len(shape)-1]])

	defer waitGroup.Done() //完了をwaitGroupに知らせる
	return
}

func isBlock(path string) bool {
	reg := regexp.MustCompile(`block/.*`)
	match := reg.FindString(path)
	return (match != "")
}
