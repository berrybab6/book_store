package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config MovieGo
type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

// LoadConfig MovieGo
func LoadConfig() (c Config, err error) {
	// viper.AddConfigPath("./pkg/common/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.SetConfigFile("./pkg/common/config/envs/.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	fmt.Println(viper.ConfigFileUsed())
	viper.Unmarshal(&c)
	return
}
