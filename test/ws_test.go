package test

import (
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/SDZZGNDRC/go-okx/ws"
	"github.com/SDZZGNDRC/go-okx/ws/public"
)

func TestBooks(t *testing.T) {
	args := &ws.Args{
		Channel: "books",
		InstId:  "APT-USDC",
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

// 测试WebSocket方式获取订单簿数据的正确性
func TestBooksCorrectness(t *testing.T) {
	instIds := []string{"APT-USDC"}
	for _, instId := range instIds {
		go BooksCorrectnessHelper(instId)
		time.Sleep(time.Millisecond * 100)
	}
	log.Println("Subscribe all...")
	select {} // stick forever
}

func BooksCorrectnessHelper(instId string) {
	LastBooks := public.Book{}
	CurrentBooks := public.Book{}
	var count int64
	args := &ws.Args{
		Channel: "books",
		InstId:  instId,
	}
	handler := func(e public.EventBooks) {
		var t [][]string
		if e.Action == "snapshot" { // 全量数据
			CurrentBooks.Asks = e.Data[0].Asks
			CurrentBooks.Bids = e.Data[0].Bids
		} else { // 增量数据
			// 更新Asks
			for _, ask := range e.Data[0].Asks {
				price := ask[0]
				price_f, err := strconv.ParseFloat(price, 64)
				if err != nil {
					panic(err)
				}
				existingIndex := -1

				// 在CurrentBooks中查找相同价格的ask
				for i, existingAsk := range CurrentBooks.Asks {
					if existingAsk[0] == price {
						existingIndex = i
						break
					}
				}

				// 如果找到相同价格的ask
				if existingIndex != -1 {
					if ask[1] == "0" {
						// 数量为0，从snapshot中删除该ask
						t = [][]string{}
						t = append(t, CurrentBooks.Asks[:existingIndex]...)
						t = append(t, CurrentBooks.Asks[existingIndex+1:]...)
						CurrentBooks.Asks = t
					} else {
						// 数量有变化，替换该ask的数据
						CurrentBooks.Asks[existingIndex] = ask
					}
				} else {
					// 如果没有相同价格的ask，按照价格升序插入
					insertIndex := sort.Search(len(CurrentBooks.Asks), func(i int) bool {
						t, err := strconv.ParseFloat(CurrentBooks.Asks[i][0], 64)
						if err != nil {
							panic(err)
						}
						return t > price_f
					})
					if insertIndex == len(CurrentBooks.Asks) {
						CurrentBooks.Asks = append(CurrentBooks.Asks, ask)
					} else {
						t = [][]string{}
						t = append(t, CurrentBooks.Asks[:insertIndex]...)
						t = append(t, ask)
						t = append(t, CurrentBooks.Asks[insertIndex:]...)
						CurrentBooks.Asks = t
						// log.Printf("Asks: insertIndex: %d, len: %d", insertIndex, len(CurrentBooks.Asks))
					}
				}
			}

			// 更新Bids
			for _, bid := range e.Data[0].Bids {
				price := bid[0]
				price_f, err := strconv.ParseFloat(price, 64)
				if err != nil {
					panic(err)
				}
				existingIndex := -1

				// 在CurrentBooks中查找相同价格的bid
				for i, existingBid := range CurrentBooks.Bids {
					if existingBid[0] == price {
						existingIndex = i
						break
					}
				}

				// 如果找到相同价格的bid
				if existingIndex != -1 {
					if bid[1] == "0" {
						// 数量为0，从snapshot中删除该bid
						t = [][]string{}
						t = append(t, CurrentBooks.Bids[:existingIndex]...)
						t = append(t, CurrentBooks.Bids[existingIndex+1:]...)
						CurrentBooks.Bids = t
					} else {
						// 数量有变化，替换该bid的数据
						CurrentBooks.Bids[existingIndex] = bid
					}
				} else {
					// 如果没有相同价格的bid，按照价格升序插入
					insertIndex := sort.Search(len(CurrentBooks.Bids), func(i int) bool {
						t, err := strconv.ParseFloat(CurrentBooks.Bids[i][0], 64)
						if err != nil {
							panic(err)
						}
						return t < price_f
					})
					if insertIndex == len(CurrentBooks.Bids) {
						CurrentBooks.Bids = append(CurrentBooks.Bids, bid)
					} else {
						t = [][]string{}
						t = append(t, CurrentBooks.Bids[:insertIndex]...)
						t = append(t, bid)
						t = append(t, CurrentBooks.Bids[insertIndex:]...)
						CurrentBooks.Bids = t
						// copy(CurrentBooks.Bids[insertIndex+1:], CurrentBooks.Bids[insertIndex:])
						// CurrentBooks.Bids[insertIndex] = bid
					}
				}
			}
		}

		CurrentBooks.Checksum = e.Data[0].Checksum
		calChecksum := public.CalculateChecksum(CurrentBooks)
		if calChecksum != CurrentBooks.Checksum {
			log.Printf("%s: calCheckSum:%d, correct:%d\n", instId, calChecksum, CurrentBooks.Checksum)
			json_LastBooks, err := json.Marshal(LastBooks)
			if err != nil {
				panic(err)
			}
			json_CurrentBooks, err := json.Marshal(CurrentBooks)
			if err != nil {
				panic(err)
			}
			json_e, err := json.Marshal(e)
			if err != nil {
				panic(err)
			}
			log.Println(string(json_LastBooks))
			log.Println(string(json_CurrentBooks))
			log.Println(string(json_e))
			panic("Stop")
		} else {
			LastBooks = CurrentBooks
			log.Printf("No %d: Correct...", count)
			count += 1
		}
	}

	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeBooks(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {} // Wait forever
}
