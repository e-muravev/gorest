package main

import (
	"net/http"
	"time"

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
	// Ctime    time.Time
}

type UserCreateSerializer struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	PasswordRepeated string `json:"password_repeated" binding:"required"`
}

type UserCreateResponseSerializer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
}

func UserDetailView(c *gin.Context) {
	c.JSON(200, gin.H{"FirstName": "Vasya1"})
}

func PasswordMinLengthValidator(value *string) bool {
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

	if !PasswordMinLengthValidator(&serializer_data.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"password": "Enter a valid password"})
		return
	}

	if !PasswordComparingValidator(&serializer_data.Password, &serializer_data.PasswordRepeated) {
		c.JSON(http.StatusBadRequest, gin.H{"password": "password and password_repeated should be the same"})
		return
	}

	user := UserModel{
		Name:     serializer_data.Name,
		Email:    serializer_data.Email,
		Password: serializer_data.Password,
		// Ctime:    time.Now(),
	}

	dbconn().Create(&user)

	// c.JSON(http.StatusCreated, gin.H{"status": "OK", "user_id": user.ID, "ctime": user.Ctime})

	resp := UserCreateResponseSerializer{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	c.JSON(http.StatusCreated, resp)
}

func dbconn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserModel{})
	return db
}

func main() {
	// Init Database
	dbconn()

	// Init http
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user/profile/", UserDetailView)
	router.POST("/user/create/", UserCreateView)

	router.Run()
}
