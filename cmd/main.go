package main

import (
	"fmt"
)

func main() {
	workingDay := 1

	fmt.Printf("Day %d\n", workingDay)
	switch workingDay {
	default:
		fmt.Println("not today")
	}

}
