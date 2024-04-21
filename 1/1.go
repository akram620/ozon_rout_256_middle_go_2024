package _

import (
	"fmt"
)

func main1() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var shipSizes [10]int

		for j := 0; j < 10; j++ {
			fmt.Scan(&shipSizes[j])
		}

		oneDeck, twoDeck, threeDeck, fourDeck := 0, 0, 0, 0
		for _, size := range shipSizes {
			switch size {
			case 1:
				oneDeck++
			case 2:
				twoDeck++
			case 3:
				threeDeck++
			case 4:
				fourDeck++
			}
		}

		if oneDeck == 4 && twoDeck == 3 && threeDeck == 2 && fourDeck == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
