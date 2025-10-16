package routes
import
( "github.com/gin-gonic/gin"
   "hng_stage1/handlers"
)


func ProfileRoute (r *gin.Engine) {
 r.GET("/me", handlers.GetProfileMe)

}