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
	// ErrorCreateSignatureFailed 签约失败
	ErrorCreateSignatureFailed
	// ErrorConsumeAwardFailed 减库存失败
	ErrorConsumeAwardFailed
	// ErrorCreateLotteryFailed 创建抽奖记录失败
	ErrorCreateLotteryFailed
	// ErrorTodayNoMoreLotteryChance 今天不能再抽奖
	ErrorTodayNoMoreLotteryChance
	// ErrorNoMoreAward 所有奖券被抽完
	ErrorNoMoreAward
	// ErrorWinLotteryFailed 抽奖状态异常
	ErrorWinLotteryFailed

	HTTPMethodGet    string = "GET"
	HTTPMethodPost   string = "POST"
	HTTPMethodPut    string = "PUT"
	HTTPMethodPatch  string = "PATCH"
	HTTPMethodDelete string = "DELETE"
	HTTPMethodHead   string = "HEAD"
)

var (
	RespCodeErrorString = map[int]string{
		ErrorCodeOK:                   "Success",
		ErrorHttpParamInvalid:         "Http param invalid",
		ErrorDbInnerError:             "Database inner error",
		ErrorCreateSignatureFailed:    "Create new signature failed, maybe it has been created?",
		ErrorConsumeAwardFailed:       "Consume award failed, please contact administrator.",
		ErrorCreateLotteryFailed:      "Create lottery failed, please try again later.",
		ErrorTodayNoMoreLotteryChance: "You have no chance to win lottery today.",
		ErrorNoMoreAward:              "All award has been used up.",
		ErrorWinLotteryFailed:         "Win lottery failed, please try again later.",
	}
)
