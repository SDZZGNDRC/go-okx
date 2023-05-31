package public

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// 计算订单簿的校验和
func CalculateChecksum(book Book) int32 {
	bookString := BookToString(book.Asks, book.Bids)

	// Calculate CRC32 checksum
	checksum := crc32.ChecksumIEEE([]byte(bookString))

	return int32(checksum)
}

func BookToString(asks [][]string, bids [][]string) string {
	var bookString string
	numbers := []int{len(asks), len(bids), 25}
	sort.Ints(numbers)
	l := numbers[0]
	for i := 0; i < l; i++ {
		bookString += fmt.Sprintf(
			"%s:%s:%s:%s",
			bids[i][0],
			bids[i][1],
			asks[i][0],
			asks[i][1],
		)
		if i != l-1 {
			bookString += ":"
		}
	}
	return bookString
}
