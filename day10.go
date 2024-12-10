package main

import (
	"strings"
)

type Day10 struct{}

func (t Day10) dfs(grid [][]byte, r, c int, rootX, rootY int, visited *[][]int, isPart1 bool) {
	if isPart1 {
		if grid[r][c] == '9' && (*visited)[r][c] == 0 {
			(*visited)[rootX][rootY]++
			(*visited)[r][c]--
			return
		}
	} else {
		if grid[r][c] == '9' {
			(*visited)[rootX][rootY]++
			return
		}
	}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, d := range directions {
		x, y := r+d[0], c+d[1]
		if isEscape(grid, x, y) {
			continue
		}
		cur := grid[r][c]
		next := grid[x][y]
		if cur+1 != next {
			continue
		}
		t.dfs(grid, x, y, rootX, rootY, visited, isPart1)
	}
}

func (t Day10) part1(raw []byte) int64 {
	var out int64
	var grid [][]byte
	var visited [][]int
	for _, line := range strings.Split(string(raw), "\n") {
		grid = append(grid, []byte(line))
		visited = append(visited, make([]int, len(line)))
	}

	for r, row := range grid {
		for c, col := range row {
			if col == '0' {
				tmp := cloneMatrix(visited)
				t.dfs(grid, r, c, r, c, &tmp, true)
				out += int64(tmp[r][c])
			}
		}
	}
	return out
}

func (t Day10) part2(raw []byte) int64 {
	var out int64
	var grid [][]byte
	var visited [][]int
	for _, line := range strings.Split(string(raw), "\n") {
		grid = append(grid, []byte(line))
		visited = append(visited, make([]int, len(line)))
	}

	for r, row := range grid {
		for c, col := range row {
			if col == '0' {
				tmp := cloneMatrix(visited)
				t.dfs(grid, r, c, r, c, &tmp, false)
				out += int64(tmp[r][c])
			}
		}
	}
	return out
}
