// package main

// import "fmt"
// import "strconv"
// import "os"
// import "bufio"

// func main() {
// 	total := 0
// 	file, _ := os.Open("./input.txt")

//     defer file.Close()
//     scanner := bufio.NewScanner(file)

//     for scanner.Scan() {
//         line := scanner.Text()

// 		first, last := "", ""

// 		for _, r := range line {
// 			c := string(r)
// 			if _, err := strconv.Atoi(c); err == nil {
// 				if first == "" {
// 					first = c
// 				} else {
// 					last = c
// 				}
// 			}
// 		}

// 		calibrationValStr := ""

// 		if last == "" {
// 			calibrationValStr = first + first
// 		} else {
// 			calibrationValStr = first + last
// 		}
		
// 		calibrationVal, _ := strconv.Atoi(calibrationValStr)
// 		total += calibrationVal
//     }

// 	fmt.Println(total)
// }