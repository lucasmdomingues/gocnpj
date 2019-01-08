package utils

import (
	"encoding/json"
)

func ParseJson(values interface{}) ([]byte, error) {

	json, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	return json, err
}
