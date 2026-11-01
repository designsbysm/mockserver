package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Sleeps for the ?delay= query value in ms, or fallback when it is absent.
func Delay(c *gin.Context, fallback int) {
	if ms := QueryInt(c, "delay", fallback); ms > 0 {
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}
