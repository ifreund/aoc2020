package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var part1, part2 int
	for _, group := range strings.Split(string(buf), "\n\n") {
		answers := map[rune]int{}
		people := strings.Fields(group)
		for _, person := range people {
			for _, answer := range person {
				answers[answer]++
			}
		}

		part1 += len(answers)

		for _, count := range answers {
			if count == len(people) {
				part2++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
