package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	north = iota
	northeast
	east
	southeast
	south
	southwest
	west
	northwest
)

func count_pattern(pattern string, grid []string, row int, col int, dir int) int {
	if len(pattern) == 0 {
		return 1
	}

	if row < 0 || row == len(grid) || col < 0 || col == len(grid[row]) || pattern[0] != grid[row][col] {
		return 0
	}

	switch dir {
	case north:
		row--
	case northeast:
		row--
		col++
	case east:
		col++
	case southeast:
		row++
		col++
	case south:
		row++
	case southwest:
		row++
		col--
	case west:
		col--
	case northwest:
		row--
		col--
	}

	return count_pattern(pattern[1:], grid, row, col, dir)
}

func is_xmas_origin(grid []string, row int, col int) bool {
	if row+2 >= len(grid) || col+2 >= len(grid[row]) {
		return false
	}

	nw := grid[row][col]
	ne := grid[row][col+2]
	mid := grid[row+1][col+1]
	sw := grid[row+2][col]
	se := grid[row+2][col+2]

	return mid == 'A' && ((nw == 'M' && se == 'S') || (nw == 'S' && se == 'M')) && ((ne == 'M' && sw == 'S') || (ne == 'S' && sw == 'M'))
}

func main() {
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	dirs := []int{
		north,
		northeast,
		east,
		southeast,
		south,
		southwest,
		west,
		northwest,
	}

	count_1 := 0
	count_2 := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range dirs {
				count_1 += count_pattern("XMAS", grid, i, j, dir)
			}
			if is_xmas_origin(grid, i, j) {
				count_2++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count_1)
	fmt.Printf("Part 2: %d\n", count_2)
}
