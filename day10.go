package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input/day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var adapters []int
	for scanner.Scan() {
		in, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		adapters = append(adapters, in)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	adapters = append(adapters, 0)
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })

	var ones, threes int
	for i := 1; i < len(adapters); i++ {
		switch adapters[i] - adapters[i-1] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	threes++

	counts := make([]uint64, len(adapters))
	counts[0] = 1

	for i, x := range adapters[1:] {
		for j, y := range adapters[:i+1] {
			if x-y == 3 || x-y == 2 || x-y == 1 {
				counts[i+1] += counts[j]
			}
		}
	}

	fmt.Printf("Part 1: %d\n", ones*threes)
	fmt.Printf("Part 2: %d\n", counts[len(counts)-1])
}

func swapRemove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
