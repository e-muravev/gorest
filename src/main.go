package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type UserCreateSerializer struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	PasswordRepeated string `json:"password_repeated" binding:"required"`
}

func UserDetailView(c *gin.Context) {
	c.JSON(200, gin.H{"FirstName": "Vasya1"})
}

func PasswordLengthValidator(value *string) bool {
	if len(*value) < 8 {
		return false
	}
	return true
}

func PasswordComparingValidator(value *string, value_repeated *string) bool {
	if *value != *value_repeated {
		return false
	}
	return true
}

func UserCreateView(c *gin.Context) {
	var serializer_data UserCreateSerializer

	if err := c.ShouldBindJSON(&serializer_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !PasswordLengthValidator(&serializer_data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"password": "Enter a valid password"})
		return
	}

	if !PasswordComparingValidator(&serializer_data.Password, &serializer_data.PasswordRepeated) {
		c.JSON(http.StatusBadRequest, gin.H{"password": "password and password_repeated should be the same"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "OK"})
}

func main() {
	// Init Database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserModel{})

	// Init http
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user/profile/", UserDetailView)
	router.POST("/user/create/", UserCreateView)

	router.Run()
}
