package public

import "go-okx/rest/api"

func NewGetInstruments(param *GetInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/public/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetInstrumentsResponse{}
}

type GetInstrumentsParam struct {
	InstType   string `url:"instType"`
	Uly        string `url:"uly"`        // 标的指数，仅适用于交割/永续/期权，期权必填
	InstFamily string `url:"instFamily"` // 交易品种, 如 BTC-USD，仅适用于交割/永续/期权
	InstId     string `url:"instId"`     // 产品ID

}

type GetInstrumentsResponse struct {
	api.Response
	Data []struct {
		InstType     string `json:"instType"`
		InstId       string `json:"instId"`       // 产品ID
		Uly          string `json:"uly"`          // 标的指数，仅适用于交割/永续/期权，期权必填
		Category     string `json:"category"`     // 已废弃
		BaseCcy      string `json:"baseCcy"`      // 交易货币币种，如 BTC-USDT 中的 BTC ，仅适用于币币/币币杠杆
		QuoteCcy     string `json:"quoteCcy"`     // 计价货币币种，如 BTC-USDT 中的USDT ，仅适用于币币/币币杠杆
		SettleCcy    string `json:"settleCcy"`    // 盈亏结算和保证金币种，如 BTC 仅适用于交割/永续/期权
		CtVal        string `json:"ctVal"`        // 合约面值，仅适用于交割/永续/期权
		CtMult       string `json:"ctMult"`       // 合约乘数，仅适用于交割/永续/期权
		CtValCcy     string `json:"ctValCcy"`     // 合约面值计价币种，仅适用于交割/永续/期权
		OptType      string `json:"optType"`      // 期权类型，C或P 仅适用于期权
		Stk          string `json:"stk"`          // 行权价格，仅适用于期权
		ListTime     string `json:"listTime"`     // 上线日期; Unix时间戳的毫秒数格式，如 1597026383085
		ExpTime      string `json:"expTime"`      // 交割/行权日期，仅适用于交割 和 期权; Unix时间戳的毫秒数格式
		Lever        string `json:"lever"`        // 该instId支持的最大杠杆倍数，不适用于币币、期权
		TickSz       string `json:"tickSz"`       // 下单价格精度，如 0.0001
		LotSz        string `json:"lotSz"`        // 下单数量精度，如 BTC-USDT-SWAP：1
		MinSz        string `json:"minSz"`        // 最小下单数量
		CtType       string `json:"ctType"`       // linear：正向合约; inverse：反向合约; 仅适用于交割/永续
		Alias        string `json:"alias"`        // 合约日期别名(本周/次周/...); 仅适用于交割
		State        string `json:"state"`        // 产品状态; live：交易中; suspend：暂停中; preopen：预上线
		MaxLmtSz     string `json:"maxLmtSz"`     // 合约或现货限价单的单笔最大委托数量
		MaxMktSz     string `json:"maxMktSz"`     // 合约或现货市价单的单笔最大委托数量
		MaxTwapSz    string `json:"maxTwapSz"`    // 合约或现货时间加权单的单笔最大委托数量
		MaxIcebergSz string `json:"maxIcebergSz"` // 合约或现货冰山委托的单笔最大委托数量
		MaxTriggerSz string `json:"maxTriggerSz"` // 合约或现货计划委托委托的单笔最大委托数量
		MaxStopSz    string `json:"maxStopSz"`    // 合约或现货止盈止损委托的单笔最大委托数量
	} `json:"data"`
}
