package api

import "gvb/api/settingsApi"

type ApiGroup struct {
	SettingsApi settingsApi.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
