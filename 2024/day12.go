package _024

import (
	"fmt"
)

type Day12 struct {
}

func (t Day12) bfs(grid Matrix[byte], r, c int, visited Matrix[byte]) (int64, int64) {
	var total, near int64
	queue := Queue[pair]{{r, c}}
	for queue.Len() > 0 {
		q := queue.Pop()
		if (visited)[q.x][q.y] == '.' {
			continue
		}
		total++
		(visited)[q.x][q.y] = '.'
		grid.JumpByFourDirections(q.x, q.y, func(x, y int, item byte) {
			if grid[r][c] == item && (visited)[x][y] != '.' {
				near++
				queue.Add(pair{x, y})
			}
		})
	}
	return total*4 - near*2, total
}

func (t Day12) part1(raw []byte) int64 {
	var out int64
	grid := convertMatrix(raw)
	visited := cloneMatrix[byte](grid)
	for r, row := range grid {
		for c, _ := range row {
			if (visited)[r][c] == '.' {
				continue
			}
			x, k := t.bfs(grid, r, c, visited)
			out += x * k
		}
	}
	return out
}
func (t Day12) bfs2(grid Matrix[byte], r, c int, visited Matrix[byte]) (int64, int64) {
	var total int64
	queue := Queue[pair]{{r, c}}
	points := Queue[pair]{}
	for queue.Len() > 0 {
		q := queue.Pop()
		if (visited)[q.x][q.y] == '.' {
			continue
		}
		total++
		(visited)[q.x][q.y] = '.'
		grid.JumpByFourDirections(q.x, q.y, func(x, y int, item byte) {
			if grid[r][c] == item && (visited)[x][y] != '.' {
				queue.Add(pair{x, y})
			}
		})
		points.Add(q)
	}
	edge := t.countEdge(grid, points)
	return edge, total
}

func (t Day12) isEdge(grid Matrix[byte], x, y, rawX, rawY int) bool {
	return grid.IsEscape(x, y) || grid[x][y] != grid[rawX][rawY]
}

func (t Day12) countEdge(grid Matrix[byte], points Queue[pair]) int64 {
	var edge int64
	mp := map[string]int{}
	for points.Len() > 0 {
		p := points.Pop()
		for i, d := range FourDirections {
			x, y := p.x+d[0], p.y+d[1]
			if !t.isEdge(grid, x, y, p.x, p.y) {
				continue
			}
			rootX, rootY := p.x, p.y
			k := fmt.Sprintf("%d-%d-%d", rootX, rootY, i)
			if mp[k] > 0 {
				continue
			}

			edge++
			mp[k] = 1
			// FourDirections  ^   >   v   <

			if i%2 == 0 {
				for _, direct := range [][]int{FourDirections[1], FourDirections[3]} {
					nextR, nextC := rootX, rootY
					for !grid.IsEscape(nextR, nextC) {
						nextR, nextC = nextR+direct[0], nextC+direct[1]
						if grid.IsEscape(nextR, nextC) || grid[rootX][rootY] != grid[nextR][nextC] {
							break
						}
						fx, fy := nextR+d[0], nextC+d[1]
						k := fmt.Sprintf("%d-%d-%d", nextR, nextC, i)
						if t.isEdge(grid, fx, fy, rootX, rootY) {
							mp[k] = 1
						} else {
							break
						}
					}
				}
			} else {
				for _, direct := range [][]int{FourDirections[0], FourDirections[2]} {
					nextR, nextC := rootX, rootY

					for !grid.IsEscape(nextR, nextC) {
						nextR, nextC = nextR+direct[0], nextC+direct[1]
						if grid.IsEscape(nextR, nextC) || grid[rootX][rootY] != grid[nextR][nextC] {
							break
						}
						fx, fy := nextR+d[0], nextC+d[1]
						k := fmt.Sprintf("%d-%d-%d", nextR, nextC, i)
						if t.isEdge(grid, fx, fy, rootX, rootY) {
							mp[k] = 1
						} else {
							break
						}
					}
				}
			}
		}
	}

	return edge
}
func (t Day12) part2(raw []byte) int64 {
	var out int64
	grid := convertMatrix(raw)
	visited := cloneMatrix[byte](grid)
	for r, row := range grid {
		for c, _ := range row {
			if (visited)[r][c] == '.' {
				continue
			}
			x, k := t.bfs2(grid, r, c, visited)
			out += x * k
		}
	}
	return out
}
