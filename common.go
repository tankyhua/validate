package validate

import (
	"reflect"
	"regexp"
	"sync"
)

type common struct {
}

var (
	mu sync.Mutex //互斥锁
	// tagMap 结构体tag Map
	tagMap = make(map[string]reflect.StructTag)

	// modelMap 结构体名称 Map string -model->fieldName->tag
	modelMap = make(map[string]map[string]reflect.StructTag)
)

type validateFuncType string

const (
	isPhone    validateFuncType = "phone"    //是否为电话
	isNum      validateFuncType = `num`      //是否为整数
	isFloat    validateFuncType = `float`    //是否是浮点数
	isEmail    validateFuncType = `email`    //是否为邮箱
	isUrl      validateFuncType = `url`      //是否为httpUrl
	isLetter   validateFuncType = `letter`   //是否为英文字母
	isPostcode validateFuncType = `postcode` //是否为邮编
	isIDCard   validateFuncType = `idcard`   //是否为身份证号
	isChinese  validateFuncType = `chinese`  //是否为中文
	isBankcard validateFuncType = `bankcard` //是否为银行号码
	isAccount  validateFuncType = `account`  //是否为帐号
)

func (this validateFuncType) ToString() string {
	return string(this)
}

//  validateFuncMap 类型验证正则表达式
var validateFuncMap = map[validateFuncType]string{
	isPhone:    `^1[3587]\d{9}$|^(0\d{2,3}-?|\(0\d{2,3}\))?[1-9]\d{4,7}(-\d{1,8})?$`,
	isNum:      `^-?\d+$`,
	isFloat:    `^(-?\d+)(\.\d+)?$`,
	isEmail:    `^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`,
	isUrl:      `^[a-zA-z]+://(\w+(-\w+)*)(\.(\w+(-\w+)*))*(\?\S*)?$`,
	isLetter:   `^[A-Za-z]+$`,
	isPostcode: `[1-9]{1}(\d+){5}`,
	isIDCard:   `^(\d{15}$|^\d{18}$|^\d{17}(\d|X|x))$`,
	isChinese:  `^[\p{Han}]+$`,
	isBankcard: `^(\d{16}|\d{19})$`,
	isAccount:  `^\w{6,20}$`,
}

func setFuncMap(tp validateFuncType, regex string) {
	mu.Lock()
	defer mu.Unlock()
	validateFuncMap[tp] = regex
}

//  ValidateFunc 通用正则验证方法
func (this *common) validateFunc(value string, regular string) (bool, error) {

	return regexp.MatchString(regular, value)

}

func (this *common) setModelMap(name, fieldName string, tag reflect.StructTag) {
	if len(string(tag)) > 0 { //防止是匿名字段
		mu.Lock()
		defer mu.Unlock()
		if modelMap[name] != nil { //有记录时
			tagMap = modelMap[name]
		} else {
			//清空tagMap数据
			tagMap = make(map[string]reflect.StructTag)
		}
		tagMap[fieldName] = tag
		modelMap[name] = tagMap
	}

}

func (this *common) getMoldeMap(name string) (map[string]reflect.StructTag, bool) {
	var res, ok = modelMap[name]
	return res, ok
}
