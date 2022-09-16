package main

import (
	"fmt"
	"log"
)

type Task struct {
	id       int
	priority int
}

func (t *Task) String() string {
	res := fmt.Sprintf("[id : %v, priority : %v]", t.id, t.priority)
	return res
}

type Queue struct {
	queue []*Task
	first int
}

func (q *Queue) SetFirst() {
	maxFunc := func(arr []*Task) int {
		max := arr[0].priority
		for _, value := range arr {
			if value.priority > max {
				max = value.priority
			}
		}
		return max
	}
	q.first = maxFunc(q.queue)
}

func (q *Queue) Next() *Task {
	return q.queue[0]
}

func (q *Queue) PopAndPush() *Task {
	t := q.queue[0]
	q.queue = q.queue[1:]
	q.queue = append(q.queue, t)
	return t
}

func (q *Queue) Pop() *Task {
	t := q.queue[0]
	q.queue = q.queue[1:]
	if t.priority == q.first {
		q.SetFirst()
	}
	return t
}

func solution(priorities []int, location int) int {
	queue := []*Task{}
	for i, v := range priorities {
		queue = append(queue, &Task{
			id:       i,
			priority: v,
		})
	}

	q := &Queue{
		queue: queue,
	}
	q.SetFirst()

	cnt := 1
	log.Println(q.Next().id, location)
	for q.Next().id != location || q.Next().priority != q.first {
		log.Println(q.queue)
		if q.Next().priority == q.first {
			q.Pop()
			cnt++
		} else {
			q.PopAndPush()
		}
	}

	return cnt
}

func main() {
	a := solution([]int{2, 1, 1, 2, 1}, 4) //
	//a := solution([]int{2, 1, 3, 2}, 2)
	//a := solution([]int{1, 1, 9, 1, 1, 1}, 0)
	//a := solution([]int{9, 9, 9, 9, 9}, 3)
	log.Println(a)
}
