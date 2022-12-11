package controllers

import "github.com/gin-gonic/gin"

func UserDetailController(c *gin.Context) {
	c.JSON(200, gin.H{"FirstName": "Vasya1"})
}
