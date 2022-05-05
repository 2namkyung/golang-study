package main

import "os"

func Defer_Panic_Recover() {
	/*
		1. 지연실행 defer
		- 특정 문장 혹은 함수를 'defer를 호출하는 함수가 Return 하기 직전에' 실행하게 한다
		 -> Java의 finally 블록처럼 마지막에 Clean-Up 작업을 위해 주로 사용된다
		 -> ex) File Open 후 Close 작업


		2. 동작그만 panic
		- 현재 함수를 즉시 멈추고 현재 함수에 defer함수들을 모두 실행 한후 즉시 리턴
		 -> 상위 함수에도 똑같이 적용되고, 계속 콜 스택을 타고 올라가며 적용된다
		 -> 마지막에는 프로그램이 에러를 내고 종료한다

		 * panic을 적극적으로 활용하지는 않는다
		  -> 에러를 반환할 수 있는 경우에는 패닉을 활용하지 않고 에러를 반환하는 방식을 이용한다

		3. 원상복구 recover
		- panic함수에 의한 상태를 다시 정상상태로 되돌리는 함수
	*/
}

// 지연실행 defer
func Defer() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	//함수 마지막에 파일 close 실행
	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

// 동작그만 panic
func Panic() {
	// 잘못된 파일명을 입력하여 파일오픈 시도
	openFile("invalid.txt")

	// openFile() 안에서 panic이 실행되면 아래 println 문장은 실행 X
	println("Done")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

// 원상복구 recover
func Recover() {
	// 잘못된 파일명을 입력하여 파일오픈 시도
	openFile2("invaild.txt")

	// recover에 의하여 println문장이 실행된다
	println("Done")
}

func openFile2(fn string) {

	// defer 함수. panic 호출 시에 실행된다
	defer func() {
		if r := recover(); r != nil {
			println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
