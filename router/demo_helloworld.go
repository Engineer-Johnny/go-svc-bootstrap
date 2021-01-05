package router

import (
	v1 "go-svc-bootstrap/api/v1"

	"github.com/gin-gonic/gin"
)

func InitHelloWorldRouter(Router *gin.RouterGroup) {
	HelloWorldRouter := Router.Group("demo")
	{
		HelloWorldRouter.POST("helloworld", v1.CreateHelloWorld)
		HelloWorldRouter.DELETE("helloworld", v1.DeleteHelloWorld)
		HelloWorldRouter.PUT("helloworld", v1.UpdateHelloWorld)
		HelloWorldRouter.GET("helloworld", v1.GetHelloWorldList)
		HelloWorldRouter.GET("helloworld/:uuid", v1.GetHelloWorldByUUID)
	}
}
