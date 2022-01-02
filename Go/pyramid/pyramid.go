package pyramid

func Pyramid(n int) [][]int {
	ret := [][]int{}
	inside := []int{}
	for i:=0; i < n; i++ {
		inside = append(inside, 1)
		ret = append(ret, inside)
	}
	return ret
}
