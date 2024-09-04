package global

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.AddConfigPath("conf")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.WatchConfig()
	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
