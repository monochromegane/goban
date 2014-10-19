package goban

import (
	"fmt"
	"strings"
)

func Render(columns []string) {

	var table table
	for _, c := range columns {
		max := getMaxLength(c)
		table.maxes = append(table.maxes, max)
	}
	fmt.Println(table)

	for i, _ := range columns {
		fmt.Printf("+")
		fmt.Printf(strings.Repeat("-", table.maxes[i]+2))
	}
	fmt.Printf("+\n")

	for i, c := range columns {
		fmt.Printf("|")
		fmt.Printf(fmt.Sprintf(" %%-%ds ", table.maxes[i]), c)
	}
	fmt.Printf("|\n")

	for i, _ := range columns {
		fmt.Printf("+")
		fmt.Printf(strings.Repeat("-", table.maxes[i]+2))
	}
	fmt.Printf("+\n")
}

func getMaxLength(c string) int {
	return len(c)
}

type table struct {
	maxes []int
}
