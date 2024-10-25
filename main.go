package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 100)
	for i := 0; i < 50; i++ {
		arr[i] = i + 1
	}
	for i := 50; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", arr[i])
	}
	arr = Sort(arr)
	fmt.Println("\nSorted array:")
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
