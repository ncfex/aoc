package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"strconv"
)

func main() {
	secretKey := "yzbqklnj"
	number := 1

	for {
		h := md5.New()
		if _, err := io.WriteString(h, secretKey+strconv.Itoa(number)); err != nil {
			log.Fatal("error writing")
		}

		hashBytes := h.Sum(nil)
		if hashBytes[0] == 0 && hashBytes[1] == 0 && (hashBytes[2]>>4) == 0 {
			fmt.Printf("Answer: %d\n", number)
			break
		}
		number++
	}
}
