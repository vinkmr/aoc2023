package main

import (
	"aoc2023/go/utils"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	gameIDNumber, minRed, minGreen, minBlue, maxRed, maxGreen, maxBlue, totalReds, totalGreens, totalBlues int
	rounds                                                                                                 []string
}

func (g game) numOfRounds() int {
	return len(g.rounds)
}

func main() {
	// Read input and load each line as array element
	fileLines := utils.ReadFile("/home/rrb/Projects/aoc2023/inputs/day2")

	// TODO: Squeeze parsed values into struct
	var games []game //?

	for num, line := range fileLines {
		gameIDNumber, maxRed, maxGreen, maxBlue, reds, greens, blues, roundScores := ParseGameDetails(line)
		games = append(games, game{gameIDNumber, maxRed, maxGreen, maxBlue, reds, greens, blues, roundScores})
		fmt.Println("Game ", gameIDNumber, ":")
		fmt.Println(games[num])
	}

	// Solution to Part - 1
	// Sum of Game IDs where games possible with 12 red, 13 green, and 14 blue cubes
	var sumOfGameIDs int
	for _, game := range games {
		if game.maxRed <= 12 && game.maxGreen <= 13 && game.maxBlue <= 14 {
			sumOfGameIDs += game.gameIDNumber
		}
	}
	fmt.Println("Solution to Part 1 = ", sumOfGameIDs)
}

func getGameIDNumber(GameID string) string {
	indexSplit := strings.Split(GameID, " ")
	return indexSplit[len(indexSplit)-1]
}

func ParseGameDetails(GameDetailsLine string) (int, int, int, int, int, int, int, []string) {
	var GameID int
	var Rounds []string
	var MinRed int
	var MinGreen int
	var MinBlue int
	var MaxRed int
	var MaxGreen int
	var MaxBlue int
	var TotalReds int
	var TotalGreens int
	var TotalBlues int
	// var GameDetails game

	gameIDMarker, gameScores, found := strings.Cut(GameDetailsLine, ":")

	if found {
		GameIDString := getGameIDNumber(gameIDMarker)
		GameIDInt, err := strconv.Atoi(GameIDString)
		if err != nil {
			fmt.Println(err)
		} else {
			GameID = GameIDInt
		}

		Rounds = strings.Split(gameScores, ";")
		fmt.Println("Rounds :")
		for _, round := range Rounds {
			fmt.Println(round)
			red, green, blue := splitByColor(round)

			// Find Highest Value drawn in game for each color
			if red > MaxRed {
				MaxRed = red
			}
			if green > MaxGreen {
				MaxGreen = green
			}
			if blue > MaxBlue {
				MaxBlue = blue
			}
			// Find Lowest Value drawn in game for each color
			if red < MinRed {
				MinRed = red
			}
			if green < MinGreen {
				MinGreen = green
			}
			if blue < MinBlue {
				MinBlue = blue
			}

			// Find cumlative value of drawn cubes in game for each color
			TotalReds += red
			TotalGreens += green
			TotalBlues += blue
		}
	} else {
		fmt.Println("Input file seems to be corrupted. Empty lines found.")
	}
	return GameID, MaxRed, MaxGreen, MaxBlue, TotalReds, TotalGreens, TotalBlues, Rounds
}

func splitByColor(RoundDetails string) (int, int, int) {
	var Red int
	var Green int
	var Blue int
	// Strip leading and trailing whitespace
	trimmedRoundDetails := strings.TrimSpace(RoundDetails)
	draws := strings.Split(trimmedRoundDetails, ",")
	for _, draw := range draws {
		trimmedDraw := strings.TrimSpace(draw)
		scoreAndColorSplit := strings.Split(trimmedDraw, " ")
		if len(scoreAndColorSplit) > 2 {
			panic("Score and Color for the selected draw could not be determined. Unwanted sequences in string.")
		}
		score := scoreAndColorSplit[0]
		color := scoreAndColorSplit[1]
		switch color {
		case "red":
			redScore, err := strconv.Atoi(score)
			if err != nil {
				fmt.Println(err)
			} else {
				Red = redScore
			}
		case "green":
			greenScore, err := strconv.Atoi(score)
			if err != nil {
				fmt.Println(err)
			} else {
				Green = greenScore
			}
		case "blue":
			blueScore, err := strconv.Atoi(score)
			if err != nil {
				fmt.Println(err)
			} else {
				Blue = blueScore
			}
		default:
			fmt.Println("Found something weird for color value: ", color)
		}

	}
	return Red, Green, Blue
}
