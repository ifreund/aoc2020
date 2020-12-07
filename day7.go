package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type content struct {
	color string
	count int
}

func main() {
	file, err := os.Open("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	bags := map[string][]content{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " bags contain ")
		bag := split[0]
		contents := split[1]

		for _, raw_content := range strings.Split(contents, ", ") {
			raw := strings.Fields(raw_content)

			if raw[0] == "no" {
				break
			}

			count, err := strconv.Atoi(raw[0])
			if err != nil {
				log.Fatal(err)
			}

			bags[bag] = append(bags[bag], content{strings.Join(raw[1:3], " "), count})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var part1 int
	for bag, _ := range bags {
		if searchBag(bag, bags) {
			part1++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", countContents("shiny gold", bags))
}

func searchBag(bag string, bags map[string][]content) bool {
	for _, con := range bags[bag] {
		if con.color == "shiny gold" {
			return true
		}
		if searchBag(con.color, bags) {
			return true
		}
	}
	return false
}

func countContents(bag string, bags map[string][]content) int {
	var count int
	for _, con := range bags[bag] {
		count += con.count * (1 + countContents(con.color, bags))
	}
	return count
}
