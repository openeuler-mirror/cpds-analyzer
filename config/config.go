package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	cobra.OnInitialize(func() {
		viper.SetConfigType(defaultConfigType)
		viper.SetConfigName(defaultConfigName)
		if c.ConfigPath != defaultConfigPath {
			// Use config file from the flag.
			viper.AddConfigPath(c.ConfigPath)
		} else {
			viper.AddConfigPath(defaultConfigPath)
		}

		c.parseConfigFile(flags)
	})
	c.installFlags(flags)
}

func (c *Config) parseConfigFile(flags *pflag.FlagSet) {
	viper.BindPFlags(flags)

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
