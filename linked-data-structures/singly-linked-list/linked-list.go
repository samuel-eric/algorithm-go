package main

import (
	"strings"
)

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() LinkedList {
	return LinkedList{
		sentinel: &Cell{
			data: "SENTINEL",
		},
	}
}

func (c *Cell) addAfter(after *Cell) {
	after.next = c.next
	c.next = after
}

func (c *Cell) deleteAfter() *Cell {
	if c.next == nil {
		panic("no cell after me")
	}
	after := c.next
	c.next = after.next
	after.next = nil
	return after
}

func (list *LinkedList) addRange(values []string) {
	lastCell := list.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}
	for _, value := range values {
		cell := &Cell{value, nil}
		lastCell.addAfter(cell)
		lastCell = cell
	}
}

func (l *LinkedList) addList(list *LinkedList) {
	lastCell := l.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}
	lastCell.next = list.sentinel.next
	list.sentinel.next = nil
}

func (list *LinkedList) append(value string) {
	lastCell := list.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}
	lastCell.addAfter(&Cell{value, nil})
}

func (list *LinkedList) remove(value string) {
	if !list.contains(value) {
		return
	}
	pointer := list.sentinel.next
	for pointer != nil && pointer.next.data != value {
		pointer = pointer.next
	}
	pointer.deleteAfter()
}

func (list *LinkedList) removeAt(index int) {
	if list.length() == 0 || index > list.length() {
		return
	}
	pointer := list.sentinel.next
	count := 0
	for pointer != nil && count < index-1 {
		pointer = pointer.next
		count++
	}
	pointer.deleteAfter()
}

func (list *LinkedList) clear() {
	list.sentinel.next = nil
}

func (list *LinkedList) toSlice() []string {
	var values []string
	pointer := list.sentinel.next
	for pointer != nil {
		values = append(values, pointer.data)
		pointer = pointer.next
	}
	return values
}

func (list *LinkedList) toString(seperator string) string {
	return strings.Join(list.toSlice(), seperator)
}

func (list *LinkedList) length() int {
	length := 0
	pointer := list.sentinel.next
	for pointer != nil {
		length += 1
		pointer = pointer.next
	}
	return length
}

func (list *LinkedList) isEmpty() bool {
	return list.sentinel.next == nil
}

func (list *LinkedList) push(value string) {
	newCell := &Cell{value, nil}
	list.sentinel.addAfter(newCell)
}

func (list *LinkedList) pop() string {
	popCell := list.sentinel.next
	if popCell == nil {
		return ""
	}
	list.sentinel.deleteAfter()
	return popCell.data
}

func (list *LinkedList) contains(value string) bool {
	if list.length() == 0 {
		return false
	}
	pointer := list.sentinel.next
	for pointer != nil {
		if pointer.data == value {
			return true
		}
		pointer = pointer.next
	}
	return false
}
