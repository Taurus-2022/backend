package handler

import (
	"github.com/gin-gonic/gin"
	"taurus-backend/api/response"
)

func CreateSignature(c *gin.Context) {
	c.String(200, "success")
}

func GetSignatureCount(c *gin.Context) {
	c.JSON(200, &response.Count{Count: 100})
}
