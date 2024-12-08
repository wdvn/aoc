package main

import (
	"fmt"
	"os"
)

func main() {
	day8 := Day8{}
	raw, _ := os.ReadFile("./input/day8.txt")
	fmt.Println(day8.part1(raw))
	fmt.Println(day8.part2(raw))

	fmt.Println(isPointOnLine(pair{0, 1}, pair{1, 1}, pair{0, 0}))
}
