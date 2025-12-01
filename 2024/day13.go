package _024

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day13 struct {
}

type I13 struct {
	x, y int64
	A    pair64
	B    pair64
}

func (t Day13) parse(raw []byte) []I13 {
	var out []I13
	reg := regexp.MustCompile(`\d+`)
	for _, line := range strings.Split(string(raw), "\n\n") {
		nums := reg.FindAllString(line, -1)
		n1, _ := strconv.ParseInt(nums[0], 10, 64)
		n2, _ := strconv.ParseInt(nums[1], 10, 64)

		n3, _ := strconv.ParseInt(nums[2], 10, 64)
		n4, _ := strconv.ParseInt(nums[3], 10, 64)

		n5, _ := strconv.ParseInt(nums[4], 10, 64)
		n6, _ := strconv.ParseInt(nums[5], 10, 64)
		out = append(out, I13{
			x: n5, y: n6, A: pair64{x: n1, y: n2}, B: pair64{x: n3, y: n4},
		})
	}
	return out
}

func (t Day13) cost(c I13) int64 {
	var out int64 = math.MaxInt64
	for i := range min(c.x, c.y) {
		for j := range min(c.x, c.y) {
			if i*c.A.x+j*c.B.x == c.x && i*c.A.y+j*c.B.y == c.y {
				out = min(out, int64(3*i+j))
			}
		}
	}
	if out == math.MaxInt64 {
		out = 0
	}
	return out
}

// https://vi.wikipedia.org/wiki/Quy_t%E1%BA%AFc_Cramer#Di%E1%BB%85n_gi%E1%BA%A3i_h%C3%ACnh_h%E1%BB%8Dc
func (t Day13) costByCramer(c I13) int64 {
	delta := func(a, b pair64) int64 {
		return a.x*b.y - a.y*b.x
	}

	cost := func(a, b, c pair64) int64 {
		deltaBC := delta(b, c)
		if deltaBC == 0 {
			return 0
		}

		deltaAC := delta(a, c)
		deltaAB := delta(b, a)

		if deltaAC%deltaBC == 0 && deltaAB%deltaBC == 0 {
			x := deltaAC / deltaBC
			y := deltaAB / deltaBC
			return x*3 + y
		}
		return 0
	}

	return cost(pair64{c.x, c.y}, c.A, c.B)

}

func (t Day13) part1(raw []byte) any {
	var out int64
	input := t.parse(raw)
	for _, item := range input {
		out += t.cost(item)
	}
	return out
}

func (t Day13) part2(raw []byte) any {
	var out int64
	input := t.parse(raw)
	for _, item := range input {
		out += t.costByCramer(I13{x: item.x + 1e13, y: item.y + 1e13, A: item.A, B: item.B})
	}
	return out
}
