package main

import (
	"strconv"
	"strings"
)

type Day18 struct {
}

func (d Day18) parse(raw []byte, k int) [][]byte {
	grid := make([][]byte, 71)
	for i := range 71 {
		arr := make([]byte, 71)
		for j := range 71 {
			arr[j] = '.'
		}
		grid[i] = arr
	}
	for i, line := range strings.Split(string(raw), "\n") {
		if line == "" || i > k {
			continue
		}
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		grid[y][x] = '#'
	}
	return grid
}

func (d Day18) search(raw []byte, k int) int {
	grid := d.parse(raw, k)
	queue := Queue[cost]{{x: 0, y: 0, cost: 0}}
	visited := map[pairWithD]int{}
	for queue.Len() > 0 {
		q := queue.Pop()
		if q.x == 70 && q.y == 70 {
			return q.cost
		}
		for i, d := range FourDirections {
			x, y := q.x+d[0], q.y+d[1]
			next := pairWithD{x, y, i}
			if isEscape(grid, y, x) || grid[y][x] == '#' || visited[next] > 0 {
				continue
			}
			visited[next] = 1
			queue.Add(cost{x: x, y: y, cost: q.cost + 1})
		}
	}
	return 0
}

func (d Day18) part1(raw []byte) any {
	return d.search(raw, 1023)
}

func (d Day18) part2(raw []byte) any {
	var out any
	lines := strings.Split(string(raw), "\n")
	arr := BinarySearch[int]{}
	for i := range len(lines) {
		arr = append(arr, i)
	}
	out = arr.BotSearch(func(mid, nextMid int) bool {
		v1 := d.search(raw, mid)
		v2 := d.search(raw, nextMid)
		if v1 > 0 && v2 > 0 {
			return true
		}
		return false
	})
	return out
}
