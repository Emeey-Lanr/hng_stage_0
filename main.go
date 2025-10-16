package main

import (
	"hng_stage1/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
 

func main () {
	r := gin.Default()

	 
	if  err := godotenv.Load(); err != nil{
	  log.Fatal("Error Loading .env", err.Error())    
	}
	// profile route
	routes.ProfileRoute(r)
	
	r.Run(":8080")
}