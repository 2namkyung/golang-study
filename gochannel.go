package main

import (
	"time"
)

func GoChannel() {
	//Go Channel
	/*
		Go Channel
		- 채널을 통하여 데이터를 주고 받는 통로
		- make() 함수를 통한 생성
		 -> 채널 생성 시 어떤 타입의 데이터를 채널에서 주고 받을 지 미리 지정 해야한다
		- '<-' 연산자를 이용
		- 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock이 필요없다
		 -> 별도의 lock이 없이 데이터 동기화 가능

		* Go Channel은 수신자와 송신자가 서로를 기다리는 속성이 있다
		  -> Go routine이 끝날 때까지 기다리는 기능 구현 가능

		1. Go Channel Buffering
		 1) Unbuffered Channel
		 - 하나의 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있게 된다
		 -> * DeadLock 주의
		  -> 수신자를 기다리고 있는데, 채널을 받는 수신자 Goroutine이 없는 경우

		 2) Buffered Channel
		 - 수신자가 받을 준비가 되어 있지 않을지라도 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행 가능
		 - make(chan type, N) 함수를 통한 생성 , N = 버퍼 갯수

		2. Channel Parameter
		 - 채널을 함수의 파라미터로 전달할 때, 일반적으로는 송수신을 모두 하는 채널을 전달
		  -> 해당 채널로 송신만할 것인지 수신만할 것인지 지정할 수 있다
		  송신 파라미터 : p chan <- int ==> 'chan<-' 사용
		  수신 파라미터 : p <- chan int ==> '<-chan' 사용

		3. Channel Close
		- close()함수를 사용하여 채널을 닫을 수 있다
		 -> 채널을 닫으면 더 이상 송신은 불가능하지만, 수신은 계속 가능하다

		4. Channel Select
		- 복수개의 채널들을 기다리면서 준비된 ( 데이터를 보내온 ) 채널을 실행하는 기능을 제공
		- 채널에 데이터를 전송할 수도 있다

		1> channel에 send할 때는 receive할 준비가 되어 있어야 한다

		2> range는 channel이 close되어야 끝이난다

		3> nil channel에 send/receive하면 영구 대기한다 , 모두 blocking

		4> select문은 무한루프가 아니지만, case가 올 때까지 기다린다
		 - default
		 - select에서 case는 순차적으로 실행되지 않는다
		 - case에 함수가 있다면 그 함수가 끝난 후, 다른 case를 검사한다
	*/
}

func Channel() {
	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123을 보낸다
	}()

	var i = <-ch // 채널로부터 123을 받는다
	println(i)
}

// Go Channel을 이용하여 Goroutine이 끝날 때까지 기다리는 기능 구현
func WaitForFinisingGoroutine() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		done <- true
	}()

	//위의 Goroutine이 끝날 때까지 대기
	<-done
}

func DeadLock() {
	c := make(chan int)
	c <- 1 // 수신 루틴이 없으므로 DeadLock
}

func BufferedChannel() {
	ch := make(chan int, 1)
	// 수신자가 없더라도 보낼 수 있다
	ch <- 101
	println(<-ch)
}

func SendChan(ch chan<- string) {
	ch <- "Data"
	// 만일, 수신 파라미터를 사용하면 에러가 발생한다
	// ex) x := <-ch
}

func ReceiveChan(ch <-chan string) {
	x := <-ch
	println(x)
	// 만일, 송신 파라미터를 사용하면 에러가 발생한다
	// ex) ch<-"data"
}

func ClosingChannel() {
	// 채널 생성
	ch := make(chan int, 2)
	// 채널에 송신
	ch <- 1
	ch <- 2

	// Channel Close
	close(ch)

	// 채널로부터 수신
	println(<-ch)
	println(<-ch)
	if _, success := <-ch; !success {
		println("더이상 데이터가 없습니다")
	}

	// Channel range
	for i := range ch {
		println(i)
	}
}

func SelectGoroutine() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
