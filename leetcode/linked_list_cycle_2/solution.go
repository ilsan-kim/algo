package linked_list_cycle_2

/*
https://fierycoding.tistory.com/45
https://mhwan.tistory.com/67
https://leetcode.com/problems/linked-list-cycle-ii/?envType=study-plan&id=level-1
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func Next(node *ListNode) *ListNode {
	if node != nil && node.Next != nil {
		return node.Next
	}
	return nil
}

func solution(head *ListNode) *ListNode {
	turtle := Next(head)
	rabbit := Next(Next(head))

	// 순환한다면 turtle 과 rabbit 은 만날것이고, 순환하지 않으면 끝까지 loop 돌고 다음으로 넘어감
	for turtle != nil && rabbit != nil {
		if turtle == rabbit {
			break
		}
		turtle = Next(turtle)
		rabbit = Next(Next(rabbit))
	}

	// 순환하지 않으면 nil
	if turtle == nil || rabbit == nil {
		return nil
	}

	// 다시 거북이를 첫 시작점으로
	turtle = head

	// 한칸씩 옮겨서 둘이 만나는지점이 시작지점
	for turtle != rabbit {
		turtle = Next(turtle)
		rabbit = Next(rabbit)
	}

	return turtle
}
