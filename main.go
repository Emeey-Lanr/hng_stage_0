package  main 

import (
	"github.com/gin-gonic/gin"
	"hng_stage1/routes"
	
)
 

func main () {
	r := gin.Default()
	 
	// profile route
	routes.ProfileRoute(r)
	
	r.Run(":8080")
}