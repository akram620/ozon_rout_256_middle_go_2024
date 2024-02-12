package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	in = bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)

		started := false
		// canceled := false
		hasError := false

		var last uint8

		for j := 0; j < len(s); j++ {
			cur := s[j]

			if cur == last {
				hasError = true
				break
			}
			last = cur

			if cur == 'M' {
				if started == false {
					started = true
					continue
				} else {
					hasError = true
					break
				}
			}

			if cur == 'R' {
				if started {
					continue
				} else {
					hasError = true
					break
				}
			}

			if cur == 'C' {
				if started {
					started = false
					continue
				} else {
					hasError = true
					break
				}
			}

			if cur == 'D' {
				if started {
					started = false
					continue
				} else {
					hasError = true
					break
				}
			}
		}

		if last != 'D' {
			fmt.Println("NO")
			continue
		}

		if !hasError {
			if started == false {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			fmt.Println("NO")
		}

	}
}
