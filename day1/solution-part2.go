package main

import "fmt"
import "strconv"
import "os"
import "bufio"
import "strings"

var letterNumbers = map[string]string {
	"one": "1", 
	"two": "2", 
	"three": "3", 
	"four": "4", 
	"five": "5", 
	"six": "6", 
	"seven": "7", 
	"eight": "8", 
	"nine": "9",
	"1": "1", 
	"2": "2", 
	"3": "3", 
	"4": "4", 
	"5": "5", 
	"6": "6", 
	"7": "7", 
	"8": "8", 
	"9": "9",
}

func main() {
	total := 0
	file, _ := os.Open("./input.txt")

    defer file.Close()
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
		first := findFirstNumber(line)
		last := findLastNumber(line)

		calibrationValStr := first + last
		fmt.Println(calibrationValStr)
		
		calibrationVal, _ := strconv.Atoi(calibrationValStr)
		total += calibrationVal
    }

	fmt.Println(total)
}

func findLastNumber(text string) string {
	subString := ""
	for i := len(text)-1; i >= 0; i-- {
		c := string(text[i])
		subString = c + subString
		if _, err := strconv.Atoi(c); err == nil {
			return c
		} else {
			for key, val := range letterNumbers {
				if strings.Contains(subString, key) {
					return val
				}
			}
		}
	}
	return "NaN"
}

func findFirstNumber(text string) string {
	subString := ""
	for i := 0; i < len(text); i++ {
		c := string(text[i])
		subString += c
		if _, err := strconv.Atoi(c); err == nil {
			return c
		} else {
			for key, val := range letterNumbers {
				if strings.Contains(subString, key) {
					return val
				}
			}
		}
	}
	return "NaN"
}