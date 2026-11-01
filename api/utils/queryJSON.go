package utils

import (
	"encoding/json"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
)

func QueryJSON(c *gin.Context, key string) map[string]any {
	raw := c.Query(key)
	if raw == "" {
		return nil
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		timber.Error(err)
		return nil
	}

	return data
}
