package _024

//
//import (
//	"fmt"
//	"os"
//	"regexp"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func parse() ([]int, []int) {
//	raw, _ := os.ReadFile("./input/day1.txt")
//
//	var l, r []int
//	reg := regexp.MustCompile(`(\d+)\s+(\d+)`)
//	for _, line := range strings.Split(string(raw), "\n") {
//		tmp := reg.FindStringSubmatch(line)
//		if len(tmp) < 3 {
//			continue
//		}
//		n1, _ := strconv.Atoi(tmp[1])
//		l = append(l, n1)
//		n2, _ := strconv.Atoi(tmp[2])
//		r = append(r, n2)
//	}
//	sort.Ints(l)
//	sort.Ints(r)
//	return l, r
//}
//
//func main() {
//	l, r := parse()
//	var p1, p2 int
//
//	abs := func(a int) int {
//		if a < 0 {
//			return -a
//		}
//		return a
//	}
//	m1 := make(map[int]int, len(r))
//	for i := range len(r) {
//		m1[r[i]]++
//	}
//	for i := range len(l) {
//		p1 += abs(l[i] - r[i])
//		p2 += l[i] * m1[l[i]]
//	}
//	fmt.Println(p1, p2)
//}
