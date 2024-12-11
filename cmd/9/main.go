package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

const (
	empty = -1
)

type node struct {
	start int
	size  int
}
type filesystem struct {
	files []node
	gaps  []node
	size  int
}

func parse_input(in_src *os.File) filesystem {
	in_bytes, _ := io.ReadAll(in_src)
	in := string(in_bytes[:])

	var files []node
	var gaps []node

	is_file := true
	blocks := 0
	for _, c := range in {
		count := int(c) - '0'
		if is_file {
			files = append(files, node{blocks, count})
		} else {
			gaps = append(gaps, node{blocks, count})
		}
		is_file = !is_file
		blocks += count
	}

	return filesystem{files, gaps, blocks}
}

func to_blocks(fs filesystem) []int {
	blocks := make([]int, fs.size)

	for i, file := range fs.files {
		for j := range file.size {
			blocks[file.start+j] = i
		}
	}

	for _, gap := range fs.gaps {
		for j := range gap.size {
			blocks[gap.start+j] = empty
		}
	}

	return blocks
}

func print_blocks(blocks []int) {
	for _, b := range blocks {
		if b == empty {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", b)
		}
	}
	fmt.Print("\n")
}

func defrag_blocks(blocks *[]int) {
	head, tail := 0, len(*blocks)-1
	for head < tail {
		if (*blocks)[head] != empty {
			head++
		} else if (*blocks)[tail] == empty {
			tail--
		} else {
			(*blocks)[head] = (*blocks)[tail]
			(*blocks)[tail] = empty
			head++
			tail--
		}
	}
}

func defrag_files(fs *filesystem) {
	for i, _ := range slices.Backward(fs.files) {
		file := &fs.files[i]
		for j := range fs.gaps {
			gap := &fs.gaps[j]
			if file.start < gap.start {
				break
			}

			if gap.size >= file.size {
				start := file.start
				file.start = gap.start
				gap.start += file.size
				gap.size -= file.size
				fs.gaps = append(fs.gaps, node{start, file.size})
				break
			}
		}
	}
}

func checksum(blocks *[]int) int {
	cs := 0
	for i, v := range *blocks {
		if v != empty {
			cs += (i * v)
		}
	}
	return cs
}
func main() {
	fs := parse_input(os.Stdin)

	blocks := to_blocks(fs)
	defrag_blocks(&blocks)
	fmt.Printf("Part 1: %d\n", checksum(&blocks))

	defrag_files(&fs)
	blocks = to_blocks(fs)
	fmt.Printf("Part 2: %d\n", checksum(&blocks))
}
