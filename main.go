package main

import (
	"github.com/designsbysm/mockserver/api"
	"github.com/designsbysm/mockserver/config"
)

func main() {
	if err := config.Environment(); err != nil {
		panic(err)
	}

	if err := config.Loggers(); err != nil {
		panic(err)
	}

	api.Serve()
}
