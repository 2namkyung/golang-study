package main

import (
	"log"
	"os"
)

func ErrorHandler() {
	//에러처리
	/*
		Golang에는 다른 언어처럼 try~catch와 같은 에러처리 문법이 없다

		error라는 interface 타입이 존재

		Go 함수가 결과와 에러를 함께 리턴한다면,
		에러가 nil인지를 체크하여 에러가 없는 지를 체크할 수 있다

		* error의 Type을 체크해서 에러 타입별로 에러 처리를 하는 방식도 있다
		 -> switch문을 이용하여 '변수명.(type)의 방식으로 타입 체크
	*/
}

func CheckError() {
	// func Open(name string) (file *File, err error)
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())
}
