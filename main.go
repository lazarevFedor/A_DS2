package main

func main() {
	arr := make([]int, 100)
	var tim Timsort
	for i := 0; i < 50; i++ {
		arr[i] = i + 1
	}
	for i := 50; i < 100; i++ {
		arr[i] = 100 - i
	}
	arr = tim.Sort(arr)
}
