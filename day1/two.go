package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartTwo() int32 {
	var sum int32 = 0

	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		value := betterDecipher(scanner.Text())
		sum += int32(value)
	}

	return sum
}

func betterDecipher(str string) int32 {
	var firstDigit, lastDigit rune

	algarisms := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	// find first digit
	for i, c := range str {
		if c >= '0' && c <= '9' {
			firstDigit = c
			break
		}

		if i+3 <= len(str) {
			substring := str[i : i+3]
			if _, ok := algarisms[substring]; ok {
				firstDigit = algarisms[substring]
				break
			}
		}

		if i+4 <= len(str) {
			substring := str[i : i+4]
			if _, ok := algarisms[substring]; ok {
				firstDigit = algarisms[substring]
				break
			}
		}

		if i+5 <= len(str) {
			substring := str[i : i+5]
			if _, ok := algarisms[substring]; ok {
				firstDigit = algarisms[substring]
				break
			}
		}
	}

	// find last digit
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] >= '0' && str[i] <= '9' {
			lastDigit = rune(str[i])
			break
		}

		if i-2 >= 0 {
			substring := str[i-2 : i+1]
			if _, ok := algarisms[substring]; ok {
				lastDigit = algarisms[substring]
				break
			}
		}

		if i-3 >= 0 {
			substring := str[i-3 : i+1]
			if _, ok := algarisms[substring]; ok {
				lastDigit = algarisms[substring]
				break
			}
		}

		if i-4 >= 0 {
			substring := str[i-4 : i+1]
			if _, ok := algarisms[substring]; ok {
				lastDigit = algarisms[substring]
				break
			}
		}
	}

	formatted := fmt.Sprintf("%c%c", firstDigit, lastDigit)

	value, err := strconv.Atoi(formatted)
	if err != nil {
		panic(err)
	}

	return int32(value)
}
