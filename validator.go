package validate

//  Valiadator 验证器
type Validator interface {
	// Validate 验证数据 返回错误信息string,是否验证成功bool
	Validate(model interface{}) (string, bool)
}
