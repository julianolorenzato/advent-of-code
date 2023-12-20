package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartOne() int32 {
	var sum int32 = 0

	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		value := decipher(scanner.Text())
		sum += int32(value)
	}

	return sum
}

func decipher(str string) int32 {
	var firstDigit, lastDigit rune

	for _, c := range str {
		if c >= '0' && c <= '9' {
			if firstDigit == 0 {
				firstDigit = c
			}

			lastDigit = c
		}
	}

	formatted := fmt.Sprintf("%c%c", firstDigit, lastDigit)

	value, err := strconv.Atoi(formatted)
	if err != nil {
		panic(err)
	}

	return int32(value)
}
