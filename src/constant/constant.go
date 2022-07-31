package constant

const (
	ErrorCodeOK      = 0
	ErrorSignInvalid = 1000 + iota
	ErrorTokenInvalid
	ErrorAuthFailed
	ErrorHttpInnerError
	ErrorHttpParamInvalid
	ErrorHttpResourceExists
	ErrorHttpResourceNotFound

	HTTPMethodGet    string = "GET"
	HTTPMethodPost   string = "POST"
	HTTPMethodPut    string = "PUT"
	HTTPMethodPatch  string = "PATCH"
	HTTPMethodDelete string = "DELETE"
	HTTPMethodHead   string = "HEAD"
)

var (
	RespCodeErrorString = map[int]string{
		ErrorCodeOK:               "Success",
		ErrorTokenInvalid:         "Token invalid.",
		ErrorHttpInnerError:       "Http inner error",
		ErrorHttpParamInvalid:     "Http param invalid",
		ErrorSignInvalid:          "Sign invalid.",
		ErrorHttpResourceExists:   "Http resource already exists.",
		ErrorHttpResourceNotFound: "Http resource not found.",
		ErrorAuthFailed:           "Authentication failed",
	}
)
