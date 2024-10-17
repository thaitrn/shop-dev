package global

import (
	"shop-dev/pkg/logger"
	"shop-dev/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config  setting.Config
	Logger  *logger.LoggerZap
	MysqlDB *gorm.DB
)
