// package main

// import "fmt"
// import "os"
// import "bufio"
// import "strings"
// import "strconv"

// func main() {
// 	limitRed, limitGreen, limitBlue := 12, 13, 14
// 	validGameIds := map[int]int {}
// 	file, _ := os.Open("./input.txt")
// 	defer file.Close()
//     scanner := bufio.NewScanner(file)

//     for scanner.Scan() {
//         line := scanner.Text()
// 		splitSemicolon := strings.Split(line, ":")
// 		gameId := strings.Split(splitSemicolon[0], " ")[1]
		
// 		cubesGame := new(CubesGame)
// 		cubesGame.Id, _ = strconv.Atoi(gameId)
// 		cubesGame.Cubes = []Cubes{}

// 		validGameIds[cubesGame.Id] = cubesGame.Id

// 		games := strings.Split(splitSemicolon[1], ";")
		
// 		for _, gameInput := range games {
// 			cubes := strings.Split(gameInput, ",")
// 			gameCube := Cubes{}
			
// 			for _, cube := range cubes {
// 				cube = strings.Trim(cube, " ")
// 				cubeNumberAndName := strings.Split(cube, " ")
// 				cubeNumber, _ := strconv.Atoi(cubeNumberAndName[0])
// 				cubeName := cubeNumberAndName[1]

// 				if cubeName == "red" {
// 					gameCube.Red = cubeNumber
// 				} else if cubeName == "green" {
// 					gameCube.Green = cubeNumber
// 				} else {
// 					gameCube.Blue = cubeNumber
// 				}
// 			}
			
// 			// cubesGame.Cubes = append(cubesGame.Cubes, gameCube)

// 			if (gameCube.Red > limitRed) || (gameCube.Green > limitGreen) || (gameCube.Blue > limitBlue) {
// 				validGameIds[cubesGame.Id] = 0
// 			}
			
// 			// fmt.Println(validGameIdCount)
// 		}
// 	}
// 	// fmt.Println(validGameIds)
// 	sum := 0
// 	for _, val := range validGameIds {
// 		sum += val
// 	}
// 	fmt.Println(sum)
// }

// type Cubes struct {  
//     Red    	int
//     Green   int
//     Blue 	int
// }

// type CubesGame struct {
// 	Id	int
// 	Cubes []Cubes
// }

