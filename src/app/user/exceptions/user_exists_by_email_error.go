package exceptions

import "errors"

var UserExistsByEmailError = errors.New("user_exists_by_email")
