package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x, y int
}

const (
	N int = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func main() {
	file, err := os.Open("input/day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	room := map[point]rune{}
	var y int
	for scanner.Scan() {
		for x, val := range scanner.Text() {
			room[point{x, y}] = val
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	room2 := map[point]rune{}
	for k, v := range room {
		room2[k] = v
	}

	changed := true
	for changed {
		room, changed = step(room)
	}

	adj := buildAdj(room2)
	changed = true
	for changed {
		room2, changed = step2(room2, adj)
	}

	var part1, part2 int
	for _, v := range room {
		if v == '#' {
			part1++
		}
	}
	for _, v := range room2 {
		if v == '#' {
			part2++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func step(room map[point]rune) (map[point]rune, bool) {
	next := map[point]rune{}
	for k, v := range room {
		next[k] = v
	}

	var changed bool
	for p, val := range room {
		switch val {
		case 'L':
			if adj(room, p) == 0 {
				next[p] = '#'
				changed = true
			}
		case '#':
			if adj(room, p) >= 4 {
				next[p] = 'L'
				changed = true
			}
		}
	}

	return next, changed
}

func adj(room map[point]rune, p point) int {
	var count int
	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+1; x++ {
			if x == p.x && y == p.y {
				continue
			}
			if room[point{x, y}] == '#' {
				count++
			}
		}
	}
	return count
}

func p(room [][]rune) {
	for _, row := range room {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func step2(room map[point]rune, adj map[point][]point) (map[point]rune, bool) {
	next := map[point]rune{}
	for k, v := range room {
		next[k] = v
	}

	var changed bool
	for p, val := range room {
		switch val {
		case 'L':
			if adj2(room, adj, p) == 0 {
				next[p] = '#'
				changed = true
			}
		case '#':
			if adj2(room, adj, p) >= 5 {
				next[p] = 'L'
				changed = true
			}
		}
	}

	return next, changed
}

func adj2(room map[point]rune, adj map[point][]point, p point) int {
	var count int
	for _, a := range adj[p] {
		if room[a] == '#' {
			count++
		}
	}
	return count
}

func buildAdj(room map[point]rune) map[point][]point {
	adj := map[point][]point{}
	for p, val := range room {
		if val == '.' {
			continue
		}
		for dir := N; dir <= NW; dir++ {
			cur := moveInDir(p, dir)
			for room[cur] == '.' {
				cur = moveInDir(cur, dir)
			}
			if room[cur] == '#' || room[cur] == 'L' {
				adj[p] = append(adj[p], cur)
			}
		}
	}
	return adj
}

func moveInDir(p point, dir int) point {
	switch dir {
	case N:
		p.y--
	case NE:
		p.y--
		p.x++
	case E:
		p.x++
	case SE:
		p.y++
		p.x++
	case S:
		p.y++
	case SW:
		p.y++
		p.x--
	case W:
		p.x--
	case NW:
		p.x--
		p.y--
	}
	return p
}
