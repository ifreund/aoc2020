package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input/day1.txt")
	exitIfErr(err)

	scanner := bufio.NewScanner(file)

	var input []int
	for scanner.Scan() {
		var in int
		n, err := fmt.Sscanf(scanner.Text(), "%d", &in)
		exitIfErr(err)
		if n != 1 {
			fmt.Fprintln(os.Stderr, "invalid input")
			os.Exit(1)
		}
		input = append(input, in)
	}
	exitIfErr(scanner.Err())

	one, two := false, false
	for _, a := range input {
		for _, b := range input {
			if !one && a+b == 2020 {
				fmt.Printf("Part 1: %d\n", a*b)
				one = true
			}
			if !two {
				for _, c := range input {
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
		fmt.Fprintln(os.Stderr, "invalid input")
		os.Exit(1)
	}
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
