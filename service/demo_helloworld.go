package service

import (
	"go-svc-bootstrap/global"
	"go-svc-bootstrap/model"
	"go-svc-bootstrap/model/request"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

// Service

//@author: [johnny](jingpeng@cn.ibm.com)
//@function: CreateHelloWorld
//@description: CreateHelloWorld
//@param: helloWorld request.HelloWorld
//@return: err error
func CreateHelloWorld(helloWorld model.HelloWorld) (err error) {
	log.Printf("helloWorld is %+v", helloWorld)

	var validate *validator.Validate
	validate = validator.New()
	//注册自定义Struct校验函数
	validate.RegisterStructValidation(checkStartTime, model.HelloWorld{})
	// _ = validate.RegisterValidation("checkValidStartTime", checkValidStartTime)
	// _ = validate.RegisterValidation("checkValidEndTime", checkValidEndTime)

	err = validate.Struct(helloWorld)
	if err != nil {
		log.Println(err)
		return
	}

	return err
}

//@author: [johnny](jingpeng@cn.ibm.com)
//@function: DeleteHelloWorldByUUIDs
//@description: DeleteHelloWorldByUUIDs
//@param: uuids request.UUIDsReq
//@return: err error
func DeleteHelloWorldByUUIDs(uuids request.UUIDsReq) (err error) {
	err = global.G_DB.Delete(&[]model.HelloWorld{}, "uuid in ?", uuids.UUIDs).Error
	return err
}

//@author: [johnny](jingpeng@cn.ibm.com)
//@function: GetHelloWorldListPagination
//@description: GetHelloWorldListPagination
//@param: info request.HelloWorldPagination
//@return: err error, list interface{}, total int64
func GetHelloWorldListPagination(info request.HelloWorldPagination) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	var helloWorldList []model.HelloWorld

	// search it in db
	db := global.G_DB.Model(&model.HelloWorld{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&helloWorldList).Error
	return err, helloWorldList, total
}

//@author: [johnny](jingpeng@cn.ibm.com)
//@function: GetHelloWorldByUUID
//@description: GetHelloWorldByUUID
//@param: uuid string
//@return: err error, bizStore model.BizStore
func GetHelloWorldByUUID(uuid string) (err error, helloWorld model.HelloWorld) {
	err = global.G_DB.Where("uuid = ?", uuid).First(&helloWorld).Error
	return
}

//@author: [johnny](jingpeng@cn.ibm.com)
//@function: UpdateHelloWorld
//@description: UpdateHelloWorld
//@param: helloWorld *model.HelloWorld
//@return: err error
func UpdateHelloWorld(helloWorld *model.HelloWorld) (err error) {
	err = global.G_DB.Where("uuid", helloWorld.UUID).Updates(helloWorld).Error
	return err
}

func checkStartTime(sl validator.StructLevel) {
	helloWorld := sl.Current().Interface().(model.HelloWorld)

	today := time.Now()
	if today.After(helloWorld.EndTime) {
		sl.ReportError(helloWorld.EndTime, "EndTime", "", "endtime > now", "")
	}

	if helloWorld.StartTime.After(helloWorld.EndTime) {
		sl.ReportError(helloWorld.StartTime, "StartTime", "", "starttime < endtime", "")
	}

}
