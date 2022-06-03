package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
}

func (adminController AdminController) AdminController(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/list", func(context *gin.Context) {
		context.String(http.StatusOK, "admin list - executed")
	})
	routerGroup.POST("/add", func(context *gin.Context) {
		context.String(http.StatusOK, "admin add - executed")
	})

}
