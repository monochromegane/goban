package main

import (
	"github.com/monochromegane/goban"
)

func main() {
	columns := []string{"columnA", "columnBBBB"}
	values := [][]string{
		[]string{"valueA1-laskdfas", "valueB1"},
		[]string{"valueA2", "valueB2"},
	}
	goban.Render(columns, values)
}
