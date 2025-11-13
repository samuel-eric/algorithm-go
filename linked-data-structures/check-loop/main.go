package main

import (
	"fmt"
	"strings"
)

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}

func (list *LinkedList) hasLoop() bool {
	fast := list.sentinel.next
	slow := list.sentinel.next

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}

	return false
}

func (list *LinkedList) toStringMax(seperator string, max int) string {
	var values []string
	pointer := list.sentinel.next
	count := 0
	for pointer != nil && count < max {
		values = append(values, pointer.data)
		pointer = pointer.next
		count++
	}
	return strings.Join(values, seperator)
}
