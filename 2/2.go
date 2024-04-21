package _

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main2() {
	var t int
	fmt.Scan(&t)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < t; i++ {
		scanner.Scan()
		input := scanner.Text()

		inputArr := strings.Split(input, " ")
		d, m, y := inputArr[0], inputArr[1], inputArr[2]
		if len(d) == 1 {
			d = "0" + d
		}
		if len(m) == 1 {
			m = "0" + m
		}
		input = d + " " + m + " " + y

		format := "02 01 2006"

		_, err := time.Parse(format, input)

		if err != nil {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
