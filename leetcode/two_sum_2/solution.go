package main

import "log"

func solution(numbers []int, target int) []int {
	leftPointer := 0
	rightPointer := len(numbers) - 1

	// 투포인터가 겹칠 필요가 없으니 leftPointer와 rightPointer의 차이가 2일 때 까지만 연산하면 된다.
	for leftPointer < rightPointer+1 {
		if numbers[leftPointer]+numbers[rightPointer] > target {
			rightPointer--
		} else if numbers[leftPointer]+numbers[rightPointer] < target {
			leftPointer++
		} else {
			return []int{leftPointer + 1, rightPointer + 1}
		}
	}
	return []int{leftPointer + 1, rightPointer + 1}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	//numbers := []int{1, 2, 3, 14, 17, 19, 71, 83}
	//target := 31
	res := solution(numbers, target)
	log.Println(res)
}
