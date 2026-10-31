package config

import (
	"os"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func Loggers() error {
	level := timber.StringToLevel(viper.GetString("timber.level"))
	if level < 1 {
		return nil
	}

	if err := timber.New(
		os.Stdout,
		level,
		viper.GetString("timber.timestamp"),
		timber.StringToFlags(viper.GetString("timber.flags")),
	); err != nil {
		return err
	}

	return nil
}
