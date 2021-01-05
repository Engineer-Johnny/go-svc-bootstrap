package v1

import (
	"go-svc-bootstrap/model/response"

	"github.com/gin-gonic/gin"
)

// @Tags Base
// @Summary Login
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param name query string false "name"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"get successfully"}"
// @Router /v1/base/login [get]
func Login(c *gin.Context) {
	name := c.Query("name")
	response.OkWithData(gin.H{"name": name}, c)
}
