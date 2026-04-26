package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Interface2String(value interface{}) string {
	str, ok := value.(string)
	if ok {
		return str
	}
	return JsonEncode(value)
}

func JsonDecode(src string, dest interface{}) error {
	return json.Unmarshal([]byte(src), dest)
}

func JsonEncode(value interface{}) string {
	marshal, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(marshal)
}
func Interface2Interface(value interface{}, dest interface{}) error {
	v := Interface2String(value)
	err := JsonDecode(v, &dest)
	if err != nil {
		return err
	}
	return nil
}

type test struct {
	Name string `json:"name"`
}

func BeautifyJsonStr(obj map[string]interface{}) string {
	rawJSON := Interface2String(&obj)
	// 美化
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(rawJSON), "", "") // 前缀为空，缩进为两个空格
	if err != nil {
		fmt.Println("BeautifyJsonStr error:", err)
	}
	return prettyJSON.String()
}

func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}

func main() {
	a := Interface2String(&test{
		Name: "tttttttttttt",
	})
	fmt.Printf("1---%T: %v\n", a, a)
	var c = test{}
	b := Interface2Interface(map[string]interface{}{
		"name": "tttttttttttt",
	}, &c)
	fmt.Printf("2---%T: %v\n", b, b)
	fmt.Printf("3---%T: %v\n", c, c)
	res := BeautifyJsonStr(map[string]interface{}{
		"name": "tttttttttttt",
	})
	fmt.Printf("4---%T: %v", res, res)
}
