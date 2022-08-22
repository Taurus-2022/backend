package request

type GetSignatureCountRequest struct {
	Street string `form:"street"`
}

type CreateSignatureCountRequest struct {
	Street string `json:"street" binding:"required"`
}

type GetUserIsSignedRequest struct {
	Phone string `form:"phone" binding:"required"`
}
