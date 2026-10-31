package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/designsbysm/gibson"
	"github.com/designsbysm/mockserver/api/handlers"
	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	// setup
	router := gin.New()

	router.Use(
		gin.Recovery(),
		gibson.Logger(),
		gibson.Error(),
	)

	router.Any("/automation", handlers.Automation)
	router.Any("/echo", handlers.Echo)

	address := viper.GetString("api.address")

	// run
	go func() {
		if err := router.Run(address); err != nil {
			timber.Error("API:", err)
		}
	}()

	if !viper.GetBool("gin.debug") {
		timber.Info(fmt.Sprintf("listening on %s", address))
	}

	// wait for ^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
