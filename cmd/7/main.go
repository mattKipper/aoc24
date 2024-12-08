package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operator int

const (
	add         = 0
	multiply    = 1
	concatenate = 2
)

type equation struct {
	result   int
	operands []int
}

func parse_input(in *os.File) []equation {
	var equations []equation
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, ":")

		result, _ := strconv.Atoi(sides[0])
		operands_raw := strings.Split(strings.TrimSpace(sides[1]), " ")

		var operands []int
		for _, raw_op := range operands_raw {
			op, _ := strconv.Atoi(raw_op)
			operands = append(operands, op)
		}

		equations = append(equations, equation{result, operands})
	}
	return equations
}

func concat(a int, b int) int {
	pow := 10
	for b >= pow {
		pow *= 10
	}
	return a*pow + b
}

func pow(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return 1
	}
	result := a
	for i := 2; i <= b; i++ {
		result *= a
	}
	return result
}

func is_solvable(eq equation, allowed_ops []operator) bool {
	op_count := len(eq.operands) - 1
	for ops := 0; ops < pow(len(allowed_ops), op_count); ops++ {
		result := eq.operands[0]
		ops_c := ops
		for i := 0; i < op_count; i++ {
			op := operator(ops_c % len(allowed_ops))
			switch op {
			case add:
				result += eq.operands[i+1]
			case multiply:
				result *= eq.operands[i+1]
			case concatenate:
				result = concat(result, eq.operands[i+1])
			}
			ops_c /= len(allowed_ops)
			if result > eq.result {
				break
			}
		}
		if result == eq.result {
			return true
		}
	}
	return false
}

func main() {
	equations := parse_input(os.Stdin)

	solvable_sum := 0
	for _, eq := range equations {
		if is_solvable(eq, []operator{add, multiply}) {
			solvable_sum += eq.result
		}
	}
	fmt.Printf("Part 1: %d\n", solvable_sum)

	solvable_sum = 0
	for _, eq := range equations {
		if is_solvable(eq, []operator{add, multiply, concatenate}) {
			solvable_sum += eq.result
		}
	}
	fmt.Printf("Part 2: %d\n", solvable_sum)
}
