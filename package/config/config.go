package config

import (
	"github.com/spf13/viper"
)

// Config stores all the configuration of the application.
// The values are read by viper from a config file or env variables.
type Config struct {
	BscScanAPIKey  string `mapstructure:"BSC_SCAN_API_KEY"`
	EthScanAPIKey  string `mapstructure:"ETH_SCAN_API_KEY"`
	TronScanAPIKey string `mapstructure:"TRON_SCAN_API_KEY"`

	BscScanUri  string `mapstructure:"BSC_SCAN_URI"`
	EthScanUri  string `mapstructure:"ETH_SCAN_URI"`
	TronScanUri string `mapstructure:"TRON_SCAN_URI"`

	BscRPCUri string `mapstructure:"BSC_RCP_URI"`
	EthRPCUri string `mapstructure:"ETH_RCP_URI"`

	ApiServer string `mapstructure:"API_SERVER"`
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
