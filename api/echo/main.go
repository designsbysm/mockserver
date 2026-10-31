package echo

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.Any("/echo", handle)
}

func handle(c *gin.Context) {
	code, err := strconv.Atoi(c.DefaultQuery("code", "200"))
	if err != nil {
		timber.Error(err)
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(c.Query("json")), &data); err != nil {
		timber.Error(err)
	}

	delay, err := strconv.Atoi(c.DefaultQuery("delay", "0"))
	if err != nil {
		timber.Error(err)
	} else if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	if len(data) > 0 {
		c.JSON(code, gin.H(data))
	} else {
		c.Status(code)
	}
}
