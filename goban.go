package goban

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Render(columns []string, rows [][]string) {
	goban := newGoban(columns, rows)
	goban.render()
}

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
		fmt.Printf(
			"| %s%s ",
			c,
			strings.Repeat(" ", g.maxes[i]-displayLength(c)),
		)
	}
	fmt.Printf("|\n")
}

func maxColumnLength(i int, column string, rows [][]string) int {
	max := displayLength(column)
	for _, row := range rows {
		if l := displayLength(row[i]); l > max {
			max = l
		}
	}
	return max
}

func displayLength(str string) int {
	length := 0
	for _, s := range str {
		if l := utf8.RuneLen(s); l > 1 {
			length += 2
		} else {
			length++
		}
	}
	return length
}
