package utility

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func JsonStrToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MapToJsonStr(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

func StructToMapReflect(obj interface{}) map[string]any {
	val := reflect.ValueOf(obj)
	typ := val.Type()
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	m := make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		m[field.Name] = val.Field(i).Interface()
	}
	return m
}
