package router

import (
	"hello-gin/m/v2/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", controller.HelloController)
}
