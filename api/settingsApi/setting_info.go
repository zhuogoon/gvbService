package settingsApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "xxx",
	})
}
