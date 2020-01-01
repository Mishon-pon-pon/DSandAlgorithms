package main

import "fmt"

//BST - binarySearchTree
type BST struct {
	value     []int
	left      []int
	right     []int
	leftFree  int
	rightFree int
	root      int
	index     int
	free      int
}

func (bst *BST) add(params ...int) {
	var val, root, index int
	switch len(params) {
	case 1:
		val = params[0]
		root = bst.root
		index = bst.index
		break
	case 2:
		val = params[0]
		root = params[1]
		index = bst.index
		break
	case 3:
		val = params[0]
		root = params[1]
		index = params[2]
		break
	}

	if bst.root == 0 {
		bst.root = val
		bst.value = append(bst.value, val)
	} else {
		if root > val {
			if bst.left[index] == 0 {
				bst.value = append(bst.value, val)
				bst.left[index] = bst.free
				bst.free++
				bst.index = 0
				return
			} else {
				bst.add(val, bst.value[bst.left[index]], bst.left[index])
			}
		} else {
			if bst.right[index] == 0 {
				bst.value = append(bst.value, val)
				bst.right[index] = bst.free
				bst.free++
				bst.index = 0
				return
			} else {
				bst.add(val, bst.value[bst.right[index]], bst.right[index])
			}
		}
	}

}

func test(params ...int) {
	fmt.Println(params)
}

func main() {
	var tree BST
	tree.right = make([]int, 7)
	tree.left = make([]int, 7)
	tree.free = 1
	tree.add(5)
	tree.add(7)
	tree.add(8)
	tree.add(6)
	tree.add(3)
	tree.add(4)
	tree.add(2)
	fmt.Println("index", []int{0, 1, 2, 3, 4, 5, 6})
	fmt.Println("value", tree.value)
	fmt.Println("right", tree.right)
	fmt.Println("left ", tree.left)
}

/*
index [0 1 2 3 4 5 6]
value [5 7 8 6 3 4 2]
right [1 2 0 0 5 0 0]
left  [4 3 0 0 6 0 0]

             5

    	 /        \

      3              7

   /     \        /     \

2          4    6          8

*/
