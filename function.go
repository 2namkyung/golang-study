package main

func Function() {
	//함수
	/*
		- 파라미터 전달 방식
		1. Pass By Value - copy value
		2. Pass By Reference - &연산자

		- 가변 인자 함수
		다양한 숫자의 파라미터를 전달하고자 할 때 사용
		' msg ...string '

		- 복수 개의 값들을 리턴할 수 있다
		 * Named Return Parameter : 리턴 파라미터명과 타입을 함께 정의 가능
		  -> return 을 써주지 않으면 에러 발생하므로 반드시 써주어야 한다

	*/

}

func Sum(number ...int) (int, int) {
	// 가변 인자 함수
	count := 0
	sum := 0
	for _, value := range number {
		count++
		sum += value
		println(value)
	}

	return count, sum
}
