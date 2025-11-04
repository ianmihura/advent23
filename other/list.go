package main

import (
	"fmt"
	"math"
)

type Node struct {
	char rune
	prev *Node
	next *Node
}

func NewNode(char rune) Node {
	var node Node
	node.char = char
	node.prev = nil
	node.next = nil
	return node
}

func NewList(word string) Node {
	head := NewNode(rune(word[0]))
	currentNode := &head
	for i := 1; i < len(word); i++ {
		newNode := NewNode(rune(word[i]))
		currentNode.next = &newNode
		newNode.prev = currentNode
		currentNode = &newNode
	}
	return head
}

func (node *Node) PrintFlat() {
	currentNode := node.getFirst()
	for range node.findLength() {
		fmt.Print(string(currentNode.char), ", ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (node *Node) Print() {
	currentNode := node.getFirst()
	for range node.findLength() {
		if currentNode.prev != nil {
			fmt.Print("prev:", string(currentNode.prev.char))
		} else {
			fmt.Print("prev: nil")
		}

		fmt.Print(", char:", string(currentNode.char))

		if currentNode.next != nil {
			fmt.Print(", next:", string(currentNode.next.char))
		} else {
			fmt.Print(", next: nil")
		}

		currentNode = currentNode.next
		fmt.Print(" // ")
	}
	fmt.Println()
}

func (node *Node) getLast() *Node {
	if node.next == nil {
		return node
	} else {
		return node.next.getLast()
	}
}

func (node *Node) getFirst() *Node {
	if node.prev == nil {
		return node
	} else {
		return node.prev.getFirst()
	}
}

func (node *Node) addToFront(char rune) {
	newNode := NewNode(char)
	first := node.getFirst()
	first.prev = &newNode
	newNode.next = first
}

func (node *Node) addToRear(char rune) {
	newNode := NewNode(char)
	last := node.getLast()
	last.next = &newNode
	newNode.prev = last
}

func (node *Node) deleteFront() *Node {
	first := node.getFirst()
	// break reference from next item that points to itself
	first.next.prev = nil
	// also cleans up current node so its not implicitly hanging
	second := first.next
	first.next = nil
	return second
}

func (node *Node) deleteRear() {
	last := node.getLast()
	// break reference from prev item that points to itself
	last.prev.next = nil
	// also cleans up current node so its not implicitly hanging
	node.prev = nil
}

func (node *Node) isEmpty() bool {
	return node.next == nil && node.prev == nil
}

func (node *Node) findLength() int {
	first := node.getFirst()
	return first._findLength(0)
}

func (node *Node) _findLength(l int) int {
	if node.next == nil {
		return 1
	} else {
		return 1 + node.next._findLength(l)
	}
}

// returns pointer to first element with this value or null
func (node *Node) findV(char rune) *Node {
	first := node.getFirst()
	return first._findV(char)
}
func (node *Node) _findV(char rune) *Node {
	if node.char == char {
		return node
	} else if node.next != nil {
		return node.next._findV(char)
	} else {
		return nil
	}
}

// deletes first element with this value or null
func (node *Node) deleteV(char rune) {
	first := node.getFirst()
	first._deleteV(char)
}
func (node *Node) _deleteV(char rune) *Node {
	if node.char == char {
		node.prev.next = node.next
		node.next.prev = node.prev
		// also cleans up current node so its not implicitly hanging
		node.next = nil
		node.prev = nil
		return nil
	} else if node.next != nil {
		return node.next._deleteV(char)
	} else {
		return nil
	}
}

// insert after ith element an element with val
func (node *Node) insertPosition(i int, char rune) {
	currentNode := node.getFirst()
	length := currentNode.findLength()

	for j := range length {
		if i == j {
			newNode := NewNode(char)
			currentNode.prev.next = &newNode
			newNode.prev = currentNode.prev
			currentNode.prev = &newNode
			newNode.next = currentNode
			return
		} else {
			currentNode = currentNode.next
		}
	}
}

// insert after ith element an element with val
func (node *Node) deletePosition(i int) {
	currentNode := node.getFirst()
	length := currentNode.findLength()

	for j := 0; j < length; j++ {
		if i == j {
			if currentNode.prev != nil {
				currentNode.prev.next = currentNode.next
			}
			if currentNode.next != nil {
				currentNode.next.prev = currentNode.prev
			}
			currentNode.next = nil
			currentNode.prev = nil
			return
		} else {
			currentNode = currentNode.next
		}
	}
}

func (node *Node) IsPalindrome(hasCaps bool) bool {
	lNode := node.getFirst()
	rNode := node.getLast()

	l := node.findLength()
	length := int(math.Floor(float64(l) / 2.0))

	for range length {
		var mod int
		if hasCaps {
			mod = 32
		} else {
			mod = 65
		}

		if int(lNode.char)%mod != int(rNode.char)%mod {
			return false
		}
		lNode = lNode.next
		rNode = rNode.prev
	}
	return true
}

/*
Here we consider the follwoing:
- That the head node = first node = front of the list => node.prev == nil
- That the tail node = last node = rear of the list => node.next == nil

We're testing every funciton, expecting a panic when there's unexpected behaviour
*/
func main() {
	testDLL()
	testPalindrome()
}

func testPalindrome() {
	list_1 := NewList("abcdefgh")
	list_2 := NewList("otto")
	list_21 := NewList("ottO")
	list_3 := NewList("ottto")
	list_4 := NewList("OtTo")

	if list_1.IsPalindrome(false) {
		panic("Error list_1")
	}
	if !list_2.IsPalindrome(false) {
		panic("Error list_2")
	}
	if list_21.IsPalindrome(false) {
		panic("Error list_21")
	}
	if !list_3.IsPalindrome(false) {
		panic("Error list_3")
	}
	if list_4.IsPalindrome(false) == list_4.IsPalindrome(true) {
		panic("Error list_4")
	}

	fmt.Println("testPalindrome OK")
}

func testDLL() {
	list := NewList("abcdefgh")

	if list.getFirst().char != 'a' {
		panic("Error getFirst")
	}
	if list.getLast().char != 'h' {
		list.Print()
		panic("Error getLast")
	}

	list.addToFront('0')
	if list.getFirst().char != '0' {
		panic("Error addToFront")
	}

	list.addToRear('9')
	if list.getLast().char != '9' {
		panic("Error addToRear")
	}

	list = NewList("abcdefgh")
	list = *list.deleteFront()
	if list.getFirst().char != 'b' {
		panic("Error deleteFront")
	}

	list = NewList("abcdefgh")
	list.deleteRear()
	if list.getLast().char != 'g' {
		panic("Error deleteRear")
	}

	list = NewList("abcdefgh")
	list.insertPosition(2, '0')
	if list.next.next.char != '0' {
		panic("Error insertPosition")
	}

	list = NewList("abcdefgh")
	list.deletePosition(2)
	if list.next.next.char != 'd' {
		panic("Error deletePosition")
	}

	list = NewList("abcdefgh")
	if list.findLength() != 8 {
		panic("Error findLength")
	}

	if list.findV('c').char != 'c' {
		panic("Error findV")
	}

	list.deleteV('c')
	if list.findLength() != 7 {
		panic("Error findV")
	}

	node := NewNode('l')
	if !node.isEmpty() || list.isEmpty() {
		panic("Error isEmpty")
	}

	fmt.Println("testDLL OK")
}
