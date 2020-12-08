package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inst struct {
	op  string
	arg int
}

func main() {
	file, err := os.Open("input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	var prog []inst
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		op := split[0]
		arg, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		prog = append(prog, inst{op, arg})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1, _ := run(prog)

	var part2 int
	for i, _ := range prog {
		modified := make([]inst, len(prog))
		copy(modified, prog)

		switch modified[i].op {
		case "acc":
			continue
		case "jmp":
			modified[i].op = "nop"
		case "nop":
			modified[i].op = "jmp"
		default:
			log.Fatal("invalid instruction")
		}

		acc, looped := run(modified)
		if !looped {
			part2 = acc
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func run(prog []inst) (int, bool) {
	var acc int
	var looped bool

	seen := map[int]bool{}

	for i := 0; i < len(prog); {
		if seen[i] {
			looped = true
			break
		}
		seen[i] = true

		switch prog[i].op {
		case "acc":
			acc += prog[i].arg
			i++
		case "jmp":
			i += prog[i].arg
		case "nop":
			i++
		default:
			log.Fatal("invalid instruction")
		}
	}

	return acc, looped
}
