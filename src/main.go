package main

import (
	db "app/storage"
	"app/user/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init Database
	db.DB()

	// Init http
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user/profile/", controllers.UserDetailController)
	router.POST("/user/create/", controllers.UserCreateController)

	router.Run()
}
