package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	router := gin.New()
	AddRoute(router)

	address := viper.GetString("api.address")

	// run
	go func() {
		if err := router.Run(address); err != nil {
			timber.Error("API:", err)
		}
	}()

	// notify
	if viper.GetBool("gin.release") {
		security := "HTTP"
		timber.Info(fmt.Sprintf("listening on %s (%s)", address, security))
	}

	// wait for ^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
