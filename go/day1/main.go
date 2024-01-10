package main

import (
	"aoc2023/go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var calibration_values []int

	fileLines := utils.ReadFile("/home/rrb/Projects/aoc2023/inputs/day1")
	for lineNumber, line := range fileLines {
		fmt.Println()
		fmt.Println(lineNumber, ":\t", line)
		line_cal_val := calibrate(line)
		fmt.Println(line_cal_val)
		calibration_values = append(calibration_values, line_cal_val)
	}

	var sum int
	for _, v := range calibration_values {
		sum += v
	}

	fmt.Println("Sum = ", sum)
}

func calibrate(cal_string string) int {
	var FirstNum, LastNum int
	var firstNumIndex int = -1
	var lastNumIndex int = -1
	var CalibrationNumberValue int
	numberSetStrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	for k := range numberSetStrings {
		currentNumIndex := strings.Index(cal_string, k)
		if currentNumIndex != -1 {
			if (firstNumIndex == -1) || (currentNumIndex < firstNumIndex) {
				firstNumIndex = currentNumIndex
				FirstNum = numberSetStrings[k]
			}
		}
	}

	for k := range numberSetStrings {
		currentNumIndex := strings.LastIndex(cal_string, k)
		if currentNumIndex != -1 {
			if (lastNumIndex == -1) || (currentNumIndex > lastNumIndex) {
				lastNumIndex = currentNumIndex
				LastNum = numberSetStrings[k]
			}
		}
	}

	fmt.Println("FirstNum = ", FirstNum)
	fmt.Println("LastNum = ", LastNum)
	CalibrationNumber := strconv.Itoa(FirstNum) + strconv.Itoa(LastNum)

	if i, err := strconv.Atoi(CalibrationNumber); err == nil {
		CalibrationNumberValue = i
	}
	return CalibrationNumberValue
}
