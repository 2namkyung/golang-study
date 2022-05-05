package main

func Map() {

	//Map
	/*
		Key에 대응하는 Value를 신속히 찾는 HashTable을 구현한 자료구조

		Map 선언
		- var 변수명 map[KeyType]ValueType
		 ex) var idMap map[int]string
		  -> int Type을 Key , string Type을 Value

		Map 초기화
		- var idMap map[int]string -> idMap은 nil 값을 가진다
		 -> idMap = make(map[int]string)
		 -> idMap := map[int]string{1:"Lee", 2:"Kim"}

		literal을 사용한 초기화
		- tickers := map[string]string{
			"GOOG": "Google Inc",
			"MSFT": "Microsoft",
			"FB": "FaceBook",
		}

		Map 사용
		- 초기화 하지 않은 map에 값을 추가하거나 사용하면 컴파일 에러 발생
		1) 추가 / 갱신
		 - idMap[10] = "Lee"
		 -> 10(Key)에 Lee(Value)를 추가하거나 값이 있다면 갱신

		2) 삭제
		- delete(idMap, Key)

		3) ContainsKey
		- val, exists := tickers["MSFT"]
		if !exists{
		  println("NO MSFT Tickers")
		}

		4) Map 요소 열거 : Using for Loop
		- for key, val := range idMap{
			fmt.Println(key, val)
		}
	*/

}
