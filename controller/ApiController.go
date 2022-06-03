package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {
}

func (apiController ApiController) ApiController(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/list", func(context *gin.Context) {
		context.String(http.StatusOK, "api list - executed")
	})

	routerGroup.POST("/add", func(context *gin.Context) {
		context.String(http.StatusOK, "api add - executed")
	})

}
