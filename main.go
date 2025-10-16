package main

import (
	"fmt"
	"hng_stage1/middleware"
	"hng_stage1/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
 

func main () {
	r := gin.Default()

	 
	if  err := godotenv.Load(); err != nil{
	  log.Println("Error Loading .env", err.Error())    
	}
	// profile route
	routes.ProfileRoute(r)
	
	// A global rate limiter
	r.Use(middleware.HandleRateLimter())
	
	port := os.Getenv("PORT")

	r.Run(fmt.Sprintf(":%s", port))
}