package entity

import (
	"time"
)

type Mysql struct {
	Username string
	Password string
	FdnUrl string
	GtyUrl string
}

type Redis struct {
	Address string
}

type Test struct {
	Mysql Mysql
	Redis Redis
}

type RouterInfo struct {
	ID int64 `gorm:"type:bigint(12);column:id;primary_key;AUTO_INCREMENT"`
	RouteType int16
	ProductType int16
	AppId int64
	Uri string
	BussinessId int64
	ProtocolType int16
	Environment int16
	Version string
	Ip string
	Port int16
	Status int16
	ServerStatus int16
	CreateTime time.Time
	UpdateTime time.Time
	LastOperator string
	Attachment string

}

type RouterInfoDto struct {
	ID int64
	RouteType int16
	ProductType int16
	AppId int64
	Uri string
	BusinessId int64
	ProtocolType int16
	Environment int16
	Version string
	Ip string
	Port int16
	Status int16
	ServerStatus int16
	CreateTime string
	UpdateTime string
	LastOperator string
}
