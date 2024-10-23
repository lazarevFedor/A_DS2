package main

import (
	"container/list"
)

func getMinRunLength(n int) int {
	flag := 0
	for n >= 64 {
		flag |= n & 1
		n >>= 1

	}
	return n + flag
}

func BinarySearch(array []int, element int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := (left + right) / 2
		if mid < element {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}

func ReverseArray(array []int) []int {
	left, right := 0, len(array)-1
	for left < right {
		if array[left] > array[right] {
			array[left], array[right] = array[right], array[left]
		}
		left++
		right--
	}
	return array
}

func MergeSort(left, right []int) ([]int, int) {
	res := make([]int, 0)
	leftCount, rightCount := 0, 0
	i, j := 0, 0
	gallopCount := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			leftCount++
			rightCount = 0
			i++
			if leftCount == 3 {
				gallopCount++
				pos := BinarySearch(left[i:], right[j])
				tmp := i
				for index := 0; index < pos; index++ {
					res = append(res, left[tmp+index])
					i++
				}
				leftCount = 0
			}
		} else {
			res = append(res, right[j])
			j++
			rightCount++
			leftCount = 0
			if rightCount == 3 {
				gallopCount++
				pos := BinarySearch(right[j:], left[i])
				tmp := j
				for index := 0; index < pos; index++ {
					res = append(res, right[tmp+index])
					j++
				}
				rightCount = 0
			}
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res, gallopCount
}

func MergeBlocs(blocs list.List) {
	//stack := list.New()
	//mergeCount := 0
	//for _, block := range blocs {
	//}
}

func Sort(slice []int) []int {
	minRun := getMinRunLength(len(slice))
	index := 0
	runs := list.New()
	run := make([]int, 0)
	FlagLenRun := false
	if len(slice) == 1 {
		return slice
	}
	for index < len(slice) {
		index += 2
		run = slice[index-2 : index]
		if index >= len(slice) {
			if slice[index-2] > slice[index-1] {
				run = ReverseArray(slice)
			}
			runs.PushBack(run)
		}
		if slice[index-2] > slice[index-1] {
			for slice[index-1] > slice[index] {
				run = append(run, slice[index])
				index++
				if index > len(slice)-1 {
					break
				}
			}
		} else if slice[index-2] <= slice[index-1] {
			for slice[index-1] <= slice[index] {
				run = append(run, slice[index])
				index++
				if index > len(slice)-1 {
					break
				}
			}
		}
		for len(run) < minRun && index <= len(slice)-1 {
			FlagLenRun = true
			run = append(run, slice[index])
			index++
		}
		if FlagLenRun == true {
			run = InsertionSort(run)
		} else {
			run = ReverseArray(run)
		}
		runs.PushBack(run)
	}
	return slice
}
