package main

import "encoding/json"

func JSON() {
	//JSON
	/*
		JSON ( JavaScript Object Notation )

		1. JSON Encoding
		- Encoding : encoding/json 패키지의 Marshal() 함수 이용
		- Decoding : encoding/json 패키지의 Unmarshal() 함수 이용
		- Go struct / map 데이터를 JSON으로 Encoding

		* JSON의 Key는 문자열이어야 한다
		-> map인 경우에 map[string]T처럼 Key가 string인 map만 지원

		* 디코딩하는 과정에서 JSON에 10개의 Key-Value Pair가 있는데, 출력 구조체는 3개 필드만 가진다면 나머지는 무시
		* Go의 JSON 패키지는 출력 구조체나 map을 미리 정의하지 않고 디코딩할 수도 있다
		* JSON 패키지는 Public 필드만 직렬화한다
	*/
}

type Member struct {
	Name   string `json:"user_name"`
	Age    int    `json:"user_age"`
	Active bool   `json:"active"`
}

func JsonEncoding() {

	// Go Data
	member := Member{"Lee", 10, true}

	// JSON Encoding -> 인코딩된 바이트 배열과 에러객체를 리턴
	jsonBytes, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}

	// JSON Byte를 문자열로 변경
	jsonString := string(jsonBytes)
	println(jsonString)
}

func JsonDecoding() {
	// JSON Data
	jsonBytes, _ := json.Marshal(Member{"Kim", 10, true})

	// JSON Decoding -> 에러객체를 리턴
	var mem Member
	err := json.Unmarshal(jsonBytes, &mem) // json.Unmarshal(JSON Data, 구조체 / Map Pointer)
	if err != nil {
		panic(err)
	}

	println(mem.Name, mem.Age, mem.Active)
}
