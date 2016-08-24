package validate

import (
	"strconv"
)

// gtInt 大于 value > str
func gtInt(str string, value int) bool {
	var v, er = strconv.Atoi(str)
	return er == nil && value > v
}

// gteInt 大于等于 value >= str
func gteInt(str string, value int) bool {
	var v, er = strconv.Atoi(str)
	return er == nil && value >= v
}

// ltInt 小于 value < str
func ltInt(str string, value int) bool {
	var v, er = strconv.Atoi(str)
	return er == nil && value < v
}

// lteInt 小于等于 value <= str
func lteInt(str string, value int) bool {
	var v, er = strconv.Atoi(str)
	return er == nil && value <= v
}

// eqInt 等于 str= value
func eqInt(str string, value int) bool {
	var v, er = strconv.Atoi(str)
	return er == nil && v == value
}

// eqInt 不等于 str != value
func neInt(str string, value int) bool {
	var v, er = strconv.Atoi(str)
	return er == nil && v != value
}

// gtgtFloatInt 大于  value > str
func gtFloat(str string, value float64) bool {
	var v, er = strconv.ParseFloat(str, 64)
	return er == nil && value > v
}

// gteFloat 大于等于  value >= str
func gteFloat(str string, value float64) bool {
	var v, er = strconv.ParseFloat(str, 64)
	return er == nil && value >= v
}

// ltFloat 小于  value < str
func ltFloat(str string, value float64) bool {
	var v, er = strconv.ParseFloat(str, 64)
	return er == nil && value < v
}

// lteFloat 小于等于  value <=
func lteFloat(str string, value float64) bool {
	var v, er = strconv.ParseFloat(str, 64)
	return er == nil && value <= v
}

// eqFloat 等于 = value
func eqFloat(str string, value float64) bool {
	var v, er = strconv.ParseFloat(str, 64)
	return er == nil && v == value
}

// neFloat 不等于 != value
func neFloat(str string, value float64) bool {
	var v, er = strconv.ParseFloat(str, 64)
	return er == nil && v != value
}
