package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/designsbysm/mockserver/api/utils"
	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Work module automation for manually testing completion emails.
//
// @query {number}  [code]           override normal response to similate an error
// @query {boolean} [success=true]   success or failure in the response body
// @query {number}  [delay=2000]     ms to wait before responding
// @body  {string}  [eventId]
func Automation(c *gin.Context) {
	code := utils.QueryInt(c, "code", -1)
	if code > 0 {
		c.String(code, http.StatusText(code))
		return
	}
	success := utils.QueryBool(c, "success", true)
	utils.Delay(c, 2000)

	results := Response{
		Success: success,
	}
	if !success {
		results.Error = &ErrorInfo{
			Code:    http.StatusInternalServerError,
			Type:    "Faux Failure",
			Message: "faux automation failed",
		}
	}
	c.JSON(http.StatusOK, results)

	var body struct {
		EventID string `json:"eventId"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		timber.Error(err)
		return
	} else if body.EventID == "" {
		timber.Warning("no eventId in dispatch body, skipping callback")
		return
	}

	go complete(body.EventID, results)
}

func complete(eventID string, results Response) {
	payload, err := json.Marshal(gin.H{"eventId": eventID, "results": results})
	if err != nil {
		timber.Error(err)
		return
	}

	url := fmt.Sprintf("%s/v1/%s/tasks/automationComplete", viper.GetString("automation.url"), viper.GetString("automation.product"))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		timber.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",
		fmt.Sprintf("Bearer %s", viper.GetString("automation.secret")))

	// bound call in case it hangs
	var client = &http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		timber.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		timber.Error(err)
		return
	}

	message := fmt.Sprintf("%s: %s %s", eventID, res.Status, body)
	if res.StatusCode >= http.StatusBadRequest {
		timber.Error(message)
	} else {
		timber.Info(message)
	}
}
