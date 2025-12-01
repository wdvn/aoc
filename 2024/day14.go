package _024

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day14 struct {
}

const W, H = 101, 103

func (d Day14) move(arr *[]int) {
	x1, y1, x2, y2 := (*arr)[0], (*arr)[1], (*arr)[2], (*arr)[3]
	ux := x1 + 1*x2
	uy := y1 + 1*y2
	if ux < 0 {
		n := (-ux / W) + 1
		ux += n * W
	}
	if uy < 0 {
		n := (-uy / H) + 1
		uy += n * H
	}
	ux %= W
	uy %= H
	(*arr)[0], (*arr)[1] = x1, y1
}

func (d Day14) parse(raw []byte) []*[]int {
	var out []*[]int
	reg := regexp.MustCompile(`-?\d+`)
	for _, line := range strings.Split(string(raw), "\n") {
		arr := reg.FindAllString(line, -1)
		var nums []int
		for _, n := range arr {
			v, _ := strconv.Atoi(n)
			nums = append(nums, v)
		}
		out = append(out, &nums)
	}

	return out
}
func isPart(x, y int) (int64, int64, int64, int64) {
	if x < W/2 && y < H/2 {
		return 1, 0, 0, 0
	} else if x < W/2 && y > H/2 {
		return 0, 1, 0, 0
	} else if x > W/2 && y < H/2 {
		return 0, 0, 1, 0
	} else if x > W/2 && y > H/2 {
		return 0, 0, 0, 1
	}
	return 0, 0, 0, 0
}
func (d Day14) part1(raw []byte) any {
	input := d.parse(raw)
	var p1, p2, p3, p4 int64

	grid := make([][]int, 0)
	for _ = range H {
		grid = append(grid, make([]int, W))
	}

	for _, c := range input {
		for _ = range 100 {
			d.move(c)
		}
		x, y := (*c)[0], (*c)[1]
		v1, v2, v3, v4 := isPart(x, y)
		p1 += v1
		p2 += v2
		p3 += v3
		p4 += v4

		grid[y][x]++
	}
	return p1 * p2 * p3 * p4
}

func (d Day14) part2(raw []byte) any {
	input := d.parse(raw)
	grid := make([][]int, 0)
	for _ = range H {
		grid = append(grid, make([]int, W))
	}
	var step int
	for step > -1 {
		mp := map[pair]int{}
		for _, c := range input {
			d.move(c)
			mp[pair{(*c)[0], (*c)[1]}]++
		}
		step++
		if step == 6475 {
			fmt.Println(len(mp), len(input))
		}
		if len(mp) == len(input) {
			fmt.Println(d.renderPixelImage(grid))
			fmt.Println(step)
			break
		}
	}
	return step
}

func (d Day14) renderPixelImage(matrix [][]int) string {
	b := strings.Builder{}
	for _, row := range matrix {
		for _, value := range row {
			if value > 0 {
				b.WriteString("1")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
