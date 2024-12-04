package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	unknown    = iota
	increasing = iota
	decreasing = iota
)

func get_reports(file *os.File) [][]int {

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		levels_raw := strings.Fields(report)

		reports = append(reports, make([]int, 0))
		for _, level_raw := range levels_raw {
			level, _ := strconv.Atoi(level_raw)
			reports[len(reports)-1] = append(reports[len(reports)-1], level)
		}
	}

	return reports
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func report_is_safe(report []int) bool {
	dir := unknown
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if dir == increasing && (diff < 1 || diff > 3) {
			return false
		} else if dir == decreasing && (diff > -1 || diff < -3) {
			return false
		} else if diff > 0 && diff <= 3 {
			dir = increasing
		} else if diff < 0 && diff >= -3 {
			dir = decreasing
		} else {
			return false
		}
	}
	return true
}

func main() {
	reports := get_reports(os.Stdin)

	safe_count := 0
	cut_safe_count := 0
	for _, report := range reports {
		if report_is_safe(report) {
			safe_count++
		} else {
			for i := range report {
				cut := make([]int, len(report))
				_ = copy(cut, report)
				cut = append(cut[:i], cut[i+1:]...)
				if report_is_safe(cut) {
					cut_safe_count++
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", safe_count)
	fmt.Printf("Part 2: %d\n", safe_count+cut_safe_count)
}
