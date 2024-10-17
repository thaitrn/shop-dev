package initalize

import (
	"fmt"
	"shop-dev/global"
)

func Run() {
	// load configuration
	LoadConfig()
	fmt.Println("Loading configuration mysql...")
	InitLogger()
	global.Logger.Debug("Config log ok")

	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8082")
}
