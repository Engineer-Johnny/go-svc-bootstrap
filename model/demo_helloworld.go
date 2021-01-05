package model

import "time"

type HelloWorld struct {
	COMMON_MODEL
	Name        string    `json:"name" form:"name" validate:"required" gorm:"column:name;comment:Name;type:varchar(100);size:100;"`
	Age         int       `json:"age" form:"age" gorm:"column:age;comment:Age;type:varchar(300);size:300;"`
	Description string    `json:"description" form:"description" gorm:"column:description;comment:Description;type:text;"`
	Email       string    `json:"email" form:"email" gorm:"column:email;comment:Email;type:varchar(100);size:100;"`
	Type        string    `json:"type" form:"type"  gorm:"column:type;comment:Typs;type:varchar(30);size:30;"`
	StartTime   time.Time `json:"start_time" form:"start_time" validate:"required" example:"2016-07-30T16:09:51.692226358+08:00" gorm:"column:start_time;comment:StartTime;type:datetime;"`
	EndTime     time.Time `json:"end_time" form:"end_time" validate:"required" example:"2016-07-30T16:09:51.692226358+08:00" gorm:"column:end_time;comment:EndTime;type:datetime;"`
	Status      string    `json:"status" form:"status" validate:"required" gorm:"column:status;comment:Status;type:varchar(30);size:30;"`
}
