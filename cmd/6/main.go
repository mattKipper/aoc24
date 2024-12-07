package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

const (
	north = iota
	east
	south
	west
)

type point struct {
	x int
	y int
}

type guard struct {
	pos point
	dir direction
}

func parse_input(in *os.File) []string {
	var input []string
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func find_guard(grid *[]string) guard {
	for i, row := range *grid {
		for j, char := range row {
			switch char {
			case '^':
				return guard{point{i, j}, north}
			case '>':
				return guard{point{i, j}, east}
			case '<':
				return guard{point{i, j}, west}
			case 'v':
				return guard{point{i, j}, south}
			}
		}
	}
	return guard{point{-1, -1}, north}
}

func mark_grid(grid *[]string, pos point, mark rune) {
	row := []rune((*grid)[pos.x])
	row[pos.y] = mark
	(*grid)[pos.x] = string(row)
}

func blocked(grid *[]string, g guard) bool {
	if g.dir == north && g.pos.x > 0 {
		return (*grid)[g.pos.x-1][g.pos.y] == '#'
	} else if g.dir == east && g.pos.y < len((*grid)[g.pos.x])-1 {
		return (*grid)[g.pos.x][g.pos.y+1] == '#'
	} else if g.dir == south && g.pos.x < len(*grid)-1 {
		return (*grid)[g.pos.x+1][g.pos.y] == '#'
	} else if g.dir == west && g.pos.y > 0 {
		return (*grid)[g.pos.x][g.pos.y-1] == '#'
	} else {
		return false
	}
}

func rotate(dir direction) direction {
	switch dir {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}
	return north
}

func dir_marker(dir direction) rune {
	switch dir {
	case north:
		return '^'
	case east:
		return '>'
	case south:
		return 'v'
	case west:
		return '<'
	}
	return '?'
}

func count_markers(grid *[]string, marker rune) int {
	count := 0
	for _, row := range *grid {
		for _, char := range row {
			if char == marker {
				count++
			}
		}
	}
	return count
}

func inbounds(grid *[]string, pos point) bool {
	return pos.x >= 0 && pos.x < len(*grid) && pos.y >= 0 && pos.y < len((*grid)[pos.x])
}

func print_grid(grid *[]string) {
	for _, row := range *grid {
		fmt.Println(row)
	}
	fmt.Println("")
}

func in_loop(prev *map[guard]struct{}, g guard) bool {
	_, looped := (*prev)[g]
	return looped
}

func traverse(grid *[]string, g guard) bool {
	prev := make(map[guard]struct{})

	for inbounds(grid, g.pos) && !in_loop(&prev, g) {
		mark_grid(grid, g.pos, 'X')
		for blocked(grid, g) {
			g.dir = rotate(g.dir)
		}

		for !blocked(grid, g) && inbounds(grid, g.pos) && !in_loop(&prev, g) {
			prev[g] = struct{}{}
			switch g.dir {
			case north:
				g.pos.x--
			case east:
				g.pos.y++
			case south:
				g.pos.x++
			case west:
				g.pos.y--
			}
			if inbounds(grid, g.pos) {
				mark_grid(grid, g.pos, 'X')
			}
		}
	}
	return in_loop(&prev, g)
}

func main() {
	original_grid := parse_input(os.Stdin)
	g := find_guard(&original_grid)

	grid := make([]string, len(original_grid))
	copy(grid, original_grid)

	traverse(&grid, g)
	fmt.Printf("Part 1: %d\n", count_markers(&grid, 'X'))

	loop_count := 0
	for i, row := range grid {
		for j, mark := range row {
			// If the tile was not hit in the original traversal, it won't
			// be hit by adding an extra barrier there. We can skip those.
			if mark == 'X' {
				modified_grid := make([]string, len(original_grid))
				copy(modified_grid, original_grid)
				mark_grid(&modified_grid, point{i, j}, '#')
				if traverse(&modified_grid, g) {
					loop_count++
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", loop_count)

}
