package initialize

import (
	"fmt"

	"github.com/DoCongThanhPhuong/go-backend/global"
	"github.com/spf13/viper"
)



func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config file
	viper.SetConfigName("local") // filename
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w", err))
	}

	// configure structure
	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}