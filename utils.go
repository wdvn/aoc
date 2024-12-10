package main

import (
	"fmt"
	"slices"
)

func cloneMatrix[T any](grid [][]T) [][]T {
	var result [][]T
	for _, row := range grid {
		result = append(result, slices.Clone(row))
	}
	return result
}
func isEscape(grid [][]byte, x, y int) bool {
	return x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0])
}
func printMatrix(grid [][]byte) {
	for row := 0; row < len(grid); row++ {
		fmt.Println(string(grid[row]))
	}
}

func printMatrixA[T any](grid [][]T) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Printf("%v  ", c)
		}
		fmt.Println()
	}
}
