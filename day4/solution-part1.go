// package main

// import "fmt"
// import "os"
// import "bufio"
// import "strings"
// import "strconv"
// import "math"

// func main() {
// 	scratchcards := readScratchcardsFromFile("./input.txt")
// 	sum := 0
// 	for _, scratchcard := range scratchcards {
// 		scratchcard.MatchingNumbers = calculateMatchingNumbers(scratchcard)
// 		fmt.Println(scratchcard)
// 		lenMatchingNumbers := len(scratchcard.MatchingNumbers)
// 		if lenMatchingNumbers > 0 {
// 			sum += int(math.Pow(2, float64(lenMatchingNumbers - 1)))
// 		}
// 	}
// 	fmt.Println(sum)
// }

// func readScratchcardsFromFile(filePath string) []Scratchcard {
// 	scratchcards := []Scratchcard {}
// 	file, _ := os.Open(filePath)
// 	defer file.Close()
//     scanner := bufio.NewScanner(file)
	
//     for scanner.Scan() {
//         line := scanner.Text()

		
// 		idAndNumbers := strings.Split(line, ":")
// 		id, _ := strconv.Atoi(strings.Split(idAndNumbers[0], " ")[1])
// 		numbers := strings.Split(idAndNumbers[1], "|")
// 		winningNumbersStr := strings.Split(strings.Trim(numbers[0], " "), " ")
// 		cardNumbersStr := strings.Split(strings.Trim(numbers[1], " "), " ")
		
// 		scratchcard := Scratchcard{}
// 		scratchcard.Id = id
// 		scratchcard.WinningNumbers = convertToIntArr(winningNumbersStr)
// 		scratchcard.CardNumbers = convertToIntArr(cardNumbersStr)

// 		scratchcards = append(scratchcards, scratchcard)
// 	}

// 	return scratchcards
// }

// func convertToIntArr(arr []string) []int {
// 	numbers := []int{}
// 	for _, numberStr := range arr {
// 		number, _ := strconv.Atoi(strings.Trim(numberStr, " "))
// 		if number != 0 {
// 			numbers = append(numbers, number)
// 		}
// 	}
// 	return numbers
// }

// func calculateMatchingNumbers(scratchcard Scratchcard) []int {
// 	matchingNumbers := []int{}
// 	for _, cardNumber := range scratchcard.CardNumbers {
// 		for _, winningNumber := range scratchcard.WinningNumbers {
// 			if cardNumber == winningNumber {
// 				matchingNumbers = append(matchingNumbers, cardNumber)
// 			}
// 		}
// 	}
// 	return matchingNumbers
// }

// type Scratchcard struct {
// 	Id				int
// 	WinningNumbers 	[]int
// 	CardNumbers 	[]int
// 	MatchingNumbers []int
// }