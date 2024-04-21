package _

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main7() {
	var t int
	fmt.Scanf("%d\n", &t)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < t; i++ {

		var k int
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		k, _ = strconv.Atoi(input)

		scanner.Scan()
		pagesArr := strings.Split(strings.TrimSpace(scanner.Text()), ",")

		pagesMap := make(map[int]struct{}, len(pagesArr))

		for _, s := range pagesArr {
			isRange := strings.Contains(s, "-")

			if isRange {
				parts := strings.Split(s, "-")
				p1Str, p2Str := parts[0], parts[1]

				p1, _ := strconv.Atoi(p1Str)
				p2, _ := strconv.Atoi(p2Str)

				for page := p1; page <= p2; page++ {
					pagesMap[page] = struct{}{}
				}

				continue
			} else {
				page, _ := strconv.Atoi(s)
				pagesMap[page] = struct{}{}
			}
		}

		var unprintedPages []int

		for page := 1; page <= k; page++ {
			_, ok := pagesMap[page]
			if !ok {
				unprintedPages = append(unprintedPages, page)
			}
		}

		sort.Ints(unprintedPages)

		if len(unprintedPages) == 1 {
			fmt.Println(unprintedPages[0])
			continue
		}

		var ranges []string
		pointer1 := 0
		pointer2 := 0

		for ind := 1; ind < len(unprintedPages); ind++ {
			if unprintedPages[ind]-unprintedPages[ind-1] == 1 {
				pointer2 = ind
				if ind == len(unprintedPages)-1 {
					if pointer1 == pointer2 {
						ranges = append(ranges, fmt.Sprintf("%d", unprintedPages[pointer1]))
					} else {
						ranges = append(ranges, fmt.Sprintf("%d-%d", unprintedPages[pointer1], unprintedPages[pointer2]))
					}
				}

			} else {
				if pointer1 == pointer2 {
					ranges = append(ranges, fmt.Sprintf("%d", unprintedPages[pointer1]))
				} else {
					ranges = append(ranges, fmt.Sprintf("%d-%d", unprintedPages[pointer1], unprintedPages[pointer2]))
				}
				pointer1 = ind
				pointer2 = ind

				if ind == len(unprintedPages)-1 {
					if pointer1 == pointer2 {
						ranges = append(ranges, fmt.Sprintf("%d", unprintedPages[pointer1]))
					} else {
						ranges = append(ranges, fmt.Sprintf("%d-%d", unprintedPages[pointer1], unprintedPages[pointer2]))
					}
				}
			}
			pointer2 = ind
		}

		fmt.Println(strings.Join(ranges, ","))
	}
}
