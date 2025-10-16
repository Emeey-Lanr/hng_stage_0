package utilis

import (
	"hng_stage1/models"

	"github.com/gin-gonic/gin"

)

func Error(c *gin.Context, method int, status string) {
  c.JSON(method, models.MeResponse{Status: status})
}

func Success (c *gin.Context, method int, status string, user models.UserData, time string, fact string){
  c.JSON(method, models.MeResponse{Status: status, User: user, Timestamp: time, Fact:  fact})

}