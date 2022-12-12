package handlers

import (
	"app/user/repositories"
)

func UserExistsByEmailHandler(email string) bool {
	user_repo := repositories.NewUserRepository()

	return user_repo.UserExistsByEmail(email)
}
