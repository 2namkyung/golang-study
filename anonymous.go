package main

// type문을 사용한 함수 원형 정의
type calculator func(int, int) int

func Anonymous() {
	//Anonymous
	/*
		 * 익명함수
		- 함수명을 갖지 않는 함수
		- 함수 전체를 변수에 할당
		- 다른 함수의 파라미터에 직접 정의되어 사용

		* 일급함수
		- Golang은 기본 일급함수로서 Go의 기본 타입과 동일하게 취급
		 -> 함수의 파라미터 , 함수의 리턴 값으로도 함수를 사용할 수 있다

		* type문을 사용한 함수 원형 정의
		- type문은 구조체(struct), 인터페이스(interface) 등 Custom Type을 정의 하기 위해 사용

	*/

	// 익명함수
	sum := func(numbers ...int) int {
		s := 0
		for _, i := range numbers {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	println(result)

	// 일급함수
	add := func(i, j int) int {
		return i + j
	}

	r := calc(add, 10, 20)
	println(r)

	// type문을 사용한 함수 원형 정의
	r2 := calcUsingType(add, 10, 20)
	println(r2)
}

// 일급함수 : 함수의 파라미터, 리턴값으로도 함수를 사용할 수 있다
func calc(f func(int, int) int, a, b int) int {
	result := f(a, b)
	return result
}

// type문을 사용한 함수 원형 정의
func calcUsingType(f calculator, a, b int) int {
	result := f(a, b)
	return result
}
