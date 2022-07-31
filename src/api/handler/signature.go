package handler

import (
	"github.com/gin-gonic/gin"
	"taurus-backend/api/request"
	"taurus-backend/api/response"
)

func CreateSignature(c *gin.Context) {
	c.String(200, "success")
}

func GetSignatureCount(c *gin.Context) {
	r := &request.GetSignatureCountRequest{}
	err := c.BindQuery(r)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, &response.Count{Count: 100})
}
