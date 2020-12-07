package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var part1, part2 int
	for scanner.Scan() {
		var lo, hi int
		var c rune
		var s string
		n, err := fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &lo, &hi, &c, &s)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("invalid input")
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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
