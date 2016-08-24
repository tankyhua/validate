package validate

import (
	"fmt"
	"reflect"
)

var validate = struct {
	precompile bool //是否已经预处理
}{}

//  StructValidator 结构体参数验证器
type StructValidator struct {
}

func (this *StructValidator) Validate(model interface{}) (string, bool) {
	var value = reflect.ValueOf(model).Elem()
	var ps parse
	var name, ok = ps.parseStruct(value)
	fmt.Println(modelMap)
	return name, ok
}
