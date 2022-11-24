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

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error, Using defalut config file path
			// Using defalut config file path
			logrus.Warnf("Config file '%s' not found, Using defalut config file path: '%s'", c.ConfigPath, defaultConfigPath)
		} else {

			// Config file was found but another error was produced
			logrus.Errorf("fatal error config file: %w", err)
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	} else {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
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
