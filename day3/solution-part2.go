package main

import "fmt"
import "os"
import "bufio"
import "unicode"
import "strconv"

func main() {
	engineSchematic := readEngineSchematicFromFile("./input.txt", 140)
	sum := calculateSumOfEngineParts(engineSchematic)
	fmt.Println(sum)
}

func readEngineSchematicFromFile(filePath string, rowCount int) [][]rune {
	engineSchematic := make([][]rune, rowCount)
	file, _ := os.Open(filePath)
	defer file.Close()
    scanner := bufio.NewScanner(file)
	
	counter := 0
    for scanner.Scan() {
        line := scanner.Text()
		engineSchematic[counter] = make([]rune, len(line))
		for i, c := range line {
			engineSchematic[counter][i] = c
		}
		counter++
	}

	return engineSchematic
}

func calculateSumOfEngineParts(engineSchematic [][]rune) int {
	sum := 0
	engineParts := []EnginePart{}
	gearRatioMap := map[string][]int{}
	digits := []rune{}
	for i, arr := range engineSchematic {
		for j, r := range arr {
			if unicode.IsDigit(r) {
				digits = append(digits, r)
			} 
			
			if len(digits) != 0 && (!unicode.IsDigit(r) || j == len(engineSchematic[i]) - 1) {
				enginePart := getEnginePart(engineSchematic, digits, i, j - len(digits))
				if enginePart.Number != -1 {
					engineParts = append(engineParts, enginePart)
				}
				digits = []rune{}
			}
		}
	}
	
	for _, enginePart := range engineParts {
		coordinate := string(enginePart.Row) + string(enginePart.Col)
		arr, exists := gearRatioMap[coordinate]
		if exists {
			gearRatioMap[coordinate] = append(arr, enginePart.Number)
		} else {
			gearRatioMap[coordinate] = []int {enginePart.Number}
		}
	}

	for _, arr := range gearRatioMap {
		if (len(arr) == 2) {
			sum += (arr[0] * arr[1])
		}
	}

	return sum
}

func getEnginePart(engineSchematic [][]rune, digits []rune, row int, col int) EnginePart {
	rStart, rEnd, cStart, cEnd := getCornersForTheNumber(engineSchematic, digits, row ,col)
	enginePart := EnginePart{}
	enginePart.Number = -1
				
	for cIndex := cStart; cIndex < cEnd; cIndex++ {
		for rIndex := rStart; rIndex < rEnd; rIndex++ {
			r := engineSchematic[rIndex][cIndex]
			if string(r) == "*" {
				enginePartNumber, _ := strconv.Atoi(string(digits))
				enginePart := EnginePart{}
				enginePart.Number = enginePartNumber
				enginePart.Row = rIndex
				enginePart.Col = cIndex
				return enginePart
			}
		}
	}
	return enginePart
}

func getCornersForTheNumber(engineSchematic [][]rune, digits []rune, row int, col int) (int, int, int, int) {
	rStart, rEnd := row, row + 2
	cStart, cEnd := col, col + len(digits) + 1
	if row > 0 {
		rStart = row - 1
	}

	if col > 0 {
		cStart = col - 1
	}
	
	if rEnd > len(engineSchematic) {
		rEnd = len(engineSchematic) - 1
	}

	if cEnd > len(engineSchematic[row]) {
		cEnd = col + len(digits)
	}

	return rStart, rEnd, cStart, cEnd
}

type EnginePart struct {
	Number 	int
	Row		int
	Col		int
}