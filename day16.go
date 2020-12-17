package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule_t struct {
	name               string
	lo1, hi1, lo2, hi2 uint32
}

func main() {
	file, err := os.Open("input/day16.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(buf), "\n\n")

	var rules []rule_t
	for _, raw_rule := range strings.Split(split[0], "\n") {
		colon := strings.Index(raw_rule, ":")

		name := raw_rule[:colon]

		var lo1, hi1, lo2, hi2 uint32
		n, err := fmt.Sscanf(raw_rule[colon:], ": %d-%d or %d-%d", &lo1, &hi1, &lo2, &hi2)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("invalid input")
		}

		rules = append(rules, rule_t{name, lo1, hi1, lo2, hi2})
	}

	var my_ticket []uint32
	for _, raw_val := range strings.Split(strings.Fields(split[1])[2], ",") {
		val, err := strconv.ParseUint(raw_val, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		my_ticket = append(my_ticket, uint32(val))
	}

	var nearby_tickets [][]uint32
	for _, raw_ticket := range strings.Split(split[2], "\n")[1:] {
		if len(raw_ticket) == 0 {
			continue
		}
		var ticket []uint32
		for _, raw_val := range strings.Split(raw_ticket, ",") {
			val, err := strconv.ParseUint(raw_val, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			ticket = append(ticket, uint32(val))
		}
		nearby_tickets = append(nearby_tickets, ticket)
	}

	var valid_tickets [][]uint32
	valid_tickets = append(valid_tickets, my_ticket)

	var part1 uint32
	for _, ticket := range nearby_tickets {
		invalid := false
	ticket_loop:
		for _, val := range ticket {
			for _, rule := range rules {
				if (rule.lo1 <= val && val <= rule.hi1) || (rule.lo2 <= val && val <= rule.hi2) {
					continue ticket_loop
				}
			}
			part1 += val
			invalid = true
			break
		}
		if !invalid {
			valid_tickets = append(valid_tickets, ticket)
		}
	}

	fmt.Printf("Part 1: %d\n", part1)

	valid := make([]bool, len(rules))
	for i := range valid {
		valid[i] = true
	}

	order := make([]int, len(rules))

	if !check(valid_tickets, rules, valid, order, 0) {
		log.Fatal("invalid input")
	}

	// Thanks a lot go for silently ignoring overflows
	var part2 uint64 = 1
	for i, j := range order {
		if strings.HasPrefix(rules[j].name, "departure") {
			part2 *= uint64(my_ticket[i])
		}
	}
	fmt.Printf("Part 2: %d\n", part2)
}

func check(tickets [][]uint32, rules []rule_t, valid []bool, order []int, idx int) bool {
	if idx >= len(rules) {
		return true
	}

	for i, rule := range rules {
		if valid[i] {
			valid[i] = false

			order[idx] = i

			found := false
			for _, ticket := range tickets {
				val := ticket[idx]
				if !((rule.lo1 <= val && val <= rule.hi1) || (rule.lo2 <= val && val <= rule.hi2)) {
					found = true
					break
				}
			}
			if !found {
				ret := check(tickets, rules, valid, order, idx+1)
				if ret {
					return true
				}
			}
			valid[i] = true
		}
	}

	return false
}
