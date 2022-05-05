package main

func Variable() {

	//변수 & 상수
	/*
		Golang에서 변수를 선언 후 사용하지 않으면 error
		변수 : var a int = 10


		상수 : const b int = 10
		-> 상수형 타입은 선언 후 사용하지 않아도 에러는 아니지만 경고 표시
		fmt.Println(a)
	*/

	//데이터 타입
	/*
			boolean -> bool
			정수형 타입 -> int , int8 , int16 , int32 , int64
			실수 및 복소수 -> float32 , float64 , complex64 , complex128
			기타 타입 -> byte : uint8과 동일하며 바이트 코드에 사용
		                rune : int32와 동일하며 유니코드 포인트에 사용

			* 데이터 타입 변환
			-> Type(value)
			-> ex) bytes := []byte(str) , str := string(bytes)
	*/

	//문자열
	/*
		1. Back Quote(``)
		-> Raw String Literal
		-> 복수 라인의 문자열 표현할 때 자주 사용한다

		2. 이중 인용 부호("")
		-> Interpreted String Literal
		-> 복수 라인 불가, 인용 부호 안의 Escape 문자열들은 특별한 의미로 해석
		-> 복수 라인을 사용하기 위해서는 + 연산자를 이용해 결합하여 사용
	*/

	//연산자
	/*
		다른 언어들과 같은 연산자를 제공하는 것 같다.
		포인터 연산자를 제공하지만, 포인터 산술은 불가능하다는 다른점이 있다.
	*/
}
