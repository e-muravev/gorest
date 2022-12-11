package handlers

func PasswordComparingValidator(value *string, value_repeated *string) bool {
	if *value != *value_repeated {
		return false
	}
	return true
}
