package main

import (
	"aoc2023/go/utils"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	gameIDNumber, maxRed, maxGreen, maxBlue, totalReds, totalGreens, totalBlues int
	rounds                                                                      []string
}

func (g game) numOfRounds() int {
	return len(g.rounds)
}

func (g game) gameScore() int {
	return g.maxRed * g.maxGreen * g.maxBlue
}

func main() {
	// Read input and load each line as array element
	fileLines := utils.ReadFile("/home/rrb/Projects/aoc2023/inputs/day2")

	// TODO: Squeeze parsed values into struct
	var games []game //?

	for _, line := range fileLines {
		game := ParseGameDetails(line)
		games = append(games, game)
		// fmt.Println("Game ", game.gameIDNumber, ":")
		// fmt.Println("Total number of rounds:", game.numOfRounds())
		// fmt.Println(games[num])
	}

	// Solution to Part - 1 & 2
	// Part 1 - Sum of Game IDs where games possible with 12 red, 13 green, and 14 blue cubes
	// Part 2 - Sum of power of games, where power is the product of the max value drawn for each color in a game
	var sumOfGameIDs int
	var sumOfGamePowers int
	for _, game := range games {
		if game.maxRed <= 12 && game.maxGreen <= 13 && game.maxBlue <= 14 {
			sumOfGameIDs += game.gameIDNumber
		}
		sumOfGamePowers += game.gameScore()
	}
	fmt.Println("Solution to Part 1 = ", sumOfGameIDs)
	fmt.Println("Solution to Part 2 = ", sumOfGamePowers)
}

func getGameIDNumber(GameID string) string {
	indexSplit := strings.Split(GameID, " ")
	return indexSplit[len(indexSplit)-1]
}

func ParseGameDetails(GameDetailsLine string) game {
	var Game game

	gameIDMarker, gameScores, found := strings.Cut(GameDetailsLine, ":")

	if found {
		GameIDString := getGameIDNumber(gameIDMarker)
		GameIDInt, err := strconv.Atoi(GameIDString)
		if err != nil {
			fmt.Println(err)
		} else {
			Game.gameIDNumber = GameIDInt
		}

		Game.rounds = strings.Split(gameScores, ";")
		// fmt.Println("Rounds :")
		for _, round := range Game.rounds {
			// fmt.Println(round)
			red, green, blue := splitByColor(round)

			// Find Highest Value drawn in game for each color
			if red > Game.maxRed {
				Game.maxRed = red
			}
			if green > Game.maxGreen {
				Game.maxGreen = green
			}
			if blue > Game.maxBlue {
				Game.maxBlue = blue
			}

			// Find cumlative value of drawn cubes in game for each color
			Game.totalReds += red
			Game.totalGreens += green
			Game.totalBlues += blue
		}
	} else {
		fmt.Println("Input file seems to be corrupted. Empty lines found.")
	}

	return Game
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
