package main

import "fmt"

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := []int{}
	right := []int{}
	middle := int((len(arr) - 1) / 2)
	if middle == 0 {
		middle = 1
	}
	for i := 0; i < middle; i++ {
		left = append(left, arr[i])
	}
	for i := middle; i < len(arr); i++ {
		right = append(right, arr[i])
	}
	left = sort(left)
	right = sort(right)
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := []int{}
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				copy(left[0:], left[1:])
				left = left[:len(left) - 1]
			} else {
				result = append(result, right[0])
				copy(right[0:], right[1:])
				right = right[:len(right) - 1]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			copy(left[0:], left[1:])
			left = left[:len(left) - 1]
		} else if len(right) > 0 {
			result = append(result, right[0])
			copy(right[0:], right[1:])
			right = right[:len(right) - 1]
		}
	}
	return result
}

func main() {
	arr := []int{2, 6, 1, 8, 3, 5, 7, 9, 10}
	z := sort(arr)
	fmt.Println(z)
}
