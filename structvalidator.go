package validate

import (
	"errors"
	"reflect"
	"strings"
)

//  StructValidator 结构体参数验证器
type structValidator struct {
}

// Validate 验证结构体
//  model: 结构体指针
// return: 不通过字段名,是否通过
func (this *structValidator) Validate(model interface{}) (string, bool) {
	var value = reflect.ValueOf(model)
	if reflect.TypeOf(model).Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return new(parse).parseStruct(value)
}

// RegisterRegex 注册正则验证
//  name: 类型名 如:验证电话号码对应 name:phone
//  regex: 正则表达式
func (this *structValidator) RegisterRegex(typeName, regex string) error {
	if len(strings.Trim(regex, " ")) <= 0 {
		return errors.New("正则表达式为空!")
	}
	var tp validateFuncType = validateFuncType(strings.ToLower(typeName))
	if _, ok := validateFuncMap[tp]; ok {
		return errors.New(tp.ToString() + "验证类型已存在!")
	}
	setFuncMap(tp, strings.Trim(regex, " "))
	return nil
}
