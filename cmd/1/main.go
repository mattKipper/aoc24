package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func get_ids(file *os.File) ([]int, []int) {
	var ids [2][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		next_ids := strings.Fields(line)

		id_a, _ := strconv.Atoi(next_ids[0])
		id_b, _ := strconv.Atoi(next_ids[1])

		ids[0] = append(ids[0], id_a)
		ids[1] = append(ids[1], id_b)
	}

	return ids[0], ids[1]
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	id_a, id_b := get_ids(os.Stdin)

	sort.Ints(id_a)
	sort.Ints(id_b)

	distance := 0
	for i := range id_a {
		distance += abs(id_a[i] - id_b[i])
	}
	fmt.Printf("Part 1: %d\n", distance)

	similarity := 0
	for i, j := 0, 0; i < len(id_a) && j < len(id_b); {
		if id_a[i] < id_b[j] {
			i++
		} else if id_a[i] > id_b[j] {
			j++
		} else {
			similarity += id_a[i]
			j++
		}
	}
	fmt.Printf("Part 2: %d\n", similarity)
}
