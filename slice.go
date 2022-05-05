package main

func Slice() {
	//슬라이스 ( Slice )
	/*
		동적인 크기 변경, 부분 배열 발췌 가능

		슬라이스 선언
		- var a []int
		  a = []int{1,2,3}

		* make()함수를 활용한 Slice 생성
		- make([]type, len, capacity) -> len, capacity 지정하지 않을 시 Default : 0
		  len : 슬라이스의 길이 , capacity : 내부 배열의 최대 길이
		  ex) s := make([]int, 5, 10) -> 길이는 5이지만, 용량은 10인 슬라이스 생성

		부분 슬라이스
		- s = s[a:b] -> a <= x < b
		- s = s[:b] -> 0 <= x < b
		- s = s[a:] -> a <= x end

		슬라이스 추가, 병합 (append) / 복사 (copy)
		- s = append(s, 2)
		- s = append(s, 2,3,4)
		-> capacity가 남아 있는 경우,  용량 내에서 슬라이스 길이(len) 변경하여 데이터 append
		capacity가 남아 있지 않는 경우, 현재 용량의 2배에 해당하는 새로운 Array 생성 후 기존 배열 값들 모두 copy

		- f1 := []string{"apple", "banana"}
		- f2 := []string{"grape", "strawberry"}
		- f3 := append(f1, f2) --> 불가능
		 -> f3 := append(f1, f2...) --> 가능

		* 슬라이스로 갖고 있는 자료를 가변인자를 두고 있는 함수로 넘기기 위해서는 ' ... ' 이용
		- lines := []string{"hello", "hi", "bye"}
		- WriteTo(w, lines...)

	*/
}
