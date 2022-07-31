package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taurus-backend/api/request"
	"taurus-backend/api/response"
	"taurus-backend/constant"
	"taurus-backend/db"
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

func GetUserIsSigned(c *gin.Context) {
	r := &request.GetUserIsSignedRequest{}
	err := c.ShouldBindQuery(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}
	total, err := db.GetSignatureCountByPhone(r.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorDbInnerError, err.Error()))
		return
	}
	isSigned := total > 0
	c.JSON(200, &response.GetUserIsSignedResponse{IsSigned: isSigned})
}
