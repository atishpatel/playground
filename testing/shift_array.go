package main

func shiftInPlace(array []int, i int) []int {
	i = i % len(array)

	for count := 1; count < i; count++ {
		tmp := array[len(array)-1]
		for n := len(array) - 2; n >= 0; n-- {
			array[n+1] = array[n]
		}
		array[0] = tmp
	}
	return array
}

func shiftNewAlloc(array []int, i int) []int {
	i = i % (len(array))
	b := append([]int{}, array[i:]...)
	a := array[:i]
	return append(b, a...)
}

func shiftCopyArray(array []int, i int) []int {
	i = i % (len(array))
	rotated := make([]int, len(array))
	for j := range array {
		newPos := (j + i - 1) % len(array)
		rotated[newPos] = array[j]
	}
	return rotated
}
