package handlers

import (
	"net/http"

	"github.com/designsbysm/mockserver/api/utils"
	"github.com/gin-gonic/gin"
)

// Wraps the response in a success/error envelope, always answering 200.
//
// @query {boolean} [success=true] success or failure in the envelope
// @query {string}  [json]         JSON object to return as results; ignored when success is false
// @query {number}  [code=500]     code and status text for the error; ignored when success is true
// @query {number}  [delay=0]      ms to wait before responding
func Envelope(c *gin.Context) {
	code := utils.QueryInt(c, "code", http.StatusInternalServerError)
	success := utils.QueryBool(c, "success", true)
	data := utils.QueryJSON(c, "json")
	utils.Delay(c, 0)

	results := Response{
		Success: success,
	}
	if success && len(data) > 0 {
		results.Results = data
	} else if !success {
		message := http.StatusText(code)
		if len(message) == 0 {
			message = "Unknown Error"
		}

		results.Error = &ErrorInfo{
			Code:    code,
			Type:    message,
			Message: message,
		}
	}

	c.JSON(http.StatusOK, results)
}
