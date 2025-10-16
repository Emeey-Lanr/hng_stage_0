package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"hng_stage1/utilis"
)

func FetchCatFact (c *gin.Context, httpClient *http.Client) {

	// context cancels the request if user leaves or10 seconds passes
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10 * time.Second)
  
	defer cancel()

	// request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://catfact.ninja/fact", nil)
	
	httpClient.Do(req) //mekes the request

	defer req.Body.Close()

	// if timeout related
	if errors.Is(err, context.DeadlineExceeded){
	 utilis.Error(c, http.StatusGatewayTimeout, "Request timeout")
	 return
	}

	// if user cancels request
	if errors.Is(err, context.Canceled){
		 utilis.Error(c, http.StatusRequestTimeout, "Request Cancelled")
	 return
	}

	




}