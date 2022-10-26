package middleware

import (
	"github.com/Abashinos/otus-msa-hw/app/cmd/server/views"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func NewDelayerMiddleware() gin.HandlerFunc {
	const maxRequestDelayMs = 1000
	return func(c *gin.Context) {
		time.Sleep(time.Duration(rand.Intn(maxRequestDelayMs)) * time.Millisecond)
		c.Next()
	}
}

var errorResponse = struct {
	Message string `json:"message"`
}{Message: "oops"}

func NewFlakyMiddleware() gin.HandlerFunc {
	const errorChancePercent = 25
	return func(c *gin.Context) {
		if c.Request.URL.Path == views.RouteHealth {
			c.Next()
			return
		}

		if rand.Intn(99) < errorChancePercent {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		}
		c.Next()
	}
}
