package json

import "encoding/json"

func ToStr(data interface{}) string {
	byteArr, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(byteArr)
}
