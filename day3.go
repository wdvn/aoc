package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	raw, _ := os.ReadFile("./input/day3.txt")

	part1 := func() {
		reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

		cal := reg.FindAllSubmatch(raw, -1)
		var out int64
		for _, sub := range cal {
			n1, _ := strconv.Atoi(string(sub[1]))
			n2, _ := strconv.Atoi(string(sub[2]))
			out += int64(n1 * n2)
		}
		fmt.Println(out)
	}
	part1()

	part2 := func() {
		enable := true
		reg := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)

		cal := reg.FindAllSubmatch(raw, -1)
		var out int64
		for _, sub := range cal {
			switch string(sub[0]) {
			case "do()":
				enable = true
			case "don't()":
				enable = false
			default:
				if !enable {
					continue
				}
				n1, _ := strconv.Atoi(string(sub[1]))
				n2, _ := strconv.Atoi(string(sub[2]))
				out += int64(n1 * n2)
			}
		}
		fmt.Println(out, enable)
	}

	part2()
}
