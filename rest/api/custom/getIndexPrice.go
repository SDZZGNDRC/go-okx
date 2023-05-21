package custom

import "github.com/SDZZGNDRC/go-okx/rest/api"

func NewGetIndexPrice(param *GetIndexPriceParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/v2/asset/earn/dcd-market-price",
		Method: api.MethodGet,
		Param:  param,
	}, &GetIndexPriceResponse{}
}

type GetIndexPriceParam struct {
	CurrencyId string `url:"currencyId,omitempty"` // 0:BTC, 2:ETH
	IsoCode    string `url:"isoCode,omitempty"`    // USD
}

type GetIndexPriceResponse struct {
	api.Response2
	Data IndexPriceData `json:"data"`
}

type IndexPriceData struct {
	CurrencyName string `json:"currencyName,omitempty"`
	IndexPrice   string `json:"indexPrice,omitempty"`
	IsoCode      string `json:"isoCode,omitempty"`
}
