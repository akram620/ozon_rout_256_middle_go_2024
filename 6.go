package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	res := make([]string, 0)

	for i := 0; i < t; i++ {
		var c string
		fmt.Scanf("%s\n", &c)

		currentTexts := make([][]string, 0)
		currentTexts = append(currentTexts, []string{})

		verPosition := 0
		horPosition := 0

		// tt := ""

		for j := 0; j < len(c); j++ {

			//tt += string(c[j])
			//
			//fmt.Println(horPosition, verPosition, string(c[j]))
			//for _, texts := range currentTexts {
			//	fmt.Println(texts)
			//}
			//fmt.Println("-----------------")

			// left
			if c[j] == 'L' {
				if horPosition > 0 {
					horPosition--
				}
				continue
			}

			// right
			if c[j] == 'R' {
				if horPosition < len(currentTexts[verPosition]) {
					horPosition++
				}
				continue
			}

			// up
			if c[j] == 'U' {
				if verPosition > 0 {
					verPosition--

					if horPosition > len(currentTexts[verPosition])-1 {
						horPosition = len(currentTexts[verPosition])
					}
				}
				continue
			}

			// down
			if c[j] == 'D' {
				if verPosition < len(currentTexts)-1 {
					verPosition++

					if horPosition > len(currentTexts[verPosition])-1 {
						horPosition = len(currentTexts[verPosition])
					}
				}
				continue
			}

			// home
			if c[j] == 'B' {
				horPosition = 0
				continue
			}

			// end
			if c[j] == 'E' {
				horPosition = len(currentTexts[verPosition])
				continue
			}

			// enter
			if c[j] == 'N' {
				if verPosition == len(currentTexts)-1 {
					el := make([]string, 0)
					for _, s := range currentTexts[verPosition][horPosition:] {
						el = append(el, s)
					}
					currentTexts = append(currentTexts, el)
				} else {
					el := make([]string, 0)
					for _, s := range currentTexts[verPosition][horPosition:] {
						el = append(el, s)
					}
					currentTexts = insertIntoSlice(currentTexts, verPosition+1, el)
				}
				currentTexts[verPosition] = currentTexts[verPosition][:horPosition]
				verPosition++
				horPosition = 0
				continue
			}

			// add symbol
			if horPosition == len(currentTexts[verPosition]) {
				currentTexts[verPosition] = append(currentTexts[verPosition], string(c[j]))
			} else {
				currentTexts[verPosition] = insertIntoSliceStr(currentTexts[verPosition], horPosition, string(c[j]))
			}
			horPosition++
		}

		for _, texts := range currentTexts {
			res = append(res, strings.Join(texts, ""))
			//fmt.Println(strings.Join(texts, ""))
		}

		res = append(res, "-")
		//fmt.Println("-")
	}
}

func insertIntoSlice(slice [][]string, index int, elements []string) [][]string {
	firstPart := slice[:index]
	secondPart := slice[index:]

	result := make([][]string, 0)
	result = append(result, firstPart...)
	result = append(result, elements)
	result = append(result, secondPart...)

	return result
}

func insertIntoSliceStr(slice []string, index int, elements string) []string {
	firstPart := slice[:index]
	secondPart := slice[index:]

	result := make([]string, 0)
	result = append(result, firstPart...)
	result = append(result, elements)
	result = append(result, secondPart...)

	return result
}
