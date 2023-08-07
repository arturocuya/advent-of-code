package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputRaw, _ := os.ReadFile("./input.txt")
	
	part1 := part1(strings.Split(string(inputRaw), "\n"))
	part2 := part2(strings.Split(string(inputRaw), "\n"))
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

func part2(lines []string) int {
	var badgesSum int = 0

	for i := 0; i < len(lines); i += 3 {
		sack1 := lines[i]
		if len(sack1) == 0 {
			continue
		}

		sack2 := lines[i + 1]
		sack3 := lines[i + 2]

		out:
		for _, sack1Item := range sack1 {
			for _, sack2Item := range sack2 {
				for _, sack3Item := range sack3 {
					if sack1Item == sack2Item && sack2Item == sack3Item {
						if unicode.IsUpper(sack1Item) {
							badgesSum += int(sack1Item - 'A' + 27)
						} else {
							badgesSum += int(sack1Item - 'a' + 1)
						}
						break out
					}
				}
			}
		}
	}

	return badgesSum
}

func part1(lines []string) int {
	var prioritySum int = 0
	for _, line := range lines {
		if (len(line) == 0) {
			continue
		}

		sack1 := line[0:len(line)/2]
		sack2 := line[len(line)/2:]

		out:
		for _, sack1Item := range sack1 {
			for _, sack2Item := range sack2 {
				if (sack1Item == sack2Item) {
					if unicode.IsUpper(sack1Item) {
						prioritySum += int(sack1Item - 'A' + 27)
					} else {
						prioritySum += int(sack1Item - 'a' + 1)
					}
					break out
				}
			}
		}
	}

	return prioritySum
}
