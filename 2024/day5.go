package _024

//
//import (
//	"fmt"
//	"os"
//	"slices"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func isCorrect(order map[int][]int, arr []int) bool {
//	for i := range len(arr) {
//		for j := i + 1; j < len(arr); j++ {
//			if slices.Contains(order[arr[i]], arr[j]) {
//				return false
//			}
//		}
//	}
//
//	return true
//}
//
//func part1(order map[int][]int, update [][]int) int {
//	var out int
//	for _, up := range update {
//		tmp := slices.Clone(up)
//		if !isCorrect(order, tmp) {
//			continue
//		}
//
//		mid := len(tmp) / 2
//		out += tmp[mid]
//	}
//	return out
//}
//
//func part2(order map[int][]int, update [][]int) int {
//	var out int
//	for _, up := range update {
//		if !isCorrect(order, up) {
//			sort.Slice(up, func(i, j int) bool {
//				if slices.Contains(order[up[i]], up[j]) {
//					return false
//				}
//				return true
//			})
//			mid := len(up) / 2
//			out += up[mid]
//		}
//
//	}
//	return out
//}
//
//func main() {
//	raw, _ := os.ReadFile("./input/day5.txt")
//	order := map[int][]int{}
//	var update [][]int
//	for _, line := range strings.Split(string(raw), "\n") {
//		if strings.Contains(line, "|") {
//			n1, _ := strconv.Atoi(line[:2])
//			n2, _ := strconv.Atoi(line[3:])
//			order[n2] = append(order[n2], n1)
//		}
//		if strings.Contains(line, ",") {
//			var tmp []int
//			arr := strings.Split(line, ",")
//			for _, n := range arr {
//				v, _ := strconv.Atoi(n)
//				tmp = append(tmp, v)
//			}
//			update = append(update, tmp)
//		}
//	}
//	v1 := part1(order, update)
//	fmt.Println(v1)
//	v2 := part2(order, update)
//	fmt.Println(v2)
//}
