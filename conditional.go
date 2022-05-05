package main

func Conditional() {
	//조건문
	/*
		반드시 Boolean 식으로 표현되어야 한다
		조건식 이전에 Optional statement , Expression 사용 가능

		< switch >
		- switch문의 경우에 case문에도 Expression을 쓸 수 있다
		- 다른 언어는 case에 break를 쓰지 않으면 다음 case로 이동하지만
		  Golang은 다음 case로 가지 않는다
		  -> 다음 case를 실행하게 하려면 'fallthrough'문을 명시해주면 된다
		- 변수의 Value 뿐만 아니라 Type에 따라 case로 분기가 가능하다
	*/

	//if문
	i := 1
	max := 10
	if val := i * 2; val < max { // Optional Statement
		println(val)
	}

	//switch문
	var name string
	var category = 1
	switch x := category << 2; x - 1 { // Expression
	case 1:
		name = "Lee"
	case 2:
		name = "Kim"
	default:
		name = "Wow"
	}
	println(name)

}
