package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Response struct {
	Success bool        `json:"success"`
	Results interface{} `json:"results,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Work module automation for manually testing completion emails.
//
// @query {string} [result=success] results success|failure|error
// @query {number} [delay=2000] ms to wait before responding
// @body  {string} [eventId]
func Automation(c *gin.Context) {
	result := c.DefaultQuery("result", "success")
	if result == "error" {
		c.Status(http.StatusInternalServerError)
		return
	}

	delay, err := strconv.Atoi(c.DefaultQuery("delay", "2000"))
	if err != nil {
		timber.Error(err)
	} else if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	var body struct {
		EventID string `json:"eventId"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		timber.Error(err)
	} else if body.EventID == "" {
		timber.Warning("no eventId in dispatch body, skipping callback")
		return
	}

	results := Response{
		Success: true,
	}
	if result == "failure" {
		results = Response{
			Success: false,
			Error:   &ErrorInfo{Code: http.StatusInternalServerError, Type: "Faux Failure", Message: "faux automation failed"},
		}
	}
	c.JSON(http.StatusOK, results)

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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		timber.Error(err)
		return
	}
	defer res.Body.Close()
}
