package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config MovieGo
type Config struct {
	Port       string `mapstructure:"PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
	ApiSecret  string `mapstructure:"API_SECRET"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`

	// SMTP configuration
	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
}

// LoadConfig MovieGo
func LoadConfig() (c Config, err error) {
	// viper.AddConfigPath("./pkg/common/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	// viper.SetConfigFile("./pkg/common/config/envs/.env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	fmt.Println(viper.ConfigFileUsed())
	viper.Unmarshal(&c)
	return
}
