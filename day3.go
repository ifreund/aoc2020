package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type slope struct {
	right, down int
}

// wtf, you can't make a constant array in go?!?!?
var slopes = [...]slope{
	slope{1, 1},
	slope{3, 1},
	slope{5, 1},
	slope{7, 1},
	slope{1, 2},
}

func main() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var input [][]bool
	for scanner.Scan() {
		var row []bool
		for _, v := range scanner.Text() {
			row = append(row, v == '#')
		}
		input = append(input, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sigh, pretty sure this is allocating at runtime.
	// can't I just have a static array ffs?
	var out [len(slopes)]int
	for i, s := range slopes {
		var x, y int
		for y < len(input) {
			if input[y][x] {
				out[i]++
			}
			x += s.right
			x %= len(input[0])
			y += s.down
		}
	}

	part2 := out[0]
	for _, v := range out[1:] {
		part2 *= v
	}

	fmt.Printf("Part 1: %d\n", out[2])
	fmt.Printf("Part 2: %d\n", part2)
}
