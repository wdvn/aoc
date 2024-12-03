package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func cal(sub [][]byte) int64 {
	n1, _ := strconv.Atoi(string(sub[1]))
	n2, _ := strconv.Atoi(string(sub[2]))
	return int64(n1 * n2)
}

func part1(s []byte) int64 {
	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	subs := reg.FindAllSubmatch(s, -1)
	var out int64
	for _, sub := range subs {
		out += cal(sub)
	}
	return out
}

func part2(s []byte) int64 {
	enable := true
	reg := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)
	var out int64
	subs := reg.FindAllSubmatch(s, -1)
	for _, sub := range subs {
		switch sub[0][3] {
		case ')':
			enable = true
		case "'"[0]:
			enable = false
		default:
			if !enable {
				continue
			}
			out += cal(sub)
		}
	}
	return out
}
func main() {
	raw, _ := os.ReadFile("./input/day3.txt")

	fmt.Println(part1(raw))
	fmt.Println(part2(raw))
}
