package main

import (
	"fmt"
	"os"
)

func main() {
	day := Day9{}
	raw, _ := os.ReadFile("./input/day9.txt")
	fmt.Println(day.part1(raw))
	fmt.Println(day.part2(raw))

	fmt.Println(isPointOnLine(pair{0, 1}, pair{1, 1}, pair{0, 0}))
}