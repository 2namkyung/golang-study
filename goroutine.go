package main

import (
	"runtime"
	"sync"
)

func Goroutine() {
	// Goroutine
	/*
		Go Runtime이 관리하는 Lightweight 논리적 Thread
		- 비동기적으로 함수 루틴을 실행하므로, 여러 코드를 동시에 실행하는데 사용
		- OS Thread보다 훨씬 가볍게 비동기 처리를 구현하기 위하여 만든 것
		- Go Channel을 통해 Go Routine간의 통신을 쉽게 할 수 있다

		* 익명함수 Goroutine
		- 익명함수를 비동기적으로 실행

		* 다중 CPU 처리
		- Go는 Default로 1개의 CPU를 시분할 처리 ( Concurrent )
		- runtime.GOMAXPROCS(CPU수)
	*/

}

func say(s string) {
	for i := 0; i < 5; i++ {
		println(s, "***", i)
	}
}

func UsingGoRoutine() {
	// 병렬 실행, 실행 순서는 일정하지 않다
	go say("Asycn1")
	go say("Async2")
	go say("Async3")
}

// 익명함수를 사용한 Goroutine
func UsingAnonymousFunc() {
	// sync.WaitGroup -> Goroutine들이 끝날 때까지 기다리는 역할
	var wait sync.WaitGroup
	wait.Add(2) // 몇 개의 Goroutine들을 기다릴 지 지정

	go func() {
		defer wait.Done()
		println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		println(msg)
	}("Hi")

	wait.Wait() // main routine에서 goroutine들이 모두 끝날 때까지 대기
}

func UsingMultiCPU() {
	// 4개의 CPU 사용 , Parallel 처리
	runtime.GOMAXPROCS(4)
}
