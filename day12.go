package main

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
	for points.Len() > 0 {
		p := points.Pop()
		for i, d := range FourDirections {
			x, y := p.x+d[0], p.y+d[1]
			if !t.isEdge(grid, x, y, p.x, p.y) {
				continue
			}
			edge++
			//check x-y-d is edge
			switch i {
			case 0: //^
				if t.isEdge(grid, x, y+1, p.x, p.y) && !t.isEdge(grid, x, y-1, p.x, p.y) {
					edge--
				}
			case 1: // >
				if t.isEdge(grid, x+1, y, p.x, p.y) && !t.isEdge(grid, x-1, y, p.x, p.y) {
					edge--
				}
			case 2: //v
				if t.isEdge(grid, x, y+1, p.x, p.y) && !t.isEdge(grid, x, y-1, p.x, p.y) {
					edge--
				}
			case 3: //<
				if t.isEdge(grid, x+1, y, p.x, p.y) && !t.isEdge(grid, x-1, y, p.x, p.y) {
					edge--
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
