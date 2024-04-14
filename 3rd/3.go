package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func provideNewMap() map[string]int {
	return map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	ans := 0
	maxMap := provideNewMap()
	gameNum := 1
	for scanner.Scan() {
		trimmed, _ := strings.CutPrefix(scanner.Text(), "Game "+fmt.Sprint(gameNum)+": ")
		draws := strings.Split(trimmed, "; ")
		for _, draw := range draws {
			colors := strings.Split(draw, ", ")
			for _, color := range colors {
				split := strings.Split(color, " ")
				num, _ := strconv.Atoi(split[0])
				maxMap[split[1]] = max(maxMap[split[1]], num)
			}
		}
		ans += maxMap["red"] * maxMap["green"] * maxMap["blue"]
		gameNum += 1
		maxMap = provideNewMap()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ans)
}
