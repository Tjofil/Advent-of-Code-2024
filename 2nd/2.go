package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var spelledDigit = [9]string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func extractRuneDigitIdx(line string) (int, int) {
	firstIdx := strings.IndexFunc(line, unicode.IsDigit)
	lastIdx := strings.LastIndexFunc(line, unicode.IsDigit)
	if firstIdx == -1 {
		firstIdx = len(line)
	}
	return firstIdx, lastIdx
}

func extractSpelledDigitIdx(line string) (int, int, int, int) {
	firstIdx, lastIdx := len(line), -1
	firstDigit, lastDigit := -1, -1
	for num, spelled := range spelledDigit {
		if idx := strings.Index(line, spelled); idx != -1 && idx < firstIdx {
			firstIdx = idx
			firstDigit = num + 1
		}
		if idx := strings.LastIndex(line, spelled); idx != -1 && idx > lastIdx {
			lastIdx = idx
			lastDigit = num + 1
		}
	}
	return firstIdx, lastIdx, firstDigit, lastDigit
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
		firstSpelledIdx, lastSpelledIdx, firstDigit, lastDigit := extractSpelledDigitIdx(scanner.Text())
		firstRuneIdx, lastRuneIdx := extractRuneDigitIdx(scanner.Text())
		// fmt.Printf("spelled out digits: (%d, %d)\n", firstDigit, lastDigit)
		if firstRuneIdx < firstSpelledIdx {
			firstDigit = int(scanner.Text()[firstRuneIdx] - '0')
		}
		if lastRuneIdx > lastSpelledIdx {
			lastDigit = int(scanner.Text()[lastRuneIdx] - '0')
		}
		// fmt.Printf("for input line: %s, digits are (%d, %d)\n", scanner.Text(), firstDigit, lastDigit)
		ans += uint64(firstDigit*10 + lastDigit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
