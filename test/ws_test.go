package test

import (
	"log"
	"testing"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/SDZZGNDRC/go-okx/ws/public"
)

func TestBooks(t *testing.T) {
	args := &ws.Args{
		Channel: "books",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventBooks) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeBooks(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {} // Wait forever
}
