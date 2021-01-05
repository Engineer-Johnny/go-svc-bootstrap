package request

import "go-svc-bootstrap/model"

type HelloWorldPagination struct {
	model.HelloWorld
	PageInfo
}
