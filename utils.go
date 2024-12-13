package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type pair struct {
	x, y int
}

type positionWithDirect struct {
	x, y int
	d    int
}

// FourDirections               ^       >        v       <
var FourDirections = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

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

func convertMatrix(raw []byte) [][]byte {
	var grid [][]byte
	for _, line := range strings.Split(string(raw), "\n") {
		grid = append(grid, []byte(line))
	}
	return grid
}

func set[T any](arr []T) []T {
	var out []T
	mp := map[string]int{}
	for _, v := range arr {
		k := fmt.Sprint(v)
		if mp[k] == 0 {
			out = append(out, v)
			mp[k] = 1
		}
	}
	return out
}

type Queue[T any] []T

func (t *Queue[T]) Len() int {
	return len(*t)
}

func (t *Queue[T]) Pop() T {
	item := (*t)[0]
	*t = (*t)[1:]
	return item
}

func (t *Queue[T]) Add(item T) {
	*t = append(*t, item)
}

type Matrix[T any] [][]T

func (t *Matrix[T]) IsEscape(r, c int) bool {
	return r < 0 || r >= len(*t) || c < 0 || c >= len((*t)[r])
}

func (t *Matrix[T]) JumpByFourDirections(r, c int, f func(x, y int, item T)) {
	for _, d := range FourDirections {
		nextR, nextC := r+d[0], c+d[1]
		if t.IsEscape(nextR, nextC) {
			continue
		}
		f(nextR, nextC, (*t)[nextR][nextC])
	}
}

func (t *Matrix[T]) JumpVertical(r, c int, f func(x, y int)) {
	for _, d := range [][]int{FourDirections[0], FourDirections[2]} {
		nextR, nextC := r, c
		for !t.IsEscape(nextR, nextC) {
			f(nextR, nextC)
			nextR, nextC = nextR+d[0], nextC+d[1]
		}
	}
}

func (t *Matrix[T]) JumpHorizontal(r, c int, f func(x, y int)) {
	for _, d := range [][]int{FourDirections[1], FourDirections[3]} {
		nextR, nextC := r, c
		for !t.IsEscape(nextR, nextC) {
			f(nextR, nextC)
			nextR, nextC = nextR+d[0], nextC+d[1]
		}
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func printSortedMap[T any](mp map[string]T) {
	var arr []string
	for k, _ := range mp {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	for _, v := range arr {
		fmt.Println(v, mp[v])
	}
}
