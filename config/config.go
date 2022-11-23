package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	defaultConfigPath = "/etc/cpds/"
	defaultConfigName = "cpds-analyzer.json"
	defaultConfigType = "json"
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
	ConfigPath       string
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(flags *pflag.FlagSet) {
	c.installFlags()
	c.parseConfigFile(flags)
}

func (c *Config) parseConfigFile(flags *pflag.FlagSet) {
	if path, err := flags.GetString("config-path"); err != nil {
		// using defalut config file path
		viper.AddConfigPath(defaultConfigPath)
	} else {
		viper.AddConfigPath(path)
	}
	viper.BindPFlags(flags)
	viper.SetConfigType(defaultConfigType)
	viper.SetConfigName(defaultConfigName)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Infof("Failed to read config file: %s", err)
		panic(fmt.Errorf("fatal error config file: %w", err))
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
