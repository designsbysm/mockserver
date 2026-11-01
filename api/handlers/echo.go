package handlers

import (
	"net/http"

	"github.com/designsbysm/mockserver/api/utils"
	"github.com/gin-gonic/gin"
)

// Echoes back the provided status and json.
//
// @query {number} [code=200]  status code to respond with
// @query {string} [json] JSON object to respond with; omitted sends the status text
// @query {number} [delay=0]   ms to wait before responding
func Echo(c *gin.Context) {
	code := utils.QueryInt(c, "code", http.StatusTeapot)
	data := utils.QueryJSON(c, "json")
	utils.Delay(c, 0)

	if len(data) > 0 {
		c.JSON(code, gin.H(data))
	} else {
		c.String(code, http.StatusText(code))
	}
}
