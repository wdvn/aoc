package main

//
//import (
//	"fmt"
//	"os"
//	"slices"
//	"strings"
//)
//
//func findStartPosition(grid [][]byte) (int, int, int) {
//	for i, row := range grid {
//		for j, col := range row {
//			switch col {
//			case '^':
//				return i, j, 0
//			case '>':
//				return i, j, 1
//			case 'v':
//				return i, j, 2
//			case '<':
//				return i, j, 3
//			}
//		}
//	}
//	return -1, -1, -1
//}
//func isEscape(grid [][]byte, x, y int) bool {
//	return x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0])
//}
//func isLoop(grid [][]byte) bool {
//	mp := make(map[string]int)
//	directions := [][]int{[]int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1}}
//	x, y, d := findStartPosition(grid)
//	for {
//		k := fmt.Sprintf("%d-%d-%d", x, y, d)
//		mp[k]++
//		if mp[k] > 1 {
//			return true
//		}
//		nextX, nextY := x+directions[d][0], y+directions[d][1]
//		if isEscape(grid, nextX, nextY) {
//			break
//		}
//		for grid[nextX][nextY] == '#' {
//			d = (d + 1) % 4
//			nextX, nextY = x+directions[d][0], y+directions[d][1]
//		}
//		x, y = nextX, nextY
//	}
//	return false
//}
//
//func part2(grid [][]byte) int {
//	var out int
//	for i, row := range grid {
//		for j, col := range row {
//			if col != '.' {
//				continue
//			}
//			tmp := cloneMatrix(grid)
//			tmp[i][j] = '#'
//			if isLoop(tmp) {
//				out++
//			}
//		}
//	}
//	return out
//}
//
//func printMatrix(matrix [][]byte) string {
//	var arr []string
//	for _, row := range matrix {
//		arr = append(arr, string(row))
//	}
//	return strings.Join(arr, "\n")
//}
//func cloneMatrix[T any](grid [][]T) [][]T {
//	var result [][]T
//	for _, row := range grid {
//		result = append(result, slices.Clone(row))
//	}
//	return result
//}
//
//func main() {
//	raw, _ := os.ReadFile("./input/day6.txt")
//	var grid [][]byte
//	for _, line := range strings.Split(string(raw), "\n") {
//		grid = append(grid, []byte(line))
//	}
//
//	fmt.Println(part2(grid))
//
//}
