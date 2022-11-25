package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Rules struct {
	Cpu_max    int `json:"cpu_max"`
	Cpu_min    int `json:"cpu_min"`
	Disk_max   int `json:"disk_max"`
	Disk_min   int `json:"disk_min"`
	Memory_max int `json:"mem_max"`
	Memory_min int `json:"mem_min"`
}

func (r *Rules) loadRules(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, r); err != nil {
		return err
	}

	return nil
}

func (r *Rules) saveRules(path string) error {
	saveData, err := json.Marshal(r)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, saveData, os.ModeAppend); err != nil {
		return err
	}

	return nil
}
