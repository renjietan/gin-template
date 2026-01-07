package utility

import (
	"fmt"
	"reflect"
)

func Tern[T any, K any](boolVal bool, a T, b K) any {
	var result any
	if boolVal {
		result = a
	} else {
		result = b
	}

	val := reflect.ValueOf(result)
	fmt.Println("tern kind============", val.Kind())
	if val.Kind() == reflect.Func {
		funcType := val.Type()
		numIn := funcType.NumIn()
		fmt.Println("tern funcType============", funcType)
		fmt.Println("tern numIn============", numIn)
		if numIn == 0 {
			results := val.Call(nil)
			if len(results) > 0 {
				return results[0].Interface()
			}
			// 无返回值的函数（如 PutUint16），返回 nil
			return nil
		}
		return result
	}

	return result
}
