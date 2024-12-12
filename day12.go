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

func (t Day12) bfs2(grid Matrix[byte], r, c int, visited *Matrix[byte]) (int64, int64) {
	var total, edge int64
	queue := Queue[pair]{}
	var points []pair
	for queue.Len() > 0 {
		q := queue.Pop()
		if (*visited)[q.x][q.y] == '.' {
			continue
		}
		total++
		points = append(points, q)
		(*visited)[q.x][q.y] = '.'
		grid.JumpByFourDirections(q.x, q.y, func(x, y int, item byte) {
			if grid[r][c] == item && (*visited)[x][y] != '.' {
				queue.Add(pair{x, y})
			}
		})
	}
	edge = t.countEdge(points)
	return edge, total
}

func (t Day12) countEdge(points []pair) int64 {
	var out int64
	mpx := map[int][]int{}
	mpy := map[int][]int{}

	for _, p := range points {
		mpx[p.x] = set(append(mpx[p.x], p.y))
		mpy[p.y] = set(append(mpy[p.y], p.x))
	}

	return out
}
func (t Day12) part2(raw []byte) int64 {
	var out int64
	//grid := convertMatrix(raw)
	//visited := cloneMatrix[byte](grid)
	//for r, row := range grid {
	//	for c, _ := range row {
	//		if (visited)[r][c] == '.' {
	//			continue
	//		}
	//		x, k := t.bfs2(grid, r, c, &visited)
	//		out += x * k
	//	}
	//}
	return out
}
