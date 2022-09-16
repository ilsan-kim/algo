package main

import (
	"log"
)

type queue struct {
	queue  []int32
	length int
}

func (s *queue) Push(v int32) {
	s.queue = append(s.queue, v)
	s.length += 1
}

func (s *queue) Pop() int32 {
	t := s.queue[0]
	s.queue = s.queue[1:]
	s.length -= 1
	return t
}

func (s *queue) Next() int32 {
	if s.length == 0 {
		return 0
	}
	return s.queue[0]
}

// https://leetcode.com/problems/is-subsequence/
func solution(s string, t string) bool {
	q := queue{
		queue:  []int32{},
		length: 0,
	}

	for _, v := range s {
		q.Push(v)
	}

	for _, v := range t {
		if q.Next() == v {
			q.Pop()
		}
	}

	if q.length == 0 {
		return true
	}
	return false
}

func main() {
	s := "acb"
	t := "ahcgdb"
	res := solution(s, t)
	log.Println(res)
}
