package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"routerWeb/constants"
	routerDao "routerWeb/dao"
	"routerWeb/entity"
	"routerWeb/logger"
	"time"
)


func AddRouterInfo(w http.ResponseWriter, req *http.Request)  {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil{

	}
	io.WriteString(w, string(bytes.NewBuffer(result).Bytes()))
}



func RouterService(w http.ResponseWriter, req *http.Request)  {
	logger := logger.GetLogger()
	id := req.FormValue("id")
	//logger.Debug("testlog logger")
	var routeInfo entity.RouterInfo
	// 根据条件查询得到满足条件的第一条记录
	db := routerDao.GetDB()
	db.Table("ai_fdn_server_route_info").Where("id = ?", id).First(&routeInfo)

	formatTime := routeInfo.CreateTime.Format(constants.TimeFormat)

	fmt.Print("\n formateTime is ", formatTime)
	logger.Info("routerService ", zap.String("formateTime", formatTime))

	newRouterInfo :=&entity.RouterInfo{
		RouteType:1,
		BussinessId:20200309165712345,
		AppId: 1000000892,
		UpdateTime:time.Now(),
		CreateTime:time.Now(),
	}

	InsertRouter(newRouterInfo)


	routerDto := &entity.RouterInfoDto{
		ID:newRouterInfo.ID,
		RouteType:newRouterInfo.RouteType,
		AppId:newRouterInfo.AppId,
		UpdateTime:newRouterInfo.UpdateTime.Format(constants.TimeFormat),
		CreateTime:newRouterInfo.CreateTime.Format(constants.TimeFormat),
	}
	retJson, _ := json.Marshal(routerDto)
	io.WriteString(w, string(retJson))
}
