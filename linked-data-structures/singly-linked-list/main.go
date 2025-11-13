package main

import "fmt"

func main() {
	// smallListTest()

	// Make a list from an array of values.
	greekLetters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := makeLinkedList()
	list.addRange(greekLetters)
	list2 := makeLinkedList()
	list2.addRange([]string{"a", "b", "c"})
	list.addList(&list2)
	list.append("d")
	list.removeAt(1)
	list.remove("a")
	fmt.Println(list.toString(" "))
	fmt.Printf("List contains 'ε': %v\n", list.contains("ε"))
	fmt.Printf("List contains 'e': %v\n", list.contains("e"))
	list.clear()
	fmt.Println("Clearing list...")
	fmt.Println("List length: ", list.length())
	fmt.Println()

	// Demonstrate a stack.
	stack := makeLinkedList()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.isEmpty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.toString(" "))
	}
}
