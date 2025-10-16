package handlers

import (
	"context"
	"errors"
	"hng_stage1/models"
	"hng_stage1/services"
	"hng_stage1/utilis"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProfileMe(c *gin.Context){
	fact, err := services.FetchCatFact(c)

	if err != nil {
		// if time out error
		if errors.Is(err, context.DeadlineExceeded){
			utilis.Error(c, http.StatusGatewayTimeout, "external service timed out")
		 return
		}

		// if user cancels request 
		if errors.Is(err, context.Canceled){
      
		  utilis.Error(c, http.StatusRequestTimeout, "request cancelled")
		    return
		}

		// any other error
		utilis.Error(c, http.StatusBadGateway, err.Error())
		return
	}

	var user  = models.UserData{Email: os.Getenv("EMAIL"), Name: os.Getenv("NAME"), Stack: "GO/GIN"}
	 
	utilis.Success(c, http.StatusOK, "success", user, time.Now().UTC().Format(time.RFC3339), fact)

}