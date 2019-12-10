package config

import (
	"github.com/samuelmjn/go-library/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig :nodoc:
func InitConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// DatabaseHost :nodoc:
func DatabaseHost() string {
	key := viper.GetString("dsn")
	return utils.UseEnvIfExists(key)
}

// Port :nodoc:
func Port() string {
	key := viper.GetString("port")
	return utils.UseEnvIfExists(key)
}
