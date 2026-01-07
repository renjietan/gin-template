package utility

import (
	"encoding/json"
	"fmt"
)

func JsonStrToMap(jsonStr string) (map[string]any, error) {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

func MapToJsonStr(m map[string]any) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

func DataToBuffer(data []string) {

}
