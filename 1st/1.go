package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func extractDigits(line string) (int, int) {
	firstDigit := strings.IndexFunc(line, unicode.IsDigit)
	lastDigit := strings.LastIndexFunc(line, unicode.IsDigit)

	return int(line[firstDigit] - '0'), int(line[lastDigit] - '0')
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	ans := uint64(0)

	for scanner.Scan() {
		firstDigit, lastDigit := extractDigits(scanner.Text())
		// fmt.Printf("for input line: %s, digits are (%d, %d)\n", scanner.Text(), firstDigit, lastDigit)
		ans += uint64(firstDigit*10 + lastDigit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
