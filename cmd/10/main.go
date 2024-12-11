package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	unknown = -1
)

type peak_set map[*node]struct{}

type node struct {
	height int
	peaks  peak_set
	trails int
	next   []*node
}

func parse_input(in *os.File) [][]node {
	var topo [][]node
	scanner := bufio.NewScanner(in)
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		topo = append(topo, []node{})
		for _, height := range line {
			height := int(height) - '0'
			peaks := make(peak_set)
			trails := unknown
			next := []*node{}
			topo[row] = append(topo[row], node{height, peaks, trails, next})
		}
	}

	rows, cols := len(topo), len(topo[0])
	for i := range topo {
		for j := range topo[i] {
			node := &topo[i][j]
			if i > 0 && (topo[i-1][j].height == node.height+1) {
				node.next = append(node.next, &topo[i-1][j])
			}
			if j > 0 && (topo[i][j-1].height == node.height+1) {
				node.next = append(node.next, &topo[i][j-1])
			}
			if (i < rows-1) && (topo[i+1][j].height == node.height+1) {
				node.next = append(node.next, &topo[i+1][j])
			}
			if (j < cols-1) && (topo[i][j+1].height == node.height+1) {
				node.next = append(node.next, &topo[i][j+1])
			}
		}
	}
	return topo
}

func find_peaks(cur *node) {
	if cur.trails != unknown {
		return
	}

	if cur.height == 9 {
		cur.peaks[cur] = struct{}{}
		cur.trails = 1
	} else {
		cur.trails = 0
		for _, next := range cur.next {
			if next.height == cur.height+1 {
				find_peaks(next)
				for peak := range next.peaks {
					cur.peaks[peak] = struct{}{}
				}
				cur.trails += next.trails
			}
		}
	}
}

func traverse(topo *[][]node) (int, int) {
	sum := 0
	trails := 0
	for i := range *topo {
		for j := range (*topo)[i] {
			node := &(*topo)[i][j]
			if node.height == 0 {
				find_peaks(node)
				sum += len(node.peaks)
				trails += node.trails
			}
		}
	}
	return sum, trails
}

func main() {
	topo := parse_input(os.Stdin)
	sum, trails := traverse(&topo)
	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", trails)
}
