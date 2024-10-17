package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	var wg sync.WaitGroup

	dimensions := strings.Split(strings.TrimSpace(string(content)), "\n")
	wrapperCh := make(chan int, len(dimensions))
	ribbonCh := make(chan int, len(dimensions))

	for _, line := range dimensions {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			calculateBoxPaper(l, wrapperCh, ribbonCh)
		}(string(line))
	}

	go func() {
		wg.Wait()
		close(wrapperCh)
		close(ribbonCh)
	}()

	total := 0
	for area := range wrapperCh {
		total += area
	}

	totalRibbon := 0
	for area := range ribbonCh {
		totalRibbon += area
	}

	fmt.Println(total, totalRibbon)
}

func calculateBoxPaper(line string, wpCh chan<- int, rbCh chan<- int) {
	dimensions := strings.Split(line, "x")
	w, err1 := strconv.Atoi(dimensions[0])
	l, err2 := strconv.Atoi(dimensions[1])
	h, err3 := strconv.Atoi(dimensions[2])
	if err1 != nil || err2 != nil || err3 != nil {
		log.Printf("Error parsing dimensions")
		wpCh <- 0
		rbCh <- 0
	}

	s1 := w * l
	s2 := h * l
	s3 := w * h
	smallest := min(s1, min(s2, s3))

	totalSurface := 2*s1 + 2*s2 + 2*s3 + smallest

	arr := []int{w, l, h}
	sort.Ints(arr)
	smallestP, secondSmallestP := arr[0], arr[1]
	totalRibbon := (smallestP*2 + secondSmallestP*2) + (w * h * l)

	wpCh <- totalSurface
	rbCh <- totalRibbon
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
