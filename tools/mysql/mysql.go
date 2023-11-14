package mysql

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _mysqlDb *gorm.DB

func initMysql() {
	username := "root"
	password := "123456"
	host := "localhost"
	port := 3306
	Dbname := "netstorage"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		klog.Error("gorm connect mysql error", err)
	}
	_mysqlDb = db
}

func GetMysql() *gorm.DB {
	if _mysqlDb != nil {
		return _mysqlDb
	}
	initMysql()
	return _mysqlDb
}
