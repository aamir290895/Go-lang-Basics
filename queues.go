package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	dequeued := q.items[0]
	q.items = q.items[1:]
	return dequeued, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func queuesTest() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Queue after enqueuing 1, 2, and 3:", queue.items)

	dequeuedItem, err := queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued item from queue:", dequeuedItem)
	}

	peekedItem, err := queue.Peek()
	if err == nil {
		fmt.Println("Peeked item from queue:", peekedItem)
	}

	fmt.Println("Is queue empty?", queue.IsEmpty())
}
