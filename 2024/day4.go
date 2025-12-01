package _024

//
//import (
//	"fmt"
//	"os"
//	"strings"
//)
//
//// XMAS
//func countXmas(grid [][]byte, r, c int) int {
//	var out int
//	var tmp []byte
//	directs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
//	for _, direct := range directs {
//		var x, y int = r, c
//		for len(tmp) < 4 {
//			tmp = append(tmp, grid[x][y])
//			x += direct[0]
//			y += direct[1]
//			if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
//				break
//			}
//		}
//		if string(tmp) == "XMAS" {
//			out++
//		}
//		tmp = nil
//	}
//	return out
//}
//
//func countMas(grid [][]byte, r, c int) int {
//	var out int
//	var tmp []byte
//	directs := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
//
//	for _, direct := range directs {
//		var x, y int = r, c
//		x += direct[0]
//		y += direct[1]
//		if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
//			break
//		}
//		tmp = append(tmp, grid[x][y])
//	}
//	/*
//		M-S   M-M  S-M  S-S
//		M-S   S-S  S-M  M-M
//	*/
//	switch string(tmp) {
//	case "MSSM", "MMSS", "SMMS", "SSMM":
//		out++
//	}
//
//	return out
//}
//func part1(grid [][]byte) int {
//	var out int
//	for i, row := range grid {
//		for j, c := range row {
//			if c == 'X' {
//				out += countXmas(grid, i, j)
//			}
//		}
//	}
//	return out
//}
//
//func part2(grid [][]byte) int {
//	var out int
//	for i, row := range grid {
//		for j, c := range row {
//			if c == 'A' {
//				out += countMas(grid, i, j)
//			}
//		}
//	}
//	return out
//}
//
//func main() {
//	raw, _ := os.ReadFile("./input/day4.txt")
//	var grid [][]byte
//	for _, line := range strings.Split(string(raw), "\n") {
//		grid = append(grid, []byte(line))
//	}
//
//	fmt.Println(part1(grid))
//	fmt.Println(part2(grid))
//}
