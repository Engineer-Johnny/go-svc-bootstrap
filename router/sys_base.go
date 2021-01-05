package router

import (
	v1 "go-svc-bootstrap/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("login", v1.Login)
	}
	return BaseRouter
}
