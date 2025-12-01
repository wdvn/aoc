package main

import (
	"fmt"
	"os"
)

func main() {
	day := Day18{}
	raw, _ := os.ReadFile("./input/day18.txt")
	fmt.Println(day.part1(raw))
	fmt.Println(day.part2(raw))
}
