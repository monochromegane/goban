package goban

import "testing"

func ExampleRender() {
	Render(
		[]string{"columnA", "カラムあ"},
		[][]string{
			[]string{"valueA", "値あ"},
			[]string{"valueB-123", "値い"},
		},
	)

	// Output:
	// +------------+----------+
	// | columnA    | カラムあ |
	// +------------+----------+
	// | valueA     | 値あ     |
	// | valueB-123 | 値い     |
	// +------------+----------+
}

func TestDisplayLength(t *testing.T) {

	type assert struct {
		str string
		len int
	}

	asserts := []assert{
		assert{"a", 1},
		assert{"あ", 2},
		assert{"aあ", 3},
	}

	for _, a := range asserts {
		if l := displayLength(a.str); l != a.len {
			t.Errorf("displayLength(%s) should return %d, got %d", a.str, a.len, l)
		}
	}
}

func TestNewGoban(t *testing.T) {

	type assert struct {
		columns []string
		rows    [][]string
		maxes   []int
	}

	a := assert{
		[]string{"a", "あ"},
		[][]string{
			[]string{"aa", "あい"},
			[]string{"aaa", "あいう"},
		},
		[]int{3, 6},
	}

	g := newGoban(a.columns, a.rows)
	for i, m := range g.maxes {
		if m != a.maxes[i] {
			t.Errorf("goban.maxes should return %v, got %v", a.maxes[i], m)
		}

	}
}
