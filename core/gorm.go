package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb/global"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql，取消grom连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		//开发环境展示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //只打印错误的sql
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDb.SetMaxOpenConns(100)              //最多可容纳
	sqlDb.SetConnMaxLifetime(time.Hour * 4) //连接最大复用的时间，不能超过MySQL的wait_timeout
	return db
}
