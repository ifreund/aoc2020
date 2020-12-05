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
		seat := scanner.Text()

		row := 0
		amount := 128
		for _, c := range seat[:7] {
			amount /= 2
			switch c {
			case 'B':
				row += amount
			case 'F':
				// do nothing
			default:
				log.Fatal("invalid input")
			}
		}

		col := 0
		amount = 8
		for _, c := range seat[7:] {
			amount /= 2
			switch c {
			case 'R':
				col += amount
			case 'L':
				// do nothing
			default:
				log.Fatal("invalid input")
			}
		}

		seat_id := row*8 + col
		seat_ids = append(seat_ids, seat_id)

		if seat_id > part1 {
			part1 = seat_id
		}
	}
	if err != nil {
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
