package public

import (
	"encoding/json"
	"time"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

// 深度频道

type HandlerBooks func(interface{})

type EventBooks struct {
	Arg     ws.Args `json:"arg"`
	Data    []Book  `json:"data"`
	Action  string  `json:"action"`
	LocalTs int64   `json:"localTs"`
}

type Book struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Ts        int64      `json:"ts,string"`
	Checksum  int32      `json:"checksum"`
	PrevSeqId int64      `json:"prevSeqId"`
	SeqId     int64      `json:"seqId"`
}

// default subscribe
func SubscribeBooks(args *ws.Args, handler HandlerFunc, handlerError ws.HandlerError, simulated bool) (*websocket.Conn, error) {
	h := func(message []byte) { // convert raw data into EventBooks
		var event EventBooks
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		event.LocalTs = time.Now().UnixMilli()
		handler(event)
	}

	return NewPublic(simulated).Subscribe(args, h, handlerError)
}
