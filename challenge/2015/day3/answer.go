package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file")
	}

	visitedPositions := make(map[string]int)
	currentPos := []int{0, 0} // x,y pos

	// init 0,0
	key := fmt.Sprintf("%d,%d", currentPos[0], currentPos[1])
	visitedPositions[key]++

	directions := string(content)
	for _, dir := range directions {
		strDir := string(dir)
		switch strDir {
		case "^":
			currentPos[0]++
		case "v":
			currentPos[0]--
		case ">":
			currentPos[1]++
		case "<":
			currentPos[1]--
		}

		key := fmt.Sprintf("%d,%d", currentPos[0], currentPos[1])
		visitedPositions[key]++
	}

	fmt.Println(len(visitedPositions))
}
