package router

import (
	"gvb/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("settings", settingsApi.SettingInfoView)
}
