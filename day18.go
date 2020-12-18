package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type state struct {
	val uint64
	op  rune
	add bool
}

func main() {
	file, err := os.Open("input/day18.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var part1 uint64
	for _, eq := range strings.Split(string(buf), "\n") {
		var states []state
		states = append(states, state{0, '+', false})
		for _, ch := range eq {
			var val uint64
			switch ch {
			case ' ':
				continue
			case '+':
				states[len(states)-1].op = '+'
				continue
			case '*':
				states[len(states)-1].op = '*'
				continue
			case '(':
				states = append(states, state{0, '+', false})
				continue
			case ')':
				val = states[len(states)-1].val
				states = states[:len(states)-1]
			default:
				val = uint64(ch - '0')
			}
			switch states[len(states)-1].op {
			case '+':
				states[len(states)-1].val += val
			case '*':
				states[len(states)-1].val *= val
			}
		}
		if len(states) != 1 {
			log.Fatal("unclosed parens")
		}
		part1 += states[0].val
	}

	var part2 uint64
	for _, eq := range strings.Split(string(buf), "\n") {
		var states []state
		states = append(states, state{0, '+', false})
		for i, ch := range eq {
			var val uint64
			switch ch {
			case ' ':
				continue
			case '+':
				states[len(states)-1].op = '+'
				continue
			case '*':
				states[len(states)-1].op = '*'
				continue
			case '(':
				states = append(states, state{0, '+', false})
				continue
			case ')':
				val = states[len(states)-1].val
				states = states[:len(states)-1]
			default:
				val = uint64(ch - '0')
			}
			if states[len(states)-1].add {
				val += states[len(states)-1].val
				states = states[:len(states)-1]
			}
			if i+2 < len(eq) && []rune(eq)[i+2] == '+' && !(states[len(states)-1].add) {
				states = append(states, state{val, '+', true})
				continue
			}
			switch states[len(states)-1].op {
			case '+':
				states[len(states)-1].val += val
			case '*':
				states[len(states)-1].val *= val
			}
		}
		if len(states) != 1 {
			log.Fatal("unclosed parens")
		}
		part2 += states[0].val
	}

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
