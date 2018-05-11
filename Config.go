package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AgsConfig struct {
	AgsURL      string `json:"ags_url"`
	AgsName     string `json:"ags_name"`
	AgsPassword string `json:"ags_password"`
}

func readConfig() (AgsConfig, error) {
	v := AgsConfig{}
	jsonfile, err := ioutil.ReadFile("Config.json")
	if err != nil {
		fmt.Println(err)
		return v, err
	}
	json.Unmarshal(jsonfile, &v)
	return v, nil

}
