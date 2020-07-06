package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5}
	// fmt.Println(permute(v))
	fmt.Println(shiftNewAlloc(v, 3))
	fmt.Println(shiftCopyArray(v, 3))
	fmt.Println(shiftInPlace(v, 3))
}

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	current := nums[0]
	ans := permute(nums[1:])
	result := make([][]int, 0)
	for _, a := range ans {
		for i := range nums {
			var r []int
			if i == 0 {
				r = append([]int{current}, a...)
			} else if i == len(nums)-1 {
				v := append([]int{}, a...)
				r = append(v, current)
			} else {
				// i is in middle of array
				v := append([]int{}, a[:i]...)
				r = append(v, current)
				r = append(r, a[i:]...)
			}
			fmt.Println(i, a, r)
			result = append(result, r)
		}
	}
	return result
}

func pop(a []int) ([]int, int) {
	if len(a) == 0 {
		return a, 0
	}
	v := a[len(a)-1]
	b := a[:len(a)-1]
	return b, v
}
