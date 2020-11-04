package util

import "encoding/json"

func ParseToObject(data []byte, structVariablePointer interface{}) error {
	errorjson := json.Unmarshal(data, structVariablePointer)

	if errorjson != nil {
		return errorjson
	}

	return nil
}
