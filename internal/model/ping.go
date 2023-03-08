package model

import (
	jsoniter "github.com/json-iterator/go"
)

type PingModel struct {
	Message string `json:"message"`
}

func GetPingResult() *PingModel {
	var resultModel PingModel
	var resultData = "{\"message\":\"ok\"}"

	err := jsoniter.Unmarshal([]byte(resultData), &resultModel)
	if err != nil {
		return nil
	}
	return &resultModel
}
