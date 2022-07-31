package constant

const (
	Stage      string = "stage"
	LocalStage string = ""
)

const (
	// MEITUAN 美团
	MEITUAN = 0 + iota
	// TENCENT 腾讯
	TENCENT
	// DIDI 青桔单车
	DIDI

	ErrorCodeOK      = 0
	ErrorSignInvalid = 1000 + iota
	ErrorTokenInvalid
	ErrorAuthFailed
	ErrorHttpInnerError
	ErrorHttpParamInvalid
	ErrorHttpResourceExists
	ErrorHttpResourceNotFound
	ErrorDbInnerError

	ErrorCreateSignatureFailed

	HTTPMethodGet    string = "GET"
	HTTPMethodPost   string = "POST"
	HTTPMethodPut    string = "PUT"
	HTTPMethodPatch  string = "PATCH"
	HTTPMethodDelete string = "DELETE"
	HTTPMethodHead   string = "HEAD"
)

var (
	RespCodeErrorString = map[int]string{
		ErrorCodeOK:                "Success",
		ErrorTokenInvalid:          "Token invalid.",
		ErrorHttpInnerError:        "Http inner error",
		ErrorHttpParamInvalid:      "Http param invalid",
		ErrorSignInvalid:           "Sign invalid.",
		ErrorHttpResourceExists:    "Http resource already exists.",
		ErrorHttpResourceNotFound:  "Http resource not found.",
		ErrorAuthFailed:            "Authentication failed",
		ErrorDbInnerError:          "Database inner error",
		ErrorCreateSignatureFailed: "Create new signature failed, maybe it has been created?",
	}
)
