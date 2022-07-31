package response

import "taurus-backend/constant"

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetErrorResponse(code int, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: constant.RespCodeErrorString[code],
		Data:    data,
	}
}
