package config

import (
	"testing"

	"github.com/spf13/pflag"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestCheckConfig(t *testing.T) {
	conf := New()
	flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
	conf.installFlags(flags)

	flags.Parse([]string{
		"--cert-file=/foo/ca",
		"--key-file=/foo/ca",
	})
	assert.Error(t, conf.CheckConfig(), "invalid flag: cert-file: /foo/ca, stat /foo/ca: no such file or directory")

	flags.Parse([]string{
		"--cert-file=/foo/ca",
		"--key-file=/foo/ca",
		"--db-address=1234",
	})
	assert.Error(t, conf.CheckConfig(), "invalid flag: db-address: 1234")
}

func TestLoadConfig(t *testing.T) {
	conf := New()
	flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)

	flags.Parse([]string{
		"--config=config/json/cpds-analyzer.json",
	})
	conf.LoadConfig(flags)

	assert.Check(t, is.Equal(conf.DatabaseUser, "root"))
	assert.Check(t, is.Equal(conf.Port, "19081"))
	assert.Equal(t, conf.Debug, false)
	assert.Equal(t, conf.LogLevel, "info")
	assert.Equal(t, conf.DatabaseAddress, "localhost")
}
