package handlers

func PasswordMinLengthValidator(value *string) bool {
	if len(*value) < 8 {
		return false
	}
	return true
}
