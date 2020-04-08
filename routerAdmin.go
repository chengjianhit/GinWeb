package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gopkg.in/gcfg.v1"
	"log"
	"net/http"
	config "routerWeb/conf"
	"routerWeb/constants"
	routerDao "routerWeb/dao"
	"routerWeb/logger"
	"routerWeb/service"

	_ "routerWeb/docs"
)


func main()  {
	var projectConfig config.ProjectConfig

	fileErr := gcfg.ReadFileInto(&projectConfig, constants.ResourceLocation)
	if fileErr != nil{
		fmt.Print("file load error", fileErr)
		log.Fatal("file load error", fileErr)
		panic("failed to load config")
	}

	go func() {
		//初始化数据库
		routerDao.InitDB(projectConfig.Mysql)
		//初始化日志
		logger.InitLoggerConfig(projectConfig.LoggerInfo)
	}()

	go func() {
		r := gin.New()
		//url := ginSwagger.URL("http://localhost:8055/swagger/doc.json") // The url pointing to API definition
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, /** url **/))
		r.POST("/ginTest", func(context *gin.Context) {
			ginMap :=context.PostFormMap("ID")
			fmt.Print(ginMap)
			inputName := context.PostForm("name")
			context.JSON(200, gin.H{
				"status":  "posted",
				"message": "success",
				"nick":    inputName,
			})
		})

		r.Run(":8055")
	}()


	http.HandleFunc("/fdn/router/getRouterInfo", service.RouterService)
	http.HandleFunc("/fdn/router/Add", service.AddRouterInfo)
	//err := http.ListenAndServe(":8089", nil)
	err := http.ListenAndServe(":"+projectConfig.AppConfig.Port, nil)
	if err != nil {
		log.Fatal("ListenAndService: ", err)
	}



}
