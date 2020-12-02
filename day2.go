package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input/day2.txt")
	exitIfErr(err)

	scanner := bufio.NewScanner(file)

	var part1 int
	var part2 int
	for scanner.Scan() {
		var lo, hi int
		var c rune
		var s string
		n, err := fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &lo, &hi, &c, &s)
		exitIfErr(err)
		if n != 4 {
			fmt.Fprintln(os.Stderr, "invalid input")
			os.Exit(1)
		}

		// Part 1
		var count int
		for _, char := range s {
			if char == c {
				count++
			}
		}

		if lo <= count && count <= hi {
			part1++
		}

		// Part 2
		r := []rune(s)
		if (lo-1 < len(r) && r[lo-1] == c) != (hi-1 < len(r) && r[hi-1] == c) {
			part2++
		}
	}
	exitIfErr(scanner.Err())

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
