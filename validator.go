package validate

//  Valiadator 验证器
type Validator interface {
	// Validate 验证数据 返回错误信息string,是否验证成功bool
	Validate(model interface{}) (string, bool)

	// RegisterRegex 注册正则验证 name: 类型名 如:验证电话号码对应 name:phone regex: 正则表达式
	RegisterRegex(name, regex string) error
}

// NewValidator 实现接口
func NewValidator() Validator {
	return new(structValidator)
}
