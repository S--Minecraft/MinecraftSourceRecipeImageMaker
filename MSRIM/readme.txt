////////////////////////////////////////////////////////////////
/MinecraftSourceRecipeImageMaker///////////////////////by S/////
////////////////////////////////////////////////////////////////

《概要》
レシピを書いたjsonから画像を出力するツールです


《使用方法》
1. cfgフォルダにマイクラのjar内にあるcrafting_table.pngを入れる
2. assetsフォルダを作り、マイクラのjar内にあるアイテムの画像を階層なしで入れる
3. assetsフォルダに**.jsonというファイル(**はなんでもいいです)を作る
4. その中に下のように書く
5. outputフォルダを作る
6. MinecraftSourceRecipeImageMaker.exeをたたく
7. outputフォルダに画像が生成される
{
	"recipes": [
		{
			"type": "craftingTableかfurnaceなど※",
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

※craftingTable, furnace, brewingStandが標準で使えます
　また、cfgフォルダにcraftingTable.jsonやfurnace.jsonを参考に
　設定を書けばMODにある機械のレシピなどにも対応できます
　typeの文字列はjsonの名前です


《注意》
まだアルファ版なのであまりうまく動作しません


《報告・連絡》
GitHubかTwitterまでお願いします


《リンク》
--GitHub
https://github.com/S--Minecraft/MinecraftSourceRecipeImageMaker

--Twitter
https://twitter.com/S__Minecraft
