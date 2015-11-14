package recipeReader

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"fmt"
)

type RecipeMain struct {
	RecipeType []RecipeType `json:"recipes"`
}

type RecipeType struct {
	Type string `json:"type"`
	Recipes []Recipe `json:"recipe"`
}
type Recipe struct {
	Shape []string `json:"shape"`
	Img map[string]string `json:"img"`
}

func Read(path string) (recipeMain RecipeMain) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Recipe read error: ", err)
	}
	json.Unmarshal(file, &recipeMain)
	return
}

func ReadAll(dir string) (recipeMains []RecipeMain) {
	if files, err := ioutil.ReadDir(dir); err != nil {
		fmt.Println("Assets folder read error:", err)
	} else {
		for _, file := range files {
			if path.Ext(file.Name()) == ".json" {
				recipeMain := Read(dir + "/" + file.Name())
				recipeMains = append(recipeMains, recipeMain)
			}
		}
	}
	return
}
