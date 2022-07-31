package handler

import (
	"github.com/gin-gonic/gin"
	"taurus-backend/api/request"
	"taurus-backend/api/response"
	"taurus-backend/constant"
)

func CreateSignature(c *gin.Context) {
	r := &request.CreateSignatureCountRequest{}
	err := c.ShouldBindJSON(r)
	if err != nil {
		c.AbortWithStatusJSON(400, response.GetErrorResponse(constant.ErrorHttpParamInvalid, nil))
	}
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
