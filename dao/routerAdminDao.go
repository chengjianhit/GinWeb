package dao

import (
	"bytes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	configInfo "routerWeb/conf"
)


var DBConnection *gorm.DB

var conf configInfo.MySqlConfig
var fdnDBUrl string

func GetDB()  *gorm.DB{
	return DBConnection
}

func setDB(db *gorm.DB) {
	DBConnection = db
}

func InitDB(dbConfig configInfo.MySqlConfig)  {
	conf = dbConfig
	fdnDBUrl = getDBUrl(conf)
	InitDBConnection()
}
func getDBUrl(dbConf configInfo.MySqlConfig) string{
	var dbUrlBuffer bytes.Buffer
	dbUrlBuffer.WriteString(dbConf.Username)
	dbUrlBuffer.WriteString(":")
	dbUrlBuffer.WriteString(dbConf.Password)
	dbUrlBuffer.WriteString("@")
	dbUrlBuffer.WriteString(dbConf.FdnUrl)
	return dbUrlBuffer.String()
}

func InitDBConnection(){
	//使用 “:=”只能对局部变量进行赋值
	_DBConnection, connErr := gorm.Open("mysql", fdnDBUrl)
	if connErr != nil {
		log.Fatal("db connection error", connErr)
		panic("failed to connect database")
	}

	_DBConnection.SingularTable(true)
	_DBConnection.DB().SetMaxOpenConns(10)
	_DBConnection.DB().SetMaxIdleConns(5)

	setDB(_DBConnection)
}




