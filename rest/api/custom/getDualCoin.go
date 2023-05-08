package custom

import "go-okx/rest/api"

func NewGetDualCoin(param *GetDualCoinParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/v2/asset/earn/v2/dcd-products",
		Method: api.MethodGet,
		Param:  param,
	}, &GetDualCoinResponse{}
}

type GetDualCoinParam struct {
	CurrencyId    string `url:"currencyId,omitempty"`    // 0:BTC, 2:ETH
	AltCurrencyId string `url:"altCurrencyId,omitempty"` // 7:USDT
	DcdOptionType string `url:"dcdOptionType,omitempty"` // # PUT/CALL, CALL表示高卖; PUT表示低买
}

type GetDualCoinResponse struct {
	api.Response2
	Data []DualCoinData `json:"data"`
}

type DualCoinData struct {
	AbsYield              string `json:"absYield,omitempty"`
	AbsYieldPercentage    string `json:"absYieldPercentage,omitempty"`
	AlternativeCurrency   string `json:"alternativeCurrency,omitempty"`
	AnnualYield           string `json:"annualYield,omitempty"`
	AnnualYieldPercentage string `json:"annualYieldPercentage,omitempty"`
	Currency              string `json:"currency,omitempty"`
	CurrencyId            int64  `json:"currencyId,omitempty"`
	DcdTerm               string `json:"dcdTerm,omitempty"`
	DcdType               string `json:"dcdType,omitempty"`
	ExpiryTime            int64  `json:"expiryTime,omitempty"`
	Id                    int64  `json:"id,omitempty"`
	InterestAccrualTime   int64  `json:"interestAccrualTime,omitempty"`
	LaunchTime            int64  `json:"launchTime,omitempty"`
	MatchCurrency         string `json:"matchCurrency,omitempty"`
	MatchCurrencyId       int64  `json:"matchCurrencyId,omitempty"`
	ProductType           string `json:"productType,omitempty"`
	RedeemEndTime         int64  `json:"redeemEndTime,omitempty"`
	RedeemStartTime       int64  `json:"redeemStartTime,omitempty"`
	SettlementTime        int64  `json:"settlementTime,omitempty"`
	ShortSymbol           string `json:"shortSymbol,omitempty"`
	Status                int64  `json:"status,omitempty"`
	StopTradingTime       int64  `json:"stopTradingTime,omitempty"`
	Strike                string `json:"strike,omitempty"`
	Symbol                string `json:"symbol,omitempty"`
	Type                  string `json:"type,omitempty"`
	UnderlyingIndex       string `json:"underlyingIndex,omitempty"`
}
