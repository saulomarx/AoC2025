package day01

import (
	"fmt"
	"strconv"
)

func Part01(input []string) {
	sum := 50
	countZero := 0
	for _, line := range input {
		d, v := getDirectionAndValue(line)
		// min 0
		if d == "L" {
			sum -= v % 100
			if sum < 0 {
				sum += 100
			}
		}
		// max 99
		if d == "R" {
			sum += v % 100
			if sum > 99 {
				sum -= 100
			}
		}

		if sum == 0 {
			countZero++
		}

	}
	fmt.Println(countZero)
}

func Part02(input []string) {
	sum := 50
	countZero := 0
	for _, line := range input {
		d, v := getDirectionAndValue(line)
		countZero += v / 100
		lastValue := sum

		// min 0
		if d == "L" {
			sum -= v % 100
			if sum < 0 {
				sum += 100
				if lastValue != 0 {
					countZero++
				}
			}
		}
		// max 99
		if d == "R" {
			sum += v % 100
			if sum > 99 {
				sum -= 100
				if sum != 0 {
					countZero++
				}
			}
		}

		if sum == 0 {
			countZero++
		}

	}
	fmt.Println(countZero)
}

func getDirectionAndValue(s string) (direction string, value int) {
	direction = s[:1]
	v, _ := strconv.Atoi(s[1:])
	value = v
	return
}
