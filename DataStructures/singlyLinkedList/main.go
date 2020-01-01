package main

import "fmt"

type node struct {
	value int
	next  *node
}
type sll struct {
	head *node
	tail int
}

func (s *sll) addAfter(i, v int) error {
	n := &node{
		value: v,
	}
	if s.head == nil {
		s.head = n
	} else {
		counter := 1
		currentNode := s.head
		if currentNode == nil {
			fmt.Println("List is empty.")
			return nil
		}
		for currentNode.next != nil {
			if i == counter {
				n.next = currentNode.next
				currentNode.next = n
			}
			currentNode = currentNode.next
			counter++

		}
		if currentNode.next == nil && counter <= i {
			currentNode.next = n
		}
	}
	s.tail++
	return nil
}

func (s *sll) push(v int) error {
	s.addAfter(s.tail, v)
	return nil
}

func (s *sll) delete(i int) error {
	if i > s.tail {
		return nil
	}
	currentNode := s.head
	if i == 1 {
		s.head = currentNode.next
	} else {
		for counter := 2; counter < i; {
			currentNode = currentNode.next
			counter++
		}
		technical := currentNode.next
		currentNode.next = technical.next
	}
	s.tail--
	return nil
}

func (s *sll) show() error {
	currentNode := s.head
	if currentNode == nil {
		fmt.Println("List is empty.")
		return nil
	}
	fmt.Println(currentNode.value)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Println(currentNode.value)
	}
	return nil
}
func main() {
	var s sll
	s.push(1)
	s.push(3)
	s.addAfter(1, 2)
	s.addAfter(5, 4)
	s.delete(2)
	s.show()
}
