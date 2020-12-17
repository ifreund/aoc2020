package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x, y, z, w int
}

func main() {
	file, err := os.Open("input/day17.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	grid := map[point]bool{}

	var start_x, start_y int
	for scanner.Scan() {
		start_x = 0
		for _, c := range scanner.Text() {
			if c == '#' {
				grid[point{start_x, start_y, 0, 0}] = true
			}
			start_x++
		}
		start_y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid2 := map[point]bool{}
	for k, v := range grid {
		grid2[k] = v
	}

	next_grid := map[point]bool{}
	for cycle := 1; cycle <= 6; cycle++ {
		for z := -cycle; z <= cycle; z++ {
			for y := -cycle; y <= start_y+cycle; y++ {
				for x := -cycle; x <= start_x+cycle; x++ {
					p := point{x, y, z, 0}
					n := neighbors(p, grid)
					if n == 3 || (n == 2 && grid[p]) {
						next_grid[p] = true
					}
				}
			}
		}
		grid = next_grid
		next_grid = map[point]bool{}
	}

	for cycle := 1; cycle <= 6; cycle++ {
		for w := -cycle; w <= cycle; w++ {
			for z := -cycle; z <= cycle; z++ {
				for y := -cycle; y <= start_y+cycle; y++ {
					for x := -cycle; x <= start_x+cycle; x++ {
						p := point{x, y, z, w}
						n := neighbors(p, grid2)
						if n == 3 || (n == 2 && grid2[p]) {
							next_grid[p] = true
						}
					}
				}
			}
		}

		grid2 = next_grid
		next_grid = map[point]bool{}
	}

	fmt.Printf("Part1: %v\n", len(grid))
	fmt.Printf("Part2: %v\n", len(grid2))
}

func neighbors(p point, grid map[point]bool) int {
	var count int
	for w := p.w - 1; w <= p.w+1; w++ {
		for z := p.z - 1; z <= p.z+1; z++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for x := p.x - 1; x <= p.x+1; x++ {
					if x == p.x && y == p.y && z == p.z && w == p.w {
						continue
					}
					if grid[point{x, y, z, w}] {
						count++
					}
				}
			}
		}
	}
	return count
}
