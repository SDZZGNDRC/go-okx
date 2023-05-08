package finance

type EarningDataItem struct {
	Ccy         string `json:"ccy"`         //收益币种
	EarningType string `json:"earningType"` //收益类型. 0: 预估收益; 1: 累计发放收益
}
