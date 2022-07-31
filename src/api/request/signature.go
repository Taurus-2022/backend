package request

type GetSignatureCountRequest struct {
	Street string `form:"street"`
}

type CreateSignatureCountRequest struct {
	Phone  string `json:"phone" binding:"required"`
	Street string `json:"street" binding:"required"`
}
