# MinecraftSourceRecipeImageMaker
jsonからレシピの画像を生成します

##Usage
`cfgフォルダ`にマイクラのjar内にある`crafting_table.png`を入れ、
`assetsフォルダ`にマイクラのjar内にあるアイテムの画像を階層なしで入れ、
さらに、`assetsフォルダ`に`**.json`というファイル(**はなんでもいいです)を作り、
その中に下のように書けば`MinecraftSourceRecipeImageMaker.exe`をたたけば、
`outputフォルダ`に画像が生成されるはずです
```
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
```
shapeの1つめから9つめは各入力スロット(空文字にすれば何もいれないことになります)
10こめは出力スロットです
imgはshapeでそれぞれ指定した文字がどの画像かを示します
ex. Xはapple.pngが作画されます
