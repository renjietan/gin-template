package utility

import (
	"github.com/valyala/fasttemplate"
)

func StringByTemplate(template string, obj interface{}) string {
	temp := fasttemplate.New(template, "{{", "}}")
	json := StructToMapReflect(obj)
	return temp.ExecuteString(json)
}
