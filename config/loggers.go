package config

import (
	"os"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func Loggers() error {
	if err := cli(); err != nil {
		return err
	}

	return nil
}

func cli() error {
	level := timber.StringToLevel(viper.GetString("timber.cli.level"))
	if level < 1 {
		return nil
	}

	return timber.New(
		os.Stdout,
		level,
		viper.GetString("timber.cli.timestamp"),
		timber.StringToFlags(viper.GetString("timber.cli.flags")),
	)
}
