package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var part1, part2 int

	var seat_ids []int
	for scanner.Scan() {
		var seat_id int
		for _, c := range scanner.Text() {
			seat_id <<= 1
			if c == 'B' || c == 'R' {
				seat_id |= 1
			} else if c != 'F' && c != 'L' {
				log.Fatal("invalid input")
			}
		}
		seat_ids = append(seat_ids, seat_id)

		if seat_id > part1 {
			part1 = seat_id
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(seat_ids, func(i, j int) bool { return seat_ids[i] < seat_ids[j] })

	var prev int
	for _, id := range seat_ids {
		if id-2 == prev {
			part2 = id - 1
			break
		}
		prev = id
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
