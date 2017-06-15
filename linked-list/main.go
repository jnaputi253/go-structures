package main

import (
	"fmt"
	"go-structures/linked-list/list"
)

func main() {
	list := list.New()
	fmt.Printf("Size: %d\n", list.Size())
	fmt.Println("Is empty? ", list.IsEmpty())

	list.AddToFront(3)
	fmt.Printf("Size: %d\n", list.Size())
	fmt.Println("Is empty? ", list.IsEmpty())

	list.PrintForward()
}
