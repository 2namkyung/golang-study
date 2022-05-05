package main

import (
	"sync"
)

// 게이트 하나의 출력이 게이트 여러 입력으로 들어가는 경우 , 분배 : 채널 공유
func FanOut() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range ch {
				println(i, n)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

// 하나의 게이트에 여러 개의 입력이 들어가는 경우 , 부하 합치기 : 채널 공유
func FanIn(ins ...<-chan int) <-chan int {
	// 보내는 고루틴에서 채널을 닫아버리면 같은 채널을 여러 번 닫는다 -> panic
	// 고루틴들이 모두 종료된 뒤 채널을 닫게하는 하나의 고루틴을 구성
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 분산처리 : 팬아웃 --> 파이프라인 --> 팬인
type IntPipe func(<-chan int) <-chan int

func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

// 사슬처럼 이어진 파이프라인을 하나로 만들기
func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func FanIn_Select(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCount := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil
		openCount--
		return openCount == 0
	}

	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()

	return out
}

func UsingFanIn_Select() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}

	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)

	for n := range FanIn_Select(c1, c2, c3) {
		print(n, ",")
	}

}
