package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, _ := os.ReadFile("./input/day2.txt")
	fmt.Println("[day 2 - part 1] - total rps points:", day2First(string(rawInput)))
	fmt.Println("[day 2 - part 2] - total rps points:", day2Second(string(rawInput)))

	rawInput, _ = os.ReadFile("./input/day1.txt")
	fmt.Println("day 1 - part 2:", day1First(string(rawInput)))
}

func day2Second(input string) int {
	totalPoints := 0
	for _, line := range strings.Split(input, "\n") {
		roundOptions := strings.Split(line, " ")
		if len(roundOptions) != 2 {
			continue
		}
		yourOption := stringToRpsOption(roundOptions[0])
		var myOutcome RspOutcome = roundOptions[1]
		var myOption RpsOption

		switch myOutcome {
			case OutComeDraw: myOption = yourOption
			case OutcomeWin: 
				switch yourOption {
					case Rock: myOption = Paper
					case Paper: myOption = Scissors
					case Scissors: myOption = Rock
				}
			case OutcomeLose:
				switch yourOption {
					case Rock: myOption = Scissors
					case Paper: myOption = Rock
					case Scissors: myOption = Paper
				}
		}

		roundPoints := playRps(yourOption, myOption)
		totalPoints += int(roundPoints)
	}
	return totalPoints
}

type RspOutcome = string

const (
	OutcomeLose RspOutcome = "X"
	OutComeDraw RspOutcome = "Y"
	OutcomeWin RspOutcome = "Z"
)

func day2First(input string) int {
	totalPoints := 0
	for _, line := range strings.Split(input, "\n") {
		roundOptions := strings.Split(line, " ")
		if len(roundOptions) != 2 {
			continue
		}

		yourOption := stringToRpsOption(roundOptions[0])
		myOption := stringToRpsOption(roundOptions[1])

		roundPoints := playRps(yourOption, myOption)
		totalPoints += int(roundPoints)
	}
	return totalPoints
}

// rock beats scissor
// scissor beats paper
// paper beats rock

type RpsResult int8
type RpsOption int8

const (
	Lose RpsResult = 0
	Draw RpsResult = 3
	Win RpsResult = 6
)

const (
	Rock RpsOption = 1
	Paper RpsOption = 2
	Scissors RpsOption = 3
)

func stringToRpsOption(input string) RpsOption {
	if (input == "A" || input == "X") {
		return Rock
	} else if input == "B" || input == "Y" {
		return Paper
	} else {
		return Scissors
	}
}

func playRps(yourOption RpsOption, myOption RpsOption) int8 {
	var outcome RpsResult

	if myOption == yourOption {
		outcome = Draw
	} else if (myOption == Rock && yourOption == Scissors) || (myOption == Scissors && yourOption == Paper) || (myOption == Paper && yourOption == Rock) {
		outcome = Win
	} else {
		outcome = Lose
	}

	return int8(myOption) + int8(outcome)
}

func day1First(input string) int {
	inputSlice := strings.Split(input, "\n")
	topThreeCalories := make([]int, 3)

	currentElfCalories := 0

	for _, line := range inputSlice {
		if line != "" {
			cals, _ := strconv.Atoi(line)
			currentElfCalories += cals
		} else {
			lowestCaloryIndex := -1
			for i := 0; i < 3; i++ {
				// (of the top 3)
				var currentLowestCalory int
				if lowestCaloryIndex == -1 {
					currentLowestCalory = math.MaxInt
				} else {
					currentLowestCalory = topThreeCalories[lowestCaloryIndex]
				}
				if topThreeCalories[i] <= currentLowestCalory {
					lowestCaloryIndex = i
				}
			}
			if topThreeCalories[lowestCaloryIndex] < currentElfCalories {
				topThreeCalories[lowestCaloryIndex] = currentElfCalories
			}
			currentElfCalories = 0
		}
	}

	sumOfTop3Cals := 0
	for i := 0; i < 3; i++ {
		sumOfTop3Cals += topThreeCalories[i]
	}
	return sumOfTop3Cals 
}
