package utils

import (
	"strconv"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
)

func QueryInt(c *gin.Context, key string, fallback int) int {
	raw := c.Query(key)
	if raw == "" {
		return fallback
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		timber.Error(err)
		return fallback
	}

	return value
}
