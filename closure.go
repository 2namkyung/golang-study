package main

func Closure() {

	// 클로저 ( Closure )
	/*
		Closure : 함수 바깥에 있는 변수를 참조하는 함수값

	*/

	next := NextValue()
	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := NextValue()
	println(anotherNext()) // 1
	println(anotherNext()) // 2

}

// nextValue : int를 리턴하는 익명함수를 리턴하는 함수
func NextValue() func() int {
	i := 0
	return func() int {
		// 익명함수가 함수 바깥에 있는 변수 i를 참조
		i++
		return i
	}
}
