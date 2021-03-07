package main

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Domain string `yaml:"domain"`

	TLS struct {
		Generate bool   `yaml:"generate"`
		Cert     string `yaml:"cert"`
		Key      string `yaml:"key"`
	} `yaml:"tls"`

	Submission struct {
		Port int `yaml:"port"`
	} `yaml:"submission"`
}

var conf Config

func Configure(*cobra.Command, []string) {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/talaria")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.talaria")

	// Defaults
	viper.SetDefault("domain", "locahost")
	viper.SetDefault("submission.port", 465)

	// Environment variables
	viper.SetEnvPrefix("TALARIA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read Config
	viper.ReadInConfig()

	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
}
