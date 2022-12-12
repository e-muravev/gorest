package exceptions

import "errors"

var PasswordComparingError = errors.New("password and password_repeated should be the same")
