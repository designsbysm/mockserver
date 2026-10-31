package api

import (
	"github.com/designsbysm/gibson"
	"github.com/designsbysm/mockserver/api/automation"
	"github.com/designsbysm/mockserver/api/echo"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		gibson.Logger(),
		gibson.Error(),
	)

	automation.AddRoute(r)
	echo.AddRoute(r)
}
