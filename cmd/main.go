package main

import (
	"fmt"

	"github.com/saulomarx/AoC2025/internal/days/day01"
	utils "github.com/saulomarx/AoC2025/pkg"
)

func main() {
	workingDay := 1

	fmt.Printf("Day %d\n", workingDay)
	switch workingDay {
	case 1:
		runsDay01()
	default:
		fmt.Println("not today")
	}

}

func runsDay01() {
	day := "day01"
	rootPath := fmt.Sprintf("./internal/days/%s/inputs/", day)
	// p := fmt.Sprintf("%s%s", rootPath, "sample.txt")
	p := fmt.Sprintf("%s%s", rootPath, "in01.txt")
	input := utils.ReadLines(p)
	day01.Part01(input)
	day01.Part02(input)
}
