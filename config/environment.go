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
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./config.yaml not found")
		} else {
			return err
		}
	}

	// setup stuff
	if viper.GetBool("gin.release") {
		gin.SetMode(gin.ReleaseMode)
	}

	return nil
}
