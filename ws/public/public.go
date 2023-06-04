package public

import (
	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/gorilla/websocket"
)

type Public struct {
	C *ws.Client
}

func NewPublic(simulated bool) *Public {
	public := &Public{
		C: ws.DefaultClientPublic,
	}
	if simulated {
		public.C = ws.DefaultClientPrivateSimulated
	}
	return public
}

// subscribe
func (p *Public) Subscribe(args interface{}, handler ws.Handler, handlerError ws.HandlerError) (*websocket.Conn, error) {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}
