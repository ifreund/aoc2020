package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input/day1.txt")
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

	one, two := false, false
	for i, a := range input {
		for j, b := range input[i:] {
			if !one && a+b == 2020 {
				fmt.Printf("Part 1: %d\n", a*b)
				one = true
			}
			if !two {
				for _, c := range input[j:] {
					if a+b+c == 2020 {
						fmt.Printf("Part 2: %d\n", a*b*c)
						two = true
						break
					}
				}
			}
		}
	}

	if !one || !two {
		log.Fatal("invalid input")
	}
}
