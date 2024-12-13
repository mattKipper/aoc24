package main

import (
	"bufio"
	"fmt"
	"os"
)

type graph struct {
	nodes int
	edges int
	adj   [][]int
}

func new_graph(nodes int) graph {
	adj := make([][]int, nodes)
	return graph{nodes, 0, adj}
}

func (g *graph) degree(v int) int {
	return len(g.adj[v])
}

func (g *graph) add_edge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.edges++
}

func dfs_cc(g *graph, v int, m *[]bool, c *[]int, cc int) {
	(*m)[v] = true
	(*c)[v] = cc
	for _, w := range g.adj[v] {
		if !(*m)[w] {
			dfs_cc(g, w, m, c, cc)
		}
	}
}

func (g *graph) connected_components() (int, []int) {
	marked := make([]bool, g.nodes)
	component := make([]int, g.nodes)
	count := 0

	for v := range g.nodes {
		if !marked[v] {
			dfs_cc(g, v, &marked, &component, count)
			count++
		}
	}

	return count, component
}

func parse_input(in *os.File) [][]rune {
	var plots [][]rune
	scanner := bufio.NewScanner(in)
	for row := 0; scanner.Scan(); row++ {
		plots = append(plots, []rune{})
		line := scanner.Text()
		for _, plant := range line {
			plots[row] = append(plots[row], plant)
		}
	}

	return plots
}

func plots_graph(plots [][]rune) graph {
	rows := len(plots)
	cols := len(plots[0])
	garden := new_graph(rows * cols)

	for i := range plots {
		for j, r := range plots[i] {
			index := (i * cols) + j
			if (i > 0) && (r == plots[i-1][j]) {
				garden.add_edge(index, index-cols)
			}
			if (j > 0) && (r == plots[i][j-1]) {
				garden.add_edge(index, index-1)
			}
		}
	}
	return garden
}

func main() {
	plots := parse_input(os.Stdin)
	garden := plots_graph(plots)

	// We can solve for both perimeter and area of each
	// region by finding connected components in the
	// garden graph.
	//
	//   - Since each component represents on region, the area
	//     of a region matches the number of nodes in its component
	//
	//   - Each node has a perimeter of 4 minus its degree, so the
	//     perimeter of a region can be calculated by applying this
	//     calculation to each node in the component
	_, components := garden.connected_components()

	area, perimeter := make(map[int]int), make(map[int]int)
	for v, c := range components {
		area[c] = area[c] + 1
		perimeter[c] = perimeter[c] + 4 - garden.degree(v)
	}

	price := make(map[int]int)
	for _, c := range components {
		price[c] = area[c] * perimeter[c]
	}

	total_price := 0
	for _, p := range price {
		total_price += p
	}

	fmt.Printf("Part 1: %d\n", total_price)
}
