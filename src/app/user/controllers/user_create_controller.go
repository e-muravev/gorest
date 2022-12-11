package controllers

import (
	"net/http"

	"app/user/handlers"
	"app/user/repositories"
	"app/user/transformers"

	"github.com/gin-gonic/gin"
)

func UserCreateController(c *gin.Context) {
	var serializer_data transformers.UserCreateTransformer

	if err := c.ShouldBindJSON(&serializer_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !handlers.PasswordMinLengthValidator(&serializer_data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"password": "Enter a valid password"})
		return
	}

	if !handlers.PasswordComparingValidator(&serializer_data.Password, &serializer_data.PasswordRepeated) {
		c.JSON(http.StatusBadRequest, gin.H{"password": "password and password_repeated should be the same"})
		return
	}

	user_repo := repositories.NewUserCreateRepository()

	user := user_repo.UserCreate(
		serializer_data.Name,
		serializer_data.Email,
		serializer_data.Password,
	)

	// c.JSON(http.StatusCreated, gin.H{"status": "OK", "user_id": user.ID, "ctime": user.Ctime})

	resp := transformers.UserCreateRespTransformer{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	c.JSON(http.StatusCreated, resp)
}
