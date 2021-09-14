package main

import (
	"fmt"

	"go_struct/queue"
	"go_struct/stack"
)

func main() {
	fmt.Println("app is executed...")
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	fmt.Println("===== before =====")
	q.Print()
	fmt.Println("===== before =====")

	fmt.Println("===== dequeue =====")
	fmt.Println(q.Dequeue())
	fmt.Println("===== dequeue =====")

	fmt.Println("===== after =====")
	q.Print()
	fmt.Println("===== after =====")

	fmt.Println("app is executed...")
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println("===== before =====")
	s.Print()
	fmt.Println("===== before =====")

	fmt.Println("===== dequeue =====")
	fmt.Println(s.Pop())
	fmt.Println("===== dequeue =====")

	fmt.Println("===== after =====")
	s.Print()
	fmt.Println("===== after =====")
}
