package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filePath string) []string {
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

	readFile.Close()

	return fileLines
}
