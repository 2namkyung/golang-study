package main

type Person struct {
	name string
	age  int
}

type Dict struct {
	data map[int]string
}

// 생성자 함수 정의
func NewDict() *Dict {
	d := Dict{}
	d.data = make(map[int]string)
	return &d // 포인터 전달 주의
}

func UsingStruct() {
	p1 := Person{name: "Lee", age: 26}
	println(p1.name)
	println(p1.age)
	// println(p1)
}

func Struct() {
	//구조체
	/*
		- Custon Data Type
		- 필드들의 집합체
		- 필드들의 컨테이너
		- 행위를 표현하는 메소드는 가지지 않는다

			* struct 선언
			type person struct{
				name string
				age int
			}

			* struct 객체 생성
			1) 빈 객체 생성 후 할당
			- var p1 person
			- p1 = person{"Lee", 20}

			2) 객체 생성 시 초기값 할당
			- p2 := person{name:"Kim", age:20}

			3) 내장함수 new() 사용
			- p := new(person)
			- new()를 사용하면 모든 필드를 Zero Value로 초기화
			- 객체의 포인터를 리턴한다
			 -> p가 포인터라도 필드 액세스 시에 .(dot)을 사용
			 -> 포인터가 자동으로 Dereference된다
			   * C에서 포인트의 경우 ->를 사용하는 문법과 다름을 주의

		    * Go에서 struct는 기본적으로 mutable 개체이다
			 -> 필드값이 변화할 경우 별도로 새로운 개체를 만들지 않고 해당 개체 메모리에서 직접 변경된다

			* struct 개체를 다른 함수의 파라미터로 넘기는 경우는 'Pass by Value' 임을 주의
			  -> 'Pass by Reference로 struct을 전달하고자 한다면, struct의 포인터를 전달

			* 생성자 ( Constructor ) 함수
			- 구조체의 필드가 사용 전에 초기화 되어야 하는 경우 사용할 수 있다 -> map
			- 생성자 함수는 struct를 리턴하는 함수로서 함수 본문에서 필요한 필드를 초기화한다

			* empty struct
			- 크기가 0이고 같은 주소를 가리킨다
			- p1 := struct{} , p2 := struct{} --> p1 == p2

			* struct{} vs &struct{}
			- &struct{}는 struct의 포인터를 받아오기 때문에 참조 횟수가 늘어나 보다 약간 느리다
			- 메소드에 포인터 리시버를 사용하게 될 경우, 함수를 실행하는 객체는 꼭 *struct 타입이어야 한다
			- 멤버변수가 변하지 않아도 된다면 포인터 리시버를 사용하지 않고 struct형으로 선언하여 실행한다

	*/

}
