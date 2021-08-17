package global

import (
	"github.com/goprogramming/blog/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
)

var (
	DBEngine *gorm.DB
)
