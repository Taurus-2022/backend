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
	err = db.CreateSignature(r.Street)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorCreateSignatureFailed, err.Error()))
		return
	}
	c.String(http.StatusOK, "success")
}

func GetSignatureCount(c *gin.Context) {
	r := &request.GetSignatureCountRequest{}
	err := c.BindQuery(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.GetErrorResponse(constant.ErrorHttpParamInvalid, err.Error()))
		return
	}

	var total int
	if "" == r.Street {
		total, err = db.GetAllSignatureCount()
	} else {
		total, err = db.GetSignatureCountByStreet(r.Street)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.GetErrorResponse(constant.ErrorDbInnerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, &response.GetSignatureCountResponse{Count: total})
}
