package _

import (
	"fmt"
)

func main4() {

	var t int
	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)

		minTemp := 30
		maxTemp := 15

		for j := 0; j < n; j++ {

			var (
				symbol string
				temp   int
			)
			fmt.Scanf("%s %d\n", &symbol, &temp)

			if symbol == ">=" {
				if temp > maxTemp {
					maxTemp = temp
				}
			} else if symbol == "<=" {
				if temp < minTemp {
					minTemp = temp
				}
			}

			if minTemp-maxTemp >= 0 {
				fmt.Println(maxTemp)
			} else {
				fmt.Println(-1)
			}
		}
		fmt.Println()
	}
}

//func saveToFile(filename string, data []string) {
//	file, err := os.Create(filename)
//	if err != nil {
//		fmt.Println("Error creating file:", err)
//		return
//	}
//	defer file.Close()
//
//	for _, value := range data {
//		if value == "0" {
//			file.WriteString("\n")
//		} else {
//			file.WriteString(fmt.Sprintf("%s", value))
//		}
//	}
//
//	fmt.Println("Output saved to", filename)
//}
//
//func compareFiles(file1 string, file2 string) {
//	content1, err := os.ReadFile(file1)
//	if err != nil {
//		fmt.Println("Error reading file:", err)
//		return
//	}
//
//	content2, err := os.ReadFile(file2)
//	if err != nil {
//		fmt.Println("Error reading file:", err)
//		return
//	}
//
//	lines1 := strings.Split(string(content1), "\n")
//	lines2 := strings.Split(string(content2), "\n")
//
//	for i := 0; i < len(lines1); i++ {
//		if lines1[i] != strings.TrimSpace(lines2[i]) {
//			fmt.Println("line:", i+1, len(lines1[i]), len(lines2[i]))
//		}
//	}
//}
