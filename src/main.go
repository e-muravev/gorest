package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Login string
	email string
}

func UserDetailView(c *gin.Context) {
	c.JSON(200, gin.H{"FirstName": "Vasya1"})
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&UserModel{})

	fmt.Println("asd")

	route := gin.Default()

	route.GET("/user/profile/", UserDetailView)

	route.Run()
}
