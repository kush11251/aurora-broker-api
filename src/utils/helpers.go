package utils

import (
	"encoding/json"
	"fmt"
)

func ToJSON(v interface{}) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "{}
"
	}
	return string(jsonBytes)
}
