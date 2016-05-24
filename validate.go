package validatestruct

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	tagRequried = "req"   //是否需要验证
	tagType     = "type"  //验证字符串类型exp: phone
	tagLen      = "len"   //是否判断长度
	tagGte      = "gte"   //大于等于 >=
	tagGt       = "gt"    //大于 >
	tagLte      = "ltg"   //小于等于 <=
	tagLt       = "lt"    //小于
	tagEq       = "eq"    //等于 =
	tagError    = "error" //自定义返回的错误信息
)

type result struct {
	IsOk         bool
	ReturnString string
	value        reflect.Value
	index        int
}

// Validate 验证结构体数据
//  model:结构体指针
func Validate(model interface{}) (bool, string) {
	var value = reflect.ValueOf(model).Elem()
	var res result
	res.IsOk = true
	res.validateStruct(value)
	return res.IsOk, res.ReturnString
}

func (this *result) validateStruct(value reflect.Value) {
	var kind = value.Kind()
	var num = 0
	if kind == reflect.Struct { //类型是否是结构体
		num = value.NumField()
		for i := 0; i < num; i++ {
			this.value = value
			this.index = i
			var strVal, isExist = this.getTag(tagRequried)
			if isExist && strings.Contains(strVal, "true") {
				if value.Field(i).Kind() == reflect.Slice { // 判断字段是否为结构体数组
					var length = value.Field(i).Len()
					// 是否切片可以为空
					if length <= 0 {
						this.IsOk = false
						this.ReturnString = this.returnFail()
						return
					}
					//遍历切片
					for j := 0; j < length; j++ {
						this.validateStruct(value.Field(i).Index(j))
					}
				} else { //不是结构体数组情况
					this.validateStructOne()
					if !this.IsOk {
						return
					}
				}
			}
		}

	} else if kind == reflect.Slice {
		var lenght = value.Len()
		if lenght == 0 {
			this.IsOk = false
			this.ReturnString = "数组不能为空!"
			return
		}
		for j := 0; j < lenght; j++ {
			this.validateStruct(value.Index(j))
			if !this.IsOk {
				break
			}
		}
	} else {
		this.IsOk = false
		this.ReturnString = "不是结构体类型"
		return
	}

}

// validateStructOne 验证单个结构体数据
func (this *result) validateStructOne() {
	var value = this.value
	var i = this.index
	var FieldType = value.Type().Field(i)
	if FieldType.Anonymous {
		//匿名组合字段,进行递归验证
		this.validateStruct(value.Field(i))
	} else {

		switch FieldType.Type.String() {
		case "string":
			//正则表达式验证字符串值是否满足类型条件
			this.isPassforStr()
			if !this.IsOk {
				return
			}
			break
		case "int", "int32", "int64":
			this.isPassforInt()
			if !this.IsOk {
				return
			}
			break
		case "float32", "float64":
			this.isPassforFloat()
			if !this.IsOk {
				return
			}
			break
		case "time.Time":
			this.isPassforTime()
			if !this.IsOk {
				return
			}
		}

	}
}

// returnFailColumnName 返回错误字段
//  paramsFail:错误信息
//  columnName:不符合条件的字段名
func (this *result) returnFail() string {
	var columnName = this.value.Type().Field(this.index).Name
	var strFail, isExist = this.getTag(tagError)
	if !isExist {
		return columnName
	}
	return columnName + "  " + strFail
}

// getTag 获取tag内容信息
//  tagName: tag名称
//  return: string-tag内容字符串 bool-该tag是否存在
func (this *result) getTag(tagName string) (string, bool) {
	var tag = this.value.Type().Field(this.index).Tag
	var val = tag.Get(tagName)
	if len(val) > 0 {
		return val, true
	}
	return val, false
}

// isPassforStr 验证字符串是否通过
func (this *result) isPassforStr() {
	var fieldValue = this.value.Field(this.index).String()
	this.ReturnString = this.returnFail()
	//字符串为空
	if strings.Trim(fieldValue, " ") == "" {
		this.IsOk = false
		return
	}
	var val = strings.ToLower(fieldValue)
	// 正则验证字符串类型
	var vType, isType = this.getTag(tagType)
	if isType {
		if !this.checkStrType(val, vType) {
			this.IsOk = false
			return
		}
	}
	this.check(fieldValue, 1)
	if this.IsOk { //当验证成功,清空返回信息数据
		this.ReturnString = ""
	}

}

// isPassforInt 验证int数值是否通过
func (this *result) isPassforInt() {
	this.ReturnString = this.returnFail() //默认初始返回信息
	var val = strconv.Itoa(int(this.value.Field(this.index).Int()))
	this.check(val, 2)
	if this.IsOk { //当验证成功,清空返回信息数据
		this.ReturnString = ""
	}
}

// isPassforFloat 验证float数值是否通过
func (this *result) isPassforFloat() {
	this.ReturnString = this.returnFail() //默认初始返回信息
	var val = strconv.FormatFloat(this.value.Field(this.index).Float(), 'f', -1, 64)
	this.check(val, 3)
	if this.IsOk { //当验证成功,清空返回信息数据
		this.ReturnString = ""
	}
}

// isPassforTime 验证时间是否大于2006-01-02 15:04:05
func (this *result) isPassforTime() {
	this.ReturnString = this.returnFail()
	var Tt time.Time
	Tt = this.value.Field(this.index).Interface().(time.Time)
	var strTime = "2006-01-02 15:04:05"
	var ti, _ = time.Parse("2006-01-02 15:04:05", strTime)
	if Tt.Before(ti) {
		this.IsOk = false
		return
	}
}

// check 数值验证
//  val: 数值字符串
//  checkType: 验证类型 1-字符串长度验证 2-int数值验证 3-float数值验证
func (this *result) check(val string, checkType int) {
	var value float64 = 0.0
	switch checkType {
	case 1: //字符串长度验证
		value = float64(len(val))
	case 2: //int数值验证
		var num, _ = strconv.ParseFloat(val, 64)
		value = float64(num)
	case 3: //float数值验证
		value, _ = strconv.ParseFloat(val, 64)
	}
	//与数值相关的验证
	this.checkforNum(value, tagEq)
	if !this.IsOk {
		return
	}
	this.checkforNum(value, tagGt)
	if !this.IsOk {
		return
	}
	this.checkforNum(value, tagGte)
	if !this.IsOk {
		return
	}
	this.checkforNum(value, tagLt)
	if !this.IsOk {
		return
	}
	this.checkforNum(value, tagLte)
	if !this.IsOk {
		return
	}

}

// checkStrType 正则验证字符串类型
//  val:字符串值
//  tagName:tag类型内容
func (this *result) checkStrType(val, tagName string) bool {
	if strings.Contains(tagName, "account") { //验证帐号
		this.IsOk = IsAccount(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "phone") { //验证电话号码
		this.IsOk = IsPhone(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "chinese") { //验证中文
		this.IsOk = IsChinese(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "idcard") { //验证身份证
		this.IsOk = IsIDCard(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "bankcard") { //验证银行卡
		this.IsOk = IsBankcard(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "email") { //验证邮箱
		this.IsOk = IsEmail(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "postcode") { //验证邮编
		this.IsOk = IsPostcode(val)
		if !this.IsOk {
			return false
		}
	}
	if strings.Contains(tagName, "url") { //验证url
		this.IsOk = IsUrl(val)
		if !this.IsOk {
			return false
		}
	}
	return true
}

// checkforNum 关于条件操作符判断
//  value: 结构体实际数值或者字符串长度
//  compareType: 条件操作符类型
func (this *result) checkforNum(value float64, compareType string) {
	var v, isExist = this.getTag(compareType)
	if !isExist {
		return
	}
	// 获取转换数值
	var num, err = strconv.ParseFloat(v, 64)
	if err != nil {
		this.ReturnString = compareType + " 数值错误!"
		this.IsOk = false
	}
	switch compareType {
	case tagEq: //等于
		if value != num {
			this.IsOk = false
		}
	case tagGte: //大于等于
		if value < num {
			this.IsOk = false
		}
	case tagGt: //大于
		if value <= num {
			this.IsOk = false
		}
	case tagLte: //小于等于
		if value > num {
			this.IsOk = false
		}
	case tagLt: //小于
		if value >= num {
			this.IsOk = false
		}
	}

}
