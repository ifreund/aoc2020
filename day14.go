package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input/day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(buf))
	fmt.Printf("Part 2: %d\n", part2(buf))
}

func part1(buf []byte) uint64 {
	mem := map[uint64]uint64{}
	var mask0, mask1 uint64
	for _, input := range strings.Split(string(buf), "\n") {
		if len(input) == 0 {
			continue
		}
		if input[:3] == "mem" {
			var addr, val uint64
			n, err := fmt.Sscanf(input, "mem[%d] = %d", &addr, &val)
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatal("invalid input")
			}

			val &= mask0
			val |= mask1

			mem[addr] = val
		} else {
			mask0 = (1 << 36) - 1
			mask1 = 0
			for i, x := range []rune(input[7:]) {
				switch x {
				case '0':
					mask0 &= ^(1 << (35 - i))
				case '1':
					mask1 |= (1 << (35 - i))
				}
			}
		}
	}

	var ret uint64
	for _, v := range mem {
		ret += v
	}
	return ret
}

func part2(buf []byte) uint64 {
	mem := map[uint64]uint64{}
	var mask1 uint64
	var maskX uint64
	for _, input := range strings.Split(string(buf), "\n") {
		if len(input) == 0 {
			continue
		}
		if input[:3] == "mem" {
			var addr, val uint64
			n, err := fmt.Sscanf(input, "mem[%d] = %d", &addr, &val)
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatal("invalid input")
			}

			addr |= mask1
			for _, a := range applyMaskX(maskX, addr) {
				mem[a] = val
			}
		} else {
			mask1 = 0
			maskX = 0
			for i, x := range []rune(input[7:]) {
				switch x {
				case '1':
					mask1 |= (1 << (35 - i))
				case 'X':
					maskX |= (1 << (35 - i))
				}
			}
		}
	}

	var ret uint64
	for _, v := range mem {
		ret += v
	}
	return ret
}

func applyMaskX(mask, addr uint64) []uint64 {
	if mask == 0 {
		return []uint64{addr}
	}
	for i := 0; i < 36; i++ {
		if mask&(1<<i) != 0 {
			part := applyMaskX(mask&(^(1 << i)), addr)
			var ret []uint64
			for _, a := range part {
				ret = append(ret, a&(^(1 << i)))
				ret = append(ret, a|(1<<i))
			}
			return ret
		}
	}
	log.Fatal("oops")
	return []uint64{}
}
