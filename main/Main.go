package main

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	type Data struct {
		A string
		B int
	}
	apiGroup := engine.Group("/api")
	adminGroup := engine.Group("/admin")

	apiController := controller.ApiController{}
	adminController := controller.AdminController{}

	apiController.ApiController(apiGroup)
	adminController.AdminController(adminGroup)

	_ = engine.Run()
}
