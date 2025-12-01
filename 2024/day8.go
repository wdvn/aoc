package _024

import (
	"fmt"
	"strings"
)

type Day8 struct{}

func getKey(p pair) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func getByMid(p, mid pair) pair {
	return pair{2*mid.x - p.x, 2*mid.y - p.y}
}

func parse(raw []byte) ([][]byte, map[rune][]pair) {
	groups := map[rune][]pair{}
	var grid [][]byte
	for i, line := range strings.Split(string(raw), "\n") {
		for j, c := range line {
			p := pair{i, j}
			if c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
				groups[c] = append(groups[c], p)
			}
		}
		grid = append(grid, []byte(line))
	}
	return grid, groups
}

func (d Day8) part1(raw []byte) any {
	grid, groups := parse(raw)
	var out int
	var maxX, maxY int = len(grid), len(grid[0])
	visited := cloneMatrix(grid)
	for _, group := range groups {
		for i := range group {
			for j := 0; j < len(group); j++ {
				if i == j {
					continue
				}
				a, b := group[i], group[j]
				v := getByMid(a, b)
				if v.x > -1 && v.x < maxX && v.y > -1 && v.y < maxY && visited[v.y][v.x] != '#' {
					visited[v.y][v.x] = '#'
					out++
				}
			}
		}
	}
	printMatrix(visited)
	return out
}

/*
^
y
|-----------
|----C------
|---B*------
|--A**------
0---------x>
*/

func isPointOnLine(a, b, c pair) bool {
	n1 := b.y - a.y
	n2 := c.x - a.x
	n3 := b.x - a.x
	return float64(c.y-a.y)-float64(n1)*float64(n2)/float64(n3) == 0
}
func (d Day8) part2(raw []byte) any {
	grid, groups := parse(raw)
	var out int
	visited := cloneMatrix(grid)

	for _, group := range groups {
		for i := range group {
			for j := 0; j < len(group); j++ {
				if i == j {
					continue
				}
				a, b := group[i], group[j]
				for x, row := range grid {
					for y, _ := range row {
						if visited[x][y] != '#' && isPointOnLine(a, b, pair{x, y}) {
							fmt.Println(a, b, pair{x, y})
							visited[x][y] = '#'
							out++
						}
					}
				}
			}
		}
	}
	printMatrix(visited)
	return out
}
