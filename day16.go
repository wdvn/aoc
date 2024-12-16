package main

import (
	"fmt"
	"math"
	"sort"
)

type Day16 struct {
}

func (d Day16) parse(raw []byte) [][]byte {
	return convertMatrix(raw)
}

type cost struct {
	x, y   int
	cost   int
	direct int
	chains int
}

type pairWithD struct {
	x, y   int
	direct int
}

func (d Day16) start(grid [][]byte) pair {
	var s pair
	for r, row := range grid {
		for c, col := range row {
			if col == 'S' {
				return pair{r, c}
			}
		}
	}
	return s
}

// Using min heap + bfs for optimize travel time
func (d Day16) part1(raw []byte) int {
	grid := d.parse(raw)
	start := d.start(grid)
	queue := Queue[cost]{}
	visited := map[pairWithD]int{}
	queue.Add(cost{x: start.x, y: start.y, cost: 0, direct: 1})
	out := math.MaxInt
	var chains int
	//	mp := cloneMatrix(grid)
	//	var chains []pair
	//defer func() {
	//	for _, n := range chains {
	//		mp[n.x][n.y] = '*'
	//	}
	//	printMatrix(mp)
	//}()
	for queue.Len() > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})
		q := queue.Pop()
		pd := pairWithD{x: q.x, y: q.y, direct: q.direct}
		if visited[pd] > 0 {
			continue
		}
		visited[pd] = 1
		for f, d := range FourDirections {

			x, y := q.x+d[0], q.y+d[1]
			if isEscape(grid, x, y) || grid[x][y] == '#' {
				continue
			}
			if grid[x][y] == 'E' {
				if out > q.cost+1 {
					chains = q.chains + 1
					out = q.cost + 1
				}
				break
			}
			c := q.cost + 1
			if f != q.direct {
				c += 1000
			}
			//next := pair{x, y}
			//blocks := slices.Clone(q.chains)
			//queue.Add(cost{x: x, y: y, cost: c, direct: f, chains: append(blocks, next)})
			queue.Add(cost{x: x, y: y, cost: c, direct: f, chains: q.chains + 1})
		}
	}
	fmt.Println("part1:", out)
	fmt.Println("part2:", chains)
	return out
}

func (d Day16) part2(raw []byte) int {
	return 0
}
