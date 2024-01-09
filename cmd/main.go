package main

import (
	"datastructure/maxheap/internal/domain"
	"fmt"
)

func main()  {
	fmt.Println("Willi3")

	m := &domain.MaxHeap{}

	m.Insert(31)
	m.Insert(44)
	m.Insert(55)
	m.Insert(45)
	m.Insert(67)
	m.Insert(2)
	m.Insert(23)
	m.Insert(123)
	m.Extract()
	m.Print()
}
