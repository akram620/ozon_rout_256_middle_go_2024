package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, p int
		fmt.Fscan(in, &n, &p)

		var sum float64

		for j := 0; j < n; j++ {
			var price int
			fmt.Fscan(in, &price)

			var c float64
			c = (float64(price) * float64(p)) / 100

			diff := c - float64(int(c))

			sum += diff
		}

		fmt.Println(fmt.Sprintf("%.2f", sum))
	}
}
