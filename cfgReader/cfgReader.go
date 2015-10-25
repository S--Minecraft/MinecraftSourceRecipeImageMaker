package cfgReader

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Config struct {
	CraftingTable CraftingTable `json:"craftingTable"`
	Furnace Furnace `json:"furnace"`
}
type CraftingTable struct {
	Gui string `json:"gui"`
	Size []uint `json:"size"`
	Place [][]int `json:"place"`
}
type Furnace struct {
	Gui string `json:"gui"`
	Size []uint `json:"size"`
	Place [][]int `json:"place"`
	Override []Override `json:"override"`
}
type Override struct {
	Before []uint `json:"before"`
	After []uint `json:"after"`
}

func Read() (config Config) {
	file, err := ioutil.ReadFile("cfg/config.json")
	if err != nil {
		fmt.Println("Config read error: ", err)
	}
	json.Unmarshal(file, &config)
	return
}
