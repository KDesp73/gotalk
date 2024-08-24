package json

import (
	"encoding/json"
)

type Json struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func (j Json) ToString() string {
	jsonData, err := json.Marshal(j)
	if err != nil {
		return "{status: 500, message: \"Internal Server Error\"}"
	}
	return string(jsonData)
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
