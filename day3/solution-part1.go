// package main

// import "fmt"
// import "os"
// import "bufio"
// import "unicode"
// import "strconv"

// func main() {
// 	engineSchematic := readEngineSchematicFromFile("./input.txt", 140)
// 	sum := calculateSumOfEngineParts(engineSchematic)
// 	fmt.Println(sum)
// }

// func readEngineSchematicFromFile(filePath string, rowCount int) [][]rune {
// 	engineSchematic := make([][]rune, rowCount)
// 	file, _ := os.Open(filePath)
// 	defer file.Close()
//     scanner := bufio.NewScanner(file)
	
// 	counter := 0
//     for scanner.Scan() {
//         line := scanner.Text()
// 		engineSchematic[counter] = make([]rune, len(line))
// 		for i, c := range line {
// 			engineSchematic[counter][i] = c
// 		}
// 		counter++
// 	}

// 	return engineSchematic
// }

// func calculateSumOfEngineParts(engineSchematic [][]rune) int {
// 	sum := 0
// 	digits := []rune{}
// 	for i, arr := range engineSchematic {
// 		for j, r := range arr {
// 			if unicode.IsDigit(r) {
// 				digits = append(digits, r)
// 			} 
			
// 			if len(digits) != 0 && (!unicode.IsDigit(r) || j == len(engineSchematic[i]) - 1) {
// 				sum += getEnginePartNumber(engineSchematic, digits, i, j - len(digits))
// 				digits = []rune{}
// 			}
// 		}
// 	}
// 	return sum
// }

// func getEnginePartNumber(engineSchematic [][]rune, digits []rune, row int, col int) int {
// 	rStart, rEnd := row, row + 2
// 	cStart, cEnd := col, col + len(digits) + 1
// 	if row > 0 {
// 		rStart = row - 1
// 	}

// 	if col > 0 {
// 		cStart = col - 1
// 	}
	
// 	if rEnd > len(engineSchematic) {
// 		rEnd = len(engineSchematic) - 1
// 	}

// 	if cEnd > len(engineSchematic[row]) {
// 		cEnd = col + len(digits)
// 	}

// 	for cIndex := cStart; cIndex < cEnd; cIndex++ {
// 		for rIndex := rStart; rIndex < rEnd; rIndex++ {
// 			r := engineSchematic[rIndex][cIndex]
// 			if r != 46 && !unicode.IsDigit(r) {
// 				enginePartNumber, _ := strconv.Atoi(string(digits))
// 				return enginePartNumber
// 			}
// 		}
// 	}

// 	return 0
// }