package main

import (
	"fmt"
	"sort"
)

// Дан массив интервалов. Интервал представлен парой [start, end].
// Необходимо объединить все пересекающиеся интервалы
// и вернуть массив непересекающихся интервалов.

// Пример:
// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
// output: [[1 6] [8 10] [15 18]]

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	type interval struct {
		val   int
		typeV rune
	}

	intervals2 := make([]interval, 0, len(intervals)*2)

	for _, v := range intervals {
		intervals2 = append(intervals2, interval{v[0], 's'})
		intervals2 = append(intervals2, interval{v[1], 'e'})
	}

	sort.Slice(intervals2, func(i, j int) bool {
		return intervals2[i].val < intervals2[j].val
	})

	res := make([][]int, 0, len(intervals))

	// fmt.Println(intervals2)

	stStack := 0
	stVal := 0
	for _, r := range intervals2 {
		if r.typeV == 's' {
			if stStack == 0 {
				stVal = r.val
			}
			stStack++
		} else {
			if stStack == 1 {
				res = append(res, []int{stVal, r.val})
			}
			stStack--
		}
	}

	fmt.Println(res)
}
