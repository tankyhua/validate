package validatestruct

import (
	"fmt"
	"regexp"
)

//  IsNum 是否为整数
//  return: 是为true，否为false
func IsNum(strValue string) bool {
	var isOk, err = regexp.MatchString(`^-?\d+$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsFloat 是否为浮点数
//  return: 是为true，否为false
func IsFloat(strValue string) bool {
	var isOk, err = regexp.MatchString(`^(-?\d+)(\.\d+)?$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsPhone 是否为电话号码（包括手机号）
//  电话格式如：021-26565621
//  return: 是为true，否为false
func IsPhone(strValue string) bool {
	//	var isOk, err = regexp.MatchString(`^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$`, strValue)
	///^1[3587]\d{9}$|^(0\d{2,3}-?|\(0\d{2,3}\))?[1-9]\d{4,7}(-\d{1,8})?$
	var isOk, err = regexp.MatchString(`^1[3587]\d{9}$|^(0\d{2,3}-?|\(0\d{2,3}\))?[1-9]\d{4,7}(-\d{1,8})?$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsEmail 是否为邮箱
//  return: 是为true，否为false
func IsEmail(strValue string) bool {
	var isOk, err = regexp.MatchString(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsUrl 是否为Url地址
//  return: 是为true，否为false
func IsUrl(strValue string) bool {
	var isOk, err = regexp.MatchString(`^[a-zA-z]+://(\w+(-\w+)*)(\.(\w+(-\w+)*))*(\?\S*)?$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsLetter 是否为大小写字母
//  return: 是为true，否为false
func IsLetter(strValue string) bool {
	var isOk, err = regexp.MatchString(`^[A-Za-z]+$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsPostcode 是否为邮编号码
//  return: 是为true，否为false
func IsPostcode(strValue string) bool {
	var isOk, err = regexp.MatchString(`[1-9]{1}(\d+){5}`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

//  IsIDCard 是否为身份证号码
//  return: 是为true，否为false
func IsIDCard(strValue string) bool {
	var isOk, err = regexp.MatchString(`^(\d{15}$|^\d{18}$|^\d{17}(\d|X|x))$`, strValue)
	if err != nil {
		return false
	}
	return isOk
}

// IsChinese 是否为汉字  /^(\d{16}|\d{19})$/
func IsChinese(strValue string) bool {
	var isOk, err = regexp.MatchString(`^[\p{Han}]+$`, strValue)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return isOk
}

// IsBankcard 是否是银行卡号
func IsBankcard(strValue string) bool {
	var isOk, err = regexp.MatchString(`^(\d{16}|\d{19})$`, strValue)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return isOk
}

// IsAccount 是否是银行卡号
func IsAccount(strValue string) bool {
	var isOk, err = regexp.MatchString(`^\w{6,20}$`, strValue)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return isOk
}
