package utils

import (
	"encoding/json"
	"errors"
)

// ConvertBytesToMap will convert the provide byte array into a [string]string map/dictionary of key/value pairs using the json provider
func ConvertBytesToMap(byteBuffer []byte) (map[string]interface{}, error) {

	if byteBuffer == nil {
		return nil, errors.New("byte[] can not be nil")
	}

	objMap := make(map[string]interface{})
	err := json.Unmarshal(byteBuffer, &objMap)

	if err == nil {
		return objMap, nil
	}

	return nil, err
}
