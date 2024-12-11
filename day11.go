package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day11 struct {
}

var mp = map[int][]int{}

func (t Day11) Extract(arr []int) [][]int {
	var out [][]int
	for _, n := range arr {
		if v := mp[n]; v != nil {
			out = append(out, v)
			continue
		}
		if n == 0 {
			out = append(out, []int{1})
		} else if len(strconv.Itoa(n))%2 == 0 {
			r := strconv.Itoa(n)
			n1, _ := strconv.Atoi(r[:len(r)/2])
			n2, _ := strconv.Atoi(r[len(r)/2:])
			out = append(out, []int{n1, n2})
		} else {
			out = append(out, []int{n * 2024})
		}
		mp[n] = out[len(out)-1]
	}
	return out
}

func (t Day11) Concat(arr [][]int) []int {
	var nums []int
	for _, r := range arr {
		for _, n := range r {
			nums = append(nums, n)
		}
	}
	return nums
}

func (t Day11) set(arr []int) []int {
	m := map[int]int{}
	var v []int
	for _, n := range arr {
		if m[n] > 0 {
			continue
		}
		v = append(v, n)
		m[n] = 1
	}
	return v
}

func (t Day11) part1(raw []byte) any {
	start := t.parse(raw)
	for _ = range 25 {
		start = t.Concat(t.Extract(start))
	}
	return len(start)
}

var mp2 = map[string]int64{}

func (t Day11) Extract2(n, blink int64) int64 {
	var out int64
	if blink == 0 {
		return 0
	}
	k := fmt.Sprintf("%d-%d", n, blink)
	if v := mp2[k]; v != 0 {
		return v
	}
	if n == 0 {
		if blink == 1 {
			out += 1
		} else {
			out += t.Extract2(1, blink-1)
		}

	} else if len(fmt.Sprint(n))%2 == 0 {
		r := fmt.Sprint(n)
		n1, _ := strconv.ParseInt(r[:len(r)/2], 10, 64)
		n2, _ := strconv.ParseInt(r[len(r)/2:], 10, 64)
		if blink == 1 {
			out += 2
		} else {
			out += t.Extract2(n1, blink-1)
			out += t.Extract2(n2, blink-1)
		}

	} else {
		if blink == 1 {
			out += 1
		} else {
			out += t.Extract2(n*2024, blink-1)
		}
	}
	mp2[k] = out
	return out
}

func (t Day11) part2(raw []byte) any {
	start := t.parse(raw)
	var out int64
	for _, n := range start {
		v := t.Extract2(int64(n), 200)
		out += v
	}
	return out
}

func (t Day11) parse(raw []byte) []int {
	var out []int
	for _, s := range strings.Split(string(raw), " ") {
		n, _ := strconv.Atoi(s)
		out = append(out, n)
	}
	return out
}
