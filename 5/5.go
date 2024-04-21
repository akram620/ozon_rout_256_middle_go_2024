package _

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main5() {
	var t int
	fmt.Scanf("%d\n", &t)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < t; i++ {

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		var res []string
		isRoot := true
		var lastEl int
		count := 0
		st := '0'

		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())

		nStr := strings.Split(input, " ")

		var cur int
		for j := 0; j < len(nStr); j++ {

			cur, _ = strconv.Atoi(nStr[j])

			if isRoot {
				isRoot = false
				res = append(res, strconv.Itoa(cur))
				lastEl = cur
				continue
			}

			if cur-lastEl == 1 && (st == '0' || st == '+') {
				st = '+'
				count++
			} else if cur-lastEl == -1 && (st == '0' || st == '-') {
				count++
				st = '-'
			} else {
				if st == '-' {
					count *= -1
				}
				res = append(res, strconv.Itoa(count))

				res = append(res, strconv.Itoa(cur))
				st = '0'
				count = 0
			}
			lastEl = cur
		}
		if st == '-' {
			count *= -1
		}
		if st == '0' {
			count = 0
		}
		res = append(res, strconv.Itoa(count))

		fmt.Println(len(res))
		fmt.Println(strings.Join(res, " "))
	}
}
