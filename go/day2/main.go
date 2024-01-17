package main

import (
	"aoc2023/go/utils"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	gameNumber, reds, greens, blues int
	rounds                          []string
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
		gameNumber, reds, greens, blues, roundScores := getGameDetails(line)
		games = append(games, game{gameNumber, reds, greens, blues, roundScores})
		fmt.Println("Game ", gameNumber, ":")
		fmt.Println(games[num])
		for _, roundScore := range roundScores {
			splitByColor(roundScore)
		}
	}
}

func getGameNumber(GameNumber string) string {
	indexSplit := strings.Split(GameNumber, " ")
	return indexSplit[len(indexSplit)-1]
}

func getGameDetails(RoundDetails string) (int, int, int, int, []string) {
	var GameNumber int
	var Rounds []string
	var Reds int
	var Greens int
	var Blues int
	gameNumber, gameScores, found := strings.Cut(RoundDetails, ":")

	if found {
		GameNumberVal := getGameNumber(gameNumber)
		GameNumberInt, err := strconv.Atoi(GameNumberVal)
		if err != nil {
			fmt.Println(err)
		} else {
			GameNumber = GameNumberInt
		}

		Rounds = strings.Split(gameScores, ";")
		fmt.Println("Rounds :")
		for _, round := range Rounds {
			fmt.Println(round)
			red, green, blue := splitByColor(round)
			Reds += red
			Greens += green
			Blues += blue
		}
	} else {
		fmt.Println("Input file seems to be corrupted. Empty lines found.")
	}
	return GameNumber, Reds, Greens, Blues, Rounds
}

func splitByColor(RoundDetails string) (int, int, int) {
	var Red int
	var Green int
	var Blue int
	// var trimmedRounds []string
	fmt.Println()
	fmt.Println("Got input:", RoundDetails)
	fmt.Println("Splitting by Color.")
	// Strip leading and trailing whitespace
	trimmedRoundDetails := strings.TrimSpace(RoundDetails)
	draws := strings.Split(trimmedRoundDetails, ",")
	for _, draw := range draws {
		trimmedDraw := strings.TrimSpace(draw)
		// fmt.Println("Debug: trimmedDraw:", trimmedDraw)
		scoreAndColorSplit := strings.Split(trimmedDraw, " ")
		if len(scoreAndColorSplit) > 2 {
			// fmt.Println("Debug: scoreAndColorSplit:", scoreAndColorSplit)
			// fmt.Println("Debug: Length:", len(scoreAndColorSplit))
			panic("Score and Color for the selected draw could not be determined. Unwanted sequences in string.")
		}
		score := scoreAndColorSplit[0]
		color := scoreAndColorSplit[1]
		switch color {
		case "red":
			fmt.Println("\nFound red.")
			redScore, err := strconv.Atoi(score)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Adding %d to red score.\n", redScore)
				Red = Red + redScore
			}
		case "green":
			fmt.Println("\nFound green.")
			greenScore, err := strconv.Atoi(score)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Adding %d to green score.\n", greenScore)
				Green = Green + greenScore
			}
		case "blue":
			fmt.Println("\nFound blue.")
			blueScore, err := strconv.Atoi(score)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Adding %d to blue score.\n", blueScore)
				Blue = Blue + blueScore
			}
		default:
			fmt.Println("Found something weird for color value: ", color)
		}

	}
	// trimmedRounds = append(trimmedRounds, strings.TrimSpace(round))
	return Red, Green, Blue
}
