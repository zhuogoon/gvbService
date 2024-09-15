package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb/core"
	"gvb/global"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warn("xixixi")
	logrus.Warn("lalal")
	//连接数据库
	global.DB = core.InitGorm()

	fmt.Println(global.Config, global.DB)
}
