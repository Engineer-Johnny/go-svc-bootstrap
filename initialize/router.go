package initialize

import (
	_ "go-svc-bootstrap/docs"
	"go-svc-bootstrap/global"
	"go-svc-bootstrap/middleware"
	"go-svc-bootstrap/router"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// init whole routers

func Routers() *gin.Engine {
	var Router = gin.Default()

	global.G_LOG.Info("use middleware logger")

	Router.Use(middleware.Cors())
	global.G_LOG.Info("use middleware cors")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.G_LOG.Info("register swagger handler")

	APIGroupV1 := Router.Group("v1")

	router.InitBaseRouter(APIGroupV1)       // Base router
	router.InitHelloWorldRouter(APIGroupV1) // hello world demo router

	global.G_LOG.Info("router register success")
	return Router
}
