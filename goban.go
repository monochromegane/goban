package goban

import (
	"fmt"
	"strings"
)

func Render(columns []string, values [][]string) {

	var table table
	for i, c := range columns {
		max := getMaxLength(i, c, values)
		table.maxes = append(table.maxes, max)
	}

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

	for _, v := range values {
		for i, c := range v {
			fmt.Printf("|")
			fmt.Printf(fmt.Sprintf(" %%-%ds ", table.maxes[i]), c)
		}
		fmt.Printf("|\n")
	}
	for i, _ := range columns {
		fmt.Printf("+")
		fmt.Printf(strings.Repeat("-", table.maxes[i]+2))
	}
	fmt.Printf("+\n")

}

func getMaxLength(i int, c string, values [][]string) int {
	max := len(c)
	for _, v := range values {
		if l := len(v[i]); l > max {
			max = l
		}
	}
	return max
}

type table struct {
	maxes []int
}
