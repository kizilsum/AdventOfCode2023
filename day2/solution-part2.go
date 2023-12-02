package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
	games := []CubesGame{}
	file, _ := os.Open("./input.txt")
	defer file.Close()
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
		splitSemicolon := strings.Split(line, ":")
		gameId := strings.Split(splitSemicolon[0], " ")[1]
		
		cubesGame := CubesGame{}
		cubesGame.Id, _ = strconv.Atoi(gameId)

		gamesStr := strings.Split(splitSemicolon[1], ";")
		
		for _, gameInput := range gamesStr {
			cubes := strings.Split(gameInput, ",")
			
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				cubeNumberAndName := strings.Split(cube, " ")
				cubeNumber, _ := strconv.Atoi(cubeNumberAndName[0])
				cubeName := cubeNumberAndName[1]

				if cubeName == "red" && cubesGame.MinRed < cubeNumber {
					cubesGame.MinRed = cubeNumber
				} else if cubeName == "green" && cubesGame.MinGreen < cubeNumber {
					cubesGame.MinGreen = cubeNumber
				} else if cubeName == "blue" && cubesGame.MinBlue < cubeNumber {
					cubesGame.MinBlue = cubeNumber
				}
			}
		}

		games = append(games, cubesGame)
	}

	sum := 0
	for _, game := range games {
		sum += (game.MinRed * game.MinGreen * game.MinBlue)
	}
	fmt.Println(sum)
}

type CubesGame struct {
	Id			int
	MinRed 		int
	MinGreen 	int
	MinBlue 	int
}

