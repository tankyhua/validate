package validate

//对外接口

// Validate
func Validate(model interface{}) (string, bool) {
	var validator StructValidator
	return validator.Validate(model)
}
