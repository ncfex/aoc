package main

import (
	"fmt"
	"log"
	"os"
)

func movePosition(pos []int, direction string) {
	switch direction {
	case "^":
		pos[0]++
	case "v":
		pos[0]--
	case ">":
		pos[1]++
	case "<":
		pos[1]--
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file")
	}

	visitedPositions := make(map[string]int)
	currentPos := []int{0, 0}     // x,y pos
	currentPosRobo := []int{0, 0} // x,y pos

	// init 0,0
	visitedPositions["0,0"]++

	directions := string(content)
	for i := 0; i < len(directions); i += 2 {
		// santa
		movePosition(currentPos, string(directions[i]))
		visitedPositions[fmt.Sprintf("%d,%d", currentPos[0], currentPos[1])]++

		// robo santa
		if i+1 < len(directions) {
			movePosition(currentPosRobo, string(directions[i+1]))
			visitedPositions[fmt.Sprintf("%d,%d", currentPosRobo[0], currentPosRobo[1])]++
		}
	}

	fmt.Println(len(visitedPositions))
}
