package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func UserExistsByNameHandler(name string) (models.UserModel, error) {
	user_repo := repositories.NewUserRepository()

	user, err := user_repo.UserExistsByName(name)

	return user, err
}
