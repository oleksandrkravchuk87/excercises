package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type matrix [][]int

func New(s string) (matrix, error) {
	var err error
	rows := strings.Split(s, "\n")
	m := make([][]int, len(rows))

	rowlen := 0
	for j := 0; j < len(rows); j++ {

		col := strings.Split(strings.TrimSpace(rows[j]), " ")

		if rowlen != 0 && len(col) != rowlen {
			return nil, errors.New("uneven")
		}
		rowlen = len(col)

		m[j] = make([]int, len(col))
		for i := 0; i < len(col); i++ {

			m[j][i], err = strconv.Atoi(col[i])

			if err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

func (m matrix) Rows() [][]int {
	mcopy := make([][]int, len([][]int(m)))
	copy(mcopy, m)
	for k := 0; k < len(mcopy); k++ {
		mcopy[k] = make([]int, len(m[k]))
		copy(mcopy[k], m[k])
	}
	return mcopy
}

func (m matrix) Cols() [][]int {
	mcopy := make([][]int, len(m[0]))
	for k := 0; k < len(mcopy); k++ {
		mcopy[k] = make([]int, len(m))
	}

	for j := 0; j < len(m); j++ {
		for i := 0; i < len(m[j]); i++ {
			mcopy[i][j] = m[j][i]
		}
	}
	return mcopy
}

func main() {
	s := "9 8 7\n19 18 17"
	m, _ := New(s)
	fmt.Println(m)

	c := m.Rows()
	fmt.Println(c)

	c2 := m.Cols()
	fmt.Println(c2)

	m.Set(1, 2, 71)

}

func (m matrix) Set(row, col, val int) bool {

	if row > len(m)-1 || col > len(m[0])-1 || row < 0 || col < 0 {
		return false
	}
	m[row][col] = val
	return true
}
