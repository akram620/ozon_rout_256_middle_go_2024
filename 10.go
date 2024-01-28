package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
1
14
75 22 I'm fine. Thank you.
84 82     Ciao!
26 22 So-so
45 26 What's wrong?
22 -1 How are you?
72 45 Maybe I got sick
81 72 I wish you a speedy recovery!
97 26   Stick it!
2 97 Thanks
47 72 I also got sick recently.
25 -1 Hi!
82 -1 Bye
17 82 Good day!
29 72 Visit the doctor

1
3
82 -1 Bye
17 82 Good day!
84 82     Ciao!



1
18
30 14 hbeuvjty
6 4 tbeewqn
32 15 hze
12 11 flk
14 6 fttenmoi
15 7 a
25 18 tezww
3 2 dpgfzcltb
11 7 phuol
20 15 uthrxhm
31 20 kmjferevx
4 3 uvuemy
17 11 bokbrdzbb
7 4 odw
23 12 pxbx
18 15 bzzit
2 -1 yzejjtcbd
35 32 rackuy




*/

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
			input = strings.TrimSpace(scanner.Text())
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
		printSlice(res, 0, 0)
		if i < t-1 {
			fmt.Println()
		}
	}
}

func printSlice(res []commentsMain, step, stack int) {
	for i, comment := range res {
		nStack := stack
		if comment.p == -1 {
			// root
			fmt.Println(comment.text)
		} else {
			var pref = ""
			for st := 0; st < step-1; st++ {
				if st < stack {
					pref += "|  "
				} else {
					pref += "   "
				}
			}
			fmt.Println(pref + "|")
			fmt.Println(pref + "|--" + comment.text)

			if i < len(res)-1 {
				nStack++
			}
		}
		printSlice(comment.child, step+1, nStack)
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
