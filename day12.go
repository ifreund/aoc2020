package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

type action struct {
	dir rune
	val int
}

func main() {
	file, err := os.Open("input/day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var actions []action
	for scanner.Scan() {
		line := []rune(scanner.Text())

		val, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatal(err)
		}
		actions = append(actions, action{line[0], val})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := part1(actions)
	part2 := part2(actions)

	fmt.Printf("Part 1: %d\n", int(math.Abs(float64(part1.x))+math.Abs(float64(part1.y))))
	fmt.Printf("Part 2: %d\n", int(math.Abs(float64(part2.x))+math.Abs(float64(part2.y))))
}

func part1(actions []action) point {
	var current point
	facing := point{1, 0} // East
	for _, a := range actions {
		switch a.dir {
		case 'N':
			current.y -= a.val
		case 'S':
			current.y += a.val
		case 'E':
			current.x += a.val
		case 'W':
			current.x -= a.val
		case 'L':
			switch a.val {
			case 270:
				facing.x, facing.y = facing.y, -facing.x
				fallthrough
			case 180:
				facing.x, facing.y = facing.y, -facing.x
				fallthrough
			case 90:
				facing.x, facing.y = facing.y, -facing.x
			}
		case 'R':
			switch a.val {
			case 270:
				facing.x, facing.y = -facing.y, facing.x
				fallthrough
			case 180:
				facing.x, facing.y = -facing.y, facing.x
				fallthrough
			case 90:
				facing.x, facing.y = -facing.y, facing.x
			}
		case 'F':
			current.x += facing.x * a.val
			current.y += facing.y * a.val
		default:
			log.Fatal("invalid action")
		}
	}
	return current
}

func part2(actions []action) point {
	current := point{0, 0}
	waypoint := point{10, -1} // 10 E 1 N
	for _, a := range actions {
		switch a.dir {
		case 'N':
			waypoint.y -= a.val
		case 'S':
			waypoint.y += a.val
		case 'E':
			waypoint.x += a.val
		case 'W':
			waypoint.x -= a.val
		case 'L':
			switch a.val {
			case 270:
				waypoint.x, waypoint.y = waypoint.y, -waypoint.x
				fallthrough
			case 180:
				waypoint.x, waypoint.y = waypoint.y, -waypoint.x
				fallthrough
			case 90:
				waypoint.x, waypoint.y = waypoint.y, -waypoint.x
			}
		case 'R':
			switch a.val {
			case 270:
				waypoint.x, waypoint.y = -waypoint.y, waypoint.x
				fallthrough
			case 180:
				waypoint.x, waypoint.y = -waypoint.y, waypoint.x
				fallthrough
			case 90:
				waypoint.x, waypoint.y = -waypoint.y, waypoint.x
			}
		case 'F':
			current.x += waypoint.x * a.val
			current.y += waypoint.y * a.val
		default:
			log.Fatal("invalid action")
		}
	}
	return current
}
