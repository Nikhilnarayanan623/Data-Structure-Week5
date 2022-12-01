package main

import (
	"fmt"
)

func main() {
	sl1 := SinglyLinkedList{}

	sl1.AddMultipleValue()
	sl1.AddValue(12)
	sl1.PrintAllValue()
	fmt.Println("The lenght of list is ", sl1.length)
}

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// to add multiple values
func (list *SinglyLinkedList) AddMultipleValue() {
	var limit int
	fmt.Print("Enter the limit: ")
	fmt.Scanf("%d", &limit)

	for i := 1; i <= limit; i++ {
		var value int
		fmt.Print("Enter the value: ")
		fmt.Scanf("%d", &value)
		list.AddValue(value)
	}
}

// add a sigle value
func (list *SinglyLinkedList) AddValue(data int) {
	temp := &Node{
		data: data,
	}

	if list.head == nil {
		list.head = temp
	} else {
		list.tail.next = temp
	}
	list.tail = temp
	list.length++ // increment the length
}

// print all values
func (list *SinglyLinkedList) PrintAllValue() {
	if list.head == nil {
		fmt.Println("There is No values in List")
		return
	}
	temp := list.head
	fmt.Println("Values in the list")
	for temp != nil {
		fmt.Print((*temp).data, " ")
		temp = temp.next
	}
	fmt.Println()
}
