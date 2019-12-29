package main

import (
	"fmt"
)

func find(elem int, arr []int) int {
	l := -1
	r := len(arr)
	for r-l > 1 {
		n := (r + l) / 2
		m := int(n)
		if arr[m] >= elem {
			r = m
		} else {
			l = m
		}
	}
	if arr[r] == elem {
		return r
	} else {
		return -1
	}
}

func main() {
	arr := []int{1, 2, 3, 5, 7, 8, 9}
	result := find(2, arr)
	fmt.Println(result)
}
