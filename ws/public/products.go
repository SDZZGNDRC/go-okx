package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
)

// 产品频道
// 当有产品状态变化时（如期货交割、期权行权、新合约/币对上线、人工暂停/恢复交易等），推送产品的增量数据。

type HandlerProducts func(EventProducts)

type EventProducts struct {
	Arg  ws.Args   `json:"arg"`
	Data []Product `json:"data"`
}

type Product struct {
	InstType     string `json:"instType,omitempty"`
	InstID       string `json:"instId,omitempty"`
	InstFamily   string `json:"instFamily,omitempty"`
	Uly          string `json:"uly,omitempty"`
	Category     string `json:"category,omitempty"`
	BaseCcy      string `json:"baseCcy,omitempty"`
	QuoteCcy     string `json:"quoteCcy,omitempty"`
	SettleCcy    string `json:"settleCcy,omitempty"`
	CtVal        string `json:"ctVal,omitempty"`
	CtMult       string `json:"ctMult,omitempty"`
	CtValCcy     string `json:"ctValCcy,omitempty"`
	OptType      string `json:"optType,omitempty"`
	Stk          string `json:"stk,omitempty"`
	ListTime     string `json:"listTime,omitempty"`
	ExpTime      string `json:"expTime,omitempty"`
	TickSz       string `json:"tickSz,omitempty"`
	LotSz        string `json:"lotSz,omitempty"`
	MinSz        string `json:"minSz,omitempty"`
	CtType       string `json:"ctType,omitempty"`
	Alias        string `json:"alias,omitempty"`
	State        string `json:"state,omitempty"`
	MaxLmtSz     string `json:"maxLmtSz,omitempty"`
	MaxMktSz     string `json:"maxMktSz,omitempty"`
	MaxTwapSz    string `json:"maxTwapSz,omitempty"`
	MaxIcebergSz string `json:"maxIcebergSz,omitempty"`
	MaxTriggerSz string `json:"maxTriggerSz,omitempty"`
	MaxStopSz    string `json:"maxStopSz,omitempty"`
}

// default subscribe
func SubscribeProducts(args *ws.Args, handler HandlerProducts, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventProducts
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
