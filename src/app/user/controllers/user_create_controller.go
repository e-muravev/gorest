package controllers

import (
	"fmt"
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

	// Check if user exists
	_, err := handlers.UserExistsByNameHandler(serializer_data.Name)

	// fmt.Println("\n", user_exists, err, "\n")
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"name": fmt.Sprintf("user with name='%s' exists", serializer_data.Name)})
		return
	}

	serializer_data.Password, _ = handlers.PasswordHashingHandler(serializer_data.Password)

	user := handlers.CreateNewUserHandler{Repository: repositories.NewUserCreateRepository()}.Run(
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
