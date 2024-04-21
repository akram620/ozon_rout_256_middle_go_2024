package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 10. Дерево комментариев

// Дано список комментариев, каждый из которых имеет уникальный идентификатор и родительский комментарий.
// Комментарии могут быть вложенными, то есть один комментарий может быть ответом на другой.
// Eсли у комментария родительский комментарий равен -1, то это корневой комментарий.
// Напишите программу, которая выводит список комментариев в виде дерева.

// Формат входных данных
// В первой строке задано целое число t — количество тестов.
// Далее следуют t тестов. Каждый тест начинается с числа n — количество комментариев.
// Далее следуют n строк, каждая из которых содержит
// - идентификатор комментария id
// - идентификатор родительского комментария p
// - текст комментария text

// Для каждого теста выведите список комментариев в виде дерева.

type commentForMap struct {
	id   int
	p    int
	text string
}

type commentsMain struct {
	id    int
	p     int
	text  string
	child []commentsMain
}

func main10() {
	var t int
	fmt.Scanf("%d\n", &t)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < t; i++ {

		var n int
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		n, _ = strconv.Atoi(input)

		commentsMap := make(map[int]commentForMap)

		for j := 0; j < n; j++ {
			scanner.Scan()
			input = scanner.Text()
			inputArr := strings.Split(input, " ")

			idStr, pStr := inputArr[0], inputArr[1]
			text := input[(len(idStr) + len(pStr) + 2):]

			id, _ := strconv.Atoi(idStr)
			p, _ := strconv.Atoi(pStr)

			commentsMap[id] = commentForMap{id: id, p: p, text: text}
		}

		res := findAllChild(&commentsMap, -1)
		//for _, i3 := range res {
		//	fmt.Println(i3)
		//}
		printSlice(res, 0, map[int]struct{}{})

		if i < t-1 {
			fmt.Println()
		}
	}
}

func printSlice(res []commentsMain, step int, m map[int]struct{}) {
	for i, comment := range res {
		if comment.p == -1 {
			// root
			fmt.Println(comment.text)
		} else {

			delete(m, step-1)
			if i < len(res)-1 {
				m[step-1] = struct{}{}
			}

			var pref = ""
			for i = 0; i < step-1; i++ {
				_, ok := m[i]
				if ok {
					pref += "|  "
				} else {
					pref += "   "
				}
			}

			fmt.Println(pref + "|")
			fmt.Println(pref + "|--" + comment.text)
		}

		printSlice(comment.child, step+1, m)

		if step == 0 && i < len(res)-1 {
			fmt.Println()
		}
	}
}

func findAllChild(commentsMap *map[int]commentForMap, id int) []commentsMain {
	var res []commentsMain

	for _, com := range *commentsMap {
		if com.p == id {
			children := findAllChild(commentsMap, com.id)

			// here I will write sort (like binary adding...)
			res = append(res, commentsMain{id: com.id, p: com.p, text: com.text, child: children})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].id < res[j].id
	})
	return res
}
