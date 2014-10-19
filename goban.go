package goban

import (
	"fmt"
	"strings"
)

type goban struct {
	columns []string
	rows    [][]string
	maxes   []int
}

func newGoban(columns []string, rows [][]string) goban {
	var maxes []int
	for i, c := range columns {
		maxes = append(maxes, maxColumnLength(i, c, rows))
	}
	goban := goban{
		columns: columns,
		rows:    rows,
		maxes:   maxes,
	}
	return goban
}

func maxColumnLength(i int, column string, rows [][]string) int {
	max := len(column)
	for _, row := range rows {
		if l := len(row[i]); l > max {
			max = l
		}
	}
	return max
}

func (g goban) render() {
	g.renderLine()
	g.renderColumn(g.columns)
	g.renderLine()
	for _, row := range g.rows {
		g.renderColumn(row)
	}
	g.renderLine()
}

func (g goban) renderLine() {
	for i, _ := range g.columns {
		fmt.Printf("+%s", strings.Repeat("-", g.maxes[i]+2))
	}
	fmt.Printf("+\n")
}

func (g goban) renderColumn(columns []string) {
	for i, c := range columns {
		fmt.Printf(fmt.Sprintf("| %%-%ds ", g.maxes[i]), c)
	}
	fmt.Printf("|\n")
}

func Render(columns []string, values [][]string) {

	goban := newGoban(columns, values)
	goban.render()

}
