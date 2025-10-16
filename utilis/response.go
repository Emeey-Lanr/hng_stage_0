package utilis

import (
"github.com/gin-gonic/gin"
"hng_stage1/models"

)

func Error(c *gin.Context, method int, status string) {
  c.JSON(method, models.MeResponse{Status: status})
}