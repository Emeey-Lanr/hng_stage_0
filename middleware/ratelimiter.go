package middleware

import (
	"hng_stage1/utilis"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)


func HandleRateLimter () gin.HandlerFunc {

	limter := rate.NewLimiter(1, 3)


	return func (c *gin.Context){
		if !limter.Allow() {
			utilis.Error(c,  http.StatusTooManyRequests, "To Many Request")
			c.Abort()
			return
		}

		c.Next()
	}
}