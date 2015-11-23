package cfgReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Gui      string     `json:"gui"`
	Trim     []uint     `json:"trim"`
	Place    [][]int    `json:"place"`
	Override []Override `json:"override"`
}
type Override struct {
	Before []uint `json:"before"`
	After  []uint `json:"after"`
}

func Read(crafter string) (config Config) {
	file, err := ioutil.ReadFile("cfg/" + crafter + ".json")
	if err != nil {
		fmt.Println("Config read error: ", err)
	}
	json.Unmarshal(file, &config)
	return
}
