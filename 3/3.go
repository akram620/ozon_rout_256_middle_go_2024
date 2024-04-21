package _

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main3() {
	var t int
	fmt.Scan(&t)

	scanner := bufio.NewScanner(os.Stdin)

	// 1. R48FA
	// 2. A9PQ

	for i := 0; i < t; i++ {
		scanner.Scan()
		input := scanner.Text()

		inputLen := len(input)

		left := 0

		// string buffer
		var sb strings.Builder
		isRight := true

		for left < inputLen {
			current := input[left:inputLen]
			if len(current) < 4 {
				fmt.Println("-")
				isRight = false
				break
			}

			if unicode.IsLetter(rune(input[left+2])) {
				r := isValidSecondType(current[:4])
				if !r {
					fmt.Println("-")
					isRight = false
					break
				} else {
					sb.WriteString(current[:4] + " ")
					left += 4
					continue
				}
			}

			if len(current) < 5 {
				fmt.Println("-")
				isRight = false
				break
			}

			r := isValidFirstType(current[:5])
			if !r {
				fmt.Println("-")
				isRight = false
				break
			} else {
				sb.WriteString(current[:5] + " ")
				left += 5
			}
		}

		if isRight {
			fmt.Println(sb.String())
		}

	}
}

// 1. R48FA
func isValidFirstType(text string) bool {
	s1 := unicode.IsLetter(rune(text[0]))
	s2 := unicode.IsDigit(rune(text[1]))
	s3 := unicode.IsDigit(rune(text[2]))
	s4 := unicode.IsLetter(rune(text[3]))
	s5 := unicode.IsLetter(rune(text[4]))

	if s1 && s2 && s3 && s4 && s5 {
		return true
	} else {
		return false
	}
}

// 2. A9PQ
func isValidSecondType(text string) bool {
	s1 := unicode.IsLetter(rune(text[0]))
	s2 := unicode.IsDigit(rune(text[1]))
	s3 := unicode.IsLetter(rune(text[2]))
	s4 := unicode.IsLetter(rune(text[3]))

	if s1 && s2 && s3 && s4 {
		return true
	} else {
		return false
	}
}
