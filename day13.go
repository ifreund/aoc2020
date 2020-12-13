package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input/day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Fields(string(buf))

	timestamp, err := strconv.ParseUint(lines[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var buses []uint64
	for _, raw_id := range strings.Split(lines[1], ",") {
		if raw_id == "x" {
			buses = append(buses, 0)
			continue
		}
		id, err := strconv.ParseUint(raw_id, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		buses = append(buses, id)
	}

	var earliest uint64 = math.MaxUint64
	var earliest_bus uint64
	for _, bus := range buses {
		if bus == 0 {
			continue
		}
		depart := bus * (timestamp/bus + 1)
		if depart < earliest {
			earliest = depart
			earliest_bus = bus
		}
	}

	var period, offset uint64 = 1, 0
	for i, bus := range buses {
		if bus == 0 {
			continue
		}
		for (offset+uint64(i))%bus != 0 {
			offset += period
		}
		period *= bus
	}

	fmt.Printf("Part 1: %d\n", (earliest-timestamp)*earliest_bus)
	fmt.Printf("Part 2: %d\n", offset)
}
