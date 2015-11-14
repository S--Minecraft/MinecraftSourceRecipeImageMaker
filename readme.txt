////////////////////////////////////////////////////////////////
/MinecraftSourceRecipeImageMaker///////////////////////by S/////
////////////////////////////////////////////////////////////////

《概要》
レシピを書いたjsonから画像を出力するツールです


《使用方法》
cfgフォルダにマイクラのjar内にあるcrafting_table.pngを入れ、
assetsフォルダにマイクラのjar内にあるアイテムの画像を階層なしで入れ、
さらに、assetsフォルダに**.jsonというファイル(**はなんでもいいです)を作り、
その中に下のように書けばMinecraftSourceRecipeImageMaker.exeをたたけば、
outputフォルダに画像が生成されるはずです
{
	"recipes": [
		{
			"type": "craftingTable",
			"recipe": [
				{
					"shape": [
						"X", "Y", "Z",
						"A", "B", "C",
						"D", "E", "F",
						"G"
					],
					"img": {
						"X": "apple",
						"Y": "apple_golden",
						"Z": "arrow",
						"A": "bed",
						"B": "beef_cooked",
						"C": "beef_raw",
						"D": "blaze_powder",
						"E": "blaze_rod",
						"F": "boat",
						"G": "beef_cooked"
					}
				}
			]
		}
}


《注意》
まだアルファ版なのであまりうまく動作しません


《報告・連絡》
GitHubかTwitterまでお願いします


《リンク》
--GitHub
https://github.com/S--Minecraft/MinecraftSourceRecipeImageMaker

--Twitter
https://twitter.com/S__Minecraft
