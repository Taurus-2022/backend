package response

type GetSignatureCountResponse struct {
	Count int `json:"count"`
}
type GetUserIsSignedResponse struct {
	IsSigned bool `json:"isSigned"`
}
