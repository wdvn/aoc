package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day11 struct {
	mp map[string]int64
}

func (t *Day11) totalStoneByBlink(n, blink int64) int64 {
	var out int64
	if blink == 0 {
		return 0
	}
	k := fmt.Sprintf("%d-%d", n, blink)
	if v := t.mp[k]; v != 0 {
		return v
	}
	if n == 0 {
		if blink == 1 {
			out += 1
		} else {
			out += t.totalStoneByBlink(1, blink-1)
		}

	} else if len(fmt.Sprint(n))%2 == 0 {
		r := fmt.Sprint(n)
		n1, _ := strconv.ParseInt(r[:len(r)/2], 10, 64)
		n2, _ := strconv.ParseInt(r[len(r)/2:], 10, 64)
		if blink == 1 {
			out += 2
		} else {
			out += t.totalStoneByBlink(n1, blink-1)
			out += t.totalStoneByBlink(n2, blink-1)
		}

	} else {
		if blink == 1 {
			out += 1
		} else {
			out += t.totalStoneByBlink(n*2024, blink-1)
		}
	}
	t.mp[k] = out
	return out
}

func (t *Day11) part(raw []byte, blink int64) any {
	start := t.parse(raw)
	var out int64
	for _, n := range start {
		v := t.totalStoneByBlink(int64(n), blink)
		out += v
	}
	return out
}

func (t *Day11) part1(raw []byte) any {
	return t.part(raw, 25)
}
func (t *Day11) part2(raw []byte) any {
	return t.part(raw, 75)
}

func (t *Day11) parse(raw []byte) []int {
	var out []int
	t.mp = map[string]int64{}
	for _, s := range strings.Split(string(raw), " ") {
		n, _ := strconv.Atoi(s)
		out = append(out, n)
	}
	return out
}
