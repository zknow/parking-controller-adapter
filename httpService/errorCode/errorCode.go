package errorCode

const (
	SetStatusError       = "設定狀態錯誤"
	LostStatusFile       = "取得狀態錯誤,狀態檔案遺失"
	ParseParameterError  = "參數解析錯誤"
	ParameterError       = "參數錯誤"
	CardNumberParseError = "卡號解析錯誤"
	NoneCardNumberForHex = "卡號不合法,非16進位卡號"
)
