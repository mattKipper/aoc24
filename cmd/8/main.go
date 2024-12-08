package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	float_eq_thresh = 1e-9
	cmp_part_1      = 0
	cmp_part_2      = 1
)

type cmp_type int

type point struct {
	row int
	col int
}

type grid struct {
	width     int
	height    int
	antennas  map[rune][]point
	antinodes map[point]struct{}
}

func parse_input(in *os.File) grid {
	antennas := make(map[rune][]point)
	width, height := 0, 0

	scanner := bufio.NewScanner(in)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, frequency := range line {
			if frequency != '.' {
				antenna, exists := antennas[frequency]
				if exists {
					antennas[frequency] = append(antenna, point{i, j})
				} else {
					antennas[frequency] = []point{{i, j}}
				}
			}
			width = j + 1
		}
		height = i + 1
	}

	antinodes := make(map[point]struct{})
	return grid{width, height, antennas, antinodes}
}

func print_grid(g grid) {
	var out []string
	for range g.height {
		out = append(out, strings.Repeat(".", g.width))
	}

	for freq, locations := range g.antennas {
		for _, location := range locations {
			updated := []rune(out[location.row])
			updated[location.col] = freq
			out[location.row] = string(updated)
		}
	}

	for location := range g.antinodes {
		updated := []rune(out[location.row])
		updated[location.col] = '#'
		out[location.row] = string(updated)
	}

	for _, row := range out {
		fmt.Println(row)
	}
}

func f64_eq(a, b float64) bool {
	return math.Abs(a-b) <= float_eq_thresh
}

func distance(a point, b point) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(a.row-b.row)), 2) + math.Pow(math.Abs(float64(a.col-b.col)), 2))
}

func slope(a point, b point) float64 {
	return float64(b.row-a.row) / float64(b.col-a.col)
}

func is_antinode(antinode point, a point, b point, cmp cmp_type) bool {
	ad := distance(antinode, a)
	bd := distance(antinode, b)

	as := slope(antinode, a)
	bs := slope(antinode, b)

	slope_match := f64_eq(as, bs)
	if cmp == cmp_part_1 {
		return slope_match && (f64_eq(ad, 2*bd) || f64_eq(bd, 2*ad))
	} else {
		return slope_match
	}
}

func add_pair_antinodes(a point, b point, g *grid, cmp cmp_type) {
	for i := range g.height {
		for j := range g.width {
			if is_antinode(point{i, j}, a, b, cmp) {
				g.antinodes[point{i, j}] = struct{}{}
			}
		}
	}
}

func add_frequency_antinodes(antennas *[]point, g *grid, cmp cmp_type) {
	for i, ant := range *antennas {
		for j := i + 1; j < len(*antennas); j++ {
			add_pair_antinodes(ant, (*antennas)[j], g, cmp)
		}
		if cmp == cmp_part_2 {
			g.antinodes[ant] = struct{}{}
		}
	}
}

func add_antinodes(g *grid, cmp cmp_type) {
	for _, locations := range g.antennas {
		add_frequency_antinodes(&locations, g, cmp)
	}
}

func main() {
	grid := parse_input(os.Stdin)

	add_antinodes(&grid, cmp_part_1)
	fmt.Printf("Part 1: %d\n", len(grid.antinodes))

	grid.antinodes = make(map[point]struct{})
	add_antinodes(&grid, cmp_part_2)
	fmt.Printf("Part 2: %d\n", len(grid.antinodes))

}
