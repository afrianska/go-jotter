package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Afrian")
	data.PushBack("Eko")
	data.PushBack("Prasetyo")
	data.PushFront("Mamat")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}