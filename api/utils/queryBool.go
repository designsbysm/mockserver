package utils

import (
	"strconv"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
)

func QueryBool(c *gin.Context, key string, fallback bool) bool {
	raw := c.Query(key)
	if raw == "" {
		return fallback
	}

	value, err := strconv.ParseBool(raw)
	if err != nil {
		timber.Error(err)
		return fallback
	}

	return value
}
