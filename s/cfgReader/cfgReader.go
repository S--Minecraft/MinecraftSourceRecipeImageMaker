package cfgReader

import (
	"github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/s/util"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
)

type Config2 struct {
	Gui      string     `json:"gui"`
	Trim     []uint     `json:"trim"`
	Place    [][]int    `json:"place"`
	Override []Override `json:"override"`
}
type Override struct {
	Before []uint `json:"before"`
	After  []uint `json:"after"`
}

type Config struct {
	Gui      string
	Trim     []image.Point
	Place    []image.Point
	Override []Override
}

func parse(cfg2 Config2) Config {
	cfg := Config{
		Gui:      cfg2.Gui,
		Trim:     util.UiToPoints(cfg2.Trim),
		Place:    util.IArrToPoints(cfg2.Place),
		Override: cfg2.Override,
	}
	return cfg
}

func Read(crafter string) Config {
	config := Config2{}
	file, err := ioutil.ReadFile("cfg/" + crafter + ".json")
	if err != nil {
		fmt.Println("Config read error: ", err)
	}
	json.Unmarshal(file, &config)
	return parse(config)
}
