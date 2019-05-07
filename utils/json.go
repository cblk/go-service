package utils

import (
	"bytes"
	"encoding/json"
	"strings"
)

func StructToString(obj interface{}) (string, error) {
	beforeValue, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	midResultMap := make(map[string]interface{}, 0)

	jen := json.NewDecoder(bytes.NewBuffer(beforeValue))
	jen.UseNumber()
	err = jen.Decode(&midResultMap)
	if err != nil {
		return "", err
	}

	eValue, err := JsonMarshal(midResultMap)
	if err != nil {
		return "", err
	}

	return string(eValue), nil
}

func JsonMarshal(value map[string]interface{}) (string, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(value)
	if err != nil {
		return "", err
	}

	result := strings.Trim(buffer.String(), "\n")

	return result, nil
}