package main

//
//type Case struct {
//	N    int
//	nums []int
//}
//
//func part1(items []Case) int64 {
//	var out int64
//	for _, item := range items {
//		if isValid(item.N, item.nums) {
//			out += int64(item.N)
//		}
//	}
//	return out
//}
//
//func isValid(n int, nums []int) bool {
//	if len(nums) == 1 {
//		return n == nums[0]
//	}
//	cloneAdd := slices.Clone(nums)[1:]
//	cloneAdd[0] += nums[0]
//	add := isValid(n, cloneAdd)
//	cloneDiv := slices.Clone(nums)[1:]
//	cloneDiv[0] *= nums[0]
//	div := isValid(n, cloneDiv)
//	return add || div
//}
//
//func part2(items []Case) int64 {
//	var out int64
//	for _, item := range items {
//		if isValid2(item.N, item.nums[0], 0, item.nums) {
//			out += int64(item.N)
//		}
//	}
//	return out
//}
//
//func isValid2(n int, k int, index int, nums []int) bool {
//	if len(nums)-1 == index {
//		return n == k
//	}
//	if k > n {
//		return false
//	}
//	v, _ := strconv.Atoi(fmt.Sprintf("%d%d", k, nums[index+1]))
//	con := isValid2(n, v, index+1, nums)
//	if con {
//		return true
//	}
//	add := isValid2(n, k+nums[index+1], index+1, nums)
//	if add {
//		return true
//	}
//	mul := isValid2(n, k*nums[index+1], index+1, nums)
//	if mul {
//		return true
//	}
//	return false
//}
//
//// 456565678667482
//func main() {
//	raw, _ := os.ReadFile("./input/day7.txt")
//	var TestCases []Case
//
//	for _, line := range strings.Split(string(raw), "\n") {
//		var arr []int
//		v := strings.Split(strings.TrimSpace(line), ":")
//		testN, _ := strconv.Atoi(v[0])
//		for _, n := range strings.Split(strings.TrimSpace(v[1]), " ") {
//			t, _ := strconv.Atoi(n)
//			arr = append(arr, t)
//		}
//		TestCases = append(TestCases, Case{testN, arr})
//	}
//
//	fmt.Println(part1(TestCases))
//	fmt.Println(part2(TestCases))
//
//}
