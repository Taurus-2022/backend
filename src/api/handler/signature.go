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
		c.AbortWithStatusJSON(400, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	c.String(200, "success")
}

func GetSignatureCount(c *gin.Context) {
	r := &request.GetSignatureCountRequest{}
	err := c.BindQuery(r)
	if err != nil {
		c.AbortWithStatusJSON(400, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	c.JSON(200, &response.GetSignatureCountResponse{Count: 100})
}
