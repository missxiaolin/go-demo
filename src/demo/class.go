package main

import (
	"fmt"
	"demo/queue"
)
func main() {
	fmt.Println("helow")
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
}
