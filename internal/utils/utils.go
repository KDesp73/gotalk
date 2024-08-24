package utils

import (
	"encoding/json"
	"time"
)

func CurrentTimestamp() string {
	currentTime := time.Now()
	return currentTime.Format("2006/01/02 15:04:05")
}

func JsonToString(jsonStruct any) string {
	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
