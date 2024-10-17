package initalize

import (
	"fmt"
	"shop-dev/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config: %w \n", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
