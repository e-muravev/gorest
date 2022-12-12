package exceptions

import "errors"

var UserExistsByNameError = errors.New("user_exists_by_name")
