package public

import (
	"encoding/json"

	"github.com/SDZZGNDRC/go-okx/ws"
)

// Status 频道
// 获取系统维护的状态，当系统维护状态改变时推送。首次订阅：”推送最新一条的变化数据“；后续每次有状态变化，推送变化的内容。

type HandlerStatus func(EventStatus)

type EventStatus struct {
	Arg  ws.Args  `json:"arg"`
	Data []Status `json:"data"`
}

type Status struct {
	Begin        string `json:"begin,omitempty"`
	End          string `json:"end,omitempty"`
	Href         string `json:"href,omitempty"`
	PreOpenBegin string `json:"preOpenBegin,omitempty"`
	ScheDesc     string `json:"scheDesc,omitempty"`
	ServiceType  string `json:"serviceType,omitempty"`
	State        string `json:"state,omitempty"`
	System       string `json:"system,omitempty"`
	MaintType    string `json:"maintType,omitempty"`
	Env          string `json:"env,omitempty"`
	Title        string `json:"title,omitempty"`
	Ts           string `json:"ts,omitempty"`
}

// default subscribe
func SubscribeStatus(args *ws.Args, handler HandlerStatus, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventStatus
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
