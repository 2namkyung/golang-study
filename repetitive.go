package main

func Repetitive() {

	// 기본적인 for문
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// for range문
	names := []string{"Lee", "Kim", "Park"}
	for index, name := range names {
		println(index, name)
	}

}
