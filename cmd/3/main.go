package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	enable  = "do()"
	disable = "don't()"
)

type mul struct {
	x int
	y int
}

func parse_instructions(memory string) []mul {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	instructions_raw := re.FindAllStringSubmatch(memory, -1)

	var instructions []mul
	for _, instruction_raw := range instructions_raw {
		x, _ := strconv.Atoi(instruction_raw[1])
		y, _ := strconv.Atoi(instruction_raw[2])
		instructions = append(instructions, mul{x, y})
	}

	return instructions
}

func mul_result_sum(memory string) int {
	instructions := parse_instructions(memory)
	result_sum := 0
	for _, instruction := range instructions {
		result_sum += (instruction.x * instruction.y)
	}
	return result_sum
}

func main() {
	memory_raw, _ := io.ReadAll(os.Stdin)
	memory := string(memory_raw[:])

	fmt.Printf("Part 1: %d\n", mul_result_sum(memory))

	enabled_result_sum := 0

	for chunk_end := strings.Index(memory, disable); chunk_end != -1; {
		memory_chunk := memory[:chunk_end]
		enabled_result_sum += mul_result_sum(memory_chunk)

		memory = memory[chunk_end+len(disable):]

		if chunk_start := strings.Index(memory, enable); chunk_start != -1 {
			memory = memory[chunk_start+len(enable):]
			chunk_end = strings.Index(memory, disable)
		} else {
			memory = ""
			break
		}
	}

	if len(memory) > 0 {
		enabled_result_sum += mul_result_sum(memory)
	}

	fmt.Printf("Part 2: %d", enabled_result_sum)
}
