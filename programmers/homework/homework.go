package homework

import (
	"fmt"
	"sort"
	"time"
)

//
//const (
//	layout = "15:04"
//)
//
//type homeworks struct {
//	list           []homework
//	ongoing        *homework
//	ongoingEndTime time.Time
//	paused         []homework
//	result         []string
//}
//
//type homework struct {
//	name          string
//	start         time.Time
//	playtime      time.Duration
//	remainingTime time.Duration
//}
//
//func load(plans [][]string) (homeworks homeworks) {
//	for _, plan := range plans {
//		h := homework{}
//		h.name = plan[0]
//		h.start, _ = time.Parse(layout, plan[1])
//		h.playtime, _ = time.ParseDuration(plan[2] + "m")
//		homeworks.list = append(homeworks.list, h)
//	}
//	return
//}
//
//func (h homeworks) sortList() homeworks {
//	sort.Slice(h.list, func(i, j int) bool {
//		return h.list[i].start.Before(h.list[j].start)
//	})
//	return h
//}
//
//func (h homework) String() string {
//	return fmt.Sprintf("과제명: %s, 과제시작 시간: %s, 과제 소요시간: %s, 과제 남은시간: %s", h.name, h.start.Format("15:04"), h.playtime, h.remainingTime)
//}
//
//func (h homeworks) String() string {
//	return fmt.Sprintf("============\n"+
//		"진행중인 과제: %v\n"+
//		"진행중인 과제 종료시점: %s\n"+
//		"일시정지 중인 과제: %v\n"+
//		"결과: %v\n=================================", h.ongoing, h.ongoingEndTime.Format("15:04"), h.paused, h.result)
//}
//
//func solution(plans [][]string) []string {
//	hs := load(plans)
//	hs = hs.sortList()
//
//	loopCnt := 1
//	for _, h := range hs.list {
//		log.Println(loopCnt, "번쨰 : ", h)
//		if hs.ongoing != nil {
//			if hs.ongoingEndTime.Before(h.start) || hs.ongoingEndTime.Equal(h.start) || hs.ongoing.remainingTime == 0 {
//				if hs.ongoingEndTime.Before(h.start) && !hs.ongoingEndTime.Equal(h.start) {
//					// 일시정지된 과제를 처리할 여유가 있다.
//				}
//				log.Println("기존에 진행중이던 과제가 정상적으로 종료됨")
//				// 결과에 넣음
//				log.Println("결과에 넣음 > ", hs.ongoing.name)
//				hs.result = append(hs.result, hs.ongoing.name)
//
//				// 초기화
//				ongoing := h
//				ongoing.remainingTime = ongoing.playtime
//				hs.ongoing = &ongoing
//				hs.ongoingEndTime = h.start.Add(h.playtime)
//
//				log.Println(hs)
//			} else {
//				log.Println("기존에 진행중이던 과제를 일시정지 하고 새로운 과제를 수행함")
//				// 진행중인 작업의 남은 시간을 진행된만큼 빼줌
//				hs.ongoing.remainingTime -= h.start.Sub(hs.ongoing.start)
//				// 진행중인게 있으면 일시정지에 넣어줌
//				hs.paused = append(hs.paused, *hs.ongoing)
//				// 새로운 과제를 진행함
//				ongoing := h
//				ongoing.remainingTime = ongoing.playtime
//				hs.ongoing = &ongoing
//				hs.ongoingEndTime = h.start.Add(h.playtime)
//				log.Println(hs)
//			}
//		} else {
//			log.Println("첫번째 과제")
//			ongoing := h
//			ongoing.remainingTime = ongoing.playtime
//			hs.ongoing = &ongoing
//			hs.ongoingEndTime = h.start.Add(h.playtime)
//			log.Println(hs)
//		}
//		loopCnt++
//	}
//
//	// 님은 과제 처리
//	hs.result = append(hs.result, hs.ongoing.name)
//
//	// 일시정지과제들 처리
//	total := len(hs.paused)
//	for i := range hs.paused {
//		hs.result = append(hs.result, hs.paused[total-i-1].name)
//	}
//
//	return hs.result
//}
//

const (
	layout = "15:04"
)

type homeworks struct {
	list           []homework
	ongoing        *homework
	ongoingEndTime time.Time
	paused         []homework
	result         []string
}

type homework struct {
	name          string
	start         time.Time
	playtime      time.Duration
	remainingTime time.Duration
}

func load(plans [][]string) (homeworks homeworks) {
	for _, plan := range plans {
		h := homework{}
		h.name = plan[0]
		h.start, _ = time.Parse(layout, plan[1])
		h.playtime, _ = time.ParseDuration(plan[2] + "m")
		homeworks.list = append(homeworks.list, h)
	}
	return
}

func (h homeworks) sortList() homeworks {
	sort.Slice(h.list, func(i, j int) bool {
		return h.list[i].start.Before(h.list[j].start)
	})
	return h
}

func (h homework) String() string {
	return fmt.Sprintf("과제명: %s, 과제시작 시간: %s, 과제 소요시간: %s, 과제 남은시간: %s", h.name, h.start.Format("15:04"), h.playtime, h.remainingTime)
}

func (h homeworks) String() string {
	return fmt.Sprintf("============\n"+
		"진행중인 과제: %v\n"+
		"진행중인 과제 종료시점: %s\n"+
		"일시정지 중인 과제: %v\n"+
		"결과: %v\n=================================", h.ongoing, h.ongoingEndTime.Format("15:04"), h.paused, h.result)
}

func solution(plans [][]string) []string {
	hs := load(plans)
	hs = hs.sortList()

	var result []string
	lastFinishedTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 0, 0, time.Local)

	for t := hs.list[0].start; t.Before(lastFinishedTime); t = t.Add(time.Minute) {
		if hs.ongoing == nil {
			hs.ongoing = &hs.list[0]
			hs.ongoingEndTime = hs.list[0].start.Add(hs.list[0].playtime)
			continue
		}

		if t.After(hs.ongoingEndTime) {
			result = append(result, hs.ongoing.name)
			if len(hs.paused) > 0 {
				hs.ongoing = &hs.paused[len(hs.paused)-1]
				hs.ongoing.remainingTime = hs.ongoing.playtime - (hs.ongoingEndTime.Sub(hs.ongoing.start))
				hs.paused = hs.paused[:len(hs.paused)-1]
			} else {
				var next *homework
				for i := range hs.list {
					if hs.list[i].start.After(hs.ongoingEndTime) {
						next = &hs.list[i]
						break
					}
				}
				if next == nil {
					result = append(result, hs.ongoing.name)
					break
				}
				hs.ongoing = next
				hs.ongoing.remainingTime = hs.ongoing.playtime
			}
			hs.ongoingEndTime = hs.ongoing.start.Add(hs.ongoing.remainingTime)
		}

		if hs.ongoing != nil {
			hs.ongoing.remainingTime -= time.Minute
			if hs.ongoing.remainingTime <= 0 {
				result = append(result, hs.ongoing.name)
				if len(hs.paused) > 0 {
					hs.ongoing = &hs.paused[len(hs.paused)-1]
					hs.ongoing.remainingTime = hs.ongoing.playtime - (hs.ongoingEndTime.Sub(hs.ongoing.start))
					hs.paused = hs.paused[:len(hs.paused)-1]
				} else {
					var next *homework
					for i := range hs.list {
						if hs.list[i].start.After(hs.ongoingEndTime) {
							next = &hs.list[i]
							break
						}
					}
					if next == nil {
						break
					}
					hs.ongoing = next
					hs.ongoing.remainingTime = hs.ongoing.playtime
				}
				hs.ongoingEndTime = hs.ongoing.start.Add(hs.ongoing.remainingTime)
			}
		}
	}

	return result
}
