package core

import (
	"fmt"
	"go-svc-bootstrap/global"
	"go-svc-bootstrap/initialize"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	Router := initialize.Routers()
	// Router.Static("/form-generator", "./resource/page")

	if global.G_CONFIG.System.HTTPS {
		address := ":443"

		// In order to ensure that the text order output can be deleted
		time.Sleep(10 * time.Microsecond)
		global.G_LOG.Info("server run success on ", zap.String("address", address))

		fmt.Printf(`
	Welcome to use go-svc-bootstrap
	default current service running at: http://127.0.0.1%s
`, address)
		// m := autocert.Manager{
		// 	Prompt:     autocert.AcceptTOS,
		// 	HostPolicy: autocert.HostWhitelist("api.jbosscafe.com", "www.jbosscafe.com"),
		// 	Cache:      autocert.DirCache("/var/www/.cache"),
		// }

		// global.G_LOG.Error(autotls.RunWithManager(Router, &m).Error())
		global.G_LOG.Error(autotls.Run(Router, "example1.com", "example2.com").Error())
	} else {

		address := fmt.Sprintf(":%d", global.G_CONFIG.System.Addr)
		s := initServer(address, Router)

		// In order to ensure that the text order output can be deleted
		time.Sleep(10 * time.Microsecond)
		global.G_LOG.Info("server run success on ", zap.String("address", address))

		fmt.Printf(`
		Welcome to use go-svc-bootstrap
		default current service running at: http://127.0.0.1%s
`, address)
		global.G_LOG.Error(s.ListenAndServe().Error())
	}
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
