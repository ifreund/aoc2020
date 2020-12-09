package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var input []int
	for scanner.Scan() {
		in, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, in)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var part1, part2 int
outer:
	for i := 25; i < len(input); i++ {
		prev25 := input[i-25 : i]
		for j, x := range prev25 {
			for _, y := range prev25[j:] {
				if x+y == input[i] {
					continue outer
				}
			}
		}
		part1 = input[i]
		break
	}

outer2:
	for i, _ := range input {
		var sum int
		for j, x := range input[i:] {
			sum += x
			if sum == part1 {
				min := math.MaxInt32
				max := 0
				for _, y := range input[i : i+j+1] {
					if y < min {
						min = y
					}
					if y > max {
						max = y
					}
				}
				part2 = min + max
				break outer2
			} else if sum > part1 {
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
