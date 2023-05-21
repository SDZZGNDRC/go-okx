package account

import "github.com/SDZZGNDRC/go-okx/rest/api"

const GetPositionsLimitNumPerSec = 5
const GetPositionsLimitRule = "UserID"

func NewGetPositions(param *GetPositionsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/positions",
		Method: api.MethodGet,
		Param:  param,
	}, &GetPositionsResponse{}
}

type GetPositionsParam struct {
	InstId   string `url:"instId,omitempty"`
	InstType string `url:"instType,omitempty"`
	PosId    string `url:"posId,omitempty"`
}

// 查看持仓信息
// 获取该账户下拥有实际持仓的信息。账户为单向持仓模式会显示净持仓（net），
// 账户为双向持仓模式下会分别返回多头（long）或空头（short）的仓位。
// 按照仓位创建时间倒序排列。
type GetPositionsResponse struct {
	api.Response
	Data []struct {
		InstType    string `json:"instType"`
		MgnMode     string `json:"mgnMode"`
		PosId       string `json:"posId"`
		PosSide     string `json:"posSide"`
		Pos         string `json:"pos"`
		BaseBal     string `json:"baseBal"`
		QuoteBal    string `json:"quoteBal"`
		PosCcy      string `json:"posCcy"`
		AvailPos    string `json:"availPos"`
		AvgPx       string `json:"avgPx"`
		Upl         string `json:"upl"`
		UplRatio    string `json:"uplRatio"`
		InstId      string `json:"instId"`
		Lever       string `json:"lever"`
		LiqPx       string `json:"liqPx"`
		MarkPx      string `json:"markPx"`
		Imr         string `json:"imr"`
		Margin      string `json:"margin"`
		MgnRatio    string `json:"mgnRatio"`
		Mmr         string `json:"mmr"`
		Liab        string `json:"liab"`
		LiabCcy     string `json:"liabCcy"`
		Interest    string `json:"interest"`
		TradeId     string `json:"tradeId"`
		OptVal      string `json:"optVal"`
		NotionalUsd string `json:"notionalUsd"`
		Adl         string `json:"adl"`
		Ccy         string `json:"ccy"`
		Last        string `json:"last"`
		UsdPx       string `json:"usdPx"`
		DeltaBS     string `json:"deltaBS"`
		DeltaPA     string `json:"deltaPA"`
		GammaBS     string `json:"gammaBS"`
		GammaPA     string `json:"gammaPA"`
		ThetaBS     string `json:"thetaBS"`
		ThetaPA     string `json:"thetaPA"`
		VegaBS      string `json:"vegaBS"`
		VegaPA      string `json:"vegaPA"`
		CTime       string `json:"cTime"`
		UTime       string `json:"uTime"`
	} `json:"data"`
}
