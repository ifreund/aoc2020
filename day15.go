package main

import "fmt"

func main() {
	mem := map[uint32]uint32{6: 1, 13: 2, 1: 3, 15: 4, 2: 5, 0: 6}

	var next uint32 = 0
	var turn uint32 = 7
	for ; turn < 30000000; turn++ {
		if turn == 2020 {
			fmt.Printf("Part 1: %d\n", next)
		}
		if mem[next] == 0 {
			mem[next] = turn
			next = 0
		} else {
			age := turn - mem[next]
			mem[next] = turn
			next = age
		}
	}

	fmt.Printf("Part 2: %d\n", next)
}
