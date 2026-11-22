package config

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Environment() error {
	// load config.yaml
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := errors.AsType[viper.ConfigFileNotFoundError](err); ok {
			return errors.New("./config.yaml not found")
		}

		return err
	}

	// setup stuff
	if !viper.GetBool("gin.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}
