package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parse_input(in *os.File) []uint64 {
	line_raw, _ := io.ReadAll(in)
	line := string(line_raw[:])

	stones_raw := strings.Split(line, " ")

	var stones []uint64
	for _, stone := range stones_raw {
		stone_val, _ := strconv.Atoi(stone)
		stones = append(stones, uint64(stone_val))
	}
	return stones
}

func digits(n uint64) int {
	digits := 1
	for n /= 10; n > 0; n /= 10 {
		digits++
	}
	return digits
}

func slice(n uint64, digits int) (uint64, uint64) {
	right := uint64(0)

	mul := uint64(1)
	for range digits / 2 {
		right += (n % 10) * mul
		mul *= 10
		n /= 10
	}

	left := uint64(0)
	mul = 1
	for range digits / 2 {
		left += (n % 10) * mul
		mul *= 10
		n /= 10
	}

	return left, right
}

func blink_stone(stone uint64) []uint64 {
	out := make([]uint64, 0, 2)
	if stone == 0 {
		out = append(out, 1)
	} else {
		digits := digits(stone)
		if digits%2 == 0 {
			left, right := slice(stone, digits)
			out = append(out, left, right)
		} else {
			out = append(out, stone*2024)
		}
	}
	return out
}

func blink_stones(counts_in map[uint64]uint64, transform *map[uint64][]uint64) map[uint64]uint64 {
	counts_out := make(map[uint64]uint64, len(counts_in)*2)

	for stone_in, count_in := range counts_in {
		blinked, has_xfrm := (*transform)[stone_in]
		if !has_xfrm {
			blinked = blink_stone(stone_in)
			(*transform)[stone_in] = blinked
		}
		for _, stone_out := range blinked {
			counts_out[stone_out] = counts_out[stone_out] + count_in
		}
	}

	return counts_out
}

func main() {
	stones_in := parse_input(os.Stdin)

	stone_counts := make(map[uint64]uint64)
	for _, stone := range stones_in {
		stone_counts[stone] = stone_counts[stone] + 1
	}

	transform := make(map[uint64][]uint64)
	for range 25 {
		stone_counts = blink_stones(stone_counts, &transform)
	}

	stones_out := uint64(0)
	for _, stone_count := range stone_counts {
		stones_out += stone_count
	}
	fmt.Printf("Part 1: %d\n", stones_out)

	for range 50 {
		stone_counts = blink_stones(stone_counts, &transform)
	}

	stones_out = uint64(0)
	for _, stone_count := range stone_counts {
		stones_out += stone_count
	}
	fmt.Printf("Part 2: %d\n", stones_out)

}
