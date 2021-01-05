package main

import (
	"go-svc-bootstrap/core"
	"go-svc-bootstrap/global"
)

// @title Swagger public APIs
// @version 0.0.1
// @description This is public APIs
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.G_VP = core.Viper() // Viper
	global.G_LOG = core.Zap()  // Zap

	core.RunServer()
}
