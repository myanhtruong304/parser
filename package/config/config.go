package config

import (
	"github.com/spf13/viper"
)

// Config stores all the configuration of the application.
// The values are read by viper from a config file or env variables.
type Config struct {
	BSC_SCAN_API_KEY  string `mapstructure:"BSC_SCAN_API_KEY"`
	ETH_SCAN_API_KEY  string `mapstructure:"ETH_SCAN_API_KEY"`
	TRON_SCAN_API_KEY string `mapstructure:"TRON_SCAN_API_KEY"`

	BSC_SCAN_URI  string `mapstructure:"BSC_SCAN_URI"`
	ETH_SCAN_URI  string `mapstructure:"ETH_SCAN_URI"`
	TRON_SCAN_URI string `mapstructure:"TRON_SCAN_URI"`

	BSC_RCP_URI string `mapstructure:"BSC_RCP_URI"`
	ETH_RCP_URI string `mapstructure:"ETH_RCP_URI"`

	API_SERVER string `mapstructure:"API_SERVER"`

	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or env variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
