package main

import (
	"log"
	"os"
	"strconv"
)

func processFloor(floors []byte) int {
	currentFloor := 0
	for _, floor := range floors {
		if floor&1 == 0 { // ( is ASCII 40, LSB is 0
			currentFloor++
		} else {
			currentFloor--
		}
	}

	return currentFloor
}

func main() {
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	answer := processFloor(inputData)
	answerStr := strconv.Itoa(answer)

	err = os.WriteFile("answer.txt", []byte(answerStr), 0600)
	if err != nil {
		log.Fatal("error writing file:", err)
	}
}
