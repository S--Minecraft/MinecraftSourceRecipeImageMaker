# MinecraftSourceRecipeImageMaker
jsonからレシピの画像を生成します  

##Usage
1. `cfgフォルダ`にマイクラのjar内にある`crafting_table.png`を入れる
1. `assetsフォルダ`を作り、マイクラのjar内にあるアイテムの画像を階層なしで入れる
1. `assets/blockフォルダ`にマイクラのjar内にあるブロックの画像を階層なしで入れる
1. `assetsフォルダ`に`**.json`というファイル(**はなんでもいいです)を作る
1. その中に下のように書く
1. `outputフォルダ`を作る
1. `MinecraftSourceRecipeImageMaker.exe`(amd64とついているほうが64bit用、386とついているほうが32bit用です)をたたく
1. `outputフォルダ`に画像が生成される
```
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
						"F": "block/dirt",
						"G": "beef_cooked"
					},
					"number": 2
				}
			]
		}
    ]
}
```
※craftingTable, furnace, brewingStandが標準で使えます  
　また、cfgフォルダにcraftingTable.jsonやfurnace.jsonを参考に  
　設定を書けばMODにある機械のレシピなどにも対応できます  
　typeの文字列はjsonの名前です  

 - craftingTableの場合  
shapeの1つめから9つめは各入力スロット(空文字にすれば何もいれないことになります)  
10こめは出力スロットです  
 - furnaceの場合  
shapeの1つ目が材料、2つ目が燃料、3つ目が出力です  
 - brewingStandの場合  
shapeの1つ目が材料、2つ目が出力です  
  
imgはshapeでそれぞれ指定した文字がどの画像かを示します  
例えば、Xはapple.pngが作画されます  
"X": "block/dirt"などと指定するとブロック用の変形した画像になります
  
numberは出力個数です、1つの場合は省略可能です  
