package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var required = [...]string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var eye_colors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func main() {
	file, err := os.Open("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var part1, part2 int
outer:
	for _, raw_pass := range strings.Split(string(buf), "\n\n") {
		pass := map[string]string{}
		for _, field := range strings.Fields(raw_pass) {
			pass[field[:3]] = field[4:]
		}

		for _, f := range required {
			if _, ok := pass[f]; !ok {
				continue outer
			}
		}
		part1++

		if y, err := strconv.Atoi(pass["byr"]); err != nil || y < 1920 || y > 2002 {
			continue
		}
		if y, err := strconv.Atoi(pass["iyr"]); err != nil || y < 2010 || y > 2020 {
			continue
		}
		if y, err := strconv.Atoi(pass["eyr"]); err != nil || y < 2020 || y > 2030 {
			continue
		}

		hgt := pass["hgt"]
		h, err := strconv.Atoi(hgt[:len(hgt)-2])
		if err != nil {
			continue
		}
		switch hgt[len(hgt)-2:] {
		case "cm":
			if h < 150 || h > 193 {
				continue
			}
		case "in":
			if h < 59 || h > 76 {
				continue
			}
		default:
			continue
		}

		hcl := pass["hcl"]
		if hcl[0] != '#' || len(hcl) != 7 {
			continue
		}
		if _, err = strconv.ParseInt(hcl[1:], 16, 64); err != nil {
			continue
		}

		if !eye_colors[pass["ecl"]] {
			continue
		}

		pid := pass["pid"]
		if len(pid) != 9 {
			continue
		}
		if _, err := strconv.Atoi(pid); err != nil {
			continue
		}

		part2++
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
