package json

import (
	"encoding/json"
	"gotalk/internal/utils"
)

type NestedJson struct {
	Key string `json:"key"`
}

type Json struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data NestedJson `json:"data"`
}

func (j Json) ToString() string {
	json := utils.JsonToString(j)
	if json == "" {
		return "{status: 500, message: \"Internal Server Error\"}"
	}

	return json
}

func (j Json) ToBytes() []byte {
	return []byte(j.ToString())
}

func ParseJson(j string) *Json {
	var newResponse *Json
	err := json.Unmarshal([]byte(j), newResponse)
	if err != nil {
		return nil
	}
	return newResponse
}
