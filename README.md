# go-backend-bootstrap
Backend bootstrap in Go

## Relase Notes

### Includes frameworks:

1. Gin - Web framework in Go
2. Viper - Complete configuration solution for Go applications
3. Zap - Uber created it. Blazing fast, structured, leveled logging in Go
4. *Gorm - The fantastic ORM library for Golang, aims to be developer friendly (We don't need it for now, keep it as compatibel mode)*
5. *Swag - Automatically generate RESTful API documentation with Swagger 2.0 for Go (We don't need it for now, keep it as compatibel mode)*


### How to use it

1. Creat struct in `model` folder
2. Creat new router in `router` folder
3. Go to `initialize/router.go` add one like `router.InitHelloWorldRouter(APIGroupV1)` to register your new router
4. Create new API method in `api/v1/` folder
5. Create new service method in `service`
6. Done

### How to run it

`make run`
