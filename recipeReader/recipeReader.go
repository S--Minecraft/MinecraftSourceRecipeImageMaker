package recipeReader

import (
	"encoding/json"
	"io/ioutil"
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
