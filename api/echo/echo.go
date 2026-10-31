// Echoes back the provided status and json.
//
// @query {number} [code=200] status code to respond with
// @query {string} [json] JSON object to respond with; omitted sends the status text
// @query {number} [delay=0] ms to wait before responding
package echo

import (
	"encoding/json"
	"net/http"
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
	strJson := c.Query("json")
	if len(strJson) > 0 {
		if err := json.Unmarshal([]byte(strJson), &data); err != nil {
			timber.Error(err)
		}
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
		c.String(code, http.StatusText(code))
	}
}
