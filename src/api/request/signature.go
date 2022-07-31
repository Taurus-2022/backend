package request

type GetSignatureCountRequest struct {
	Street string `form:"street"`
}

type CreateSignatureCountRequest struct {
	Phone  string `json:"phone" binding:"required"`
	Street string `json:"street" binding:"required"`
}

type GetUserIsSignedRequest struct {
	Phone string `form:"phone" binding:"required"`
}
