package finance

type InvestDataItem struct {
	Ccy    string `json:"ccy"`    //投资币种
	Bal    string `json:"bal"`    //可投数列
	MinAmt string `json:"minAmt"` //最小申购量
	MaxAmt string `json:"MaxAmt"` //最大申购量
}
