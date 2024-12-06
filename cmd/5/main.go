package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

func parse_input(file *os.File) ([]rule, [][]int) {
	var rules []rule
	var updates [][]int

	rules_parsed := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !rules_parsed {
			if len(line) <= 1 {
				rules_parsed = true
			} else {
				fields := strings.Split(line, "|")
				before, _ := strconv.Atoi(fields[0])
				after, _ := strconv.Atoi(fields[1])
				rules = append(rules, rule{before, after})
			}
		} else {
			fields := strings.Split(line, ",")
			var update []int
			for _, field := range fields {
				item, _ := strconv.Atoi(field)
				update = append(update, item)
			}
			updates = append(updates, update)
		}
	}
	return rules, updates
}

func update_is_valid(update []int, rules *map[int][]int) bool {
	unseen := make(map[int]struct{}, len(update))
	for _, item := range update {
		unseen[item] = struct{}{}
	}

	// Iterate through the update backwards. An item is invalid
	// if a rule exists for that item such that an unseen element
	// (i.e. an item earlier in the update) is required to be after
	// the current item
	for _, item := range slices.Backward(update) {
		rule, has_rule := (*rules)[item]
		if has_rule {
			for _, must_be_after := range rule {
				_, is_before := unseen[must_be_after]
				if is_before {
					return false
				}
			}
		}
		delete(unseen, item)
	}
	return true
}

func fix_update(update []int, rules *map[int][]int) []int {
	set := make(map[int]struct{}, len(update))
	for _, item := range update {
		set[item] = struct{}{}
	}

	after_counts := make([]int, len(update))
	for i, item := range update {
		rule_list := (*rules)[item]
		for _, rule := range rule_list {
			_, matches := set[rule]
			if matches {
				after_counts[i] = after_counts[i] + 1
			}
		}
	}

	fixed := make([]int, len(update))
	for i, count := range after_counts {
		fixed[len(fixed)-1-count] = update[i]
	}

	return fixed
}

func main() {
	rules, updates := parse_input(os.Stdin)

	rule_map := make(map[int][]int)
	for _, rule := range rules {
		afters, has_afters := rule_map[rule.before]
		if has_afters {
			rule_map[rule.before] = append(afters, rule.after)
		} else {
			rule_map[rule.before] = []int{rule.after}
		}
	}

	valid_update_sum := 0
	fixed_update_sum := 0
	for _, update := range updates {
		if update_is_valid(update, &rule_map) {
			valid_update_sum += update[len(update)/2]
		} else {
			fixed := fix_update(update, &rule_map)
			fixed_update_sum += fixed[len(fixed)/2]
		}
	}

	fmt.Printf("Part 1: %d\n", valid_update_sum)
	fmt.Printf("Part 2: %d\n", fixed_update_sum)
}
