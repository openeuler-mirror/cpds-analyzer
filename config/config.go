package config

import (
	"fmt"
	"net"
	"os"
	"strconv"

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
	CertFile         string
	KeyFile          string
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
		CertFile:         viper.GetString("cert-file"),
		KeyFile:          viper.GetString("key-file"),
	}
}

func (c *Config) CheckConfig() error {
	if net.ParseIP(c.DatabaseAddress) == nil {
		return fmt.Errorf("invalid flag: db-address: %s", c.DatabaseAddress)
	}

	if p, err := strconv.Atoi(c.DatabasePort); err != nil {
		return fmt.Errorf("invalid flag: db-port: %s, %w", c.DatabasePort, err)
	} else if p < 0 || p > 65535 {
		return fmt.Errorf("invalid port number range: %s, should be 0 - 65535", c.DatabasePort)
	}

	if net.ParseIP(c.BindAddress) == nil {
		return fmt.Errorf("invalid flag: bind-address: %s", c.BindAddress)
	}

	if p, err := strconv.Atoi(c.Port); err != nil {
		return err
	} else if p < 0 || p > 65535 {
		return fmt.Errorf("invalid port number range: %s, should be 0 - 65535", c.Port)
	}

	if _, err := os.Stat(c.CertFile); err != nil {
		return fmt.Errorf("invalid flag: cert-file: %s, %w", c.CertFile, err)
	}

	if _, err := os.Stat(c.KeyFile); err != nil {
		return fmt.Errorf("invalid flag: key-file: %s, %w", c.KeyFile, err)
	}

	return nil
}
