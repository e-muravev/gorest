package handlers

import (
	"app/user/models"
	"app/user/repositories"
)

func CreateNewUserHandler(name string, email string, passowrd string) models.UserModel {
	user_repo := repositories.NewUserCreateRepository()

	return user_repo.UserCreate(name, email, passowrd)
}
