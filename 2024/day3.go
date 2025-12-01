package _024

//import (
//	"fmt"
//	"os"
//	"regexp"
//	"strconv"
//)
//
//func cal(sub [][]byte) int64 {
//	n1, _ := strconv.Atoi(string(sub[1]))
//	n2, _ := strconv.Atoi(string(sub[2]))
//	return int64(n1 * n2)
//}
//
//func part1(s []byte) int64 {
//	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
//
//	subs := reg.FindAllSubmatch(s, -1)
//	var out int64
//	for _, sub := range subs {
//		out += cal(sub)
//	}
//	return out
//}
//
//func cal2(sub [][]byte, enable bool) (int64, bool) {
//	switch sub[0][3] {
//	case ')':
//		return 0, true
//	case "'"[0]:
//		return 0, false
//	default:
//		if !enable {
//			return 0, false
//		}
//		return cal(sub), true
//	}
//}
//
//func part2(s []byte) int64 {
//	enable := true
//	reg := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)
//	var out, v int64
//	subs := reg.FindAllSubmatch(s, -1)
//	for _, sub := range subs {
//		v, enable = cal2(sub, enable)
//		out += v
//	}
//	return out
//}
//
//func main() {
//	raw, _ := os.ReadFile("./input/day3.txt")
//
//	fmt.Println(part1(raw))
//	fmt.Println(part2(raw))
//}
