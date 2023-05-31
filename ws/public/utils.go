package public

import (
	"fmt"
	"hash/crc32"
)

// 计算订单簿的校验和
func CalculateChecksum(book Book) int32 {
	var bids, asks []string
	for _, bid := range book.Bids {
		bids = append(bids, bookToString(bid))
	}
	for _, ask := range book.Asks {
		asks = append(asks, bookToString(ask))
	}

	// Trim bids and asks to 25 levels
	if len(bids) > 25 {
		bids = bids[:25]
	}
	if len(asks) > 25 {
		asks = asks[:25]
	}

	// Concatenate bids and asks strings
	var checksumString string
	for i := 0; i < len(bids) || i < len(asks); i++ {
		if i < len(bids) {
			checksumString += bids[i] + ":"
		}
		if i < len(asks) {
			checksumString += asks[i] + ":"
		}
	}
	checksumString = checksumString[:len(checksumString)-1] // 去除末尾的":"
	// log.Println(checksumString)
	// Calculate CRC32 checksum
	checksum := crc32.ChecksumIEEE([]byte(checksumString))

	return int32(checksum)
}

func bookToString(book []string) string {
	return fmt.Sprintf("%s:%s", book[0], book[1])
}
