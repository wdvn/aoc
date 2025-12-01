package _024

//
//import (
//	"fmt"
//	"os"
//	"slices"
//	"strconv"
//	"strings"
//)
//
//func main() {
//
//	abs := func(a int) int {
//		if a <= 0 {
//			return -a
//		}
//		return a
//	}
//
//	conv := func(arr []string) []int {
//		t := make([]int, 0, len(arr))
//		for _, n := range arr {
//			v, _ := strconv.Atoi(n)
//			t = append(t, v)
//		}
//		return t
//	}
//
//	check := func(arr []int) bool {
//		base := arr[0] - arr[1]
//		for i := 0; i < len(arr)-1; i++ {
//			space := arr[i] - arr[i+1]
//			if abs(space) > 3 || abs(space) == 0 || base*space <= 0 {
//				return false
//			}
//		}
//		return true
//	}
//
//	raw, _ := os.ReadFile("./input/day2.txt")
//	var out int
//	for _, line := range strings.Split(string(raw), "\n") {
//		nums := strings.Split(line, " ")
//		if len(nums) < 2 {
//			out++
//			continue
//		}
//		numbers := conv(nums)
//		for i := range len(numbers) {
//			if check(slices.Delete(slices.Clone(numbers), i, i+1)) {
//				out++
//				break
//			}
//		}
//
//	}
//	fmt.Println(out)
//}
