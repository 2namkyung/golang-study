package main

func Method() {
	//메소드
	/*
		Golang은 OOP를 고유의 방식으로 지원
		-> struct는 필드만을 가지고, 메소드는 별도로 분리되어 정의

		* Receiver
		메소드는
		함수 정의에서 func 키워드와 함수명 사이에
		"함수가 어떤 struct를 위한 메소드인지" 명시한다
		-> 메소드가 속한 struct 타입과 struct 변수명을 지정한다

		* Value Receiver vs Pointer Receiver
		1) Value Receiver
		- struct의 데이터를 copy하여 전달
		- 메소드 내에서 해당 struct의 필드값이 변경되더라도 호출자의 데이터는 변경 X

		2) Pointer Receiver
		- struct의 포인터를 전달
		- 메소드 내에서 해당 struct의 필드값이 변경되면 호출자의 데이터는 변경 O

	*/
}

// Rect struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메소드
func (r Rect) Area() int {
	return r.width * r.height
}
