package v1

import (
	"go-svc-bootstrap/global"
	"go-svc-bootstrap/model"
	"go-svc-bootstrap/model/request"
	"go-svc-bootstrap/model/response"
	"go-svc-bootstrap/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Public APIs

// @Tags HelloWorld
// @Summary Create BizStore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelloWorld true "Create HelloWorld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"create successfully"}"
// @Router /v1/demo/helloworld [post]
func CreateHelloWorld(c *gin.Context) {
	var helloWorld model.HelloWorld
	_ = c.ShouldBindJSON(&helloWorld)

	if err := service.CreateHelloWorld(helloWorld); err != nil {
		global.G_LOG.Error("create fail", zap.Any("err", err))
		response.FailWithMessage("create successfully", c)
	} else {
		response.OkWithMessage("create successfully", c)
	}
}

// @Tags HelloWorld
// @Summary Delete HelloWorld by UUIDs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UUIDsReq true "Delete HelloWorld by UUIDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"delete successfully"}"
// @Router /v1/demo/helloworld [delete]
func DeleteHelloWorld(c *gin.Context) {
	var uuids request.UUIDsReq
	_ = c.ShouldBindJSON(&uuids)
	if err := service.DeleteHelloWorldByUUIDs(uuids); err != nil {
		global.G_LOG.Error("delete fail", zap.Any("err", err))
		response.FailWithMessage("delete fail", c)
	} else {
		response.OkWithMessage("delete successfully", c)
	}
}

// @Tags HelloWorld
// @Summary Get HelloWorld per pages
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page query int false "page number"
// @Param pageSize query int false "items per page"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"get successfully"}"
// @Router /v1/demo/helloworld [get]
func GetHelloWorldList(c *gin.Context) {
	var pageInfo request.HelloWorldPagination
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetHelloWorldListPagination(pageInfo); err != nil {
		global.G_LOG.Error("get fail", zap.Any("err", err))
		response.FailWithMessage("get fail", c)
	} else {
		if total <= 0 {
			response.FailWithMessage("no results", c)
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "get successfully", c)
		}
	}
}

// @Tags HelloWorld
// @Summary Get helloworld by UUID
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param uuid path string true "Get helloworld by UUID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"get successfully"}"
// @Router /v1/demo/helloworld/{uuid} [get]
func GetHelloWorldByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if err, reHelloWorld := service.GetHelloWorldByUUID(uuid); err != nil {
		global.G_LOG.Error("get fail", zap.Any("err", err))
		response.FailWithMessage("get fail", c)
	} else {
		if reHelloWorld.UUID == "" {
			response.FailWithMessage("no results", c)
		} else {
			response.OkWithData(gin.H{"helloWorld": reHelloWorld}, c)
		}
	}
}

// @Tags HelloWorld
// @Summary Update helloworld
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HelloWorld true "Update helloworld"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update successfully"}"
// @Router /v1/demo/helloworld [put]
func UpdateHelloWorld(c *gin.Context) {
	var helloWorld model.HelloWorld
	_ = c.ShouldBindJSON(&helloWorld)
	if err := service.UpdateHelloWorld(&helloWorld); err != nil {
		global.G_LOG.Error("update fail", zap.Any("err", err))
		response.FailWithMessage("update fail", c)
	} else {
		response.OkWithMessage("update successfully", c)
	}
}
