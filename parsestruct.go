package validate

import (
	"reflect"
	"strings"
)

// parse 解析验证器
type parse struct {
	pValue reflect.Value
	pType  reflect.Type
	tag    reflect.StructTag
	index  int
	com    common
}

var (
	// parseValue 记录需要验证数据
	parseValue = make(map[string]string)
	// parseType 判断类型
	parseType = []string{"val", "gt", "gte", "lt", "lte", "ne", "eq"}
)

// validateType tag验证类型
type validateType string

var (
	validateTypeGt   validateType = "gt"   // >
	validateTypeGte  validateType = "gte"  // >=
	validateTypeLt   validateType = "lt"   // <
	validateTypeLte  validateType = "lte"  // <=
	validateTypeNe   validateType = "ne"   // !=
	validateTypeEq   validateType = "eq"   // =
	validateTypeVal  validateType = "val"  // val
	validateTypeReq  validateType = "req"  // val
	validateTypeFail validateType = "fail" //错误信息
)

func (this validateType) ToString() string {
	return string(this)
}

func (this *parse) parseStruct(value reflect.Value) (string, bool) {
	var kind = value.Kind()
	if kind == reflect.Struct {
		var modelName = value.Type().Name()
		// 判断是否已存在tag记录
		var res, ok = this.com.getMoldeMap(modelName)
		var num = value.NumField()
		for i := 0; i < num; i++ {
			this.index = i
			this.tag = ""
			var fieldName = value.Type().Field(i).Name
			if ok { //是否已存在tag记录
				var tag, isOk = res[fieldName]
				if isOk {
					this.tag = tag
				}
			}
			if len(this.tag) <= 0 {
				//获取tag
				if len(this.getTag(value.Type())) <= 0 {
					if !value.Type().Field(i).Anonymous {
						continue //当tag长度=0且不是匿名字段时,说明无需验证
					}
				}
				//记录全局tag,以便多次使用
				this.com.setModelMap(modelName, fieldName, this.getTag(value.Type()))
			}
			// 属否是切片数组
			if value.Field(i).Kind() == reflect.Slice {
				var length = value.Field(i).Len()
				if length <= 0 { //数组为空
					return this.getFail(value), false
				}
				// 遍历
				for j := 0; j < length; j++ {
					if res, isOk := this.parseStruct(value.Field(i).Index(j)); !isOk {
						return res, isOk
					}
				}

			} else {
				if name, ok := this.parseStructOne(value); !ok {
					return name, ok
				}
			}

		}
	} else if kind == reflect.Slice { //结构体数组
		var lenght = value.Len()
		if lenght == 0 {
			return "数组内容为空!", false
		}
		for j := 0; j < lenght; j++ {
			var res, isOk = this.parseStruct(value.Index(j))
			if !isOk {
				return res, isOk
			}
		}
	} else {
		return "不是结构体类型", false
	}
	return "", true
}

// validateStructOne 验证单个结构体数据
func (this *parse) parseStructOne(value reflect.Value) (string, bool) {
	var ok = true
	var FieldType = value.Type().Field(this.index)
	if FieldType.Anonymous {
		//匿名组合字段,进行递归验证
		return this.parseStruct(value.Field(this.index))
	} else {
		this.parseTag()
		switch FieldType.Type.String() {
		case "string":
			//正则表达式验证字符串值是否满足类型条件
			ok = this.validateString(value.Field(this.index).String())
			break
		case "int", "int32", "int64":
			ok = this.validateInt(int(value.Field(this.index).Int()))
			break
		case "float32", "float64":
			ok = this.validateFloat(value.Field(this.index).Float())
			break
		case "time.Time":

		}

	}
	if ok {
		return "", ok
	}
	return this.getFail(value), ok
}

// validateString 验证字符串
func (this parse) validateString(value string) bool {
	var l = len(strings.Trim(value, " "))
	return this.validateValueInt(l, value)
}

// validateInt 验证整型
func (this parse) validateInt(value int) bool {
	return this.validateValueInt(value, "")
}

// validateFloat 验证浮点型
func (this parse) validateFloat(l float64) bool {
	for k, v := range parseValue {
		switch k {
		case validateTypeEq.ToString(): // =
			if !eqFloat(v, l) {
				return false
			}
		case validateTypeNe.ToString(): // !=
			if !neFloat(v, l) {
				return false
			}
		case validateTypeGt.ToString(): // >
			if !gtFloat(v, l) {
				return false
			}
		case validateTypeGte.ToString(): // >=
			if !gteFloat(v, l) {
				return false
			}

		case validateTypeLt.ToString(): // >=
			if !ltFloat(v, l) {
				return false
			}
		case validateTypeLte.ToString(): //<=
			if !lteFloat(v, l) {
				return false
			}
		}

	}
	return true
}

// validateValueInt 验证int
func (this parse) validateValueInt(l int, value string) bool {
	for k, v := range parseValue {
		switch k {
		case validateTypeVal.ToString(): //判断是否允许为空
			if !this.validateStringType(v, value) {
				return false
			}

			if strings.Contains(v, validateTypeReq.ToString()) {
				if l > 0 {
					continue
				} else {
					return false
				}
			}

		case validateTypeEq.ToString(): // =
			if !eqInt(v, l) {
				return false
			}
		case validateTypeNe.ToString(): // !=
			if !neInt(v, l) {
				return false
			}
		case validateTypeGt.ToString(): // >
			if !gtInt(v, l) {
				return false
			}
		case validateTypeGte.ToString(): // >=
			if !gteInt(v, l) {
				return false
			}

		case validateTypeLt.ToString(): // >=
			if !ltInt(v, l) {
				return false
			}
		case validateTypeLte.ToString(): // <=
			if !lteInt(v, l) {
				return false
			}
		}

	}
	return true
}

// validateStringType 验证字符串类型
//  str: tag字符串
//  value: 实际值
func (this parse) validateStringType(str string, value string) bool {
	for k, v := range validateFuncMap {
		if strings.HasPrefix(str, k.ToString()) {
			var isok, _ = this.com.validateFunc(value, v)
			return isok
		}
	}
	return true

}

// getFail 获取显示错误信息
func (this *parse) getFail(value reflect.Value) string {
	var res = this.tag.Get(validateTypeFail.ToString())
	if len(res) > 0 {
		return res
	}
	return value.Type().Field(this.index).Name
}

// getTag 获取tag字符串
func (this *parse) getTag(pType reflect.Type) reflect.StructTag {
	this.tag = pType.Field(this.index).Tag
	return this.tag
}

// parseTag 解析tag数据
func (this *parse) parseTag() {
	parseValue = make(map[string]string)
	for _, v := range parseType {
		var vl = strings.ToLower(this.tag.Get(v))
		if len(vl) > 0 {
			parseValue[v] = vl
		}
	}
}

// getTagValue 获取tag对应的值
//  tagName: tag名称
//  return: tag值
func (this *parse) getTagValue(tagName string) string {
	return this.tag.Get(tagName)
}
