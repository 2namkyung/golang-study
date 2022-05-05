package main

import (
	"math"
)

func Interface() {
	//인터페이스 '확장성'
	/*
		구초제는
		'필드' 들의 집합체

		인터페이스는
		'메소드' 들의 집합체
		- Type이 구현해야 하는 메소드 원형(PrototType)들을 정의
		- 사용자 정의 타입이 interface를 구현하기 위해서는 가진 '모든' 메소드를 구현해야한다

		* Interface 선언
		- struct와 마찬가지로 type문을 사용하여 정의

		* Interface Type
		- func Marshal(v interface{}) ([]byte, error);
		 -> 빈 interface = interface type : interface{}로 표현 -> 어떠한 타입도 담을 수 있는 컨테이너
		 -> Java : Object , C/C++ : void*
		- func Println(a ...interface{}) (n int, err error);

		* Type Assertion
		- Interface Type의 x와 타입 T에 대하여 x.(T)로 표현했을 때,
		  x가 nil이 아니며, x는 T Type에 속하는 다는 점을 확인하는 것

		* Interface 내부에는 Pointer가 저장된다
		- value를 넣어주어도 내부에서는 주소를 저장하기 때문에, 주소를 넣어준 것과 같다

	*/
}

// 인터페이스 정의
type Shape interface {
	Area() float64
	Perimether() float64
}

type Circle struct {
	radius float64
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimether() float64 {
	return 2 * math.Pi * c.radius
}

// 인터페이스 사용
func UsingInterface(shapes ...Shape) {
	for _, s := range shapes {
		a := s.Area()
		println(a)
	}
}

// 빈 인터페이스 사용
func EmptyInterface(v interface{}) {
	println(v) //Tom
}

// Type Assertion
func TypeAssertion() {
	var a interface{} = 1
	i := a       // a와 i는 Dynamic Type , 값은 1
	j := a.(int) // j는 int Type , 값은 1
	println(i)   // 포인터 주소 출력
	println(j)   // 1 출력
}
