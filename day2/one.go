package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decodeLine(line string) (gameID, red, green, blue int) {
	line = strings.ReplaceAll(line, ",", "")

	splitted := strings.Split(line, ":")

	header := splitted[0]
	data := splitted[1]

	gameID, err := strconv.Atoi(strings.Fields(header)[1])
	if err != nil {
		panic(err)
	}

	rounds := strings.Split(data, ";")
	for _, round := range rounds {

		results := strings.Fields(round)
		for i, result := range results {
			if result == "red" {
				if n, _ := strconv.Atoi(results[i-1]); n > 20 {

				}
			}
		}
	}

	return
}

func PartOne() int32 {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		decodeLine(line)
	}

	return -2
}
