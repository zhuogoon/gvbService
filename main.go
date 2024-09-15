package main

import (
	"gvb/core"
	"gvb/global"
	"gvb/router"
)

func main() {
	//读取配置文件
	core.InitConfig()

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()

	//初始化路由
	r := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_service运行在: %s", addr)
	r.Run(addr)

}
