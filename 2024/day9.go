package _024

type Day9 struct{}

func (d Day9) part1(raw []byte) int64 {
	var out int64 = 0
	define, _ := genDisk(raw)
	l, r := 0, len(define)-1
	for isFree(define) {
		for define[l] >= 0 {
			l++
		}
		for define[r] < 0 {
			r--
		}
		define[l], define[r] = define[r], define[l]
	}
	for i := range len(define) {
		if define[i] == -1 {
			break
		}
		out += int64(i * define[i])
	}

	return out
}

func (d Day9) part2(raw []byte) int64 {
	var out int64 = 0

	disk, M := genDisk(raw)

	for id := M; id > -1; id-- {
		var l, r []int
		for i := range len(disk) {
			if disk[i] == id {
				r = append(r, i)
			}
		}
		if len(r) == 0 {
			continue
		}
		for i := range len(disk) {
			if disk[i] == id {
				break
			}
			if disk[i] == -1 {
				l = append(l, i)
				if len(l) >= len(r) {
					for x, _ := range r {
						disk[l[x]], disk[r[x]] = disk[r[x]], disk[l[x]]
					}
					l = nil
					continue
				}
			} else {
				l = nil
			}
		}

	}
	//fmt.Println(disk)
	for i := range len(disk) {
		if disk[i] == -1 {
			continue
		}
		out += int64(i * disk[i])
	}
	return out
}

func genDisk(raw []byte) ([]int, int) {
	var define []int
	var id int
	var swap bool
	for _, b := range raw {
		k := int(b - '0')
		for k > 0 {
			if !swap {
				define = append(define, id)
			} else {
				define = append(define, -1)
			}
			k--
		}

		swap = !swap
		if swap {
			id++
		}
	}
	return define, id
}
func isFree(arr []int) bool {
	for i := range len(arr) - 1 {
		if arr[i] == -1 && arr[i+1] > 0 {
			return true
		}
	}
	return false
}
