package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJsonFromFile(path string, obj interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &obj); err != nil {
		return err
	}

	return nil
}

func SaveAsJsonFile(path string, obj interface{}) error {
	saveData, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, saveData, os.ModeAppend|os.FileMode(0666)); err != nil {
		return err
	}

	return nil
}
