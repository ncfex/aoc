package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// aeiou
var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}
var bannedSeqs []string = []string{"ab", "cd", "pq", "xy"}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error reading file")
	}

	// 3 vowels
	// one letter twice in a row

	words := strings.Split(strings.TrimSpace(string(content)), "\n")
	niceStringsCount := 0

	var wg sync.WaitGroup
	resultCh := make(chan bool, len(words))

	for _, w := range words {
		wg.Add(1)
		go func(wrd string) {
			defer wg.Done()
			r1 := hasEnoughVowels(wrd)
			r2 := hasBannedSeq(wrd)
			r3 := hasRecurringLetter(wrd)

			if r1 && !r2 && r3 {
				resultCh <- true
			} else {
				resultCh <- false
			}
		}(w)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		if res {
			niceStringsCount++
		}
	}

	fmt.Printf("Answer: %d\n", niceStringsCount)
}

func hasEnoughVowels(word string) bool {
	vowelCount := 0
	for _, r := range word {
		if vowels[r] {
			vowelCount++
			if vowelCount == 3 {
				return true
			}
		}
	}
	return false
}

func hasBannedSeq(word string) bool {
	for _, bW := range bannedSeqs {
		if strings.Contains(word, bW) {
			return true
		}
	}
	return false
}

func hasRecurringLetter(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			return true
		}
	}
	return false
}
