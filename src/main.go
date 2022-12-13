package main

import (
	"app/dbstorage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"app/user"
	"app/auth"
)

func main() {
	// Init Database
	db := dbstorage.NewDB()
	db.Migrate()

	// Init http
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user/profile/", user.controllers.UserDetailController)
	router.POST("/user/create/", user.controllers.UserCreateController)
	router.POST("/user/login/", auth.controllers.UserLoginController)

	router.Run()
}
