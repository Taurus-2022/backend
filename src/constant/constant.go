package constant

const (
	Stage      string = "stage"
	LocalStage string = ""
)

const (
	// NONE 未中奖
	NONE = 0 + iota
	// MEITUAN 美团
	MEITUAN
	// TENCENT 腾讯
	TENCENT
	// DIDI 青桔单车
	DIDI

	ErrorCodeOK           = 0
	ErrorHttpParamInvalid = 1000 + iota
	ErrorDbInnerError
	ErrorCreateSignatureFailed
	ErrorConsumeAwardFailed
	ErrorCreateLotteryFailed
	ErrorLotteryNoChance

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
		ErrorHttpParamInvalid:      "Http param invalid",
		ErrorDbInnerError:          "Database inner error",
		ErrorCreateSignatureFailed: "Create new signature failed, maybe it has been created?",
		ErrorConsumeAwardFailed:    "Consume award failed, please contact administrator.",
		ErrorCreateLotteryFailed:   "Create lottery failed, please try again later.",
		ErrorLotteryNoChance:       "You have no chance to win lottery.",
	}
)
