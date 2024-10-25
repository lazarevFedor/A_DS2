package main

import "fmt"

func main() {
	arr := make([]int, 100)
	for i := 0; i < 50; i++ {
		arr[i] = i + 1
	}
	for i := 50; i < 100; i++ {
		arr[i] = 100 - i
	}
	arr = Sort(arr)
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
