package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var gridRows = 1000
var gridCols = 1000

var grid = make([][]bool, gridRows)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file")
	}

	for i := 0; i < gridRows; i++ {
		grid[i] = make([]bool, gridCols)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")
		event := split[:1][0]

		var from []string
		var to []string
		switch event {
		case "turn":
			from = strings.Split(split[2], ",")
			to = strings.Split(split[4], ",")
			switch split[1] {
			case "on":
				event = "on"
			case "off":
				event = "off"
			}
		case "toggle":
			from = strings.Split(split[1], ",")
			to = strings.Split(split[len(split)-1], ",")
			event = "toggle"
		}

		rStart, err := strconv.Atoi(from[0])
		rEnd, err2 := strconv.Atoi(to[0])
		cStart, err3 := strconv.Atoi(from[1])
		cEnd, err4 := strconv.Atoi(to[1])
		if err != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Fatal("error converting")
		}

		for r := rStart; r <= rEnd; r++ {
			for c := cStart; c <= cEnd; c++ {
				switch event {
				case "on":
					grid[r][c] = true
				case "off":
					grid[r][c] = false
				case "toggle":
					grid[r][c] = !grid[r][c]
				}
			}
		}
	}

	lightCount := 0
	for r := 0; r < gridRows; r++ {
		for c := 0; c < gridCols; c++ {
			if grid[r][c] {
				lightCount++
			}
		}
	}

	fmt.Println(lightCount)
}
