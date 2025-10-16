package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"hng_stage1/models"
	"hng_stage1/utilis"

	"github.com/gin-gonic/gin"
)

func FetchCatFact (c *gin.Context) (string, error){
 var catFact models.CatFact
 var url = os.Getenv("CATURL")
	// context cancels the request if user leaves or10 seconds passes
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10 * time.Second)
  
	defer cancel()

	// request
	request , err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
 
	 if err != nil {
	return "", err
 }
	 
//  Header with the Request
request.Header.Set("Accept", "application/json")
	 //makes the request

   	 response, err := utilis.HttpClient.Do(request)
	 
	 if err != nil {
		return "", err
	 }
	
	defer response.Body.Close()

//   decode body and bind to struct
  if err := json.NewDecoder(response.Body).Decode(&catFact); err != nil {
	log.Printf("decode error: %v\n", err)
	return "", err
  }

  return catFact.Fact, nil
	
}