package main

import (
	"myproject/src/api/controller"
	"myproject/src/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	router.POST("/create", controller.CreatePost )
	router.GET("/get",controller.FindPost)
	router.GET("/get/:id",controller.FindPost )
	router.PATCH("/patch/:id", controller.UpdatePost)
	router.DELETE("/delete/:id", controller.DeletePost)
	router.Run(":8080")

}