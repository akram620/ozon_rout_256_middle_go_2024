package _

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main9() {
	var t int
	fmt.Scanf("%d\n", &t)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < t; i++ {

		res = []int{}

		var n int
		scanner.Scan()
		inputArr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		n, _ = strconv.Atoi(inputArr[0])

		board := make([][]string, 0, n)

		for j := 0; j < n; j++ {
			scanner.Scan()
			inputM := strings.Split(strings.TrimSpace(scanner.Text()), "")
			board = append(board, inputM)
		}

		for rI, r := range board {
			for cI, c := range r {
				if c == "*" {
					board = findEndOfRectangle(board, rI, cI, 0)
				}
			}
		}

		//for _, i2 := range board {
		//	fmt.Println(i2)
		//}
		sort.Ints(res)

		resStr := make([]string, len(res))
		for j := 0; j < len(res); j++ {
			resStr[j] = strconv.Itoa(res[j])
		}
		fmt.Println(strings.Join(resStr, " "))
	}
}

var res []int

func findStar(board [][]string, rowStart, rowEnd, colStart, colEnd, level int) [][]string {

	for i := rowStart; i <= rowEnd; i++ {
		for j := colStart; j <= colEnd; j++ {
			if board[i][j] == "*" {
				// fmt.Println("find * at", rowStart, colStart)

				// find end of this rectangle and clear the body
				board = findEndOfRectangle(board, i, j, level)
			}
		}
	}

	return board
}

func findEndOfRectangle(board [][]string, rowStart, colStart int, level int) [][]string {

	// fmt.Println("findEndOfRectangle:", rowStart, colStart, "level:", level)

	res = append(res, level)

	var rowEnd = rowStart
	var colEnd = colStart

	for i := colStart; i < len(board[rowStart]); i++ {
		if board[rowStart][i] == "*" {
			board[rowStart][i] = "."
			colEnd = i
		} else {
			break
		}
	}

	// fmt.Println("cE", colEnd)

	rowEnd++
	for i := rowEnd; i < len(board); i++ {
		if board[i][colStart] == "*" {
			board[i][colStart] = "."
			board[i][colEnd] = "."
			rowEnd = i
		} else {
			break
		}
	}
	for i := colStart; i <= colEnd; i++ {
		board[rowEnd][i] = "."
	}

	//fmt.Println("end of findEndOfRectangle:", rowEnd, colEnd)
	//fmt.Println("--------------------------------")
	//for _, i2 := range board {
	//	fmt.Println(i2)
	//}

	if rowEnd-rowStart > 3 && colEnd-colStart > 3 {
		findStar(board, rowStart, rowEnd, colStart, colEnd, level+1)
	}

	return board
}
