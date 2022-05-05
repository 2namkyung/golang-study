package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait()

	//graph.go test
	// var adjList [][]int

	// if err := ReadFrom(os.Stdin, &adjList); err != nil {
	// 	fmt.Println(err)
	// }

	// err := WriteTo(os.Stdout, adjList)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
