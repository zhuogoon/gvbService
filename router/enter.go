package router

import (
	"github.com/gin-gonic/gin"
	"gvb/global"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	routerGroup := RouterGroup{router}

	//系统配置api
	routerGroup.SettingsRouter()

	return router
}
