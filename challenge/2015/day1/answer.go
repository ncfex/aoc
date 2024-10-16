package main

import (
	"fmt"
	"log"
	"os"
)

func processFloor(floors []byte) (int, int) {
	currentFloor := 0
	firstBasement := 0
	found := false
	for i, floor := range floors {
		if floor&1 == 0 { // ( is ASCII 40, LSB is 0
			currentFloor++
		} else {
			currentFloor--
		}

		if currentFloor == -1 && !found {
			firstBasement = i + 1
		}
	}

	return currentFloor, firstBasement
}

func main() {
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	answer, firstBasement := processFloor(inputData)
	output := fmt.Sprintf("%d %d", answer, firstBasement)

	err = os.WriteFile("answer-out.txt", []byte(output), 0600)
	if err != nil {
		log.Fatal("error writing file:", err)
	}
}
