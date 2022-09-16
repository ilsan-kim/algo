package main

import "log"

type QueueWithSum struct {
	queue []int
	sum   int
}

func (q *QueueWithSum) Pop() int {
	t := q.queue[0]
	q.queue = q.queue[1:]
	q.sum -= t
	return t
}

func (q *QueueWithSum) Push(t int) {
	q.queue = append(q.queue, t)
	q.sum += t
}

func solution(queue1, queue2 []int) int {
	sum := func(arr []int) int {
		res := 0
		for _, v := range arr {
			res += v
		}
		return res
	}

	cnt := 0
	q1 := QueueWithSum{sum: sum(queue1), queue: queue1}
	q2 := QueueWithSum{sum: sum(queue2), queue: queue2}

	for q1.sum != q2.sum {
		if q1.sum > q2.sum {
			t := q1.Pop()
			q2.Push(t)
			cnt += 1

		} else if q2.sum > q1.sum {
			t := q2.Pop()
			q1.Push(t)
			cnt += 1
		}
		log.Printf("Q1 :::: cnt >> %v queue >> %v", cnt, q1.queue)
		log.Printf("Q2 :::: cnt >> %v queue >> %v", cnt, q2.queue)
		if cnt > len(queue1)*3 {
			return -1
		}
	}

	return cnt
}

func main() {
	queue1 := []int{1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9, 1, 1, 1, 8, 10, 9}
	queue2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	a := solution(queue1, queue2)
	log.Println(a)
}
