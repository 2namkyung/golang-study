package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	Encoding & Decoding
	- 많은 데이터를 처리할 때 스트림을 사용하는 것이 효율적

	1. json.NewEncoder
	- GoValue --> JSON

	type Encoder struct
	func NewEncoder(w io.Writer) *Encoder
	func(enc *Encoder) Encode(v interface{}) error

	2. json.NewDecoder
	- JSON --> GoValue

	type Decoder struct
	func NewDecoder(r io.Reader) *Decoder
	func(dec *Decoder) Decode(v interface{}) error


	* Marshal/Unmarshal vs Encoding/Decoding
	- byte slice나 string을 사용하려면 Marshal/Unmarshal 함수가 적합하다
	- 표준 입출력/파일 같인 Reader/Writer 인터페이스를 사용하여 스트림 기반으로 동작하기에는 Encoder/Decoder가 적합하다
	- 처리 속도는 스트림 방식이 더 빠르다


*/
type User struct{}

// convert Govalue to JSON Data
func Encoding() {
	user := new(User)
	// 입력할 데이터를 표준 출력으로 연결하는 스트림
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(user)
	// -> JSON Type Data
}

// convert JSON Data to Govalue
func Decoding() {
	var u User
	dec := json.NewDecoder(os.Stdin)
	dec.Decode(&u)
	fmt.Printf("%+v\n", u)
}
