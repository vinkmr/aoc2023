package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")

	readFile, err := os.Open("/home/rrb/Projects/aoc2023/inputs/day2")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	gameNumber := map[string]int{}

	for _, line := range fileLines {
		currentGameNumber, currentGameNumberVal := GetGameNumber(line)
		gameNumber[currentGameNumber] = currentGameNumberVal
		// fmt.Println("Game ", currentGameNumberVal)
	}
}

func GetGameNumber(GameRoundDetails string) (string, int) {
	indexSplit := strings.Split(GameRoundDetails, ":")
	gameNumberSplit := strings.Split(indexSplit[0], " ")
	gameNumber, err := strconv.Atoi(gameNumberSplit[1])
	if err != nil {
		fmt.Println(err)
	}
	return gameNumberSplit[0], gameNumber
}
