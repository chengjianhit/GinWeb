package service

import (
	"routerWeb/entity"
	routerDao "routerWeb/dao"
	)

func InsertRouter(router *entity.RouterInfo){
	db := routerDao.GetDB()
	db.Table("ai_fdn_server_route_info").Create(router)

}
