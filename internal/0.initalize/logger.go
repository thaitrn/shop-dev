package initalize

import (
	"shop-dev/global"
	"shop-dev/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
