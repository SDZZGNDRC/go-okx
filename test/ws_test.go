package test

import (
	"log"
	"sort"
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

// 测试WebSocket方式获取订单簿数据的正确性
func TestBooksCorrectness(t *testing.T) {
	CurrentBooks := public.Book{}
	args := &ws.Args{
		Channel: "books",
		InstId:  "BTC-USDT",
	}
	handler := func(e public.EventBooks) {
		if e.Action == "snapshot" { // 全量数据
			CurrentBooks.Asks = e.Data[0].Asks
			CurrentBooks.Bids = e.Data[0].Bids
		} else { // 增量数据
			// 更新Asks
			for _, ask := range e.Data[0].Asks {
				price := ask[0]
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
						CurrentBooks.Asks = append(
							CurrentBooks.Asks[:existingIndex],
							CurrentBooks.Asks[existingIndex+1:]...,
						)
					} else {
						// 数量有变化，替换该ask的数据
						CurrentBooks.Asks[existingIndex] = ask
					}
				} else {
					// 如果没有相同价格的ask，按照价格升序插入
					insertIndex := sort.Search(len(CurrentBooks.Asks), func(i int) bool {
						return CurrentBooks.Asks[i][0] > price
					})
					CurrentBooks.Asks = append(CurrentBooks.Asks, nil)
					copy(CurrentBooks.Asks[insertIndex+1:], CurrentBooks.Asks[insertIndex:])
					CurrentBooks.Asks[insertIndex] = ask
				}
			}

			// 更新Bids
			for _, bid := range e.Data[0].Bids {
				price := bid[0]
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
						CurrentBooks.Bids = append(
							CurrentBooks.Bids[:existingIndex],
							CurrentBooks.Bids[existingIndex+1:]...,
						)
					} else {
						// 数量有变化，替换该bid的数据
						CurrentBooks.Bids[existingIndex] = bid
					}
				} else {
					// 如果没有相同价格的bid，按照价格升序插入
					insertIndex := sort.Search(len(CurrentBooks.Bids), func(i int) bool {
						return CurrentBooks.Bids[i][0] < price
					})
					CurrentBooks.Bids = append(CurrentBooks.Bids, nil)
					copy(CurrentBooks.Bids[insertIndex+1:], CurrentBooks.Bids[insertIndex:])
					CurrentBooks.Bids[insertIndex] = bid
				}
			}
		}
		CurrentBooks.Checksum = e.Data[0].Checksum
		calChecksum := public.CalculateChecksum(CurrentBooks)
		if calChecksum != CurrentBooks.Checksum {
			log.Printf("calCheckSum:%d, correct:%d\n", calChecksum, CurrentBooks.Checksum)
			log.Println(public.BookToString(CurrentBooks.Asks, CurrentBooks.Bids))
			t.Fatal()
		} else {
			log.Println("Correct...")
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
