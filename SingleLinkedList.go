package main

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	sentinel *Cell
}

type Cell struct {
	data string
	next *Cell
}

func make_linked_list() LinkedList {
	cell := Cell{}
	list := LinkedList{sentinel: &cell}
	return list
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	me.next = after
}

func (me *Cell) delete_after() *Cell {
	if me.next == nil {
		panic("No following cell to delete")
	}
	deleted_cell := me.next
	me.next = deleted_cell.next
	return deleted_cell
}

func (list *LinkedList) add_range(values []string) {
	last_cell := list.sentinel
	for ; last_cell.next != nil; last_cell = last_cell.next {
	}

	for _, v := range values {
		cell := Cell{data: v}
		last_cell.add_after(&cell)
		last_cell = &cell
	}
}

func (list *LinkedList) to_string(separator string) string {
	sbuf := make([]string, 0)
	for p := list.sentinel.next; p != nil; p = p.next {
		sbuf = append(sbuf, p.data)
	}
	return strings.Join(sbuf[:], ",")
}

func (list *LinkedList) length() int {
	len := 0
	for p := list.sentinel.next; p != nil; p = p.next {
		len++
	}
	return len
}

func (list *LinkedList) is_empty() bool {
	return list.sentinel.next == nil
}

func (list *LinkedList) push(v string) {
	cell := Cell{data: v}
	list.sentinel.add_after(&cell)
}

func (list *LinkedList) pop() string {
	return list.sentinel.delete_after().data
}

func main() {
	// small_list_test()

	// Make a list from a slice of values.
	greek_letters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := make_linked_list()
	list.add_range(greek_letters)
	fmt.Println(list.to_string(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := make_linked_list()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.is_empty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.to_string(" "))
	}
}
