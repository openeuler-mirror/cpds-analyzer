package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Debug            bool
	LogLevel         string
	DatabaseAddress  string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	BindAddress      string
	Port             string
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(flags *pflag.FlagSet) {
	c.installFlags(flags)
	viper.BindPFlags(flags)
	c.parseConfigFile(flags)
}

func (c *Config) parseConfigFile(flags *pflag.FlagSet) {
	viper.AddConfigPath(".")
	viper.SetConfigName("cpds-analyzer.json")
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Infof("Failed to read config file: %s", err)
		os.Exit(1)
	}
	*c = Config{
		Debug:            viper.GetBool("debug"),
		LogLevel:         viper.GetString("log-level"),
		DatabaseAddress:  viper.GetString("db-address"),
		DatabasePort:     viper.GetString("db-port"),
		DatabaseUser:     viper.GetString("db-user"),
		DatabasePassword: viper.GetString("db-password"),
		BindAddress:      viper.GetString("bind-address"),
		Port:             viper.GetString("port"),
	}
}
