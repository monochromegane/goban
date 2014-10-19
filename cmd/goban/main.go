package main

import (
	"github.com/monochromegane/goban"
)

func main() {
	columns := []string{"columnA", "columnBBBB"}
	goban.Render(columns)
}
