package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gitee.com/cpds/cpds-analyzer/pkgs/rules"
)

func LoadRules(r *rules.Rules, path string) error {
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

func SaveRules(r *rules.Rules, path string) error {
	saveData, err := json.Marshal(r)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, saveData, os.ModeAppend); err != nil {
		return err
	}

	return nil
}
